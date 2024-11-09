# SetProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Duration** | Pointer to **int32** | Duration of a pause, in milliseconds.  Enabled should be false. | [optional] 

## Methods

### NewSetProtectionRequest

`func NewSetProtectionRequest(enabled bool, ) *SetProtectionRequest`

NewSetProtectionRequest instantiates a new SetProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetProtectionRequestWithDefaults

`func NewSetProtectionRequestWithDefaults() *SetProtectionRequest`

NewSetProtectionRequestWithDefaults instantiates a new SetProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetProtectionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetProtectionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetProtectionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDuration

`func (o *SetProtectionRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SetProtectionRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SetProtectionRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SetProtectionRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


