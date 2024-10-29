# OrganizationWithSubscriptionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the account owner. | 
**OrganizationName** | **string** | Name of the organization. | 
**FirstName** | **string** | First name of the account owner. | 
**LastName** | Pointer to **string** | Last name of the account owner. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment in which the organization should be created. | 
**TrialPlanPeriod** | Pointer to **int32** | Specify the duration of the Trial plan. If not specified, your subscription plan will be used for the created organization. | [optional] 
**Baselines** | Pointer to [**Baselines**](Baselines.md) |  | [optional] 

## Methods

### NewOrganizationWithSubscriptionDetails

`func NewOrganizationWithSubscriptionDetails(email string, organizationName string, firstName string, deploymentId string, ) *OrganizationWithSubscriptionDetails`

NewOrganizationWithSubscriptionDetails instantiates a new OrganizationWithSubscriptionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithSubscriptionDetailsWithDefaults

`func NewOrganizationWithSubscriptionDetailsWithDefaults() *OrganizationWithSubscriptionDetails`

NewOrganizationWithSubscriptionDetailsWithDefaults instantiates a new OrganizationWithSubscriptionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationWithSubscriptionDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationWithSubscriptionDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationWithSubscriptionDetails) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationName

`func (o *OrganizationWithSubscriptionDetails) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationWithSubscriptionDetails) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationWithSubscriptionDetails) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetFirstName

`func (o *OrganizationWithSubscriptionDetails) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrganizationWithSubscriptionDetails) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrganizationWithSubscriptionDetails) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *OrganizationWithSubscriptionDetails) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrganizationWithSubscriptionDetails) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrganizationWithSubscriptionDetails) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrganizationWithSubscriptionDetails) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDeploymentId

`func (o *OrganizationWithSubscriptionDetails) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *OrganizationWithSubscriptionDetails) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *OrganizationWithSubscriptionDetails) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetTrialPlanPeriod

`func (o *OrganizationWithSubscriptionDetails) GetTrialPlanPeriod() int32`

GetTrialPlanPeriod returns the TrialPlanPeriod field if non-nil, zero value otherwise.

### GetTrialPlanPeriodOk

`func (o *OrganizationWithSubscriptionDetails) GetTrialPlanPeriodOk() (*int32, bool)`

GetTrialPlanPeriodOk returns a tuple with the TrialPlanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlanPeriod

`func (o *OrganizationWithSubscriptionDetails) SetTrialPlanPeriod(v int32)`

SetTrialPlanPeriod sets TrialPlanPeriod field to given value.

### HasTrialPlanPeriod

`func (o *OrganizationWithSubscriptionDetails) HasTrialPlanPeriod() bool`

HasTrialPlanPeriod returns a boolean if a field has been set.

### GetBaselines

`func (o *OrganizationWithSubscriptionDetails) GetBaselines() Baselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *OrganizationWithSubscriptionDetails) GetBaselinesOk() (*Baselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *OrganizationWithSubscriptionDetails) SetBaselines(v Baselines)`

SetBaselines sets Baselines field to given value.

### HasBaselines

`func (o *OrganizationWithSubscriptionDetails) HasBaselines() bool`

HasBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


