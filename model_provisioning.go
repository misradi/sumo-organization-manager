/*
Sumo Logic Organizations Management API

Welcome to the Sumo Logic's API Reference for Organizations Management. You can use these APIs to interact with the Sumo Logic platform to manage accounts and subscription. Refer to [API Authentication](https://help.sumologic.com/APIs/General-API-Information/API-Authentication) for more information about authentication. You can also look at [other APIs](https://help.sumologic.com/APIs) for more information about some other API endpoints. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Provisioning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Provisioning{}

// Provisioning struct for Provisioning
type Provisioning struct {
	Cse *ProvisioningStateForGroup `json:"cse,omitempty"`
}

// NewProvisioning instantiates a new Provisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioning() *Provisioning {
	this := Provisioning{}
	return &this
}

// NewProvisioningWithDefaults instantiates a new Provisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningWithDefaults() *Provisioning {
	this := Provisioning{}
	return &this
}

// GetCse returns the Cse field value if set, zero value otherwise.
func (o *Provisioning) GetCse() ProvisioningStateForGroup {
	if o == nil || IsNil(o.Cse) {
		var ret ProvisioningStateForGroup
		return ret
	}
	return *o.Cse
}

// GetCseOk returns a tuple with the Cse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetCseOk() (*ProvisioningStateForGroup, bool) {
	if o == nil || IsNil(o.Cse) {
		return nil, false
	}
	return o.Cse, true
}

// HasCse returns a boolean if a field has been set.
func (o *Provisioning) HasCse() bool {
	if o != nil && !IsNil(o.Cse) {
		return true
	}

	return false
}

// SetCse gets a reference to the given ProvisioningStateForGroup and assigns it to the Cse field.
func (o *Provisioning) SetCse(v ProvisioningStateForGroup) {
	o.Cse = &v
}

func (o Provisioning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Provisioning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cse) {
		toSerialize["cse"] = o.Cse
	}
	return toSerialize, nil
}

type NullableProvisioning struct {
	value *Provisioning
	isSet bool
}

func (v NullableProvisioning) Get() *Provisioning {
	return v.value
}

func (v *NullableProvisioning) Set(val *Provisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioning(val *Provisioning) *NullableProvisioning {
	return &NullableProvisioning{value: val, isSet: true}
}

func (v NullableProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

