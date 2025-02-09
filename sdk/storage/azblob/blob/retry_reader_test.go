//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package blob

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Testings for RetryReader
// This reader return one byte through each Read call
// nolint
type perByteReader struct {
	RandomBytes []byte // Random generated bytes

	byteCount              int // Bytes can be returned before EOF
	currentByteIndex       int // Bytes that have already been returned.
	doInjectError          bool
	doInjectErrorByteIndex int
	doInjectTimes          int
	injectedError          error

	// sleepDuration and closeChannel are only use in "forced cancellation" tests
	sleepDuration time.Duration
	closeChannel  chan struct{}
}

// nolint
func newPerByteReader(byteCount int) *perByteReader {
	perByteReader := perByteReader{
		byteCount:    byteCount,
		closeChannel: nil,
	}

	perByteReader.RandomBytes = make([]byte, byteCount)
	_, _ = rand.Read(perByteReader.RandomBytes)

	return &perByteReader
}

// nolint
func newSingleUsePerByteReader(contents []byte) *perByteReader {
	perByteReader := perByteReader{
		byteCount:    len(contents),
		closeChannel: make(chan struct{}, 10),
	}

	perByteReader.RandomBytes = contents

	return &perByteReader
}

// nolint
func (r *perByteReader) Read(b []byte) (n int, err error) {
	if r.doInjectError && r.doInjectErrorByteIndex == r.currentByteIndex && r.doInjectTimes > 0 {
		r.doInjectTimes--
		return 0, r.injectedError
	}

	if r.currentByteIndex < r.byteCount {
		n = copy(b, r.RandomBytes[r.currentByteIndex:r.currentByteIndex+1])
		r.currentByteIndex += n

		// simulate a delay, which may be successful or, if we're closed from another go-routine, may return an
		// error
		select {
		case <-r.closeChannel:
			return n, errors.New(ReadOnClosedBodyMessage)
		case <-time.After(r.sleepDuration):
			return n, nil
		}
	}

	return 0, io.EOF
}

// nolint
func (r *perByteReader) Close() error {
	if r.closeChannel != nil {
		r.closeChannel <- struct{}{}
	}
	return nil
}

// Test normal retry succeed, note initial response not provided.
// Tests both with and without notification of failures
// nolint
func TestRetryReaderReadWithRetry(t *testing.T) {
	// Test twice, the second time using the optional "logging"/notification callback for failed tries
	// We must test both with and without the callback, since be testing without
	// we are testing that it is, indeed, optional to provide the callback
	for _, logThisRun := range []bool{false, true} {

		// Extra setup for testing notification of failures (i.e. of unsuccessful tries)
		failureMethodNumCalls := 0
		failureWillRetryCount := 0
		failureLastReportedFailureCount := int32(-1)
		var failureLastReportedError error = nil
		failureMethod := func(failureCount int32, lastError error, offset int64, count int64, willRetry bool) {
			failureMethodNumCalls++
			if willRetry {
				failureWillRetryCount++
			}
			failureLastReportedFailureCount = failureCount
			failureLastReportedError = lastError
		}

		// Main test setup
		byteCount := 1
		body := newPerByteReader(byteCount)
		body.doInjectError = true
		body.doInjectErrorByteIndex = 0
		body.doInjectTimes = 1
		body.injectedError = &net.DNSError{IsTemporary: true}

		getter := func(ctx context.Context, info httpGetterInfo) (io.ReadCloser, error) {
			r := http.Response{}
			body.currentByteIndex = int(info.Offset)
			r.Body = body

			return r.Body, nil
		}

		httpGetterInfo := httpGetterInfo{Offset: 0, Count: int64(byteCount)}
		initResponse, err := getter(context.Background(), httpGetterInfo)
		require.NoError(t, err)

		rrOptions := RetryReaderOptions{MaxRetries: 1}
		if logThisRun {
			rrOptions.OnFailedRead = failureMethod
		}
		retryReader := newRetryReader(context.Background(), initResponse, httpGetterInfo, getter, rrOptions)

		// should fail and succeed through retry
		can := make([]byte, 1)
		n, err := retryReader.Read(can)
		require.Equal(t, n, 1)
		require.NoError(t, err)

		// check "logging", if it was enabled
		if logThisRun {
			// We only expect one failed try in this test
			// And the notification method is not called for successes
			require.Equal(t, failureMethodNumCalls, 1)                  // this is the number of calls we counted
			require.Equal(t, failureWillRetryCount, 1)                  // the sole failure was retried
			require.Equal(t, failureLastReportedFailureCount, int32(1)) // this is the number of failures reported by the notification method
			require.Error(t, failureLastReportedError)
		}
		// should return EOF
		n, err = retryReader.Read(can)
		require.Equal(t, n, 0)
		require.Equal(t, err, io.EOF)
	}
}

// Test normal retry succeed, note initial response not provided.
// Tests both with and without notification of failures
// nolint
func TestRetryReaderWithRetryIoUnexpectedEOF(t *testing.T) {
	// Test twice, the second time using the optional "logging"/notification callback for failed tries
	// We must test both with and without the callback, since be testing without
	// we are testing that it is, indeed, optional to provide the callback
	for _, logThisRun := range []bool{false, true} {

		// Extra setup for testing notification of failures (i.e. of unsuccessful tries)
		failureMethodNumCalls := 0
		failureWillRetryCount := 0
		failureLastReportedFailureCount := int32(-1)
		var failureLastReportedError error = nil
		failureMethod := func(failureCount int32, lastError error, offset int64, count int64, willRetry bool) {
			failureMethodNumCalls++
			if willRetry {
				failureWillRetryCount++
			}
			failureLastReportedFailureCount = failureCount
			failureLastReportedError = lastError
		}

		// Main test setup
		byteCount := 1
		body := newPerByteReader(byteCount)
		body.doInjectError = true
		body.doInjectErrorByteIndex = 0
		body.doInjectTimes = 1
		body.injectedError = io.ErrUnexpectedEOF

		getter := func(ctx context.Context, info httpGetterInfo) (io.ReadCloser, error) {
			r := http.Response{}
			body.currentByteIndex = int(info.Offset)
			r.Body = body

			return r.Body, nil
		}

		httpGetterInfo := httpGetterInfo{Offset: 0, Count: int64(byteCount)}
		initResponse, err := getter(context.Background(), httpGetterInfo)
		require.NoError(t, err)

		rrOptions := RetryReaderOptions{MaxRetries: 1}
		if logThisRun {
			rrOptions.OnFailedRead = failureMethod
		}
		retryReader := newRetryReader(context.Background(), initResponse, httpGetterInfo, getter, rrOptions)

		// should fail and succeed through retry
		can := make([]byte, 1)
		n, err := retryReader.Read(can)
		require.Equal(t, n, 1)
		require.NoError(t, err)

		// check "logging", if it was enabled
		if logThisRun {
			// We only expect one failed try in this test
			// And the notification method is not called for successes
			require.Equal(t, failureMethodNumCalls, 1)                  // this is the number of calls we counted
			require.Equal(t, failureWillRetryCount, 1)                  // the sole failure was retried
			require.Equal(t, failureLastReportedFailureCount, int32(1)) // this is the number of failures reported by the notification method
			require.Error(t, failureLastReportedError)
		}
		// should return EOF
		n, err = retryReader.Read(can)
		require.Equal(t, n, 0)
		require.Equal(t, err, io.EOF)
	}
}

// Test normal retry fail as retry Count not enough.
// nolint
func TestRetryReaderReadNegativeNormalFail(t *testing.T) {
	// Extra setup for testing notification of failures (i.e. of unsuccessful tries)
	failureMethodNumCalls := 0
	failureWillRetryCount := 0
	failureLastReportedFailureCount := int32(-1)
	var failureLastReportedError error = nil
	failureMethod := func(failureCount int32, lastError error, offset int64, count int64, willRetry bool) {
		failureMethodNumCalls++
		if willRetry {
			failureWillRetryCount++
		}
		failureLastReportedFailureCount = failureCount
		failureLastReportedError = lastError
	}

	// Main test setup
	byteCount := 1
	body := newPerByteReader(byteCount)
	body.doInjectError = true
	body.doInjectErrorByteIndex = 0
	body.doInjectTimes = 2
	body.injectedError = &net.DNSError{IsTemporary: true}

	startResponse := body

	getter := func(ctx context.Context, info httpGetterInfo) (io.ReadCloser, error) {
		r := http.Response{}
		body.currentByteIndex = int(info.Offset)
		r.Body = body

		return r.Body, nil
	}

	rrOptions := RetryReaderOptions{
		MaxRetries:   1,
		OnFailedRead: failureMethod}
	retryReader := newRetryReader(context.Background(), startResponse, httpGetterInfo{Offset: 0, Count: int64(byteCount)}, getter, rrOptions)

	// should fail
	can := make([]byte, 1)
	n, err := retryReader.Read(can)
	require.Equal(t, n, 0)
	require.Equal(t, err, body.injectedError)

	// Check that we received the right notification callbacks
	// We only expect two failed tries in this test, but only one
	// of the would have had willRetry = true
	require.Equal(t, failureMethodNumCalls, 2)                  // this is the number of calls we counted
	require.Equal(t, failureWillRetryCount, 1)                  // only the first failure was retried
	require.Equal(t, failureLastReportedFailureCount, int32(2)) // this is the number of failures reported by the notification method
	require.Error(t, failureLastReportedError)
}

// Test boundary case when Count equals to 0 and fail.
// nolint
func TestRetryReaderReadCount0(t *testing.T) {
	byteCount := 1
	body := newPerByteReader(byteCount)
	body.doInjectError = true
	body.doInjectErrorByteIndex = 1
	body.doInjectTimes = 1
	body.injectedError = &net.DNSError{IsTemporary: true}

	startResponseBody := body

	getter := func(ctx context.Context, info httpGetterInfo) (io.ReadCloser, error) {
		r := http.Response{}
		body.currentByteIndex = int(info.Offset)
		r.Body = body

		return r.Body, nil
	}

	retryReader := newRetryReader(context.Background(), startResponseBody, httpGetterInfo{Offset: 0, Count: int64(byteCount)}, getter, RetryReaderOptions{MaxRetries: 1})

	// should consume the only byte
	can := make([]byte, 1)
	n, err := retryReader.Read(can)
	require.Equal(t, n, 1)
	require.NoError(t, err)

	// should not read when Count=0, and should return EOF
	n, err = retryReader.Read(can)
	require.Equal(t, n, 0)
	require.Equal(t, err, io.EOF)
}

// nolint
func TestRetryReaderReadNegativeNonRetriableError(t *testing.T) {
	byteCount := 1
	body := newPerByteReader(byteCount)
	body.doInjectError = true
	body.doInjectErrorByteIndex = 0
	body.doInjectTimes = 1
	body.injectedError = fmt.Errorf("not retriable error")

	startResponseBody := body

	getter := func(ctx context.Context, info httpGetterInfo) (io.ReadCloser, error) {
		r := http.Response{}
		body.currentByteIndex = int(info.Offset)
		r.Body = body

		return r.Body, nil
	}

	retryReader := newRetryReader(context.Background(), startResponseBody, httpGetterInfo{Offset: 0, Count: int64(byteCount)}, getter, RetryReaderOptions{MaxRetries: 2})

	dest := make([]byte, 1)
	_, err := retryReader.Read(dest)
	require.Equal(t, err, body.injectedError)
}

// Test the case where we programmatically force a retry to happen, via closing the body early from another goroutine
// Unlike the retries orchestrated elsewhere in this test file, which simulate network failures for the
// purposes of unit testing, here we are testing the cancellation mechanism that is exposed to
// consumers of the API, to allow programmatic forcing of retries (e.g. if the consumer deems
// the read to be taking too long, they may force a retry in the hope of better performance next time).
// nolint
func TestRetryReaderReadWithForcedRetry(t *testing.T) {
	for _, enableRetryOnEarlyClose := range []bool{false, true} {

		// use the notification callback, so we know that the retry really did happen
		failureMethodNumCalls := 0
		failureMethod := func(failureCount int32, lastError error, offset int64, count int64, willRetry bool) {
			failureMethodNumCalls++
		}

		// Main test setup
		byteCount := 10 // so multiple passes through read loop will be required
		sleepDuration := 100 * time.Millisecond
		randBytes := make([]byte, byteCount)
		_, _ = rand.Read(randBytes)
		getter := func(ctx context.Context, info httpGetterInfo) (io.ReadCloser, error) {
			body := newSingleUsePerByteReader(randBytes) // make new one every time, since we force closes in this test, and it is unusable after a close
			body.sleepDuration = sleepDuration
			r := http.Response{}
			body.currentByteIndex = int(info.Offset)
			r.Body = body

			return r.Body, nil
		}

		httpGetterInfo := httpGetterInfo{Offset: 0, Count: int64(byteCount)}
		initResponse, err := getter(context.Background(), httpGetterInfo)
		require.NoError(t, err)

		rrOptions := RetryReaderOptions{MaxRetries: 2, EarlyCloseAsError: !enableRetryOnEarlyClose}
		rrOptions.OnFailedRead = failureMethod
		retryReader := newRetryReader(context.Background(), initResponse, httpGetterInfo, getter, rrOptions)

		// set up timed cancellation from separate goroutine
		go func() {
			time.Sleep(sleepDuration * 5)
			err := retryReader.Close()
			if err != nil {
				return
			}
		}()

		// do the read (should fail, due to forced cancellation, and succeed through retry)
		output := make([]byte, byteCount)
		n, err := io.ReadFull(retryReader, output)
		if enableRetryOnEarlyClose {
			require.Equal(t, n, byteCount)
			require.NoError(t, err)
			require.EqualValues(t, output, randBytes)
			require.Equal(t, failureMethodNumCalls, 1) // assert that the cancellation did indeed happen
		} else {
			require.Error(t, err)
		}
	}
}
