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

// checks if the ParentUsageAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentUsageAllOf{}

// ParentUsageAllOf struct for ParentUsageAllOf
type ParentUsageAllOf struct {
	// Denotes the total number of credits that have been allocated to the child organizations.
	CreditsAllocated float64 `json:"creditsAllocated"`
}

type _ParentUsageAllOf ParentUsageAllOf

// NewParentUsageAllOf instantiates a new ParentUsageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentUsageAllOf(creditsAllocated float64) *ParentUsageAllOf {
	this := ParentUsageAllOf{}
	this.CreditsAllocated = creditsAllocated
	return &this
}

// NewParentUsageAllOfWithDefaults instantiates a new ParentUsageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentUsageAllOfWithDefaults() *ParentUsageAllOf {
	this := ParentUsageAllOf{}
	return &this
}

// GetCreditsAllocated returns the CreditsAllocated field value
func (o *ParentUsageAllOf) GetCreditsAllocated() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CreditsAllocated
}

// GetCreditsAllocatedOk returns a tuple with the CreditsAllocated field value
// and a boolean to check if the value has been set.
func (o *ParentUsageAllOf) GetCreditsAllocatedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsAllocated, true
}

// SetCreditsAllocated sets field value
func (o *ParentUsageAllOf) SetCreditsAllocated(v float64) {
	o.CreditsAllocated = v
}

func (o ParentUsageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentUsageAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creditsAllocated"] = o.CreditsAllocated
	return toSerialize, nil
}

func (o *ParentUsageAllOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"creditsAllocated",
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

	varParentUsageAllOf := _ParentUsageAllOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParentUsageAllOf)

	if err != nil {
		return err
	}

	*o = ParentUsageAllOf(varParentUsageAllOf)

	return err
}

type NullableParentUsageAllOf struct {
	value *ParentUsageAllOf
	isSet bool
}

func (v NullableParentUsageAllOf) Get() *ParentUsageAllOf {
	return v.value
}

func (v *NullableParentUsageAllOf) Set(val *ParentUsageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableParentUsageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableParentUsageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentUsageAllOf(val *ParentUsageAllOf) *NullableParentUsageAllOf {
	return &NullableParentUsageAllOf{value: val, isSet: true}
}

func (v NullableParentUsageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentUsageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


