/*
Sumo Logic Organizations Management API

Welcome to the Sumo Logic's API Reference for Organizations Management. You can use these APIs to interact with the Sumo Logic platform to manage accounts and subscription. Refer to [API Authentication](https://help.sumologic.com/APIs/General-API-Information/API-Authentication) for more information about authentication. You can also look at [other APIs](https://help.sumologic.com/APIs) for more information about some other API endpoints.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ReadOrganizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadOrganizationResponse{}

// ReadOrganizationResponse struct for ReadOrganizationResponse
type ReadOrganizationResponse struct {
	// Email address of the account owner.
	Email string `json:"email"`
	// Name of the organization.
	OrganizationName string `json:"organizationName" validate:"regexp=^([a-zA-Z0-9 +%\\\\-@.,_()]+)$"`
	// First name of the account owner.
	FirstName string `json:"firstName"`
	// Last name of the account owner.
	LastName *string `json:"lastName,omitempty"`
	Subscription CreditsSubscription `json:"subscription"`
	// The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash `-` character.
	OrgId string `json:"orgId"`
	// Identifier of the deployment in which the organization is present.
	DeploymentId *string `json:"deploymentId,omitempty" validate:"regexp=^(mb|stag|long|prod|us2|dub|syd|mum|fra|tky|mon|fed|au|ca|de|eu|in|jp|us1)$"`
	// Only applicable for flex orgs
	IsSSOSetup bool `json:"isSSOSetup"`
}

type _ReadOrganizationResponse ReadOrganizationResponse

// NewReadOrganizationResponse instantiates a new ReadOrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOrganizationResponse(email string, organizationName string, firstName string, subscription CreditsSubscription, orgId string) *ReadOrganizationResponse {
	this := ReadOrganizationResponse{}
	this.Email = email
	this.OrganizationName = organizationName
	this.FirstName = firstName
	this.Subscription = subscription
	this.OrgId = orgId
	return &this
}

// NewReadOrganizationResponseWithDefaults instantiates a new ReadOrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOrganizationResponseWithDefaults() *ReadOrganizationResponse {
	this := ReadOrganizationResponse{}
	return &this
}

// GetEmail returns the Email field value
func (o *ReadOrganizationResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ReadOrganizationResponse) SetEmail(v string) {
	o.Email = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *ReadOrganizationResponse) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *ReadOrganizationResponse) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetFirstName returns the FirstName field value
func (o *ReadOrganizationResponse) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ReadOrganizationResponse) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ReadOrganizationResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ReadOrganizationResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ReadOrganizationResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetSubscription returns the Subscription field value
func (o *ReadOrganizationResponse) GetSubscription() CreditsSubscription {
	if o == nil {
		var ret CreditsSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetSubscriptionOk() (*CreditsSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *ReadOrganizationResponse) SetSubscription(v CreditsSubscription) {
	o.Subscription = v
}

// GetOrgId returns the OrgId field value
func (o *ReadOrganizationResponse) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ReadOrganizationResponse) SetOrgId(v string) {
	o.OrgId = v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *ReadOrganizationResponse) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOrganizationResponse) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *ReadOrganizationResponse) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *ReadOrganizationResponse) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

func (o ReadOrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadOrganizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["organizationName"] = o.OrganizationName
	toSerialize["firstName"] = o.FirstName
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["subscription"] = o.Subscription
	toSerialize["orgId"] = o.OrgId
	if !IsNil(o.DeploymentId) {
		toSerialize["deploymentId"] = o.DeploymentId
	}
	return toSerialize, nil
}

func (o *ReadOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"organizationName",
		"firstName",
		"subscription",
		"orgId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReadOrganizationResponse := _ReadOrganizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReadOrganizationResponse)

	if err != nil {
		return err
	}

	*o = ReadOrganizationResponse(varReadOrganizationResponse)

	return err
}

type NullableReadOrganizationResponse struct {
	value *ReadOrganizationResponse
	isSet bool
}

func (v NullableReadOrganizationResponse) Get() *ReadOrganizationResponse {
	return v.value
}

func (v *NullableReadOrganizationResponse) Set(val *ReadOrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOrganizationResponse(val *ReadOrganizationResponse) *NullableReadOrganizationResponse {
	return &NullableReadOrganizationResponse{value: val, isSet: true}
}

func (v NullableReadOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
