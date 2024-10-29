# TrialPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysLeft** | **int32** | The number of days left before the Trial period expires. Post expiry, the Trial plan will default to Free. | 

## Methods

### NewTrialPlan

`func NewTrialPlan(daysLeft int32, ) *TrialPlan`

NewTrialPlan instantiates a new TrialPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlanWithDefaults

`func NewTrialPlanWithDefaults() *TrialPlan`

NewTrialPlanWithDefaults instantiates a new TrialPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysLeft

`func (o *TrialPlan) GetDaysLeft() int32`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *TrialPlan) GetDaysLeftOk() (*int32, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *TrialPlan) SetDaysLeft(v int32)`

SetDaysLeft sets DaysLeft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


