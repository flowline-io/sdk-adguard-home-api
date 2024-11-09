# GetQueryLogConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Is query log enabled | 
**Interval** | **float32** | Time period for query log rotation in milliseconds.  | 
**AnonymizeClientIp** | **bool** | Anonymize clients&#39; IP addresses | 
**Ignored** | **[]string** | List of host names, which should not be written to log | 

## Methods

### NewGetQueryLogConfigResponse

`func NewGetQueryLogConfigResponse(enabled bool, interval float32, anonymizeClientIp bool, ignored []string, ) *GetQueryLogConfigResponse`

NewGetQueryLogConfigResponse instantiates a new GetQueryLogConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQueryLogConfigResponseWithDefaults

`func NewGetQueryLogConfigResponseWithDefaults() *GetQueryLogConfigResponse`

NewGetQueryLogConfigResponseWithDefaults instantiates a new GetQueryLogConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetQueryLogConfigResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetQueryLogConfigResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetQueryLogConfigResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInterval

`func (o *GetQueryLogConfigResponse) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetQueryLogConfigResponse) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetQueryLogConfigResponse) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetAnonymizeClientIp

`func (o *GetQueryLogConfigResponse) GetAnonymizeClientIp() bool`

GetAnonymizeClientIp returns the AnonymizeClientIp field if non-nil, zero value otherwise.

### GetAnonymizeClientIpOk

`func (o *GetQueryLogConfigResponse) GetAnonymizeClientIpOk() (*bool, bool)`

GetAnonymizeClientIpOk returns a tuple with the AnonymizeClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeClientIp

`func (o *GetQueryLogConfigResponse) SetAnonymizeClientIp(v bool)`

SetAnonymizeClientIp sets AnonymizeClientIp field to given value.


### GetIgnored

`func (o *GetQueryLogConfigResponse) GetIgnored() []string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *GetQueryLogConfigResponse) GetIgnoredOk() (*[]string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *GetQueryLogConfigResponse) SetIgnored(v []string)`

SetIgnored sets Ignored field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


