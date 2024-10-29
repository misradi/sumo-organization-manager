/*
Sumo Logic Organizations Management API

Welcome to the Sumo Logic's API Reference for Organizations Management. You can use these APIs to interact with the Sumo Logic platform to manage accounts and subscription. Refer to [API Authentication](https://help.sumologic.com/APIs/General-API-Information/API-Authentication) for more information about authentication. You can also look at [other APIs](https://help.sumologic.com/APIs) for more information about some other API endpoints. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AccessKeyPublic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessKeyPublic{}

// AccessKeyPublic struct for AccessKeyPublic
type AccessKeyPublic struct {
	// Identifier of the access key.
	Id string `json:"id"`
	// The name of the access key.
	Label string `json:"label"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
	// Indicates whether the access key is disabled or not.
	Disabled bool `json:"disabled"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the access key.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Last used timestamp in UTC.  <br> **Note:** Property not in use, it is part of an upcoming feature.
	LastUsed *time.Time `json:"lastUsed,omitempty"`
	// Scopes assigned to the key. ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - viewConnections   - manageConnections  ### Metrics   - metricsTransformation   - runMetricsQuery   ### Logs   - runLogSearch  ### Security   - managePasswordPolicy   - ipAllowlisting   - ipWhitelisting   - manageAccessKeys   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Organizations   - viewOrganizations  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse   - cseViewAutomations   - cseManageContextActions   - cseViewNetworkBlocks   - cseManageInsightTags   - cseViewRules   - cseViewThreatIntelligence   - cseCommentOnInsights   - cseViewEntityGroups   - cseManageEntityConfiguration   - cseManageNetworkBlocks   - cseManageMatchLists   - cseViewCustomInsights   - cseManageActions   - cseManageAutomations   - cseManageMappings   - cseManageThreatIntelligence   - cseViewActions   - cseCreateInsights   - cseManageTagSchemas   - cseInvokeInsights   - cseManageCustomEntityType   - cseViewTagSchemas   - cseDeleteInsights   - cseManageCustomInsights   - cseViewFileAnalysis   - cseManageFileAnalysis   - cseManageEntityCriticality   - cseViewEntityCriticality   - cseViewEntity   - cseManageCustomInsightStatuses   - cseViewContextActions   - cseViewMappings   - cseViewCustomEntityType   - cseManageEntityGroups   - cseViewCustomInsightStatuses   - cseViewEnrichments   - cseManageInsightSignals   - cseManageRules   - cseManageArtifacts   - cseViewMatchLists   - cseManageInsightPolicy   - cseManageEnrichments   - cseViewEntityConfiguration   - cseManageEntity   - cseExecuteAutomations   - cseManageSuppressedEntities   - cseManageInsightStatus     - cseManageInsightAssignee   - cseManageFavoriteFields   - cseViewSuppressedEntities  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts   - viewMutingSchedules   - manageMutingSchedules   - adminMonitorsV2  ### SLO   - viewSlos   - manageSlos  ### CloudSoar   - cloudSoarPlaybooksAccess   - cloudSoarNotificationConfigure   - cloudSoarReportAll   - cloudSoarIncidentTriageAccess   - cloudSoarIncidentTaskView   - cloudSoarIncidentChangeOwnership   - cloudSoarIncidentNotesEdit   - cloudSoarAPIEmailEdit   - cloudSoarIncidentTemplatesAccess   - cloudSoarIncidentPlaybooksManage   - cloudSoarGeneralConfigure   - cloudSoarEntitiesAccess   - cloudSoarEntitiesBulkPhysicalDelete   - cloudSoarIncidentAttachmentsAccess   - cloudSoarAppCentralAccess   - cloudSoarBridgeMonitoringAccess   - viewCloudSoar   - cloudSoarIncidentView   - cloudSoarObservabilityAccess   - cloudSoarAPIEmailRead   - cloudSoarAppCentralExport   - cloudSoarWidgetsAll   - cloudSoarIncidentTaskReassign   - cloudSoarIntegrationsAccess   - cloudSoarCustomizationIncidentLabels   - cloudSoarAutomationRulesConfigure   - cloudSoarIncidentTaskAccessAll   - cloudSoarAuditAndInformationConfigureAuditTrail   - cloudSoarIncidentTriageEdit   - cloudSoarIncidentEdit   - cloudSoarNotificationTriage   - cloudSoarIncidentTriageBulkPhysicalDelete   - cloudSoarIncidentNotesAccess   - cloudSoarAPIUse   - cloudSoarIncidentPlaybooksEdit   - cloudSoarDashboardAll   - cloudSoarEntitiesManage   - cloudSoarIncidentTemplatesConfigure   - cloudSoarIncidentTriageAccessAll   - cloudSoarPlaybooksConfigure   - cloudSoarIncidentAccessAll   - cloudSoarCustomizationLogo   - cloudSoarIncidentTaskAccess   - cloudSoarIncidentTriageView   - cloudSoarIntegrationsConfigure   - cloudSoarIncidentManageInvestigators   - cloudSoarIncidentAccess   - cloudSoarAuditAndInformationLicenseInformation   - cloudSoarIncidentBulkOperations   - cloudSoarCustomizationFields   - cloudSoarIncidentTaskEdit   - cloudSoarDashboardAccess   - cloudSoarIncidentAttachmentsEdit   - cloudSoarIncidentFoldersEdit   - cloudSoarUserManagementGroups   - cloudSoarIncidentPlaybooksAccess   - cloudSoarIncidentWarRoomUse   - cloudSoarReportAccess   - cloudSoarAuditAndInformationAuditTrail   - cloudSoarAutomationRulesAccess   - cloudSoarIncidentTriageChangeOwnership   - cloudSoarObservabilityManagement
	Scopes []string `json:"scopes,omitempty"`
	// Effective scopes based on the intersection of the user's RBAC capabilities and the assigned scopes.
	EffectiveScopes []string `json:"effectiveScopes,omitempty"`
}

type _AccessKeyPublic AccessKeyPublic

// NewAccessKeyPublic instantiates a new AccessKeyPublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyPublic(id string, label string, disabled bool, createdAt time.Time, createdBy string, modifiedAt time.Time) *AccessKeyPublic {
	this := AccessKeyPublic{}
	this.Id = id
	this.Label = label
	this.Disabled = disabled
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAccessKeyPublicWithDefaults instantiates a new AccessKeyPublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyPublicWithDefaults() *AccessKeyPublic {
	this := AccessKeyPublic{}
	return &this
}

// GetId returns the Id field value
func (o *AccessKeyPublic) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessKeyPublic) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *AccessKeyPublic) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AccessKeyPublic) SetLabel(v string) {
	o.Label = v
}

// GetCorsHeaders returns the CorsHeaders field value if set, zero value otherwise.
func (o *AccessKeyPublic) GetCorsHeaders() []string {
	if o == nil || IsNil(o.CorsHeaders) {
		var ret []string
		return ret
	}
	return o.CorsHeaders
}

// GetCorsHeadersOk returns a tuple with the CorsHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetCorsHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsHeaders) {
		return nil, false
	}
	return o.CorsHeaders, true
}

// HasCorsHeaders returns a boolean if a field has been set.
func (o *AccessKeyPublic) HasCorsHeaders() bool {
	if o != nil && !IsNil(o.CorsHeaders) {
		return true
	}

	return false
}

// SetCorsHeaders gets a reference to the given []string and assigns it to the CorsHeaders field.
func (o *AccessKeyPublic) SetCorsHeaders(v []string) {
	o.CorsHeaders = v
}

// GetDisabled returns the Disabled field value
func (o *AccessKeyPublic) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *AccessKeyPublic) SetDisabled(v bool) {
	o.Disabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccessKeyPublic) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccessKeyPublic) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *AccessKeyPublic) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *AccessKeyPublic) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AccessKeyPublic) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AccessKeyPublic) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *AccessKeyPublic) GetLastUsed() time.Time {
	if o == nil || IsNil(o.LastUsed) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetLastUsedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsed) {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *AccessKeyPublic) HasLastUsed() bool {
	if o != nil && !IsNil(o.LastUsed) {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given time.Time and assigns it to the LastUsed field.
func (o *AccessKeyPublic) SetLastUsed(v time.Time) {
	o.LastUsed = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AccessKeyPublic) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AccessKeyPublic) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AccessKeyPublic) SetScopes(v []string) {
	o.Scopes = v
}

// GetEffectiveScopes returns the EffectiveScopes field value if set, zero value otherwise.
func (o *AccessKeyPublic) GetEffectiveScopes() []string {
	if o == nil || IsNil(o.EffectiveScopes) {
		var ret []string
		return ret
	}
	return o.EffectiveScopes
}

// GetEffectiveScopesOk returns a tuple with the EffectiveScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyPublic) GetEffectiveScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.EffectiveScopes) {
		return nil, false
	}
	return o.EffectiveScopes, true
}

// HasEffectiveScopes returns a boolean if a field has been set.
func (o *AccessKeyPublic) HasEffectiveScopes() bool {
	if o != nil && !IsNil(o.EffectiveScopes) {
		return true
	}

	return false
}

// SetEffectiveScopes gets a reference to the given []string and assigns it to the EffectiveScopes field.
func (o *AccessKeyPublic) SetEffectiveScopes(v []string) {
	o.EffectiveScopes = v
}

func (o AccessKeyPublic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessKeyPublic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	if !IsNil(o.CorsHeaders) {
		toSerialize["corsHeaders"] = o.CorsHeaders
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	if !IsNil(o.LastUsed) {
		toSerialize["lastUsed"] = o.LastUsed
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.EffectiveScopes) {
		toSerialize["effectiveScopes"] = o.EffectiveScopes
	}
	return toSerialize, nil
}

func (o *AccessKeyPublic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"label",
		"disabled",
		"createdAt",
		"createdBy",
		"modifiedAt",
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

	varAccessKeyPublic := _AccessKeyPublic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessKeyPublic)

	if err != nil {
		return err
	}

	*o = AccessKeyPublic(varAccessKeyPublic)

	return err
}

type NullableAccessKeyPublic struct {
	value *AccessKeyPublic
	isSet bool
}

func (v NullableAccessKeyPublic) Get() *AccessKeyPublic {
	return v.value
}

func (v *NullableAccessKeyPublic) Set(val *AccessKeyPublic) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeyPublic) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeyPublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeyPublic(val *AccessKeyPublic) *NullableAccessKeyPublic {
	return &NullableAccessKeyPublic{value: val, isSet: true}
}

func (v NullableAccessKeyPublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeyPublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

