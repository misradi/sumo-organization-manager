# AccessKeyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the access key. | 
**Label** | **string** | The name of the access key. | 
**CorsHeaders** | Pointer to **[]string** | An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don&#39;t match any entry in the allowlist. | [optional] 
**Disabled** | **bool** | Indicates whether the access key is disabled or not. | 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the access key. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**LastUsed** | Pointer to **time.Time** | Last used timestamp in UTC.  &lt;br&gt; **Note:** Property not in use, it is part of an upcoming feature. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes assigned to the key. ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - viewConnections   - manageConnections  ### Metrics   - metricsTransformation   - runMetricsQuery   ### Logs   - runLogSearch  ### Security   - managePasswordPolicy   - ipAllowlisting   - ipWhitelisting   - manageAccessKeys   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Organizations   - viewOrganizations  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse   - cseViewAutomations   - cseManageContextActions   - cseViewNetworkBlocks   - cseManageInsightTags   - cseViewRules   - cseViewThreatIntelligence   - cseCommentOnInsights   - cseViewEntityGroups   - cseManageEntityConfiguration   - cseManageNetworkBlocks   - cseManageMatchLists   - cseViewCustomInsights   - cseManageActions   - cseManageAutomations   - cseManageMappings   - cseManageThreatIntelligence   - cseViewActions   - cseCreateInsights   - cseManageTagSchemas   - cseInvokeInsights   - cseManageCustomEntityType   - cseViewTagSchemas   - cseDeleteInsights   - cseManageCustomInsights   - cseViewFileAnalysis   - cseManageFileAnalysis   - cseManageEntityCriticality   - cseViewEntityCriticality   - cseViewEntity   - cseManageCustomInsightStatuses   - cseViewContextActions   - cseViewMappings   - cseViewCustomEntityType   - cseManageEntityGroups   - cseViewCustomInsightStatuses   - cseViewEnrichments   - cseManageInsightSignals   - cseManageRules   - cseManageArtifacts   - cseViewMatchLists   - cseManageInsightPolicy   - cseManageEnrichments   - cseViewEntityConfiguration   - cseManageEntity   - cseExecuteAutomations   - cseManageSuppressedEntities   - cseManageInsightStatus     - cseManageInsightAssignee   - cseManageFavoriteFields   - cseViewSuppressedEntities  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts   - viewMutingSchedules   - manageMutingSchedules   - adminMonitorsV2  ### SLO   - viewSlos   - manageSlos  ### CloudSoar   - cloudSoarPlaybooksAccess   - cloudSoarNotificationConfigure   - cloudSoarReportAll   - cloudSoarIncidentTriageAccess   - cloudSoarIncidentTaskView   - cloudSoarIncidentChangeOwnership   - cloudSoarIncidentNotesEdit   - cloudSoarAPIEmailEdit   - cloudSoarIncidentTemplatesAccess   - cloudSoarIncidentPlaybooksManage   - cloudSoarGeneralConfigure   - cloudSoarEntitiesAccess   - cloudSoarEntitiesBulkPhysicalDelete   - cloudSoarIncidentAttachmentsAccess   - cloudSoarAppCentralAccess   - cloudSoarBridgeMonitoringAccess   - viewCloudSoar   - cloudSoarIncidentView   - cloudSoarObservabilityAccess   - cloudSoarAPIEmailRead   - cloudSoarAppCentralExport   - cloudSoarWidgetsAll   - cloudSoarIncidentTaskReassign   - cloudSoarIntegrationsAccess   - cloudSoarCustomizationIncidentLabels   - cloudSoarAutomationRulesConfigure   - cloudSoarIncidentTaskAccessAll   - cloudSoarAuditAndInformationConfigureAuditTrail   - cloudSoarIncidentTriageEdit   - cloudSoarIncidentEdit   - cloudSoarNotificationTriage   - cloudSoarIncidentTriageBulkPhysicalDelete   - cloudSoarIncidentNotesAccess   - cloudSoarAPIUse   - cloudSoarIncidentPlaybooksEdit   - cloudSoarDashboardAll   - cloudSoarEntitiesManage   - cloudSoarIncidentTemplatesConfigure   - cloudSoarIncidentTriageAccessAll   - cloudSoarPlaybooksConfigure   - cloudSoarIncidentAccessAll   - cloudSoarCustomizationLogo   - cloudSoarIncidentTaskAccess   - cloudSoarIncidentTriageView   - cloudSoarIntegrationsConfigure   - cloudSoarIncidentManageInvestigators   - cloudSoarIncidentAccess   - cloudSoarAuditAndInformationLicenseInformation   - cloudSoarIncidentBulkOperations   - cloudSoarCustomizationFields   - cloudSoarIncidentTaskEdit   - cloudSoarDashboardAccess   - cloudSoarIncidentAttachmentsEdit   - cloudSoarIncidentFoldersEdit   - cloudSoarUserManagementGroups   - cloudSoarIncidentPlaybooksAccess   - cloudSoarIncidentWarRoomUse   - cloudSoarReportAccess   - cloudSoarAuditAndInformationAuditTrail   - cloudSoarAutomationRulesAccess   - cloudSoarIncidentTriageChangeOwnership   - cloudSoarObservabilityManagement | [optional] 
**EffectiveScopes** | Pointer to **[]string** | Effective scopes based on the intersection of the user&#39;s RBAC capabilities and the assigned scopes. | [optional] 

## Methods

### NewAccessKeyPublic

`func NewAccessKeyPublic(id string, label string, disabled bool, createdAt time.Time, createdBy string, modifiedAt time.Time, ) *AccessKeyPublic`

NewAccessKeyPublic instantiates a new AccessKeyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyPublicWithDefaults

`func NewAccessKeyPublicWithDefaults() *AccessKeyPublic`

NewAccessKeyPublicWithDefaults instantiates a new AccessKeyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessKeyPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKeyPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKeyPublic) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *AccessKeyPublic) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AccessKeyPublic) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AccessKeyPublic) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCorsHeaders

`func (o *AccessKeyPublic) GetCorsHeaders() []string`

GetCorsHeaders returns the CorsHeaders field if non-nil, zero value otherwise.

### GetCorsHeadersOk

`func (o *AccessKeyPublic) GetCorsHeadersOk() (*[]string, bool)`

GetCorsHeadersOk returns a tuple with the CorsHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHeaders

`func (o *AccessKeyPublic) SetCorsHeaders(v []string)`

SetCorsHeaders sets CorsHeaders field to given value.

### HasCorsHeaders

`func (o *AccessKeyPublic) HasCorsHeaders() bool`

HasCorsHeaders returns a boolean if a field has been set.

### GetDisabled

`func (o *AccessKeyPublic) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccessKeyPublic) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccessKeyPublic) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetCreatedAt

`func (o *AccessKeyPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessKeyPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessKeyPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AccessKeyPublic) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccessKeyPublic) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccessKeyPublic) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *AccessKeyPublic) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AccessKeyPublic) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AccessKeyPublic) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetLastUsed

`func (o *AccessKeyPublic) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *AccessKeyPublic) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *AccessKeyPublic) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *AccessKeyPublic) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetScopes

`func (o *AccessKeyPublic) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessKeyPublic) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessKeyPublic) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AccessKeyPublic) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetEffectiveScopes

`func (o *AccessKeyPublic) GetEffectiveScopes() []string`

GetEffectiveScopes returns the EffectiveScopes field if non-nil, zero value otherwise.

### GetEffectiveScopesOk

`func (o *AccessKeyPublic) GetEffectiveScopesOk() (*[]string, bool)`

GetEffectiveScopesOk returns a tuple with the EffectiveScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveScopes

`func (o *AccessKeyPublic) SetEffectiveScopes(v []string)`

SetEffectiveScopes sets EffectiveScopes field to given value.

### HasEffectiveScopes

`func (o *AccessKeyPublic) HasEffectiveScopes() bool`

HasEffectiveScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


