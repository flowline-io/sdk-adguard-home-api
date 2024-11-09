# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Ids** | Pointer to **[]string** | IP, CIDR, MAC, or ClientID. | [optional] 
**UseGlobalSettings** | Pointer to **bool** |  | [optional] 
**FilteringEnabled** | Pointer to **bool** |  | [optional] 
**ParentalEnabled** | Pointer to **bool** |  | [optional] 
**SafebrowsingEnabled** | Pointer to **bool** |  | [optional] 
**SafesearchEnabled** | Pointer to **bool** |  | [optional] 
**SafeSearch** | Pointer to [**SafeSearchConfig**](SafeSearchConfig.md) |  | [optional] 
**UseGlobalBlockedServices** | Pointer to **bool** |  | [optional] 
**BlockedServicesSchedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**BlockedServices** | Pointer to **[]string** |  | [optional] 
**Upstreams** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IgnoreQuerylog** | Pointer to **bool** | NOTE: If &#x60;ignore_querylog&#x60; is not set in HTTP API &#x60;GET /clients/add&#x60; request then default value (false) will be used.  If &#x60;ignore_querylog&#x60; is not set in HTTP API &#x60;GET /clients/update&#x60; request then the existing value will not be changed.  This behaviour can be changed in the future versions.  | [optional] 
**IgnoreStatistics** | Pointer to **bool** | NOTE: If &#x60;ignore_statistics&#x60; is not set in HTTP API &#x60;GET /clients/add&#x60; request then default value (false) will be used.  If &#x60;ignore_statistics&#x60; is not set in HTTP API &#x60;GET /clients/update&#x60; request then the existing value will not be changed.  This behaviour can be changed in the future versions.  | [optional] 
**UpstreamsCacheEnabled** | Pointer to **bool** | NOTE: If &#x60;upstreams_cache_enabled&#x60; is not set in HTTP API &#x60;GET /clients/add&#x60; request then default value (false) will be used.  If &#x60;upstreams_cache_enabled&#x60; is not set in HTTP API &#x60;GET /clients/update&#x60; request then the existing value will not be changed.  This behaviour can be changed in the future versions.  | [optional] 
**UpstreamsCacheSize** | Pointer to **int32** | NOTE: If &#x60;upstreams_cache_enabled&#x60; is not set in HTTP API &#x60;GET /clients/update&#x60; request then the existing value will not be changed.  This behaviour can be changed in the future versions.  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Client) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Client) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Client) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Client) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIds

`func (o *Client) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *Client) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *Client) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *Client) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetUseGlobalSettings

`func (o *Client) GetUseGlobalSettings() bool`

GetUseGlobalSettings returns the UseGlobalSettings field if non-nil, zero value otherwise.

### GetUseGlobalSettingsOk

`func (o *Client) GetUseGlobalSettingsOk() (*bool, bool)`

GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalSettings

`func (o *Client) SetUseGlobalSettings(v bool)`

SetUseGlobalSettings sets UseGlobalSettings field to given value.

### HasUseGlobalSettings

`func (o *Client) HasUseGlobalSettings() bool`

HasUseGlobalSettings returns a boolean if a field has been set.

### GetFilteringEnabled

`func (o *Client) GetFilteringEnabled() bool`

GetFilteringEnabled returns the FilteringEnabled field if non-nil, zero value otherwise.

### GetFilteringEnabledOk

`func (o *Client) GetFilteringEnabledOk() (*bool, bool)`

GetFilteringEnabledOk returns a tuple with the FilteringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringEnabled

`func (o *Client) SetFilteringEnabled(v bool)`

SetFilteringEnabled sets FilteringEnabled field to given value.

### HasFilteringEnabled

`func (o *Client) HasFilteringEnabled() bool`

HasFilteringEnabled returns a boolean if a field has been set.

### GetParentalEnabled

`func (o *Client) GetParentalEnabled() bool`

GetParentalEnabled returns the ParentalEnabled field if non-nil, zero value otherwise.

### GetParentalEnabledOk

`func (o *Client) GetParentalEnabledOk() (*bool, bool)`

GetParentalEnabledOk returns a tuple with the ParentalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentalEnabled

`func (o *Client) SetParentalEnabled(v bool)`

SetParentalEnabled sets ParentalEnabled field to given value.

### HasParentalEnabled

`func (o *Client) HasParentalEnabled() bool`

HasParentalEnabled returns a boolean if a field has been set.

### GetSafebrowsingEnabled

`func (o *Client) GetSafebrowsingEnabled() bool`

GetSafebrowsingEnabled returns the SafebrowsingEnabled field if non-nil, zero value otherwise.

### GetSafebrowsingEnabledOk

`func (o *Client) GetSafebrowsingEnabledOk() (*bool, bool)`

GetSafebrowsingEnabledOk returns a tuple with the SafebrowsingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafebrowsingEnabled

`func (o *Client) SetSafebrowsingEnabled(v bool)`

SetSafebrowsingEnabled sets SafebrowsingEnabled field to given value.

### HasSafebrowsingEnabled

`func (o *Client) HasSafebrowsingEnabled() bool`

HasSafebrowsingEnabled returns a boolean if a field has been set.

### GetSafesearchEnabled

`func (o *Client) GetSafesearchEnabled() bool`

GetSafesearchEnabled returns the SafesearchEnabled field if non-nil, zero value otherwise.

### GetSafesearchEnabledOk

`func (o *Client) GetSafesearchEnabledOk() (*bool, bool)`

GetSafesearchEnabledOk returns a tuple with the SafesearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafesearchEnabled

`func (o *Client) SetSafesearchEnabled(v bool)`

SetSafesearchEnabled sets SafesearchEnabled field to given value.

### HasSafesearchEnabled

`func (o *Client) HasSafesearchEnabled() bool`

HasSafesearchEnabled returns a boolean if a field has been set.

### GetSafeSearch

`func (o *Client) GetSafeSearch() SafeSearchConfig`

GetSafeSearch returns the SafeSearch field if non-nil, zero value otherwise.

### GetSafeSearchOk

`func (o *Client) GetSafeSearchOk() (*SafeSearchConfig, bool)`

GetSafeSearchOk returns a tuple with the SafeSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeSearch

`func (o *Client) SetSafeSearch(v SafeSearchConfig)`

SetSafeSearch sets SafeSearch field to given value.

### HasSafeSearch

`func (o *Client) HasSafeSearch() bool`

HasSafeSearch returns a boolean if a field has been set.

### GetUseGlobalBlockedServices

`func (o *Client) GetUseGlobalBlockedServices() bool`

GetUseGlobalBlockedServices returns the UseGlobalBlockedServices field if non-nil, zero value otherwise.

### GetUseGlobalBlockedServicesOk

`func (o *Client) GetUseGlobalBlockedServicesOk() (*bool, bool)`

GetUseGlobalBlockedServicesOk returns a tuple with the UseGlobalBlockedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalBlockedServices

`func (o *Client) SetUseGlobalBlockedServices(v bool)`

SetUseGlobalBlockedServices sets UseGlobalBlockedServices field to given value.

### HasUseGlobalBlockedServices

`func (o *Client) HasUseGlobalBlockedServices() bool`

HasUseGlobalBlockedServices returns a boolean if a field has been set.

### GetBlockedServicesSchedule

`func (o *Client) GetBlockedServicesSchedule() Schedule`

GetBlockedServicesSchedule returns the BlockedServicesSchedule field if non-nil, zero value otherwise.

### GetBlockedServicesScheduleOk

`func (o *Client) GetBlockedServicesScheduleOk() (*Schedule, bool)`

GetBlockedServicesScheduleOk returns a tuple with the BlockedServicesSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServicesSchedule

`func (o *Client) SetBlockedServicesSchedule(v Schedule)`

SetBlockedServicesSchedule sets BlockedServicesSchedule field to given value.

### HasBlockedServicesSchedule

`func (o *Client) HasBlockedServicesSchedule() bool`

HasBlockedServicesSchedule returns a boolean if a field has been set.

### GetBlockedServices

`func (o *Client) GetBlockedServices() []string`

GetBlockedServices returns the BlockedServices field if non-nil, zero value otherwise.

### GetBlockedServicesOk

`func (o *Client) GetBlockedServicesOk() (*[]string, bool)`

GetBlockedServicesOk returns a tuple with the BlockedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServices

`func (o *Client) SetBlockedServices(v []string)`

SetBlockedServices sets BlockedServices field to given value.

### HasBlockedServices

`func (o *Client) HasBlockedServices() bool`

HasBlockedServices returns a boolean if a field has been set.

### GetUpstreams

`func (o *Client) GetUpstreams() []string`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *Client) GetUpstreamsOk() (*[]string, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *Client) SetUpstreams(v []string)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *Client) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetTags

`func (o *Client) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Client) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Client) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Client) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIgnoreQuerylog

`func (o *Client) GetIgnoreQuerylog() bool`

GetIgnoreQuerylog returns the IgnoreQuerylog field if non-nil, zero value otherwise.

### GetIgnoreQuerylogOk

`func (o *Client) GetIgnoreQuerylogOk() (*bool, bool)`

GetIgnoreQuerylogOk returns a tuple with the IgnoreQuerylog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreQuerylog

`func (o *Client) SetIgnoreQuerylog(v bool)`

SetIgnoreQuerylog sets IgnoreQuerylog field to given value.

### HasIgnoreQuerylog

`func (o *Client) HasIgnoreQuerylog() bool`

HasIgnoreQuerylog returns a boolean if a field has been set.

### GetIgnoreStatistics

`func (o *Client) GetIgnoreStatistics() bool`

GetIgnoreStatistics returns the IgnoreStatistics field if non-nil, zero value otherwise.

### GetIgnoreStatisticsOk

`func (o *Client) GetIgnoreStatisticsOk() (*bool, bool)`

GetIgnoreStatisticsOk returns a tuple with the IgnoreStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreStatistics

`func (o *Client) SetIgnoreStatistics(v bool)`

SetIgnoreStatistics sets IgnoreStatistics field to given value.

### HasIgnoreStatistics

`func (o *Client) HasIgnoreStatistics() bool`

HasIgnoreStatistics returns a boolean if a field has been set.

### GetUpstreamsCacheEnabled

`func (o *Client) GetUpstreamsCacheEnabled() bool`

GetUpstreamsCacheEnabled returns the UpstreamsCacheEnabled field if non-nil, zero value otherwise.

### GetUpstreamsCacheEnabledOk

`func (o *Client) GetUpstreamsCacheEnabledOk() (*bool, bool)`

GetUpstreamsCacheEnabledOk returns a tuple with the UpstreamsCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamsCacheEnabled

`func (o *Client) SetUpstreamsCacheEnabled(v bool)`

SetUpstreamsCacheEnabled sets UpstreamsCacheEnabled field to given value.

### HasUpstreamsCacheEnabled

`func (o *Client) HasUpstreamsCacheEnabled() bool`

HasUpstreamsCacheEnabled returns a boolean if a field has been set.

### GetUpstreamsCacheSize

`func (o *Client) GetUpstreamsCacheSize() int32`

GetUpstreamsCacheSize returns the UpstreamsCacheSize field if non-nil, zero value otherwise.

### GetUpstreamsCacheSizeOk

`func (o *Client) GetUpstreamsCacheSizeOk() (*int32, bool)`

GetUpstreamsCacheSizeOk returns a tuple with the UpstreamsCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamsCacheSize

`func (o *Client) SetUpstreamsCacheSize(v int32)`

SetUpstreamsCacheSize sets UpstreamsCacheSize field to given value.

### HasUpstreamsCacheSize

`func (o *Client) HasUpstreamsCacheSize() bool`

HasUpstreamsCacheSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


