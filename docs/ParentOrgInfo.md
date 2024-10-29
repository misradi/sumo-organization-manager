# ParentOrgInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEligibleForTrialOrgs** | Pointer to **bool** | Tells whether the parent org can set up trial child orgs subscriptions. | [optional] 
**IsEligibleForDeploymentCharge** | Pointer to **bool** | Tells whether the org is subject to deployment charges. | [optional] 
**DeploymentCharges** | Pointer to [**[]DeploymentCharge**](DeploymentCharge.md) | List of deployment charges for the customer for setting up child org in each deployment. | [optional] 
**PlanName** | Pointer to [**String**](String.md) | Plan name of the account. | [optional] 

## Methods

### NewParentOrgInfo

`func NewParentOrgInfo() *ParentOrgInfo`

NewParentOrgInfo instantiates a new ParentOrgInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentOrgInfoWithDefaults

`func NewParentOrgInfoWithDefaults() *ParentOrgInfo`

NewParentOrgInfoWithDefaults instantiates a new ParentOrgInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEligibleForTrialOrgs

`func (o *ParentOrgInfo) GetIsEligibleForTrialOrgs() bool`

GetIsEligibleForTrialOrgs returns the IsEligibleForTrialOrgs field if non-nil, zero value otherwise.

### GetIsEligibleForTrialOrgsOk

`func (o *ParentOrgInfo) GetIsEligibleForTrialOrgsOk() (*bool, bool)`

GetIsEligibleForTrialOrgsOk returns a tuple with the IsEligibleForTrialOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForTrialOrgs

`func (o *ParentOrgInfo) SetIsEligibleForTrialOrgs(v bool)`

SetIsEligibleForTrialOrgs sets IsEligibleForTrialOrgs field to given value.

### HasIsEligibleForTrialOrgs

`func (o *ParentOrgInfo) HasIsEligibleForTrialOrgs() bool`

HasIsEligibleForTrialOrgs returns a boolean if a field has been set.

### GetIsEligibleForDeploymentCharge

`func (o *ParentOrgInfo) GetIsEligibleForDeploymentCharge() bool`

GetIsEligibleForDeploymentCharge returns the IsEligibleForDeploymentCharge field if non-nil, zero value otherwise.

### GetIsEligibleForDeploymentChargeOk

`func (o *ParentOrgInfo) GetIsEligibleForDeploymentChargeOk() (*bool, bool)`

GetIsEligibleForDeploymentChargeOk returns a tuple with the IsEligibleForDeploymentCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForDeploymentCharge

`func (o *ParentOrgInfo) SetIsEligibleForDeploymentCharge(v bool)`

SetIsEligibleForDeploymentCharge sets IsEligibleForDeploymentCharge field to given value.

### HasIsEligibleForDeploymentCharge

`func (o *ParentOrgInfo) HasIsEligibleForDeploymentCharge() bool`

HasIsEligibleForDeploymentCharge returns a boolean if a field has been set.

### GetDeploymentCharges

`func (o *ParentOrgInfo) GetDeploymentCharges() []DeploymentCharge`

GetDeploymentCharges returns the DeploymentCharges field if non-nil, zero value otherwise.

### GetDeploymentChargesOk

`func (o *ParentOrgInfo) GetDeploymentChargesOk() (*[]DeploymentCharge, bool)`

GetDeploymentChargesOk returns a tuple with the DeploymentCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCharges

`func (o *ParentOrgInfo) SetDeploymentCharges(v []DeploymentCharge)`

SetDeploymentCharges sets DeploymentCharges field to given value.

### HasDeploymentCharges

`func (o *ParentOrgInfo) HasDeploymentCharges() bool`

HasDeploymentCharges returns a boolean if a field has been set.

### GetPlanName

`func (o *ParentOrgInfo) GetPlanName() String`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *ParentOrgInfo) GetPlanNameOk() (*String, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *ParentOrgInfo) SetPlanName(v String)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *ParentOrgInfo) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


