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

// checks if the CreditsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditsSubscription{}

// CreditsSubscription struct for CreditsSubscription
type CreditsSubscription struct {
	// Start date of the contract.
	StartDate string `json:"startDate"`
	// End date of the contract.
	EndDate string `json:"endDate"`
	// Status of the subscription.
	Status string `json:"status" validate:"regexp=^(Draft|Active|Canceled|Error)$"`
	Plan Plan `json:"plan"`
	// The total number of credits allocated to the organization.
	Credits int64 `json:"credits"`
	// The number of credits allocated to the organization in form of deployment charge.
	DeploymentChargeCredits *int64 `json:"deploymentChargeCredits,omitempty"`
	// The number of credits allocated to the organization excluding deployment charge.
	AllocatedCredits *int64 `json:"allocatedCredits,omitempty"`
	Baselines Baselines `json:"baselines"`
}

type _CreditsSubscription CreditsSubscription

// NewCreditsSubscription instantiates a new CreditsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditsSubscription(startDate string, endDate string, status string, plan Plan, credits int64, baselines Baselines) *CreditsSubscription {
	this := CreditsSubscription{}
	this.StartDate = startDate
	this.EndDate = endDate
	this.Status = status
	this.Plan = plan
	this.Credits = credits
	this.Baselines = baselines
	return &this
}

// NewCreditsSubscriptionWithDefaults instantiates a new CreditsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditsSubscriptionWithDefaults() *CreditsSubscription {
	this := CreditsSubscription{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *CreditsSubscription) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CreditsSubscription) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *CreditsSubscription) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *CreditsSubscription) SetEndDate(v string) {
	o.EndDate = v
}

// GetStatus returns the Status field value
func (o *CreditsSubscription) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreditsSubscription) SetStatus(v string) {
	o.Status = v
}

// GetPlan returns the Plan field value
func (o *CreditsSubscription) GetPlan() Plan {
	if o == nil {
		var ret Plan
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetPlanOk() (*Plan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *CreditsSubscription) SetPlan(v Plan) {
	o.Plan = v
}

// GetCredits returns the Credits field value
func (o *CreditsSubscription) GetCredits() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetCreditsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credits, true
}

// SetCredits sets field value
func (o *CreditsSubscription) SetCredits(v int64) {
	o.Credits = v
}

// GetDeploymentChargeCredits returns the DeploymentChargeCredits field value if set, zero value otherwise.
func (o *CreditsSubscription) GetDeploymentChargeCredits() int64 {
	if o == nil || IsNil(o.DeploymentChargeCredits) {
		var ret int64
		return ret
	}
	return *o.DeploymentChargeCredits
}

// GetDeploymentChargeCreditsOk returns a tuple with the DeploymentChargeCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetDeploymentChargeCreditsOk() (*int64, bool) {
	if o == nil || IsNil(o.DeploymentChargeCredits) {
		return nil, false
	}
	return o.DeploymentChargeCredits, true
}

// HasDeploymentChargeCredits returns a boolean if a field has been set.
func (o *CreditsSubscription) HasDeploymentChargeCredits() bool {
	if o != nil && !IsNil(o.DeploymentChargeCredits) {
		return true
	}

	return false
}

// SetDeploymentChargeCredits gets a reference to the given int64 and assigns it to the DeploymentChargeCredits field.
func (o *CreditsSubscription) SetDeploymentChargeCredits(v int64) {
	o.DeploymentChargeCredits = &v
}

// GetAllocatedCredits returns the AllocatedCredits field value if set, zero value otherwise.
func (o *CreditsSubscription) GetAllocatedCredits() int64 {
	if o == nil || IsNil(o.AllocatedCredits) {
		var ret int64
		return ret
	}
	return *o.AllocatedCredits
}

// GetAllocatedCreditsOk returns a tuple with the AllocatedCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetAllocatedCreditsOk() (*int64, bool) {
	if o == nil || IsNil(o.AllocatedCredits) {
		return nil, false
	}
	return o.AllocatedCredits, true
}

// HasAllocatedCredits returns a boolean if a field has been set.
func (o *CreditsSubscription) HasAllocatedCredits() bool {
	if o != nil && !IsNil(o.AllocatedCredits) {
		return true
	}

	return false
}

// SetAllocatedCredits gets a reference to the given int64 and assigns it to the AllocatedCredits field.
func (o *CreditsSubscription) SetAllocatedCredits(v int64) {
	o.AllocatedCredits = &v
}

// GetBaselines returns the Baselines field value
func (o *CreditsSubscription) GetBaselines() Baselines {
	if o == nil {
		var ret Baselines
		return ret
	}

	return o.Baselines
}

// GetBaselinesOk returns a tuple with the Baselines field value
// and a boolean to check if the value has been set.
func (o *CreditsSubscription) GetBaselinesOk() (*Baselines, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Baselines, true
}

// SetBaselines sets field value
func (o *CreditsSubscription) SetBaselines(v Baselines) {
	o.Baselines = v
}

func (o CreditsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	toSerialize["status"] = o.Status
	toSerialize["plan"] = o.Plan
	toSerialize["credits"] = o.Credits
	if !IsNil(o.DeploymentChargeCredits) {
		toSerialize["deploymentChargeCredits"] = o.DeploymentChargeCredits
	}
	if !IsNil(o.AllocatedCredits) {
		toSerialize["allocatedCredits"] = o.AllocatedCredits
	}
	toSerialize["baselines"] = o.Baselines
	return toSerialize, nil
}

func (o *CreditsSubscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startDate",
		"endDate",
		"status",
		"plan",
		"credits",
		"baselines",
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

	varCreditsSubscription := _CreditsSubscription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditsSubscription)

	if err != nil {
		return err
	}

	*o = CreditsSubscription(varCreditsSubscription)

	return err
}

type NullableCreditsSubscription struct {
	value *CreditsSubscription
	isSet bool
}

func (v NullableCreditsSubscription) Get() *CreditsSubscription {
	return v.value
}

func (v *NullableCreditsSubscription) Set(val *CreditsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditsSubscription(val *CreditsSubscription) *NullableCreditsSubscription {
	return &NullableCreditsSubscription{value: val, isSet: true}
}

func (v NullableCreditsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


