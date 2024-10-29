# BasicUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCapacity** | **float64** | Denotes the total number of credits provisioned for the child organization to use. | 
**TotalCreditsUsed** | **float64** | Denotes the total number of credits that have been utilized. | 
**DeploymentChargeCreditsUsed** | **float64** | The amount of credits used by the organization in form of deployment charge. | 
**AllocatedCreditsUsed** | **float64** | The amount of credits used by the organization excluding deployment charge. | 
**OrgId** | **string** | The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash &#x60;-&#x60; character. | 

## Methods

### NewBasicUsage

`func NewBasicUsage(totalCapacity float64, totalCreditsUsed float64, deploymentChargeCreditsUsed float64, allocatedCreditsUsed float64, orgId string, ) *BasicUsage`

NewBasicUsage instantiates a new BasicUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicUsageWithDefaults

`func NewBasicUsageWithDefaults() *BasicUsage`

NewBasicUsageWithDefaults instantiates a new BasicUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCapacity

`func (o *BasicUsage) GetTotalCapacity() float64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *BasicUsage) GetTotalCapacityOk() (*float64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *BasicUsage) SetTotalCapacity(v float64)`

SetTotalCapacity sets TotalCapacity field to given value.


### GetTotalCreditsUsed

`func (o *BasicUsage) GetTotalCreditsUsed() float64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *BasicUsage) GetTotalCreditsUsedOk() (*float64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *BasicUsage) SetTotalCreditsUsed(v float64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetDeploymentChargeCreditsUsed

`func (o *BasicUsage) GetDeploymentChargeCreditsUsed() float64`

GetDeploymentChargeCreditsUsed returns the DeploymentChargeCreditsUsed field if non-nil, zero value otherwise.

### GetDeploymentChargeCreditsUsedOk

`func (o *BasicUsage) GetDeploymentChargeCreditsUsedOk() (*float64, bool)`

GetDeploymentChargeCreditsUsedOk returns a tuple with the DeploymentChargeCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChargeCreditsUsed

`func (o *BasicUsage) SetDeploymentChargeCreditsUsed(v float64)`

SetDeploymentChargeCreditsUsed sets DeploymentChargeCreditsUsed field to given value.


### GetAllocatedCreditsUsed

`func (o *BasicUsage) GetAllocatedCreditsUsed() float64`

GetAllocatedCreditsUsed returns the AllocatedCreditsUsed field if non-nil, zero value otherwise.

### GetAllocatedCreditsUsedOk

`func (o *BasicUsage) GetAllocatedCreditsUsedOk() (*float64, bool)`

GetAllocatedCreditsUsedOk returns a tuple with the AllocatedCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCreditsUsed

`func (o *BasicUsage) SetAllocatedCreditsUsed(v float64)`

SetAllocatedCreditsUsed sets AllocatedCreditsUsed field to given value.


### GetOrgId

`func (o *BasicUsage) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BasicUsage) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BasicUsage) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


