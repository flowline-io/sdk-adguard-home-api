# DNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapDns** | Pointer to **[]string** | Bootstrap servers, port is optional after colon.  Empty value will reset it to default values.  | [optional] 
**UpstreamDns** | Pointer to **[]string** | Upstream servers, port is optional after colon.  Empty value will reset it to default values.  | [optional] 
**FallbackDns** | Pointer to **[]string** | List of fallback DNS servers used when upstream DNS servers are not responding.  Empty value will clear the list.  | [optional] 
**UpstreamDnsFile** | Pointer to **string** |  | [optional] 
**ProtectionEnabled** | Pointer to **bool** |  | [optional] 
**Ratelimit** | Pointer to **int32** |  | [optional] 
**RatelimitSubnetSubnetLenIpv4** | Pointer to **int32** | Length of the subnet mask for IPv4 addresses. | [optional] [default to 24]
**RatelimitSubnetSubnetLenIpv6** | Pointer to **int32** | Length of the subnet mask for IPv6 addresses. | [optional] [default to 56]
**RatelimitWhitelist** | Pointer to **[]string** | List of IP addresses excluded from rate limiting. | [optional] 
**BlockingMode** | Pointer to **string** |  | [optional] 
**BlockingIpv4** | Pointer to **string** |  | [optional] 
**BlockingIpv6** | Pointer to **string** |  | [optional] 
**BlockedResponseTtl** | Pointer to **int32** | TTL for blocked responses. | [optional] 
**ProtectionDisabledUntil** | Pointer to **string** | Protection is pause until this time.  Nullable. | [optional] 
**EdnsCsEnabled** | Pointer to **bool** |  | [optional] 
**EdnsCsUseCustom** | Pointer to **bool** |  | [optional] 
**EdnsCsCustomIp** | Pointer to **string** |  | [optional] 
**DisableIpv6** | Pointer to **bool** |  | [optional] 
**DnssecEnabled** | Pointer to **bool** |  | [optional] 
**CacheSize** | Pointer to **int32** |  | [optional] 
**CacheTtlMin** | Pointer to **int32** |  | [optional] 
**CacheTtlMax** | Pointer to **int32** |  | [optional] 
**CacheOptimistic** | Pointer to **bool** |  | [optional] 
**UpstreamMode** | Pointer to **string** | Upstream modes enumeration. | [optional] 
**UsePrivatePtrResolvers** | Pointer to **bool** |  | [optional] 
**ResolveClients** | Pointer to **bool** |  | [optional] 
**LocalPtrUpstreams** | Pointer to **[]string** | Upstream servers, port is optional after colon.  Empty value will reset it to default values.  | [optional] 

## Methods

### NewDNSConfig

`func NewDNSConfig() *DNSConfig`

NewDNSConfig instantiates a new DNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSConfigWithDefaults

`func NewDNSConfigWithDefaults() *DNSConfig`

NewDNSConfigWithDefaults instantiates a new DNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapDns

`func (o *DNSConfig) GetBootstrapDns() []string`

GetBootstrapDns returns the BootstrapDns field if non-nil, zero value otherwise.

### GetBootstrapDnsOk

`func (o *DNSConfig) GetBootstrapDnsOk() (*[]string, bool)`

GetBootstrapDnsOk returns a tuple with the BootstrapDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapDns

`func (o *DNSConfig) SetBootstrapDns(v []string)`

SetBootstrapDns sets BootstrapDns field to given value.

### HasBootstrapDns

`func (o *DNSConfig) HasBootstrapDns() bool`

HasBootstrapDns returns a boolean if a field has been set.

### GetUpstreamDns

`func (o *DNSConfig) GetUpstreamDns() []string`

GetUpstreamDns returns the UpstreamDns field if non-nil, zero value otherwise.

### GetUpstreamDnsOk

`func (o *DNSConfig) GetUpstreamDnsOk() (*[]string, bool)`

GetUpstreamDnsOk returns a tuple with the UpstreamDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamDns

`func (o *DNSConfig) SetUpstreamDns(v []string)`

SetUpstreamDns sets UpstreamDns field to given value.

### HasUpstreamDns

`func (o *DNSConfig) HasUpstreamDns() bool`

HasUpstreamDns returns a boolean if a field has been set.

### GetFallbackDns

`func (o *DNSConfig) GetFallbackDns() []string`

GetFallbackDns returns the FallbackDns field if non-nil, zero value otherwise.

### GetFallbackDnsOk

`func (o *DNSConfig) GetFallbackDnsOk() (*[]string, bool)`

GetFallbackDnsOk returns a tuple with the FallbackDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackDns

`func (o *DNSConfig) SetFallbackDns(v []string)`

SetFallbackDns sets FallbackDns field to given value.

### HasFallbackDns

`func (o *DNSConfig) HasFallbackDns() bool`

HasFallbackDns returns a boolean if a field has been set.

### GetUpstreamDnsFile

`func (o *DNSConfig) GetUpstreamDnsFile() string`

GetUpstreamDnsFile returns the UpstreamDnsFile field if non-nil, zero value otherwise.

### GetUpstreamDnsFileOk

`func (o *DNSConfig) GetUpstreamDnsFileOk() (*string, bool)`

GetUpstreamDnsFileOk returns a tuple with the UpstreamDnsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamDnsFile

`func (o *DNSConfig) SetUpstreamDnsFile(v string)`

SetUpstreamDnsFile sets UpstreamDnsFile field to given value.

### HasUpstreamDnsFile

`func (o *DNSConfig) HasUpstreamDnsFile() bool`

HasUpstreamDnsFile returns a boolean if a field has been set.

### GetProtectionEnabled

`func (o *DNSConfig) GetProtectionEnabled() bool`

GetProtectionEnabled returns the ProtectionEnabled field if non-nil, zero value otherwise.

### GetProtectionEnabledOk

`func (o *DNSConfig) GetProtectionEnabledOk() (*bool, bool)`

GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnabled

`func (o *DNSConfig) SetProtectionEnabled(v bool)`

SetProtectionEnabled sets ProtectionEnabled field to given value.

### HasProtectionEnabled

`func (o *DNSConfig) HasProtectionEnabled() bool`

HasProtectionEnabled returns a boolean if a field has been set.

### GetRatelimit

`func (o *DNSConfig) GetRatelimit() int32`

GetRatelimit returns the Ratelimit field if non-nil, zero value otherwise.

### GetRatelimitOk

`func (o *DNSConfig) GetRatelimitOk() (*int32, bool)`

GetRatelimitOk returns a tuple with the Ratelimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimit

`func (o *DNSConfig) SetRatelimit(v int32)`

SetRatelimit sets Ratelimit field to given value.

### HasRatelimit

`func (o *DNSConfig) HasRatelimit() bool`

HasRatelimit returns a boolean if a field has been set.

### GetRatelimitSubnetSubnetLenIpv4

`func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv4() int32`

GetRatelimitSubnetSubnetLenIpv4 returns the RatelimitSubnetSubnetLenIpv4 field if non-nil, zero value otherwise.

### GetRatelimitSubnetSubnetLenIpv4Ok

`func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv4Ok() (*int32, bool)`

GetRatelimitSubnetSubnetLenIpv4Ok returns a tuple with the RatelimitSubnetSubnetLenIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimitSubnetSubnetLenIpv4

`func (o *DNSConfig) SetRatelimitSubnetSubnetLenIpv4(v int32)`

SetRatelimitSubnetSubnetLenIpv4 sets RatelimitSubnetSubnetLenIpv4 field to given value.

### HasRatelimitSubnetSubnetLenIpv4

`func (o *DNSConfig) HasRatelimitSubnetSubnetLenIpv4() bool`

HasRatelimitSubnetSubnetLenIpv4 returns a boolean if a field has been set.

### GetRatelimitSubnetSubnetLenIpv6

`func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv6() int32`

GetRatelimitSubnetSubnetLenIpv6 returns the RatelimitSubnetSubnetLenIpv6 field if non-nil, zero value otherwise.

### GetRatelimitSubnetSubnetLenIpv6Ok

`func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv6Ok() (*int32, bool)`

GetRatelimitSubnetSubnetLenIpv6Ok returns a tuple with the RatelimitSubnetSubnetLenIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimitSubnetSubnetLenIpv6

`func (o *DNSConfig) SetRatelimitSubnetSubnetLenIpv6(v int32)`

SetRatelimitSubnetSubnetLenIpv6 sets RatelimitSubnetSubnetLenIpv6 field to given value.

### HasRatelimitSubnetSubnetLenIpv6

`func (o *DNSConfig) HasRatelimitSubnetSubnetLenIpv6() bool`

HasRatelimitSubnetSubnetLenIpv6 returns a boolean if a field has been set.

### GetRatelimitWhitelist

`func (o *DNSConfig) GetRatelimitWhitelist() []string`

GetRatelimitWhitelist returns the RatelimitWhitelist field if non-nil, zero value otherwise.

### GetRatelimitWhitelistOk

`func (o *DNSConfig) GetRatelimitWhitelistOk() (*[]string, bool)`

GetRatelimitWhitelistOk returns a tuple with the RatelimitWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimitWhitelist

`func (o *DNSConfig) SetRatelimitWhitelist(v []string)`

SetRatelimitWhitelist sets RatelimitWhitelist field to given value.

### HasRatelimitWhitelist

`func (o *DNSConfig) HasRatelimitWhitelist() bool`

HasRatelimitWhitelist returns a boolean if a field has been set.

### GetBlockingMode

`func (o *DNSConfig) GetBlockingMode() string`

GetBlockingMode returns the BlockingMode field if non-nil, zero value otherwise.

### GetBlockingModeOk

`func (o *DNSConfig) GetBlockingModeOk() (*string, bool)`

GetBlockingModeOk returns a tuple with the BlockingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingMode

`func (o *DNSConfig) SetBlockingMode(v string)`

SetBlockingMode sets BlockingMode field to given value.

### HasBlockingMode

`func (o *DNSConfig) HasBlockingMode() bool`

HasBlockingMode returns a boolean if a field has been set.

### GetBlockingIpv4

`func (o *DNSConfig) GetBlockingIpv4() string`

GetBlockingIpv4 returns the BlockingIpv4 field if non-nil, zero value otherwise.

### GetBlockingIpv4Ok

`func (o *DNSConfig) GetBlockingIpv4Ok() (*string, bool)`

GetBlockingIpv4Ok returns a tuple with the BlockingIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv4

`func (o *DNSConfig) SetBlockingIpv4(v string)`

SetBlockingIpv4 sets BlockingIpv4 field to given value.

### HasBlockingIpv4

`func (o *DNSConfig) HasBlockingIpv4() bool`

HasBlockingIpv4 returns a boolean if a field has been set.

### GetBlockingIpv6

`func (o *DNSConfig) GetBlockingIpv6() string`

GetBlockingIpv6 returns the BlockingIpv6 field if non-nil, zero value otherwise.

### GetBlockingIpv6Ok

`func (o *DNSConfig) GetBlockingIpv6Ok() (*string, bool)`

GetBlockingIpv6Ok returns a tuple with the BlockingIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv6

`func (o *DNSConfig) SetBlockingIpv6(v string)`

SetBlockingIpv6 sets BlockingIpv6 field to given value.

### HasBlockingIpv6

`func (o *DNSConfig) HasBlockingIpv6() bool`

HasBlockingIpv6 returns a boolean if a field has been set.

### GetBlockedResponseTtl

`func (o *DNSConfig) GetBlockedResponseTtl() int32`

GetBlockedResponseTtl returns the BlockedResponseTtl field if non-nil, zero value otherwise.

### GetBlockedResponseTtlOk

`func (o *DNSConfig) GetBlockedResponseTtlOk() (*int32, bool)`

GetBlockedResponseTtlOk returns a tuple with the BlockedResponseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedResponseTtl

`func (o *DNSConfig) SetBlockedResponseTtl(v int32)`

SetBlockedResponseTtl sets BlockedResponseTtl field to given value.

### HasBlockedResponseTtl

`func (o *DNSConfig) HasBlockedResponseTtl() bool`

HasBlockedResponseTtl returns a boolean if a field has been set.

### GetProtectionDisabledUntil

`func (o *DNSConfig) GetProtectionDisabledUntil() string`

GetProtectionDisabledUntil returns the ProtectionDisabledUntil field if non-nil, zero value otherwise.

### GetProtectionDisabledUntilOk

`func (o *DNSConfig) GetProtectionDisabledUntilOk() (*string, bool)`

GetProtectionDisabledUntilOk returns a tuple with the ProtectionDisabledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDisabledUntil

`func (o *DNSConfig) SetProtectionDisabledUntil(v string)`

SetProtectionDisabledUntil sets ProtectionDisabledUntil field to given value.

### HasProtectionDisabledUntil

`func (o *DNSConfig) HasProtectionDisabledUntil() bool`

HasProtectionDisabledUntil returns a boolean if a field has been set.

### GetEdnsCsEnabled

`func (o *DNSConfig) GetEdnsCsEnabled() bool`

GetEdnsCsEnabled returns the EdnsCsEnabled field if non-nil, zero value otherwise.

### GetEdnsCsEnabledOk

`func (o *DNSConfig) GetEdnsCsEnabledOk() (*bool, bool)`

GetEdnsCsEnabledOk returns a tuple with the EdnsCsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsCsEnabled

`func (o *DNSConfig) SetEdnsCsEnabled(v bool)`

SetEdnsCsEnabled sets EdnsCsEnabled field to given value.

### HasEdnsCsEnabled

`func (o *DNSConfig) HasEdnsCsEnabled() bool`

HasEdnsCsEnabled returns a boolean if a field has been set.

### GetEdnsCsUseCustom

`func (o *DNSConfig) GetEdnsCsUseCustom() bool`

GetEdnsCsUseCustom returns the EdnsCsUseCustom field if non-nil, zero value otherwise.

### GetEdnsCsUseCustomOk

`func (o *DNSConfig) GetEdnsCsUseCustomOk() (*bool, bool)`

GetEdnsCsUseCustomOk returns a tuple with the EdnsCsUseCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsCsUseCustom

`func (o *DNSConfig) SetEdnsCsUseCustom(v bool)`

SetEdnsCsUseCustom sets EdnsCsUseCustom field to given value.

### HasEdnsCsUseCustom

`func (o *DNSConfig) HasEdnsCsUseCustom() bool`

HasEdnsCsUseCustom returns a boolean if a field has been set.

### GetEdnsCsCustomIp

`func (o *DNSConfig) GetEdnsCsCustomIp() string`

GetEdnsCsCustomIp returns the EdnsCsCustomIp field if non-nil, zero value otherwise.

### GetEdnsCsCustomIpOk

`func (o *DNSConfig) GetEdnsCsCustomIpOk() (*string, bool)`

GetEdnsCsCustomIpOk returns a tuple with the EdnsCsCustomIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsCsCustomIp

`func (o *DNSConfig) SetEdnsCsCustomIp(v string)`

SetEdnsCsCustomIp sets EdnsCsCustomIp field to given value.

### HasEdnsCsCustomIp

`func (o *DNSConfig) HasEdnsCsCustomIp() bool`

HasEdnsCsCustomIp returns a boolean if a field has been set.

### GetDisableIpv6

`func (o *DNSConfig) GetDisableIpv6() bool`

GetDisableIpv6 returns the DisableIpv6 field if non-nil, zero value otherwise.

### GetDisableIpv6Ok

`func (o *DNSConfig) GetDisableIpv6Ok() (*bool, bool)`

GetDisableIpv6Ok returns a tuple with the DisableIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIpv6

`func (o *DNSConfig) SetDisableIpv6(v bool)`

SetDisableIpv6 sets DisableIpv6 field to given value.

### HasDisableIpv6

`func (o *DNSConfig) HasDisableIpv6() bool`

HasDisableIpv6 returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *DNSConfig) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *DNSConfig) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *DNSConfig) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *DNSConfig) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetCacheSize

`func (o *DNSConfig) GetCacheSize() int32`

GetCacheSize returns the CacheSize field if non-nil, zero value otherwise.

### GetCacheSizeOk

`func (o *DNSConfig) GetCacheSizeOk() (*int32, bool)`

GetCacheSizeOk returns a tuple with the CacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSize

`func (o *DNSConfig) SetCacheSize(v int32)`

SetCacheSize sets CacheSize field to given value.

### HasCacheSize

`func (o *DNSConfig) HasCacheSize() bool`

HasCacheSize returns a boolean if a field has been set.

### GetCacheTtlMin

`func (o *DNSConfig) GetCacheTtlMin() int32`

GetCacheTtlMin returns the CacheTtlMin field if non-nil, zero value otherwise.

### GetCacheTtlMinOk

`func (o *DNSConfig) GetCacheTtlMinOk() (*int32, bool)`

GetCacheTtlMinOk returns a tuple with the CacheTtlMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtlMin

`func (o *DNSConfig) SetCacheTtlMin(v int32)`

SetCacheTtlMin sets CacheTtlMin field to given value.

### HasCacheTtlMin

`func (o *DNSConfig) HasCacheTtlMin() bool`

HasCacheTtlMin returns a boolean if a field has been set.

### GetCacheTtlMax

`func (o *DNSConfig) GetCacheTtlMax() int32`

GetCacheTtlMax returns the CacheTtlMax field if non-nil, zero value otherwise.

### GetCacheTtlMaxOk

`func (o *DNSConfig) GetCacheTtlMaxOk() (*int32, bool)`

GetCacheTtlMaxOk returns a tuple with the CacheTtlMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtlMax

`func (o *DNSConfig) SetCacheTtlMax(v int32)`

SetCacheTtlMax sets CacheTtlMax field to given value.

### HasCacheTtlMax

`func (o *DNSConfig) HasCacheTtlMax() bool`

HasCacheTtlMax returns a boolean if a field has been set.

### GetCacheOptimistic

`func (o *DNSConfig) GetCacheOptimistic() bool`

GetCacheOptimistic returns the CacheOptimistic field if non-nil, zero value otherwise.

### GetCacheOptimisticOk

`func (o *DNSConfig) GetCacheOptimisticOk() (*bool, bool)`

GetCacheOptimisticOk returns a tuple with the CacheOptimistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOptimistic

`func (o *DNSConfig) SetCacheOptimistic(v bool)`

SetCacheOptimistic sets CacheOptimistic field to given value.

### HasCacheOptimistic

`func (o *DNSConfig) HasCacheOptimistic() bool`

HasCacheOptimistic returns a boolean if a field has been set.

### GetUpstreamMode

`func (o *DNSConfig) GetUpstreamMode() string`

GetUpstreamMode returns the UpstreamMode field if non-nil, zero value otherwise.

### GetUpstreamModeOk

`func (o *DNSConfig) GetUpstreamModeOk() (*string, bool)`

GetUpstreamModeOk returns a tuple with the UpstreamMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamMode

`func (o *DNSConfig) SetUpstreamMode(v string)`

SetUpstreamMode sets UpstreamMode field to given value.

### HasUpstreamMode

`func (o *DNSConfig) HasUpstreamMode() bool`

HasUpstreamMode returns a boolean if a field has been set.

### GetUsePrivatePtrResolvers

`func (o *DNSConfig) GetUsePrivatePtrResolvers() bool`

GetUsePrivatePtrResolvers returns the UsePrivatePtrResolvers field if non-nil, zero value otherwise.

### GetUsePrivatePtrResolversOk

`func (o *DNSConfig) GetUsePrivatePtrResolversOk() (*bool, bool)`

GetUsePrivatePtrResolversOk returns a tuple with the UsePrivatePtrResolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePrivatePtrResolvers

`func (o *DNSConfig) SetUsePrivatePtrResolvers(v bool)`

SetUsePrivatePtrResolvers sets UsePrivatePtrResolvers field to given value.

### HasUsePrivatePtrResolvers

`func (o *DNSConfig) HasUsePrivatePtrResolvers() bool`

HasUsePrivatePtrResolvers returns a boolean if a field has been set.

### GetResolveClients

`func (o *DNSConfig) GetResolveClients() bool`

GetResolveClients returns the ResolveClients field if non-nil, zero value otherwise.

### GetResolveClientsOk

`func (o *DNSConfig) GetResolveClientsOk() (*bool, bool)`

GetResolveClientsOk returns a tuple with the ResolveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveClients

`func (o *DNSConfig) SetResolveClients(v bool)`

SetResolveClients sets ResolveClients field to given value.

### HasResolveClients

`func (o *DNSConfig) HasResolveClients() bool`

HasResolveClients returns a boolean if a field has been set.

### GetLocalPtrUpstreams

`func (o *DNSConfig) GetLocalPtrUpstreams() []string`

GetLocalPtrUpstreams returns the LocalPtrUpstreams field if non-nil, zero value otherwise.

### GetLocalPtrUpstreamsOk

`func (o *DNSConfig) GetLocalPtrUpstreamsOk() (*[]string, bool)`

GetLocalPtrUpstreamsOk returns a tuple with the LocalPtrUpstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPtrUpstreams

`func (o *DNSConfig) SetLocalPtrUpstreams(v []string)`

SetLocalPtrUpstreams sets LocalPtrUpstreams field to given value.

### HasLocalPtrUpstreams

`func (o *DNSConfig) HasLocalPtrUpstreams() bool`

HasLocalPtrUpstreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


