//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package container_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Example_container_NewClient() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)
	fmt.Println(containerClient.URL())
}

func Example_container_NewClientWithSharedKeyCredential() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	accountKey, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_KEY")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_KEY could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)
	containerClient, err := container.NewClientWithSharedKeyCredential(containerURL, cred, nil)
	handleError(err)
	fmt.Println(containerClient.URL())
}

func Example_container_NewClientWithNoCredential() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	sharedAccessSignature, ok := os.LookupEnv("AZURE_STORAGE_SHARED_ACCESS_SIGNATURE")
	if !ok {
		panic("AZURE_STORAGE_SHARED_ACCESS_SIGNATURE could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s?%s", accountName, containerName, sharedAccessSignature)

	containerClient, err := container.NewClientWithNoCredential(containerURL, nil)
	handleError(err)
	fmt.Println(containerClient.URL())
}

func Example_container_NewClientFromConnectionString() {
	// Your connection string can be obtained from the Azure Portal.
	connectionString, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		log.Fatal("the environment variable 'AZURE_STORAGE_CONNECTION_STRING' could not be found")
	}
	containerName := "testcontainer"
	containerClient, err := container.NewClientFromConnectionString(connectionString, containerName, nil)
	handleError(err)
	fmt.Println(containerClient.URL())
}

func Example_container_ClientNewAppendBlobClient() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	appendBlobClient := containerClient.NewAppendBlobClient("test_append_blob")
	handleError(err)
	fmt.Println(appendBlobClient.URL())
}

func Example_container_ClientNewBlobClient() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	blobClient := containerClient.NewBlobClient("test_blob")
	handleError(err)
	fmt.Println(blobClient.URL())
}

func Example_container_ClientNewBlockBlobClient() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	blockBlobClient := containerClient.NewBlockBlobClient("test_block_blob")
	handleError(err)
	fmt.Println(blockBlobClient.URL())
}

func Example_container_ClientNewPageBlobClient() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	pageBlobClient := containerClient.NewPageBlobClient("test_page_blob")
	handleError(err)
	fmt.Println(pageBlobClient.URL())
}

func Example_container_ClientCreate() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	containerCreateResponse, err := containerClient.Create(context.TODO(), &container.CreateOptions{
		Metadata: map[string]string{"Foo": "Bar"},
	})
	handleError(err)
	fmt.Println(containerCreateResponse)
}

func Example_container_ClientDelete() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	containerDeleteResponse, err := containerClient.Delete(context.TODO(), nil)
	handleError(err)
	fmt.Println(containerDeleteResponse)
}

func Example_container_ClientListBlobsFlat() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	pager := containerClient.NewListBlobsFlatPager(&container.ListBlobsFlatOptions{
		Include: []container.ListBlobsIncludeItem{container.ListBlobsIncludeItemSnapshots, container.ListBlobsIncludeItemVersions},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}

func Example_container_ClientListBlobsHierarchy() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	maxResults := int32(5)
	pager := containerClient.NewListBlobsHierarchyPager("/", &container.ListBlobsHierarchyOptions{
		Include: []container.ListBlobsIncludeItem{
			container.ListBlobsIncludeItemMetadata,
			container.ListBlobsIncludeItemTags,
		},
		MaxResults: &maxResults,
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, blob := range resp.ListBlobsHierarchySegmentResponse.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}

func Example_container_ClientGetSASURL() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	permission := sas.ContainerPermissions{Read: true}
	start := time.Now()
	expiry := start.AddDate(1, 0, 0)
	sasURL, err := containerClient.GetSASURL(permission, start, expiry)
	handleError(err)
	_ = sasURL
}

// This example shows how to manipulate a container's permissions.
func Example_container_ClientSetAccessPolicy() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	// Create the container
	_, err = containerClient.Create(context.TODO(), nil)
	handleError(err)

	// Upload a simple blob.
	blockBlobClient := containerClient.NewBlockBlobClient("HelloWorld.txt")
	handleError(err)

	_, err = blockBlobClient.Upload(context.TODO(), streaming.NopCloser(strings.NewReader("Hello World!")), nil)
	handleError(err)

	// Attempt to read the blob
	get, err := http.Get(blockBlobClient.URL())
	handleError(err)
	if get.StatusCode == http.StatusNotFound {
		// ChangeLease the blob to be public access blob
		_, err := containerClient.SetAccessPolicy(
			context.TODO(),
			nil,
			&container.SetAccessPolicyOptions{
				Access: to.Ptr(container.PublicAccessTypeBlob),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		// Now, this works
		get, err = http.Get(blockBlobClient.URL())
		if err != nil {
			log.Fatal(err)
		}
		var text bytes.Buffer
		_, err = text.ReadFrom(get.Body)
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(get.Body)

		fmt.Println("Public access blob data: ", text.String())
	}
}

func Example_container_ClientSetMetadata() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	containerName := "testcontainer"
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	containerClient, err := container.NewClient(containerURL, cred, nil)
	handleError(err)

	// Create a container with some metadata, key names are converted to lowercase before being sent to the service.
	// You should always use lowercase letters, especially when querying a map for a metadata key.
	creatingApp, err := os.Executable()
	handleError(err)
	_, err = containerClient.Create(context.TODO(), &container.CreateOptions{Metadata: map[string]string{"author": "azblob", "app": creatingApp}})
	handleError(err)

	// Query the container's metadata
	containerGetPropertiesResponse, err := containerClient.GetProperties(context.TODO(), nil)
	handleError(err)

	if containerGetPropertiesResponse.Metadata == nil {
		log.Fatal("metadata is empty!")
	}

	for k, v := range containerGetPropertiesResponse.Metadata {
		fmt.Printf("%s=%s\n", k, v)
	}

	// Update the metadata and write it back to the container
	containerGetPropertiesResponse.Metadata["author"] = "Mohit"
	_, err = containerClient.SetMetadata(context.TODO(), &container.SetMetadataOptions{Metadata: containerGetPropertiesResponse.Metadata})
	handleError(err)

	// NOTE: SetMetadata & SetProperties methods update the container's ETag & LastModified properties
}
