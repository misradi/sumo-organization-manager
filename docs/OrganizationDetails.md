# OrganizationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the account owner. | 
**OrganizationName** | **string** | Name of the organization. | 
**FirstName** | **string** | First name of the account owner. | 
**LastName** | Pointer to **string** | Last name of the account owner. | [optional] 

## Methods

### NewOrganizationDetails

`func NewOrganizationDetails(email string, organizationName string, firstName string, ) *OrganizationDetails`

NewOrganizationDetails instantiates a new OrganizationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsWithDefaults

`func NewOrganizationDetailsWithDefaults() *OrganizationDetails`

NewOrganizationDetailsWithDefaults instantiates a new OrganizationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationDetails) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationName

`func (o *OrganizationDetails) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationDetails) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationDetails) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetFirstName

`func (o *OrganizationDetails) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrganizationDetails) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrganizationDetails) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *OrganizationDetails) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrganizationDetails) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrganizationDetails) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrganizationDetails) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


