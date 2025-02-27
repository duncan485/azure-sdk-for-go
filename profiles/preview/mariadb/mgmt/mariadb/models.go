//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package mariadb

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreateMode = original.CreateMode

const (
	CreateModeDefault                   CreateMode = original.CreateModeDefault
	CreateModeGeoRestore                CreateMode = original.CreateModeGeoRestore
	CreateModePointInTimeRestore        CreateMode = original.CreateModePointInTimeRestore
	CreateModeReplica                   CreateMode = original.CreateModeReplica
	CreateModeServerPropertiesForCreate CreateMode = original.CreateModeServerPropertiesForCreate
)

type GeoRedundantBackup = original.GeoRedundantBackup

const (
	Disabled GeoRedundantBackup = original.Disabled
	Enabled  GeoRedundantBackup = original.Enabled
)

type OperationOrigin = original.OperationOrigin

const (
	NotSpecified OperationOrigin = original.NotSpecified
	System       OperationOrigin = original.System
	User         OperationOrigin = original.User
)

type PrivateEndpointProvisioningState = original.PrivateEndpointProvisioningState

const (
	Approving PrivateEndpointProvisioningState = original.Approving
	Dropping  PrivateEndpointProvisioningState = original.Dropping
	Failed    PrivateEndpointProvisioningState = original.Failed
	Ready     PrivateEndpointProvisioningState = original.Ready
	Rejecting PrivateEndpointProvisioningState = original.Rejecting
)

type PrivateLinkServiceConnectionStateActionsRequire = original.PrivateLinkServiceConnectionStateActionsRequire

const (
	None PrivateLinkServiceConnectionStateActionsRequire = original.None
)

type PrivateLinkServiceConnectionStateStatus = original.PrivateLinkServiceConnectionStateStatus

const (
	Approved     PrivateLinkServiceConnectionStateStatus = original.Approved
	Disconnected PrivateLinkServiceConnectionStateStatus = original.Disconnected
	Pending      PrivateLinkServiceConnectionStateStatus = original.Pending
	Rejected     PrivateLinkServiceConnectionStateStatus = original.Rejected
)

type PublicNetworkAccessEnum = original.PublicNetworkAccessEnum

const (
	PublicNetworkAccessEnumDisabled PublicNetworkAccessEnum = original.PublicNetworkAccessEnumDisabled
	PublicNetworkAccessEnumEnabled  PublicNetworkAccessEnum = original.PublicNetworkAccessEnumEnabled
)

type ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyState

const (
	ServerSecurityAlertPolicyStateDisabled ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyStateDisabled
	ServerSecurityAlertPolicyStateEnabled  ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyStateEnabled
)

type ServerState = original.ServerState

const (
	ServerStateDisabled ServerState = original.ServerStateDisabled
	ServerStateDropping ServerState = original.ServerStateDropping
	ServerStateReady    ServerState = original.ServerStateReady
)

type ServerVersion = original.ServerVersion

const (
	FiveFullStopSeven ServerVersion = original.FiveFullStopSeven
	FiveFullStopSix   ServerVersion = original.FiveFullStopSix
)

type SkuTier = original.SkuTier

const (
	Basic           SkuTier = original.Basic
	GeneralPurpose  SkuTier = original.GeneralPurpose
	MemoryOptimized SkuTier = original.MemoryOptimized
)

type SslEnforcementEnum = original.SslEnforcementEnum

const (
	SslEnforcementEnumDisabled SslEnforcementEnum = original.SslEnforcementEnumDisabled
	SslEnforcementEnumEnabled  SslEnforcementEnum = original.SslEnforcementEnumEnabled
)

type StorageAutogrow = original.StorageAutogrow

const (
	StorageAutogrowDisabled StorageAutogrow = original.StorageAutogrowDisabled
	StorageAutogrowEnabled  StorageAutogrow = original.StorageAutogrowEnabled
)

type VirtualNetworkRuleState = original.VirtualNetworkRuleState

const (
	VirtualNetworkRuleStateDeleting     VirtualNetworkRuleState = original.VirtualNetworkRuleStateDeleting
	VirtualNetworkRuleStateInitializing VirtualNetworkRuleState = original.VirtualNetworkRuleStateInitializing
	VirtualNetworkRuleStateInProgress   VirtualNetworkRuleState = original.VirtualNetworkRuleStateInProgress
	VirtualNetworkRuleStateReady        VirtualNetworkRuleState = original.VirtualNetworkRuleStateReady
	VirtualNetworkRuleStateUnknown      VirtualNetworkRuleState = original.VirtualNetworkRuleStateUnknown
)

type Advisor = original.Advisor
type AdvisorsClient = original.AdvisorsClient
type AdvisorsResultList = original.AdvisorsResultList
type AdvisorsResultListIterator = original.AdvisorsResultListIterator
type AdvisorsResultListPage = original.AdvisorsResultListPage
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BasicServerPropertiesForCreate = original.BasicServerPropertiesForCreate
type CheckNameAvailabilityClient = original.CheckNameAvailabilityClient
type CloudError = original.CloudError
type Configuration = original.Configuration
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsCreateOrUpdateFuture = original.ConfigurationsCreateOrUpdateFuture
type CreateRecommendedActionSessionFuture = original.CreateRecommendedActionSessionFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabasesClient = original.DatabasesClient
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type FirewallRulesCreateOrUpdateFuture = original.FirewallRulesCreateOrUpdateFuture
type FirewallRulesDeleteFuture = original.FirewallRulesDeleteFuture
type LocationBasedPerformanceTierClient = original.LocationBasedPerformanceTierClient
type LocationBasedRecommendedActionSessionsOperationStatusClient = original.LocationBasedRecommendedActionSessionsOperationStatusClient
type LocationBasedRecommendedActionSessionsResultClient = original.LocationBasedRecommendedActionSessionsResultClient
type LogFile = original.LogFile
type LogFileListResult = original.LogFileListResult
type LogFileProperties = original.LogFileProperties
type LogFilesClient = original.LogFilesClient
type NameAvailability = original.NameAvailability
type NameAvailabilityRequest = original.NameAvailabilityRequest
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PerformanceTierListResult = original.PerformanceTierListResult
type PerformanceTierProperties = original.PerformanceTierProperties
type PerformanceTierServiceLevelObjectives = original.PerformanceTierServiceLevelObjectives
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointConnectionsUpdateTagsFuture = original.PrivateEndpointConnectionsUpdateTagsFuture
type PrivateEndpointProperty = original.PrivateEndpointProperty
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceListResultIterator = original.PrivateLinkResourceListResultIterator
type PrivateLinkResourceListResultPage = original.PrivateLinkResourceListResultPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionStateProperty = original.PrivateLinkServiceConnectionStateProperty
type ProxyResource = original.ProxyResource
type QueryStatistic = original.QueryStatistic
type QueryStatisticProperties = original.QueryStatisticProperties
type QueryText = original.QueryText
type QueryTextProperties = original.QueryTextProperties
type QueryTextsClient = original.QueryTextsClient
type QueryTextsResultList = original.QueryTextsResultList
type QueryTextsResultListIterator = original.QueryTextsResultListIterator
type QueryTextsResultListPage = original.QueryTextsResultListPage
type RecommendationAction = original.RecommendationAction
type RecommendationActionProperties = original.RecommendationActionProperties
type RecommendationActionsResultList = original.RecommendationActionsResultList
type RecommendationActionsResultListIterator = original.RecommendationActionsResultListIterator
type RecommendationActionsResultListPage = original.RecommendationActionsResultListPage
type RecommendedActionSessionsOperationStatus = original.RecommendedActionSessionsOperationStatus
type RecommendedActionsClient = original.RecommendedActionsClient
type ReplicasClient = original.ReplicasClient
type Resource = original.Resource
type SecurityAlertPolicyProperties = original.SecurityAlertPolicyProperties
type Server = original.Server
type ServerForCreate = original.ServerForCreate
type ServerListResult = original.ServerListResult
type ServerPrivateEndpointConnection = original.ServerPrivateEndpointConnection
type ServerPrivateEndpointConnectionProperties = original.ServerPrivateEndpointConnectionProperties
type ServerPrivateLinkServiceConnectionStateProperty = original.ServerPrivateLinkServiceConnectionStateProperty
type ServerProperties = original.ServerProperties
type ServerPropertiesForCreate = original.ServerPropertiesForCreate
type ServerPropertiesForDefaultCreate = original.ServerPropertiesForDefaultCreate
type ServerPropertiesForGeoRestore = original.ServerPropertiesForGeoRestore
type ServerPropertiesForReplica = original.ServerPropertiesForReplica
type ServerPropertiesForRestore = original.ServerPropertiesForRestore
type ServerSecurityAlertPoliciesClient = original.ServerSecurityAlertPoliciesClient
type ServerSecurityAlertPoliciesCreateOrUpdateFuture = original.ServerSecurityAlertPoliciesCreateOrUpdateFuture
type ServerSecurityAlertPolicy = original.ServerSecurityAlertPolicy
type ServerUpdateParameters = original.ServerUpdateParameters
type ServerUpdateParametersProperties = original.ServerUpdateParametersProperties
type ServersClient = original.ServersClient
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersRestartFuture = original.ServersRestartFuture
type ServersStartFuture = original.ServersStartFuture
type ServersStopFuture = original.ServersStopFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type Sku = original.Sku
type StorageProfile = original.StorageProfile
type TagsObject = original.TagsObject
type TopQueryStatisticsClient = original.TopQueryStatisticsClient
type TopQueryStatisticsInput = original.TopQueryStatisticsInput
type TopQueryStatisticsInputProperties = original.TopQueryStatisticsInputProperties
type TopQueryStatisticsResultList = original.TopQueryStatisticsResultList
type TopQueryStatisticsResultListIterator = original.TopQueryStatisticsResultListIterator
type TopQueryStatisticsResultListPage = original.TopQueryStatisticsResultListPage
type TrackedResource = original.TrackedResource
type VirtualNetworkRule = original.VirtualNetworkRule
type VirtualNetworkRuleListResult = original.VirtualNetworkRuleListResult
type VirtualNetworkRuleListResultIterator = original.VirtualNetworkRuleListResultIterator
type VirtualNetworkRuleListResultPage = original.VirtualNetworkRuleListResultPage
type VirtualNetworkRuleProperties = original.VirtualNetworkRuleProperties
type VirtualNetworkRulesClient = original.VirtualNetworkRulesClient
type VirtualNetworkRulesCreateOrUpdateFuture = original.VirtualNetworkRulesCreateOrUpdateFuture
type VirtualNetworkRulesDeleteFuture = original.VirtualNetworkRulesDeleteFuture
type WaitStatistic = original.WaitStatistic
type WaitStatisticProperties = original.WaitStatisticProperties
type WaitStatisticsClient = original.WaitStatisticsClient
type WaitStatisticsInput = original.WaitStatisticsInput
type WaitStatisticsInputProperties = original.WaitStatisticsInputProperties
type WaitStatisticsResultList = original.WaitStatisticsResultList
type WaitStatisticsResultListIterator = original.WaitStatisticsResultListIterator
type WaitStatisticsResultListPage = original.WaitStatisticsResultListPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAdvisorsClient(subscriptionID string) AdvisorsClient {
	return original.NewAdvisorsClient(subscriptionID)
}
func NewAdvisorsClientWithBaseURI(baseURI string, subscriptionID string) AdvisorsClient {
	return original.NewAdvisorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAdvisorsResultListIterator(page AdvisorsResultListPage) AdvisorsResultListIterator {
	return original.NewAdvisorsResultListIterator(page)
}
func NewAdvisorsResultListPage(cur AdvisorsResultList, getNextPage func(context.Context, AdvisorsResultList) (AdvisorsResultList, error)) AdvisorsResultListPage {
	return original.NewAdvisorsResultListPage(cur, getNextPage)
}
func NewCheckNameAvailabilityClient(subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClient(subscriptionID)
}
func NewCheckNameAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedPerformanceTierClient(subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClient(subscriptionID)
}
func NewLocationBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedRecommendedActionSessionsOperationStatusClient(subscriptionID string) LocationBasedRecommendedActionSessionsOperationStatusClient {
	return original.NewLocationBasedRecommendedActionSessionsOperationStatusClient(subscriptionID)
}
func NewLocationBasedRecommendedActionSessionsOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedRecommendedActionSessionsOperationStatusClient {
	return original.NewLocationBasedRecommendedActionSessionsOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedRecommendedActionSessionsResultClient(subscriptionID string) LocationBasedRecommendedActionSessionsResultClient {
	return original.NewLocationBasedRecommendedActionSessionsResultClient(subscriptionID)
}
func NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedRecommendedActionSessionsResultClient {
	return original.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(baseURI, subscriptionID)
}
func NewLogFilesClient(subscriptionID string) LogFilesClient {
	return original.NewLogFilesClient(subscriptionID)
}
func NewLogFilesClientWithBaseURI(baseURI string, subscriptionID string) LogFilesClient {
	return original.NewLogFilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(cur PrivateEndpointConnectionListResult, getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListResultIterator(page PrivateLinkResourceListResultPage) PrivateLinkResourceListResultIterator {
	return original.NewPrivateLinkResourceListResultIterator(page)
}
func NewPrivateLinkResourceListResultPage(cur PrivateLinkResourceListResult, getNextPage func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)) PrivateLinkResourceListResultPage {
	return original.NewPrivateLinkResourceListResultPage(cur, getNextPage)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueryTextsClient(subscriptionID string) QueryTextsClient {
	return original.NewQueryTextsClient(subscriptionID)
}
func NewQueryTextsClientWithBaseURI(baseURI string, subscriptionID string) QueryTextsClient {
	return original.NewQueryTextsClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueryTextsResultListIterator(page QueryTextsResultListPage) QueryTextsResultListIterator {
	return original.NewQueryTextsResultListIterator(page)
}
func NewQueryTextsResultListPage(cur QueryTextsResultList, getNextPage func(context.Context, QueryTextsResultList) (QueryTextsResultList, error)) QueryTextsResultListPage {
	return original.NewQueryTextsResultListPage(cur, getNextPage)
}
func NewRecommendationActionsResultListIterator(page RecommendationActionsResultListPage) RecommendationActionsResultListIterator {
	return original.NewRecommendationActionsResultListIterator(page)
}
func NewRecommendationActionsResultListPage(cur RecommendationActionsResultList, getNextPage func(context.Context, RecommendationActionsResultList) (RecommendationActionsResultList, error)) RecommendationActionsResultListPage {
	return original.NewRecommendationActionsResultListPage(cur, getNextPage)
}
func NewRecommendedActionsClient(subscriptionID string) RecommendedActionsClient {
	return original.NewRecommendedActionsClient(subscriptionID)
}
func NewRecommendedActionsClientWithBaseURI(baseURI string, subscriptionID string) RecommendedActionsClient {
	return original.NewRecommendedActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicasClient(subscriptionID string) ReplicasClient {
	return original.NewReplicasClient(subscriptionID)
}
func NewReplicasClientWithBaseURI(baseURI string, subscriptionID string) ReplicasClient {
	return original.NewReplicasClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerSecurityAlertPoliciesClient(subscriptionID string) ServerSecurityAlertPoliciesClient {
	return original.NewServerSecurityAlertPoliciesClient(subscriptionID)
}
func NewServerSecurityAlertPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ServerSecurityAlertPoliciesClient {
	return original.NewServerSecurityAlertPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopQueryStatisticsClient(subscriptionID string) TopQueryStatisticsClient {
	return original.NewTopQueryStatisticsClient(subscriptionID)
}
func NewTopQueryStatisticsClientWithBaseURI(baseURI string, subscriptionID string) TopQueryStatisticsClient {
	return original.NewTopQueryStatisticsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopQueryStatisticsResultListIterator(page TopQueryStatisticsResultListPage) TopQueryStatisticsResultListIterator {
	return original.NewTopQueryStatisticsResultListIterator(page)
}
func NewTopQueryStatisticsResultListPage(cur TopQueryStatisticsResultList, getNextPage func(context.Context, TopQueryStatisticsResultList) (TopQueryStatisticsResultList, error)) TopQueryStatisticsResultListPage {
	return original.NewTopQueryStatisticsResultListPage(cur, getNextPage)
}
func NewVirtualNetworkRuleListResultIterator(page VirtualNetworkRuleListResultPage) VirtualNetworkRuleListResultIterator {
	return original.NewVirtualNetworkRuleListResultIterator(page)
}
func NewVirtualNetworkRuleListResultPage(cur VirtualNetworkRuleListResult, getNextPage func(context.Context, VirtualNetworkRuleListResult) (VirtualNetworkRuleListResult, error)) VirtualNetworkRuleListResultPage {
	return original.NewVirtualNetworkRuleListResultPage(cur, getNextPage)
}
func NewVirtualNetworkRulesClient(subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClient(subscriptionID)
}
func NewVirtualNetworkRulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWaitStatisticsClient(subscriptionID string) WaitStatisticsClient {
	return original.NewWaitStatisticsClient(subscriptionID)
}
func NewWaitStatisticsClientWithBaseURI(baseURI string, subscriptionID string) WaitStatisticsClient {
	return original.NewWaitStatisticsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWaitStatisticsResultListIterator(page WaitStatisticsResultListPage) WaitStatisticsResultListIterator {
	return original.NewWaitStatisticsResultListIterator(page)
}
func NewWaitStatisticsResultListPage(cur WaitStatisticsResultList, getNextPage func(context.Context, WaitStatisticsResultList) (WaitStatisticsResultList, error)) WaitStatisticsResultListPage {
	return original.NewWaitStatisticsResultListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return original.PossibleGeoRedundantBackupValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossiblePrivateEndpointProvisioningStateValues() []PrivateEndpointProvisioningState {
	return original.PossiblePrivateEndpointProvisioningStateValues()
}
func PossiblePrivateLinkServiceConnectionStateActionsRequireValues() []PrivateLinkServiceConnectionStateActionsRequire {
	return original.PossiblePrivateLinkServiceConnectionStateActionsRequireValues()
}
func PossiblePrivateLinkServiceConnectionStateStatusValues() []PrivateLinkServiceConnectionStateStatus {
	return original.PossiblePrivateLinkServiceConnectionStateStatusValues()
}
func PossiblePublicNetworkAccessEnumValues() []PublicNetworkAccessEnum {
	return original.PossiblePublicNetworkAccessEnumValues()
}
func PossibleServerSecurityAlertPolicyStateValues() []ServerSecurityAlertPolicyState {
	return original.PossibleServerSecurityAlertPolicyStateValues()
}
func PossibleServerStateValues() []ServerState {
	return original.PossibleServerStateValues()
}
func PossibleServerVersionValues() []ServerVersion {
	return original.PossibleServerVersionValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSslEnforcementEnumValues() []SslEnforcementEnum {
	return original.PossibleSslEnforcementEnumValues()
}
func PossibleStorageAutogrowValues() []StorageAutogrow {
	return original.PossibleStorageAutogrowValues()
}
func PossibleVirtualNetworkRuleStateValues() []VirtualNetworkRuleState {
	return original.PossibleVirtualNetworkRuleStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
