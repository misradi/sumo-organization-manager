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

// checks if the BasicUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicUsage{}

// BasicUsage struct for BasicUsage
type BasicUsage struct {
	// Denotes the total number of credits provisioned for the child organization to use.
	TotalCapacity float64 `json:"totalCapacity"`
	// Denotes the total number of credits that have been utilized.
	TotalCreditsUsed float64 `json:"totalCreditsUsed"`
	// The amount of credits used by the organization in form of deployment charge.
	DeploymentChargeCreditsUsed float64 `json:"deploymentChargeCreditsUsed"`
	// The amount of credits used by the organization excluding deployment charge.
	AllocatedCreditsUsed float64 `json:"allocatedCreditsUsed"`
	// The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash `-` character.
	OrgId string `json:"orgId"`
}

type _BasicUsage BasicUsage

// NewBasicUsage instantiates a new BasicUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicUsage(totalCapacity float64, totalCreditsUsed float64, deploymentChargeCreditsUsed float64, allocatedCreditsUsed float64, orgId string) *BasicUsage {
	this := BasicUsage{}
	this.TotalCapacity = totalCapacity
	this.TotalCreditsUsed = totalCreditsUsed
	this.DeploymentChargeCreditsUsed = deploymentChargeCreditsUsed
	this.AllocatedCreditsUsed = allocatedCreditsUsed
	this.OrgId = orgId
	return &this
}

// NewBasicUsageWithDefaults instantiates a new BasicUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicUsageWithDefaults() *BasicUsage {
	this := BasicUsage{}
	return &this
}

// GetTotalCapacity returns the TotalCapacity field value
func (o *BasicUsage) GetTotalCapacity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value
// and a boolean to check if the value has been set.
func (o *BasicUsage) GetTotalCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCapacity, true
}

// SetTotalCapacity sets field value
func (o *BasicUsage) SetTotalCapacity(v float64) {
	o.TotalCapacity = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *BasicUsage) GetTotalCreditsUsed() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *BasicUsage) GetTotalCreditsUsedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *BasicUsage) SetTotalCreditsUsed(v float64) {
	o.TotalCreditsUsed = v
}

// GetDeploymentChargeCreditsUsed returns the DeploymentChargeCreditsUsed field value
func (o *BasicUsage) GetDeploymentChargeCreditsUsed() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DeploymentChargeCreditsUsed
}

// GetDeploymentChargeCreditsUsedOk returns a tuple with the DeploymentChargeCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *BasicUsage) GetDeploymentChargeCreditsUsedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentChargeCreditsUsed, true
}

// SetDeploymentChargeCreditsUsed sets field value
func (o *BasicUsage) SetDeploymentChargeCreditsUsed(v float64) {
	o.DeploymentChargeCreditsUsed = v
}

// GetAllocatedCreditsUsed returns the AllocatedCreditsUsed field value
func (o *BasicUsage) GetAllocatedCreditsUsed() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AllocatedCreditsUsed
}

// GetAllocatedCreditsUsedOk returns a tuple with the AllocatedCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *BasicUsage) GetAllocatedCreditsUsedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedCreditsUsed, true
}

// SetAllocatedCreditsUsed sets field value
func (o *BasicUsage) SetAllocatedCreditsUsed(v float64) {
	o.AllocatedCreditsUsed = v
}

// GetOrgId returns the OrgId field value
func (o *BasicUsage) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *BasicUsage) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *BasicUsage) SetOrgId(v string) {
	o.OrgId = v
}

func (o BasicUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalCapacity"] = o.TotalCapacity
	toSerialize["totalCreditsUsed"] = o.TotalCreditsUsed
	toSerialize["deploymentChargeCreditsUsed"] = o.DeploymentChargeCreditsUsed
	toSerialize["allocatedCreditsUsed"] = o.AllocatedCreditsUsed
	toSerialize["orgId"] = o.OrgId
	return toSerialize, nil
}

func (o *BasicUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"totalCapacity",
		"totalCreditsUsed",
		"deploymentChargeCreditsUsed",
		"allocatedCreditsUsed",
		"orgId",
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

	varBasicUsage := _BasicUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBasicUsage)

	if err != nil {
		return err
	}

	*o = BasicUsage(varBasicUsage)

	return err
}

type NullableBasicUsage struct {
	value *BasicUsage
	isSet bool
}

func (v NullableBasicUsage) Get() *BasicUsage {
	return v.value
}

func (v *NullableBasicUsage) Set(val *BasicUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicUsage(val *BasicUsage) *NullableBasicUsage {
	return &NullableBasicUsage{value: val, isSet: true}
}

func (v NullableBasicUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


