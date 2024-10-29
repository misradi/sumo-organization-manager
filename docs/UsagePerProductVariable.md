# UsagePerProductVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductVariable** | **string** | A Product Variable is a unique service performance feature that is tracked through credit utilization. Valid values are &#39;continuousIngest&#39;, &#39;frequentIngest&#39;, &#39;storage&#39;, &#39;metrics&#39;, &#39;infrequentScan&#39;, &#39;infrequentIngest&#39;, &#39;inFrequentStorage&#39;, &#39;cseIngest&#39;, &#39;cseStorage&#39;. | 
**CreditsUsed** | **float64** | Denotes the total number of actual credits that have been used. | 
**DeploymentChargeCredits** | **float64** | Denotes the total number of credits that have been used in form of deployment charges. | 
**CreditsDeducted** | **float64** | Denotes the total number of credits that have been used including deployment charges. | 
**Utilization** | **float64** | The native utilization of the product variable. | 
**Unit** | **string** | The unit in which the native utilization is measured. | 

## Methods

### NewUsagePerProductVariable

`func NewUsagePerProductVariable(productVariable string, creditsUsed float64, deploymentChargeCredits float64, creditsDeducted float64, utilization float64, unit string, ) *UsagePerProductVariable`

NewUsagePerProductVariable instantiates a new UsagePerProductVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagePerProductVariableWithDefaults

`func NewUsagePerProductVariableWithDefaults() *UsagePerProductVariable`

NewUsagePerProductVariableWithDefaults instantiates a new UsagePerProductVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductVariable

`func (o *UsagePerProductVariable) GetProductVariable() string`

GetProductVariable returns the ProductVariable field if non-nil, zero value otherwise.

### GetProductVariableOk

`func (o *UsagePerProductVariable) GetProductVariableOk() (*string, bool)`

GetProductVariableOk returns a tuple with the ProductVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariable

`func (o *UsagePerProductVariable) SetProductVariable(v string)`

SetProductVariable sets ProductVariable field to given value.


### GetCreditsUsed

`func (o *UsagePerProductVariable) GetCreditsUsed() float64`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *UsagePerProductVariable) GetCreditsUsedOk() (*float64, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *UsagePerProductVariable) SetCreditsUsed(v float64)`

SetCreditsUsed sets CreditsUsed field to given value.


### GetDeploymentChargeCredits

`func (o *UsagePerProductVariable) GetDeploymentChargeCredits() float64`

GetDeploymentChargeCredits returns the DeploymentChargeCredits field if non-nil, zero value otherwise.

### GetDeploymentChargeCreditsOk

`func (o *UsagePerProductVariable) GetDeploymentChargeCreditsOk() (*float64, bool)`

GetDeploymentChargeCreditsOk returns a tuple with the DeploymentChargeCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChargeCredits

`func (o *UsagePerProductVariable) SetDeploymentChargeCredits(v float64)`

SetDeploymentChargeCredits sets DeploymentChargeCredits field to given value.


### GetCreditsDeducted

`func (o *UsagePerProductVariable) GetCreditsDeducted() float64`

GetCreditsDeducted returns the CreditsDeducted field if non-nil, zero value otherwise.

### GetCreditsDeductedOk

`func (o *UsagePerProductVariable) GetCreditsDeductedOk() (*float64, bool)`

GetCreditsDeductedOk returns a tuple with the CreditsDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsDeducted

`func (o *UsagePerProductVariable) SetCreditsDeducted(v float64)`

SetCreditsDeducted sets CreditsDeducted field to given value.


### GetUtilization

`func (o *UsagePerProductVariable) GetUtilization() float64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *UsagePerProductVariable) GetUtilizationOk() (*float64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *UsagePerProductVariable) SetUtilization(v float64)`

SetUtilization sets Utilization field to given value.


### GetUnit

`func (o *UsagePerProductVariable) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UsagePerProductVariable) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UsagePerProductVariable) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


