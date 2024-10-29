# DeploymentCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | Identifier of the deployment for the child org for which deployment charge is applied. | [optional] 
**DeploymentCharge** | Pointer to **float64** | Deployment charge is a charge that applies to child orgs deployed in different regions. This number is a percentage applied to the total credits being allocated to the child org. | [optional] 

## Methods

### NewDeploymentCharge

`func NewDeploymentCharge() *DeploymentCharge`

NewDeploymentCharge instantiates a new DeploymentCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentChargeWithDefaults

`func NewDeploymentChargeWithDefaults() *DeploymentCharge`

NewDeploymentChargeWithDefaults instantiates a new DeploymentCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DeploymentCharge) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentCharge) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentCharge) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentCharge) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetDeploymentCharge

`func (o *DeploymentCharge) GetDeploymentCharge() float64`

GetDeploymentCharge returns the DeploymentCharge field if non-nil, zero value otherwise.

### GetDeploymentChargeOk

`func (o *DeploymentCharge) GetDeploymentChargeOk() (*float64, bool)`

GetDeploymentChargeOk returns a tuple with the DeploymentCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCharge

`func (o *DeploymentCharge) SetDeploymentCharge(v float64)`

SetDeploymentCharge sets DeploymentCharge field to given value.

### HasDeploymentCharge

`func (o *DeploymentCharge) HasDeploymentCharge() bool`

HasDeploymentCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


