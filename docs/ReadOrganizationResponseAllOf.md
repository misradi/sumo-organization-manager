# ReadOrganizationResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**CreditsSubscription**](CreditsSubscription.md) |  | 
**OrgId** | **string** | The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash &#x60;-&#x60; character. | 
**DeploymentId** | Pointer to **string** | Identifier of the deployment in which the organization is present. | [optional] 

## Methods

### NewReadOrganizationResponseAllOf

`func NewReadOrganizationResponseAllOf(subscription CreditsSubscription, orgId string, ) *ReadOrganizationResponseAllOf`

NewReadOrganizationResponseAllOf instantiates a new ReadOrganizationResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOrganizationResponseAllOfWithDefaults

`func NewReadOrganizationResponseAllOfWithDefaults() *ReadOrganizationResponseAllOf`

NewReadOrganizationResponseAllOfWithDefaults instantiates a new ReadOrganizationResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *ReadOrganizationResponseAllOf) GetSubscription() CreditsSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ReadOrganizationResponseAllOf) GetSubscriptionOk() (*CreditsSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ReadOrganizationResponseAllOf) SetSubscription(v CreditsSubscription)`

SetSubscription sets Subscription field to given value.


### GetOrgId

`func (o *ReadOrganizationResponseAllOf) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ReadOrganizationResponseAllOf) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ReadOrganizationResponseAllOf) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetDeploymentId

`func (o *ReadOrganizationResponseAllOf) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ReadOrganizationResponseAllOf) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ReadOrganizationResponseAllOf) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ReadOrganizationResponseAllOf) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


