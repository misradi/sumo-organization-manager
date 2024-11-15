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

// checks if the EnterprisePlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterprisePlan{}

// EnterprisePlan struct for EnterprisePlan
type EnterprisePlan struct {
	Plan
}

type _EnterprisePlan EnterprisePlan

// NewEnterprisePlan instantiates a new EnterprisePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterprisePlan(planName string) *EnterprisePlan {
	this := EnterprisePlan{}
	this.PlanName = planName
	return &this
}

// NewEnterprisePlanWithDefaults instantiates a new EnterprisePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterprisePlanWithDefaults() *EnterprisePlan {
	this := EnterprisePlan{}
	return &this
}

func (o EnterprisePlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterprisePlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPlan, errPlan := json.Marshal(o.Plan)
	if errPlan != nil {
		return map[string]interface{}{}, errPlan
	}
	errPlan = json.Unmarshal([]byte(serializedPlan), &toSerialize)
	if errPlan != nil {
		return map[string]interface{}{}, errPlan
	}
	return toSerialize, nil
}

func (o *EnterprisePlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planName",
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

	varEnterprisePlan := _EnterprisePlan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnterprisePlan)

	if err != nil {
		return err
	}

	*o = EnterprisePlan(varEnterprisePlan)

	return err
}

type NullableEnterprisePlan struct {
	value *EnterprisePlan
	isSet bool
}

func (v NullableEnterprisePlan) Get() *EnterprisePlan {
	return v.value
}

func (v *NullableEnterprisePlan) Set(val *EnterprisePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterprisePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterprisePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterprisePlan(val *EnterprisePlan) *NullableEnterprisePlan {
	return &NullableEnterprisePlan{value: val, isSet: true}
}

func (v NullableEnterprisePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterprisePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


