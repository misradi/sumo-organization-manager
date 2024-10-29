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

// checks if the OrganizationWithSubscriptionDetailsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationWithSubscriptionDetailsAllOf{}

// OrganizationWithSubscriptionDetailsAllOf struct for OrganizationWithSubscriptionDetailsAllOf
type OrganizationWithSubscriptionDetailsAllOf struct {
	// Identifier of the deployment in which the organization should be created.
	DeploymentId string `json:"deploymentId" validate:"regexp=^(mb|stag|long|prod|us2|dub|syd|mum|fra|tky|mon|fed|au|ca|de|eu|in|jp|us1)$"`
	// Specify the duration of the Trial plan. If not specified, your subscription plan will be used for the created organization.
	TrialPlanPeriod *int32 `json:"trialPlanPeriod,omitempty"`
	Baselines *Baselines `json:"baselines,omitempty"`
}

type _OrganizationWithSubscriptionDetailsAllOf OrganizationWithSubscriptionDetailsAllOf

// NewOrganizationWithSubscriptionDetailsAllOf instantiates a new OrganizationWithSubscriptionDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationWithSubscriptionDetailsAllOf(deploymentId string) *OrganizationWithSubscriptionDetailsAllOf {
	this := OrganizationWithSubscriptionDetailsAllOf{}
	this.DeploymentId = deploymentId
	return &this
}

// NewOrganizationWithSubscriptionDetailsAllOfWithDefaults instantiates a new OrganizationWithSubscriptionDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithSubscriptionDetailsAllOfWithDefaults() *OrganizationWithSubscriptionDetailsAllOf {
	this := OrganizationWithSubscriptionDetailsAllOf{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value
func (o *OrganizationWithSubscriptionDetailsAllOf) GetDeploymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value
// and a boolean to check if the value has been set.
func (o *OrganizationWithSubscriptionDetailsAllOf) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentId, true
}

// SetDeploymentId sets field value
func (o *OrganizationWithSubscriptionDetailsAllOf) SetDeploymentId(v string) {
	o.DeploymentId = v
}

// GetTrialPlanPeriod returns the TrialPlanPeriod field value if set, zero value otherwise.
func (o *OrganizationWithSubscriptionDetailsAllOf) GetTrialPlanPeriod() int32 {
	if o == nil || IsNil(o.TrialPlanPeriod) {
		var ret int32
		return ret
	}
	return *o.TrialPlanPeriod
}

// GetTrialPlanPeriodOk returns a tuple with the TrialPlanPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWithSubscriptionDetailsAllOf) GetTrialPlanPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPlanPeriod) {
		return nil, false
	}
	return o.TrialPlanPeriod, true
}

// HasTrialPlanPeriod returns a boolean if a field has been set.
func (o *OrganizationWithSubscriptionDetailsAllOf) HasTrialPlanPeriod() bool {
	if o != nil && !IsNil(o.TrialPlanPeriod) {
		return true
	}

	return false
}

// SetTrialPlanPeriod gets a reference to the given int32 and assigns it to the TrialPlanPeriod field.
func (o *OrganizationWithSubscriptionDetailsAllOf) SetTrialPlanPeriod(v int32) {
	o.TrialPlanPeriod = &v
}

// GetBaselines returns the Baselines field value if set, zero value otherwise.
func (o *OrganizationWithSubscriptionDetailsAllOf) GetBaselines() Baselines {
	if o == nil || IsNil(o.Baselines) {
		var ret Baselines
		return ret
	}
	return *o.Baselines
}

// GetBaselinesOk returns a tuple with the Baselines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWithSubscriptionDetailsAllOf) GetBaselinesOk() (*Baselines, bool) {
	if o == nil || IsNil(o.Baselines) {
		return nil, false
	}
	return o.Baselines, true
}

// HasBaselines returns a boolean if a field has been set.
func (o *OrganizationWithSubscriptionDetailsAllOf) HasBaselines() bool {
	if o != nil && !IsNil(o.Baselines) {
		return true
	}

	return false
}

// SetBaselines gets a reference to the given Baselines and assigns it to the Baselines field.
func (o *OrganizationWithSubscriptionDetailsAllOf) SetBaselines(v Baselines) {
	o.Baselines = &v
}

func (o OrganizationWithSubscriptionDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationWithSubscriptionDetailsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deploymentId"] = o.DeploymentId
	if !IsNil(o.TrialPlanPeriod) {
		toSerialize["trialPlanPeriod"] = o.TrialPlanPeriod
	}
	if !IsNil(o.Baselines) {
		toSerialize["baselines"] = o.Baselines
	}
	return toSerialize, nil
}

func (o *OrganizationWithSubscriptionDetailsAllOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deploymentId",
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

	varOrganizationWithSubscriptionDetailsAllOf := _OrganizationWithSubscriptionDetailsAllOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationWithSubscriptionDetailsAllOf)

	if err != nil {
		return err
	}

	*o = OrganizationWithSubscriptionDetailsAllOf(varOrganizationWithSubscriptionDetailsAllOf)

	return err
}

type NullableOrganizationWithSubscriptionDetailsAllOf struct {
	value *OrganizationWithSubscriptionDetailsAllOf
	isSet bool
}

func (v NullableOrganizationWithSubscriptionDetailsAllOf) Get() *OrganizationWithSubscriptionDetailsAllOf {
	return v.value
}

func (v *NullableOrganizationWithSubscriptionDetailsAllOf) Set(val *OrganizationWithSubscriptionDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWithSubscriptionDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWithSubscriptionDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWithSubscriptionDetailsAllOf(val *OrganizationWithSubscriptionDetailsAllOf) *NullableOrganizationWithSubscriptionDetailsAllOf {
	return &NullableOrganizationWithSubscriptionDetailsAllOf{value: val, isSet: true}
}

func (v NullableOrganizationWithSubscriptionDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWithSubscriptionDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

