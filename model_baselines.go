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

// checks if the Baselines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Baselines{}

// Baselines Details of consumable and its quantity.
type Baselines struct {
	// The average daily amount of continuous logs this child org is expected to ingest, in GB.
	ContinuousIngest *int64 `json:"continuousIngest,omitempty"`
	// The average daily amount of frequent logs this child org is expected to ingest, in GB.
	FrequentIngest *int64 `json:"frequentIngest,omitempty"`
	// The average daily amount of infrequent logs this child org is expected to ingest, in GB.
	InfrequentIngest *int64 `json:"infrequentIngest,omitempty"`
	// The average daily amount of metrics this child org is expected to ingest, in DPMs(Data points per minute).
	Metrics *int64 `json:"metrics,omitempty"`
	// The average daily amount of tracing data this child org is expected to ingest, in GB. This is currently not available for all customers. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	TracingIngest *int64 `json:"tracingIngest,omitempty"`
	// The average daily amount of cse data this child org is expected to ingest, in GB.  This is currently not available for all customers. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	CseIngest *int64 `json:"cseIngest,omitempty"`
	FeatureActivation *FeatureActivation `json:"featureActivation,omitempty"`
	// Total number of credits to be allocated to the child organization. This field is applicable only in CREATE a new organization and UPDATE an organization requests.
	TotalCreditsAllocated *int64 `json:"totalCreditsAllocated,omitempty"`
}

// NewBaselines instantiates a new Baselines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaselines() *Baselines {
	this := Baselines{}
	var continuousIngest int64 = 0
	this.ContinuousIngest = &continuousIngest
	var frequentIngest int64 = 0
	this.FrequentIngest = &frequentIngest
	var infrequentIngest int64 = 0
	this.InfrequentIngest = &infrequentIngest
	var metrics int64 = 0
	this.Metrics = &metrics
	var tracingIngest int64 = 0
	this.TracingIngest = &tracingIngest
	var cseIngest int64 = 0
	this.CseIngest = &cseIngest
	return &this
}

// NewBaselinesWithDefaults instantiates a new Baselines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaselinesWithDefaults() *Baselines {
	this := Baselines{}
	var continuousIngest int64 = 0
	this.ContinuousIngest = &continuousIngest
	var frequentIngest int64 = 0
	this.FrequentIngest = &frequentIngest
	var infrequentIngest int64 = 0
	this.InfrequentIngest = &infrequentIngest
	var metrics int64 = 0
	this.Metrics = &metrics
	var tracingIngest int64 = 0
	this.TracingIngest = &tracingIngest
	var cseIngest int64 = 0
	this.CseIngest = &cseIngest
	return &this
}

// GetContinuousIngest returns the ContinuousIngest field value if set, zero value otherwise.
func (o *Baselines) GetContinuousIngest() int64 {
	if o == nil || IsNil(o.ContinuousIngest) {
		var ret int64
		return ret
	}
	return *o.ContinuousIngest
}

// GetContinuousIngestOk returns a tuple with the ContinuousIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetContinuousIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.ContinuousIngest) {
		return nil, false
	}
	return o.ContinuousIngest, true
}

// HasContinuousIngest returns a boolean if a field has been set.
func (o *Baselines) HasContinuousIngest() bool {
	if o != nil && !IsNil(o.ContinuousIngest) {
		return true
	}

	return false
}

// SetContinuousIngest gets a reference to the given int64 and assigns it to the ContinuousIngest field.
func (o *Baselines) SetContinuousIngest(v int64) {
	o.ContinuousIngest = &v
}

// GetFrequentIngest returns the FrequentIngest field value if set, zero value otherwise.
func (o *Baselines) GetFrequentIngest() int64 {
	if o == nil || IsNil(o.FrequentIngest) {
		var ret int64
		return ret
	}
	return *o.FrequentIngest
}

// GetFrequentIngestOk returns a tuple with the FrequentIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetFrequentIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.FrequentIngest) {
		return nil, false
	}
	return o.FrequentIngest, true
}

// HasFrequentIngest returns a boolean if a field has been set.
func (o *Baselines) HasFrequentIngest() bool {
	if o != nil && !IsNil(o.FrequentIngest) {
		return true
	}

	return false
}

// SetFrequentIngest gets a reference to the given int64 and assigns it to the FrequentIngest field.
func (o *Baselines) SetFrequentIngest(v int64) {
	o.FrequentIngest = &v
}

// GetInfrequentIngest returns the InfrequentIngest field value if set, zero value otherwise.
func (o *Baselines) GetInfrequentIngest() int64 {
	if o == nil || IsNil(o.InfrequentIngest) {
		var ret int64
		return ret
	}
	return *o.InfrequentIngest
}

// GetInfrequentIngestOk returns a tuple with the InfrequentIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetInfrequentIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.InfrequentIngest) {
		return nil, false
	}
	return o.InfrequentIngest, true
}

// HasInfrequentIngest returns a boolean if a field has been set.
func (o *Baselines) HasInfrequentIngest() bool {
	if o != nil && !IsNil(o.InfrequentIngest) {
		return true
	}

	return false
}

// SetInfrequentIngest gets a reference to the given int64 and assigns it to the InfrequentIngest field.
func (o *Baselines) SetInfrequentIngest(v int64) {
	o.InfrequentIngest = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Baselines) GetMetrics() int64 {
	if o == nil || IsNil(o.Metrics) {
		var ret int64
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetMetricsOk() (*int64, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Baselines) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given int64 and assigns it to the Metrics field.
func (o *Baselines) SetMetrics(v int64) {
	o.Metrics = &v
}

// GetTracingIngest returns the TracingIngest field value if set, zero value otherwise.
func (o *Baselines) GetTracingIngest() int64 {
	if o == nil || IsNil(o.TracingIngest) {
		var ret int64
		return ret
	}
	return *o.TracingIngest
}

// GetTracingIngestOk returns a tuple with the TracingIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetTracingIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.TracingIngest) {
		return nil, false
	}
	return o.TracingIngest, true
}

// HasTracingIngest returns a boolean if a field has been set.
func (o *Baselines) HasTracingIngest() bool {
	if o != nil && !IsNil(o.TracingIngest) {
		return true
	}

	return false
}

// SetTracingIngest gets a reference to the given int64 and assigns it to the TracingIngest field.
func (o *Baselines) SetTracingIngest(v int64) {
	o.TracingIngest = &v
}

// GetCseIngest returns the CseIngest field value if set, zero value otherwise.
func (o *Baselines) GetCseIngest() int64 {
	if o == nil || IsNil(o.CseIngest) {
		var ret int64
		return ret
	}
	return *o.CseIngest
}

// GetCseIngestOk returns a tuple with the CseIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetCseIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.CseIngest) {
		return nil, false
	}
	return o.CseIngest, true
}

// HasCseIngest returns a boolean if a field has been set.
func (o *Baselines) HasCseIngest() bool {
	if o != nil && !IsNil(o.CseIngest) {
		return true
	}

	return false
}

// SetCseIngest gets a reference to the given int64 and assigns it to the CseIngest field.
func (o *Baselines) SetCseIngest(v int64) {
	o.CseIngest = &v
}

// GetFeatureActivation returns the FeatureActivation field value if set, zero value otherwise.
func (o *Baselines) GetFeatureActivation() FeatureActivation {
	if o == nil || IsNil(o.FeatureActivation) {
		var ret FeatureActivation
		return ret
	}
	return *o.FeatureActivation
}

// GetFeatureActivationOk returns a tuple with the FeatureActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetFeatureActivationOk() (*FeatureActivation, bool) {
	if o == nil || IsNil(o.FeatureActivation) {
		return nil, false
	}
	return o.FeatureActivation, true
}

// HasFeatureActivation returns a boolean if a field has been set.
func (o *Baselines) HasFeatureActivation() bool {
	if o != nil && !IsNil(o.FeatureActivation) {
		return true
	}

	return false
}

// SetFeatureActivation gets a reference to the given FeatureActivation and assigns it to the FeatureActivation field.
func (o *Baselines) SetFeatureActivation(v FeatureActivation) {
	o.FeatureActivation = &v
}

// GetTotalCreditsAllocated returns the TotalCreditsAllocated field value if set, zero value otherwise.
func (o *Baselines) GetTotalCreditsAllocated() int64 {
	if o == nil || IsNil(o.TotalCreditsAllocated) {
		var ret int64
		return ret
	}
	return *o.TotalCreditsAllocated
}

// GetTotalCreditsAllocatedOk returns a tuple with the TotalCreditsAllocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetTotalCreditsAllocatedOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCreditsAllocated) {
		return nil, false
	}
	return o.TotalCreditsAllocated, true
}

// HasTotalCreditsAllocated returns a boolean if a field has been set.
func (o *Baselines) HasTotalCreditsAllocated() bool {
	if o != nil && !IsNil(o.TotalCreditsAllocated) {
		return true
	}

	return false
}

// SetTotalCreditsAllocated gets a reference to the given int64 and assigns it to the TotalCreditsAllocated field.
func (o *Baselines) SetTotalCreditsAllocated(v int64) {
	o.TotalCreditsAllocated = &v
}

func (o Baselines) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Baselines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContinuousIngest) {
		toSerialize["continuousIngest"] = o.ContinuousIngest
	}
	if !IsNil(o.FrequentIngest) {
		toSerialize["frequentIngest"] = o.FrequentIngest
	}
	if !IsNil(o.InfrequentIngest) {
		toSerialize["infrequentIngest"] = o.InfrequentIngest
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.TracingIngest) {
		toSerialize["tracingIngest"] = o.TracingIngest
	}
	if !IsNil(o.CseIngest) {
		toSerialize["cseIngest"] = o.CseIngest
	}
	if !IsNil(o.FeatureActivation) {
		toSerialize["featureActivation"] = o.FeatureActivation
	}
	if !IsNil(o.TotalCreditsAllocated) {
		toSerialize["totalCreditsAllocated"] = o.TotalCreditsAllocated
	}
	return toSerialize, nil
}

type NullableBaselines struct {
	value *Baselines
	isSet bool
}

func (v NullableBaselines) Get() *Baselines {
	return v.value
}

func (v *NullableBaselines) Set(val *Baselines) {
	v.value = val
	v.isSet = true
}

func (v NullableBaselines) IsSet() bool {
	return v.isSet
}

func (v *NullableBaselines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaselines(val *Baselines) *NullableBaselines {
	return &NullableBaselines{value: val, isSet: true}
}

func (v NullableBaselines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaselines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

