# ProvisioningStateForGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProvisioned** | Pointer to **bool** | Is the product group enabled for the child organization. | [optional] [default to false]
**IsProvisioningToggleAllowed** | Pointer to **bool** | Whether updating the provisioning state of the child organization is allowed or not. | [optional] [default to false]

## Methods

### NewProvisioningStateForGroup

`func NewProvisioningStateForGroup() *ProvisioningStateForGroup`

NewProvisioningStateForGroup instantiates a new ProvisioningStateForGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningStateForGroupWithDefaults

`func NewProvisioningStateForGroupWithDefaults() *ProvisioningStateForGroup`

NewProvisioningStateForGroupWithDefaults instantiates a new ProvisioningStateForGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsProvisioned

`func (o *ProvisioningStateForGroup) GetIsProvisioned() bool`

GetIsProvisioned returns the IsProvisioned field if non-nil, zero value otherwise.

### GetIsProvisionedOk

`func (o *ProvisioningStateForGroup) GetIsProvisionedOk() (*bool, bool)`

GetIsProvisionedOk returns a tuple with the IsProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvisioned

`func (o *ProvisioningStateForGroup) SetIsProvisioned(v bool)`

SetIsProvisioned sets IsProvisioned field to given value.

### HasIsProvisioned

`func (o *ProvisioningStateForGroup) HasIsProvisioned() bool`

HasIsProvisioned returns a boolean if a field has been set.

### GetIsProvisioningToggleAllowed

`func (o *ProvisioningStateForGroup) GetIsProvisioningToggleAllowed() bool`

GetIsProvisioningToggleAllowed returns the IsProvisioningToggleAllowed field if non-nil, zero value otherwise.

### GetIsProvisioningToggleAllowedOk

`func (o *ProvisioningStateForGroup) GetIsProvisioningToggleAllowedOk() (*bool, bool)`

GetIsProvisioningToggleAllowedOk returns a tuple with the IsProvisioningToggleAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvisioningToggleAllowed

`func (o *ProvisioningStateForGroup) SetIsProvisioningToggleAllowed(v bool)`

SetIsProvisioningToggleAllowed sets IsProvisioningToggleAllowed field to given value.

### HasIsProvisioningToggleAllowed

`func (o *ProvisioningStateForGroup) HasIsProvisioningToggleAllowed() bool`

HasIsProvisioningToggleAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


