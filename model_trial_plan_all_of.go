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

// checks if the TrialPlanAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrialPlanAllOf{}

// TrialPlanAllOf struct for TrialPlanAllOf
type TrialPlanAllOf struct {
	// The number of days left before the Trial period expires. Post expiry, the Trial plan will default to Free.
	DaysLeft int32 `json:"daysLeft"`
}

type _TrialPlanAllOf TrialPlanAllOf

// NewTrialPlanAllOf instantiates a new TrialPlanAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialPlanAllOf(daysLeft int32) *TrialPlanAllOf {
	this := TrialPlanAllOf{}
	this.DaysLeft = daysLeft
	return &this
}

// NewTrialPlanAllOfWithDefaults instantiates a new TrialPlanAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialPlanAllOfWithDefaults() *TrialPlanAllOf {
	this := TrialPlanAllOf{}
	return &this
}

// GetDaysLeft returns the DaysLeft field value
func (o *TrialPlanAllOf) GetDaysLeft() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysLeft
}

// GetDaysLeftOk returns a tuple with the DaysLeft field value
// and a boolean to check if the value has been set.
func (o *TrialPlanAllOf) GetDaysLeftOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysLeft, true
}

// SetDaysLeft sets field value
func (o *TrialPlanAllOf) SetDaysLeft(v int32) {
	o.DaysLeft = v
}

func (o TrialPlanAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrialPlanAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["daysLeft"] = o.DaysLeft
	return toSerialize, nil
}

func (o *TrialPlanAllOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"daysLeft",
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

	varTrialPlanAllOf := _TrialPlanAllOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTrialPlanAllOf)

	if err != nil {
		return err
	}

	*o = TrialPlanAllOf(varTrialPlanAllOf)

	return err
}

type NullableTrialPlanAllOf struct {
	value *TrialPlanAllOf
	isSet bool
}

func (v NullableTrialPlanAllOf) Get() *TrialPlanAllOf {
	return v.value
}

func (v *NullableTrialPlanAllOf) Set(val *TrialPlanAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialPlanAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialPlanAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialPlanAllOf(val *TrialPlanAllOf) *NullableTrialPlanAllOf {
	return &NullableTrialPlanAllOf{value: val, isSet: true}
}

func (v NullableTrialPlanAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialPlanAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


