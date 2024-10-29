# ReadOrganizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the account owner. | 
**OrganizationName** | **string** | Name of the organization. | 
**FirstName** | **string** | First name of the account owner. | 
**LastName** | Pointer to **string** | Last name of the account owner. | [optional] 
**Subscription** | [**CreditsSubscription**](CreditsSubscription.md) |  | 
**OrgId** | **string** | The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash &#x60;-&#x60; character. | 
**DeploymentId** | Pointer to **string** | Identifier of the deployment in which the organization is present. | [optional] 

## Methods

### NewReadOrganizationResponse

`func NewReadOrganizationResponse(email string, organizationName string, firstName string, subscription CreditsSubscription, orgId string, ) *ReadOrganizationResponse`

NewReadOrganizationResponse instantiates a new ReadOrganizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOrganizationResponseWithDefaults

`func NewReadOrganizationResponseWithDefaults() *ReadOrganizationResponse`

NewReadOrganizationResponseWithDefaults instantiates a new ReadOrganizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ReadOrganizationResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReadOrganizationResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReadOrganizationResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationName

`func (o *ReadOrganizationResponse) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ReadOrganizationResponse) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ReadOrganizationResponse) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetFirstName

`func (o *ReadOrganizationResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ReadOrganizationResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ReadOrganizationResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ReadOrganizationResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ReadOrganizationResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ReadOrganizationResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ReadOrganizationResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetSubscription

`func (o *ReadOrganizationResponse) GetSubscription() CreditsSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ReadOrganizationResponse) GetSubscriptionOk() (*CreditsSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ReadOrganizationResponse) SetSubscription(v CreditsSubscription)`

SetSubscription sets Subscription field to given value.


### GetOrgId

`func (o *ReadOrganizationResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ReadOrganizationResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ReadOrganizationResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetDeploymentId

`func (o *ReadOrganizationResponse) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ReadOrganizationResponse) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ReadOrganizationResponse) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ReadOrganizationResponse) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


