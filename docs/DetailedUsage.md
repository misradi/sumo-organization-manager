# DetailedUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCapacity** | **float64** | Denotes the total number of credits provisioned for the child organization to use. | 
**TotalCreditsUsed** | **float64** | Denotes the total number of credits that have been utilized. | 
**DeploymentChargeCreditsUsed** | **float64** | The amount of credits used by the organization in form of deployment charge. | 
**AllocatedCreditsUsed** | **float64** | The amount of credits used by the organization excluding deployment charge. | 
**Usages** | [**[]UsagePerProductVariable**](UsagePerProductVariable.md) | Contains details of the credits used per product variable. | 

## Methods

### NewDetailedUsage

`func NewDetailedUsage(totalCapacity float64, totalCreditsUsed float64, deploymentChargeCreditsUsed float64, allocatedCreditsUsed float64, usages []UsagePerProductVariable, ) *DetailedUsage`

NewDetailedUsage instantiates a new DetailedUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedUsageWithDefaults

`func NewDetailedUsageWithDefaults() *DetailedUsage`

NewDetailedUsageWithDefaults instantiates a new DetailedUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCapacity

`func (o *DetailedUsage) GetTotalCapacity() float64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *DetailedUsage) GetTotalCapacityOk() (*float64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *DetailedUsage) SetTotalCapacity(v float64)`

SetTotalCapacity sets TotalCapacity field to given value.


### GetTotalCreditsUsed

`func (o *DetailedUsage) GetTotalCreditsUsed() float64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *DetailedUsage) GetTotalCreditsUsedOk() (*float64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *DetailedUsage) SetTotalCreditsUsed(v float64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetDeploymentChargeCreditsUsed

`func (o *DetailedUsage) GetDeploymentChargeCreditsUsed() float64`

GetDeploymentChargeCreditsUsed returns the DeploymentChargeCreditsUsed field if non-nil, zero value otherwise.

### GetDeploymentChargeCreditsUsedOk

`func (o *DetailedUsage) GetDeploymentChargeCreditsUsedOk() (*float64, bool)`

GetDeploymentChargeCreditsUsedOk returns a tuple with the DeploymentChargeCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChargeCreditsUsed

`func (o *DetailedUsage) SetDeploymentChargeCreditsUsed(v float64)`

SetDeploymentChargeCreditsUsed sets DeploymentChargeCreditsUsed field to given value.


### GetAllocatedCreditsUsed

`func (o *DetailedUsage) GetAllocatedCreditsUsed() float64`

GetAllocatedCreditsUsed returns the AllocatedCreditsUsed field if non-nil, zero value otherwise.

### GetAllocatedCreditsUsedOk

`func (o *DetailedUsage) GetAllocatedCreditsUsedOk() (*float64, bool)`

GetAllocatedCreditsUsedOk returns a tuple with the AllocatedCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCreditsUsed

`func (o *DetailedUsage) SetAllocatedCreditsUsed(v float64)`

SetAllocatedCreditsUsed sets AllocatedCreditsUsed field to given value.


### GetUsages

`func (o *DetailedUsage) GetUsages() []UsagePerProductVariable`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *DetailedUsage) GetUsagesOk() (*[]UsagePerProductVariable, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *DetailedUsage) SetUsages(v []UsagePerProductVariable)`

SetUsages sets Usages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


