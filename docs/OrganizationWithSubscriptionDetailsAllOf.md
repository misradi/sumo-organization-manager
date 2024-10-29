# OrganizationWithSubscriptionDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Identifier of the deployment in which the organization should be created. | 
**TrialPlanPeriod** | Pointer to **int32** | Specify the duration of the Trial plan. If not specified, your subscription plan will be used for the created organization. | [optional] 
**Baselines** | Pointer to [**Baselines**](Baselines.md) |  | [optional] 

## Methods

### NewOrganizationWithSubscriptionDetailsAllOf

`func NewOrganizationWithSubscriptionDetailsAllOf(deploymentId string, ) *OrganizationWithSubscriptionDetailsAllOf`

NewOrganizationWithSubscriptionDetailsAllOf instantiates a new OrganizationWithSubscriptionDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithSubscriptionDetailsAllOfWithDefaults

`func NewOrganizationWithSubscriptionDetailsAllOfWithDefaults() *OrganizationWithSubscriptionDetailsAllOf`

NewOrganizationWithSubscriptionDetailsAllOfWithDefaults instantiates a new OrganizationWithSubscriptionDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *OrganizationWithSubscriptionDetailsAllOf) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *OrganizationWithSubscriptionDetailsAllOf) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *OrganizationWithSubscriptionDetailsAllOf) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetTrialPlanPeriod

`func (o *OrganizationWithSubscriptionDetailsAllOf) GetTrialPlanPeriod() int32`

GetTrialPlanPeriod returns the TrialPlanPeriod field if non-nil, zero value otherwise.

### GetTrialPlanPeriodOk

`func (o *OrganizationWithSubscriptionDetailsAllOf) GetTrialPlanPeriodOk() (*int32, bool)`

GetTrialPlanPeriodOk returns a tuple with the TrialPlanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlanPeriod

`func (o *OrganizationWithSubscriptionDetailsAllOf) SetTrialPlanPeriod(v int32)`

SetTrialPlanPeriod sets TrialPlanPeriod field to given value.

### HasTrialPlanPeriod

`func (o *OrganizationWithSubscriptionDetailsAllOf) HasTrialPlanPeriod() bool`

HasTrialPlanPeriod returns a boolean if a field has been set.

### GetBaselines

`func (o *OrganizationWithSubscriptionDetailsAllOf) GetBaselines() Baselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *OrganizationWithSubscriptionDetailsAllOf) GetBaselinesOk() (*Baselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *OrganizationWithSubscriptionDetailsAllOf) SetBaselines(v Baselines)`

SetBaselines sets Baselines field to given value.

### HasBaselines

`func (o *OrganizationWithSubscriptionDetailsAllOf) HasBaselines() bool`

HasBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


