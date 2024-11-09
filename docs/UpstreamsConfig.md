# UpstreamsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapDns** | **[]string** | Bootstrap DNS servers, port is optional after colon.  | 
**UpstreamDns** | **[]string** | Upstream DNS servers, port is optional after colon.  | 
**FallbackDns** | Pointer to **[]string** | Fallback DNS servers, port is optional after colon.  | [optional] 
**PrivateUpstream** | Pointer to **[]string** | Local PTR resolvers, port is optional after colon.  | [optional] 

## Methods

### NewUpstreamsConfig

`func NewUpstreamsConfig(bootstrapDns []string, upstreamDns []string, ) *UpstreamsConfig`

NewUpstreamsConfig instantiates a new UpstreamsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamsConfigWithDefaults

`func NewUpstreamsConfigWithDefaults() *UpstreamsConfig`

NewUpstreamsConfigWithDefaults instantiates a new UpstreamsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapDns

`func (o *UpstreamsConfig) GetBootstrapDns() []string`

GetBootstrapDns returns the BootstrapDns field if non-nil, zero value otherwise.

### GetBootstrapDnsOk

`func (o *UpstreamsConfig) GetBootstrapDnsOk() (*[]string, bool)`

GetBootstrapDnsOk returns a tuple with the BootstrapDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapDns

`func (o *UpstreamsConfig) SetBootstrapDns(v []string)`

SetBootstrapDns sets BootstrapDns field to given value.


### GetUpstreamDns

`func (o *UpstreamsConfig) GetUpstreamDns() []string`

GetUpstreamDns returns the UpstreamDns field if non-nil, zero value otherwise.

### GetUpstreamDnsOk

`func (o *UpstreamsConfig) GetUpstreamDnsOk() (*[]string, bool)`

GetUpstreamDnsOk returns a tuple with the UpstreamDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamDns

`func (o *UpstreamsConfig) SetUpstreamDns(v []string)`

SetUpstreamDns sets UpstreamDns field to given value.


### GetFallbackDns

`func (o *UpstreamsConfig) GetFallbackDns() []string`

GetFallbackDns returns the FallbackDns field if non-nil, zero value otherwise.

### GetFallbackDnsOk

`func (o *UpstreamsConfig) GetFallbackDnsOk() (*[]string, bool)`

GetFallbackDnsOk returns a tuple with the FallbackDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackDns

`func (o *UpstreamsConfig) SetFallbackDns(v []string)`

SetFallbackDns sets FallbackDns field to given value.

### HasFallbackDns

`func (o *UpstreamsConfig) HasFallbackDns() bool`

HasFallbackDns returns a boolean if a field has been set.

### GetPrivateUpstream

`func (o *UpstreamsConfig) GetPrivateUpstream() []string`

GetPrivateUpstream returns the PrivateUpstream field if non-nil, zero value otherwise.

### GetPrivateUpstreamOk

`func (o *UpstreamsConfig) GetPrivateUpstreamOk() (*[]string, bool)`

GetPrivateUpstreamOk returns a tuple with the PrivateUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateUpstream

`func (o *UpstreamsConfig) SetPrivateUpstream(v []string)`

SetPrivateUpstream sets PrivateUpstream field to given value.

### HasPrivateUpstream

`func (o *UpstreamsConfig) HasPrivateUpstream() bool`

HasPrivateUpstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


