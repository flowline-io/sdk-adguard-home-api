# GetStatsConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Are statistics enabled | 
**Interval** | **float32** | Statistics rotation interval in milliseconds | 
**Ignored** | **[]string** | List of host names, which should not be counted | 

## Methods

### NewGetStatsConfigResponse

`func NewGetStatsConfigResponse(enabled bool, interval float32, ignored []string, ) *GetStatsConfigResponse`

NewGetStatsConfigResponse instantiates a new GetStatsConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatsConfigResponseWithDefaults

`func NewGetStatsConfigResponseWithDefaults() *GetStatsConfigResponse`

NewGetStatsConfigResponseWithDefaults instantiates a new GetStatsConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetStatsConfigResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetStatsConfigResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetStatsConfigResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInterval

`func (o *GetStatsConfigResponse) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetStatsConfigResponse) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetStatsConfigResponse) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetIgnored

`func (o *GetStatsConfigResponse) GetIgnored() []string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *GetStatsConfigResponse) GetIgnoredOk() (*[]string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *GetStatsConfigResponse) SetIgnored(v []string)`

SetIgnored sets Ignored field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


