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

// checks if the ErrorDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDescription{}

// ErrorDescription struct for ErrorDescription
type ErrorDescription struct {
	// An error code describing the type of error.
	Code string `json:"code"`
	// A short English-language description of the error.
	Message string `json:"message"`
	// An optional fuller English-language description of the error.
	Detail *string `json:"detail,omitempty"`
	// An optional list of metadata about the error.
	Meta map[string]interface{} `json:"meta,omitempty"`
}

type _ErrorDescription ErrorDescription

// NewErrorDescription instantiates a new ErrorDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDescription(code string, message string) *ErrorDescription {
	this := ErrorDescription{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorDescriptionWithDefaults instantiates a new ErrorDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDescriptionWithDefaults() *ErrorDescription {
	this := ErrorDescription{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorDescription) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorDescription) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorDescription) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorDescription) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorDescription) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorDescription) SetMessage(v string) {
	o.Message = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorDescription) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDescription) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorDescription) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorDescription) SetDetail(v string) {
	o.Detail = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ErrorDescription) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDescription) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ErrorDescription) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ErrorDescription) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o ErrorDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *ErrorDescription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varErrorDescription := _ErrorDescription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorDescription)

	if err != nil {
		return err
	}

	*o = ErrorDescription(varErrorDescription)

	return err
}

type NullableErrorDescription struct {
	value *ErrorDescription
	isSet bool
}

func (v NullableErrorDescription) Get() *ErrorDescription {
	return v.value
}

func (v *NullableErrorDescription) Set(val *ErrorDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDescription(val *ErrorDescription) *NullableErrorDescription {
	return &NullableErrorDescription{value: val, isSet: true}
}

func (v NullableErrorDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


