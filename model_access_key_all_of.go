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

// checks if the AccessKeyAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessKeyAllOf{}

// AccessKeyAllOf struct for AccessKeyAllOf
type AccessKeyAllOf struct {
	// The key for the created access key. This field will have values only in the response for an access key create request. The value will be an empty string while listing all keys.
	Key string `json:"key"`
}

type _AccessKeyAllOf AccessKeyAllOf

// NewAccessKeyAllOf instantiates a new AccessKeyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyAllOf(key string) *AccessKeyAllOf {
	this := AccessKeyAllOf{}
	this.Key = key
	return &this
}

// NewAccessKeyAllOfWithDefaults instantiates a new AccessKeyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyAllOfWithDefaults() *AccessKeyAllOf {
	this := AccessKeyAllOf{}
	return &this
}

// GetKey returns the Key field value
func (o *AccessKeyAllOf) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AccessKeyAllOf) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AccessKeyAllOf) SetKey(v string) {
	o.Key = v
}

func (o AccessKeyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessKeyAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *AccessKeyAllOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varAccessKeyAllOf := _AccessKeyAllOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessKeyAllOf)

	if err != nil {
		return err
	}

	*o = AccessKeyAllOf(varAccessKeyAllOf)

	return err
}

type NullableAccessKeyAllOf struct {
	value *AccessKeyAllOf
	isSet bool
}

func (v NullableAccessKeyAllOf) Get() *AccessKeyAllOf {
	return v.value
}

func (v *NullableAccessKeyAllOf) Set(val *AccessKeyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeyAllOf(val *AccessKeyAllOf) *NullableAccessKeyAllOf {
	return &NullableAccessKeyAllOf{value: val, isSet: true}
}

func (v NullableAccessKeyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

