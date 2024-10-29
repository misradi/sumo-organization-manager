# Baselines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousIngest** | Pointer to **int64** | The average daily amount of continuous logs this child org is expected to ingest, in GB. | [optional] [default to 0]
**FrequentIngest** | Pointer to **int64** | The average daily amount of frequent logs this child org is expected to ingest, in GB. | [optional] [default to 0]
**InfrequentIngest** | Pointer to **int64** | The average daily amount of infrequent logs this child org is expected to ingest, in GB. | [optional] [default to 0]
**Metrics** | Pointer to **int64** | The average daily amount of metrics this child org is expected to ingest, in DPMs(Data points per minute). | [optional] [default to 0]
**TracingIngest** | Pointer to **int64** | The average daily amount of tracing data this child org is expected to ingest, in GB. This is currently not available for all customers. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to 0]
**CseIngest** | Pointer to **int64** | The average daily amount of cse data this child org is expected to ingest, in GB.  This is currently not available for all customers. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to 0]
**FeatureActivation** | Pointer to [**FeatureActivation**](FeatureActivation.md) |  | [optional] 
**TotalCreditsAllocated** | Pointer to **int64** | Total number of credits to be allocated to the child organization. This field is applicable only in CREATE a new organization and UPDATE an organization requests. | [optional] 

## Methods

### NewBaselines

`func NewBaselines() *Baselines`

NewBaselines instantiates a new Baselines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaselinesWithDefaults

`func NewBaselinesWithDefaults() *Baselines`

NewBaselinesWithDefaults instantiates a new Baselines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousIngest

`func (o *Baselines) GetContinuousIngest() int64`

GetContinuousIngest returns the ContinuousIngest field if non-nil, zero value otherwise.

### GetContinuousIngestOk

`func (o *Baselines) GetContinuousIngestOk() (*int64, bool)`

GetContinuousIngestOk returns a tuple with the ContinuousIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousIngest

`func (o *Baselines) SetContinuousIngest(v int64)`

SetContinuousIngest sets ContinuousIngest field to given value.

### HasContinuousIngest

`func (o *Baselines) HasContinuousIngest() bool`

HasContinuousIngest returns a boolean if a field has been set.

### GetFrequentIngest

`func (o *Baselines) GetFrequentIngest() int64`

GetFrequentIngest returns the FrequentIngest field if non-nil, zero value otherwise.

### GetFrequentIngestOk

`func (o *Baselines) GetFrequentIngestOk() (*int64, bool)`

GetFrequentIngestOk returns a tuple with the FrequentIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentIngest

`func (o *Baselines) SetFrequentIngest(v int64)`

SetFrequentIngest sets FrequentIngest field to given value.

### HasFrequentIngest

`func (o *Baselines) HasFrequentIngest() bool`

HasFrequentIngest returns a boolean if a field has been set.

### GetInfrequentIngest

`func (o *Baselines) GetInfrequentIngest() int64`

GetInfrequentIngest returns the InfrequentIngest field if non-nil, zero value otherwise.

### GetInfrequentIngestOk

`func (o *Baselines) GetInfrequentIngestOk() (*int64, bool)`

GetInfrequentIngestOk returns a tuple with the InfrequentIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequentIngest

`func (o *Baselines) SetInfrequentIngest(v int64)`

SetInfrequentIngest sets InfrequentIngest field to given value.

### HasInfrequentIngest

`func (o *Baselines) HasInfrequentIngest() bool`

HasInfrequentIngest returns a boolean if a field has been set.

### GetMetrics

`func (o *Baselines) GetMetrics() int64`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Baselines) GetMetricsOk() (*int64, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Baselines) SetMetrics(v int64)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Baselines) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTracingIngest

`func (o *Baselines) GetTracingIngest() int64`

GetTracingIngest returns the TracingIngest field if non-nil, zero value otherwise.

### GetTracingIngestOk

`func (o *Baselines) GetTracingIngestOk() (*int64, bool)`

GetTracingIngestOk returns a tuple with the TracingIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingIngest

`func (o *Baselines) SetTracingIngest(v int64)`

SetTracingIngest sets TracingIngest field to given value.

### HasTracingIngest

`func (o *Baselines) HasTracingIngest() bool`

HasTracingIngest returns a boolean if a field has been set.

### GetCseIngest

`func (o *Baselines) GetCseIngest() int64`

GetCseIngest returns the CseIngest field if non-nil, zero value otherwise.

### GetCseIngestOk

`func (o *Baselines) GetCseIngestOk() (*int64, bool)`

GetCseIngestOk returns a tuple with the CseIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseIngest

`func (o *Baselines) SetCseIngest(v int64)`

SetCseIngest sets CseIngest field to given value.

### HasCseIngest

`func (o *Baselines) HasCseIngest() bool`

HasCseIngest returns a boolean if a field has been set.

### GetFeatureActivation

`func (o *Baselines) GetFeatureActivation() FeatureActivation`

GetFeatureActivation returns the FeatureActivation field if non-nil, zero value otherwise.

### GetFeatureActivationOk

`func (o *Baselines) GetFeatureActivationOk() (*FeatureActivation, bool)`

GetFeatureActivationOk returns a tuple with the FeatureActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureActivation

`func (o *Baselines) SetFeatureActivation(v FeatureActivation)`

SetFeatureActivation sets FeatureActivation field to given value.

### HasFeatureActivation

`func (o *Baselines) HasFeatureActivation() bool`

HasFeatureActivation returns a boolean if a field has been set.

### GetTotalCreditsAllocated

`func (o *Baselines) GetTotalCreditsAllocated() int64`

GetTotalCreditsAllocated returns the TotalCreditsAllocated field if non-nil, zero value otherwise.

### GetTotalCreditsAllocatedOk

`func (o *Baselines) GetTotalCreditsAllocatedOk() (*int64, bool)`

GetTotalCreditsAllocatedOk returns a tuple with the TotalCreditsAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsAllocated

`func (o *Baselines) SetTotalCreditsAllocated(v int64)`

SetTotalCreditsAllocated sets TotalCreditsAllocated field to given value.

### HasTotalCreditsAllocated

`func (o *Baselines) HasTotalCreditsAllocated() bool`

HasTotalCreditsAllocated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


