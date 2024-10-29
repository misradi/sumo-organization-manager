# CreditsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | Start date of the contract. | 
**EndDate** | **string** | End date of the contract. | 
**Status** | **string** | Status of the subscription. | 
**Plan** | [**Plan**](Plan.md) |  | 
**Credits** | **int64** | The total number of credits allocated to the organization. | 
**DeploymentChargeCredits** | Pointer to **int64** | The number of credits allocated to the organization in form of deployment charge. | [optional] 
**AllocatedCredits** | Pointer to **int64** | The number of credits allocated to the organization excluding deployment charge. | [optional] 
**Baselines** | [**Baselines**](Baselines.md) |  | 

## Methods

### NewCreditsSubscription

`func NewCreditsSubscription(startDate string, endDate string, status string, plan Plan, credits int64, baselines Baselines, ) *CreditsSubscription`

NewCreditsSubscription instantiates a new CreditsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditsSubscriptionWithDefaults

`func NewCreditsSubscriptionWithDefaults() *CreditsSubscription`

NewCreditsSubscriptionWithDefaults instantiates a new CreditsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *CreditsSubscription) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreditsSubscription) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreditsSubscription) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CreditsSubscription) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreditsSubscription) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreditsSubscription) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetStatus

`func (o *CreditsSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditsSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditsSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlan

`func (o *CreditsSubscription) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CreditsSubscription) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CreditsSubscription) SetPlan(v Plan)`

SetPlan sets Plan field to given value.


### GetCredits

`func (o *CreditsSubscription) GetCredits() int64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *CreditsSubscription) GetCreditsOk() (*int64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *CreditsSubscription) SetCredits(v int64)`

SetCredits sets Credits field to given value.


### GetDeploymentChargeCredits

`func (o *CreditsSubscription) GetDeploymentChargeCredits() int64`

GetDeploymentChargeCredits returns the DeploymentChargeCredits field if non-nil, zero value otherwise.

### GetDeploymentChargeCreditsOk

`func (o *CreditsSubscription) GetDeploymentChargeCreditsOk() (*int64, bool)`

GetDeploymentChargeCreditsOk returns a tuple with the DeploymentChargeCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChargeCredits

`func (o *CreditsSubscription) SetDeploymentChargeCredits(v int64)`

SetDeploymentChargeCredits sets DeploymentChargeCredits field to given value.

### HasDeploymentChargeCredits

`func (o *CreditsSubscription) HasDeploymentChargeCredits() bool`

HasDeploymentChargeCredits returns a boolean if a field has been set.

### GetAllocatedCredits

`func (o *CreditsSubscription) GetAllocatedCredits() int64`

GetAllocatedCredits returns the AllocatedCredits field if non-nil, zero value otherwise.

### GetAllocatedCreditsOk

`func (o *CreditsSubscription) GetAllocatedCreditsOk() (*int64, bool)`

GetAllocatedCreditsOk returns a tuple with the AllocatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCredits

`func (o *CreditsSubscription) SetAllocatedCredits(v int64)`

SetAllocatedCredits sets AllocatedCredits field to given value.

### HasAllocatedCredits

`func (o *CreditsSubscription) HasAllocatedCredits() bool`

HasAllocatedCredits returns a boolean if a field has been set.

### GetBaselines

`func (o *CreditsSubscription) GetBaselines() Baselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *CreditsSubscription) GetBaselinesOk() (*Baselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *CreditsSubscription) SetBaselines(v Baselines)`

SetBaselines sets Baselines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


