/*
Sumo Logic Organizations Management API

Welcome to the Sumo Logic's API Reference for Organizations Management. You can use these APIs to interact with the Sumo Logic platform to manage accounts and subscription. Refer to [API Authentication](https://help.sumologic.com/APIs/General-API-Information/API-Authentication) for more information about authentication. You can also look at [other APIs](https://help.sumologic.com/APIs) for more information about some other API endpoints. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccessKeyCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessKeyCreateRequest{}

// AccessKeyCreateRequest struct for AccessKeyCreateRequest
type AccessKeyCreateRequest struct {
	// A name for the access key to be created.
	Label string `json:"label"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
	// Scopes assigned to the key. ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - viewConnections   - manageConnections  ### Metrics   - metricsTransformation   - runMetricsQuery   ### Logs   - runLogSearch  ### Security   - managePasswordPolicy   - ipAllowlisting   - ipWhitelisting   - manageAccessKeys   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Organizations   - viewOrganizations  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse   - cseViewAutomations   - cseManageContextActions   - cseViewNetworkBlocks   - cseManageInsightTags   - cseViewRules   - cseViewThreatIntelligence   - cseCommentOnInsights   - cseViewEntityGroups   - cseManageEntityConfiguration   - cseManageNetworkBlocks   - cseManageMatchLists   - cseViewCustomInsights   - cseManageActions   - cseManageAutomations   - cseManageMappings   - cseManageThreatIntelligence   - cseViewActions   - cseCreateInsights   - cseManageTagSchemas   - cseInvokeInsights   - cseManageCustomEntityType   - cseViewTagSchemas   - cseDeleteInsights   - cseManageCustomInsights   - cseViewFileAnalysis   - cseManageFileAnalysis   - cseManageEntityCriticality   - cseViewEntityCriticality   - cseViewEntity   - cseManageCustomInsightStatuses   - cseViewContextActions   - cseViewMappings   - cseViewCustomEntityType   - cseManageEntityGroups   - cseViewCustomInsightStatuses   - cseViewEnrichments   - cseManageInsightSignals   - cseManageRules   - cseManageArtifacts   - cseViewMatchLists   - cseManageInsightPolicy   - cseManageEnrichments   - cseViewEntityConfiguration   - cseManageEntity   - cseExecuteAutomations   - cseManageSuppressedEntities   - cseManageInsightStatus     - cseManageInsightAssignee   - cseManageFavoriteFields   - cseViewSuppressedEntities  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts   - viewMutingSchedules   - manageMutingSchedules   - adminMonitorsV2  ### SLO   - viewSlos   - manageSlos  ### CloudSoar   - cloudSoarPlaybooksAccess   - cloudSoarNotificationConfigure   - cloudSoarReportAll   - cloudSoarIncidentTriageAccess   - cloudSoarIncidentTaskView   - cloudSoarIncidentChangeOwnership   - cloudSoarIncidentNotesEdit   - cloudSoarAPIEmailEdit   - cloudSoarIncidentTemplatesAccess   - cloudSoarIncidentPlaybooksManage   - cloudSoarGeneralConfigure   - cloudSoarEntitiesAccess   - cloudSoarEntitiesBulkPhysicalDelete   - cloudSoarIncidentAttachmentsAccess   - cloudSoarAppCentralAccess   - cloudSoarBridgeMonitoringAccess   - viewCloudSoar   - cloudSoarIncidentView   - cloudSoarObservabilityAccess   - cloudSoarAPIEmailRead   - cloudSoarAppCentralExport   - cloudSoarWidgetsAll   - cloudSoarIncidentTaskReassign   - cloudSoarIntegrationsAccess   - cloudSoarCustomizationIncidentLabels   - cloudSoarAutomationRulesConfigure   - cloudSoarIncidentTaskAccessAll   - cloudSoarAuditAndInformationConfigureAuditTrail   - cloudSoarIncidentTriageEdit   - cloudSoarIncidentEdit   - cloudSoarNotificationTriage   - cloudSoarIncidentTriageBulkPhysicalDelete   - cloudSoarIncidentNotesAccess   - cloudSoarAPIUse   - cloudSoarIncidentPlaybooksEdit   - cloudSoarDashboardAll   - cloudSoarEntitiesManage   - cloudSoarIncidentTemplatesConfigure   - cloudSoarIncidentTriageAccessAll   - cloudSoarPlaybooksConfigure   - cloudSoarIncidentAccessAll   - cloudSoarCustomizationLogo   - cloudSoarIncidentTaskAccess   - cloudSoarIncidentTriageView   - cloudSoarIntegrationsConfigure   - cloudSoarIncidentManageInvestigators   - cloudSoarIncidentAccess   - cloudSoarAuditAndInformationLicenseInformation   - cloudSoarIncidentBulkOperations   - cloudSoarCustomizationFields   - cloudSoarIncidentTaskEdit   - cloudSoarDashboardAccess   - cloudSoarIncidentAttachmentsEdit   - cloudSoarIncidentFoldersEdit   - cloudSoarUserManagementGroups   - cloudSoarIncidentPlaybooksAccess   - cloudSoarIncidentWarRoomUse   - cloudSoarReportAccess   - cloudSoarAuditAndInformationAuditTrail   - cloudSoarAutomationRulesAccess   - cloudSoarIncidentTriageChangeOwnership   - cloudSoarObservabilityManagement
	Scopes []string `json:"scopes,omitempty"`
}

type _AccessKeyCreateRequest AccessKeyCreateRequest

// NewAccessKeyCreateRequest instantiates a new AccessKeyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyCreateRequest(label string) *AccessKeyCreateRequest {
	this := AccessKeyCreateRequest{}
	this.Label = label
	return &this
}

// NewAccessKeyCreateRequestWithDefaults instantiates a new AccessKeyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyCreateRequestWithDefaults() *AccessKeyCreateRequest {
	this := AccessKeyCreateRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *AccessKeyCreateRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AccessKeyCreateRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AccessKeyCreateRequest) SetLabel(v string) {
	o.Label = v
}

// GetCorsHeaders returns the CorsHeaders field value if set, zero value otherwise.
func (o *AccessKeyCreateRequest) GetCorsHeaders() []string {
	if o == nil || IsNil(o.CorsHeaders) {
		var ret []string
		return ret
	}
	return o.CorsHeaders
}

// GetCorsHeadersOk returns a tuple with the CorsHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyCreateRequest) GetCorsHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsHeaders) {
		return nil, false
	}
	return o.CorsHeaders, true
}

// HasCorsHeaders returns a boolean if a field has been set.
func (o *AccessKeyCreateRequest) HasCorsHeaders() bool {
	if o != nil && !IsNil(o.CorsHeaders) {
		return true
	}

	return false
}

// SetCorsHeaders gets a reference to the given []string and assigns it to the CorsHeaders field.
func (o *AccessKeyCreateRequest) SetCorsHeaders(v []string) {
	o.CorsHeaders = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AccessKeyCreateRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyCreateRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AccessKeyCreateRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AccessKeyCreateRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o AccessKeyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessKeyCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !IsNil(o.CorsHeaders) {
		toSerialize["corsHeaders"] = o.CorsHeaders
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

func (o *AccessKeyCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessKeyCreateRequest := _AccessKeyCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessKeyCreateRequest)

	if err != nil {
		return err
	}

	*o = AccessKeyCreateRequest(varAccessKeyCreateRequest)

	return err
}

type NullableAccessKeyCreateRequest struct {
	value *AccessKeyCreateRequest
	isSet bool
}

func (v NullableAccessKeyCreateRequest) Get() *AccessKeyCreateRequest {
	return v.value
}

func (v *NullableAccessKeyCreateRequest) Set(val *AccessKeyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeyCreateRequest(val *AccessKeyCreateRequest) *NullableAccessKeyCreateRequest {
	return &NullableAccessKeyCreateRequest{value: val, isSet: true}
}

func (v NullableAccessKeyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

