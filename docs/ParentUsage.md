# ParentUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCapacity** | **float64** | Denotes the total number of credits provisioned for the child organization to use. | 
**TotalCreditsUsed** | **float64** | Denotes the total number of credits that have been utilized. | 
**CreditsAllocated** | **float64** | Denotes the total number of credits that have been allocated to the child organizations. | 

## Methods

### NewParentUsage

`func NewParentUsage(totalCapacity float64, totalCreditsUsed float64, creditsAllocated float64, ) *ParentUsage`

NewParentUsage instantiates a new ParentUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentUsageWithDefaults

`func NewParentUsageWithDefaults() *ParentUsage`

NewParentUsageWithDefaults instantiates a new ParentUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCapacity

`func (o *ParentUsage) GetTotalCapacity() float64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *ParentUsage) GetTotalCapacityOk() (*float64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *ParentUsage) SetTotalCapacity(v float64)`

SetTotalCapacity sets TotalCapacity field to given value.


### GetTotalCreditsUsed

`func (o *ParentUsage) GetTotalCreditsUsed() float64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *ParentUsage) GetTotalCreditsUsedOk() (*float64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *ParentUsage) SetTotalCreditsUsed(v float64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetCreditsAllocated

`func (o *ParentUsage) GetCreditsAllocated() float64`

GetCreditsAllocated returns the CreditsAllocated field if non-nil, zero value otherwise.

### GetCreditsAllocatedOk

`func (o *ParentUsage) GetCreditsAllocatedOk() (*float64, bool)`

GetCreditsAllocatedOk returns a tuple with the CreditsAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAllocated

`func (o *ParentUsage) SetCreditsAllocated(v float64)`

SetCreditsAllocated sets CreditsAllocated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


