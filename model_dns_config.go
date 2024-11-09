/*
AdGuard Home

AdGuard Home REST-ish API.  Our admin web interface is built on top of this REST-ish API. 

API version: 0.107
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DNSConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSConfig{}

// DNSConfig DNS server configuration
type DNSConfig struct {
	// Bootstrap servers, port is optional after colon.  Empty value will reset it to default values. 
	BootstrapDns []string `json:"bootstrap_dns,omitempty"`
	// Upstream servers, port is optional after colon.  Empty value will reset it to default values. 
	UpstreamDns []string `json:"upstream_dns,omitempty"`
	// List of fallback DNS servers used when upstream DNS servers are not responding.  Empty value will clear the list. 
	FallbackDns []string `json:"fallback_dns,omitempty"`
	UpstreamDnsFile *string `json:"upstream_dns_file,omitempty"`
	ProtectionEnabled *bool `json:"protection_enabled,omitempty"`
	Ratelimit *int32 `json:"ratelimit,omitempty"`
	// Length of the subnet mask for IPv4 addresses.
	RatelimitSubnetSubnetLenIpv4 *int32 `json:"ratelimit_subnet_subnet_len_ipv4,omitempty"`
	// Length of the subnet mask for IPv6 addresses.
	RatelimitSubnetSubnetLenIpv6 *int32 `json:"ratelimit_subnet_subnet_len_ipv6,omitempty"`
	// List of IP addresses excluded from rate limiting.
	RatelimitWhitelist []string `json:"ratelimit_whitelist,omitempty"`
	BlockingMode *string `json:"blocking_mode,omitempty"`
	BlockingIpv4 *string `json:"blocking_ipv4,omitempty"`
	BlockingIpv6 *string `json:"blocking_ipv6,omitempty"`
	// TTL for blocked responses.
	BlockedResponseTtl *int32 `json:"blocked_response_ttl,omitempty"`
	// Protection is pause until this time.  Nullable.
	ProtectionDisabledUntil *string `json:"protection_disabled_until,omitempty"`
	EdnsCsEnabled *bool `json:"edns_cs_enabled,omitempty"`
	EdnsCsUseCustom *bool `json:"edns_cs_use_custom,omitempty"`
	EdnsCsCustomIp *string `json:"edns_cs_custom_ip,omitempty"`
	DisableIpv6 *bool `json:"disable_ipv6,omitempty"`
	DnssecEnabled *bool `json:"dnssec_enabled,omitempty"`
	CacheSize *int32 `json:"cache_size,omitempty"`
	CacheTtlMin *int32 `json:"cache_ttl_min,omitempty"`
	CacheTtlMax *int32 `json:"cache_ttl_max,omitempty"`
	CacheOptimistic *bool `json:"cache_optimistic,omitempty"`
	// Upstream modes enumeration.
	UpstreamMode *string `json:"upstream_mode,omitempty"`
	UsePrivatePtrResolvers *bool `json:"use_private_ptr_resolvers,omitempty"`
	ResolveClients *bool `json:"resolve_clients,omitempty"`
	// Upstream servers, port is optional after colon.  Empty value will reset it to default values. 
	LocalPtrUpstreams []string `json:"local_ptr_upstreams,omitempty"`
}

// NewDNSConfig instantiates a new DNSConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSConfig() *DNSConfig {
	this := DNSConfig{}
	var ratelimitSubnetSubnetLenIpv4 int32 = 24
	this.RatelimitSubnetSubnetLenIpv4 = &ratelimitSubnetSubnetLenIpv4
	var ratelimitSubnetSubnetLenIpv6 int32 = 56
	this.RatelimitSubnetSubnetLenIpv6 = &ratelimitSubnetSubnetLenIpv6
	return &this
}

// NewDNSConfigWithDefaults instantiates a new DNSConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSConfigWithDefaults() *DNSConfig {
	this := DNSConfig{}
	var ratelimitSubnetSubnetLenIpv4 int32 = 24
	this.RatelimitSubnetSubnetLenIpv4 = &ratelimitSubnetSubnetLenIpv4
	var ratelimitSubnetSubnetLenIpv6 int32 = 56
	this.RatelimitSubnetSubnetLenIpv6 = &ratelimitSubnetSubnetLenIpv6
	return &this
}

// GetBootstrapDns returns the BootstrapDns field value if set, zero value otherwise.
func (o *DNSConfig) GetBootstrapDns() []string {
	if o == nil || IsNil(o.BootstrapDns) {
		var ret []string
		return ret
	}
	return o.BootstrapDns
}

// GetBootstrapDnsOk returns a tuple with the BootstrapDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetBootstrapDnsOk() ([]string, bool) {
	if o == nil || IsNil(o.BootstrapDns) {
		return nil, false
	}
	return o.BootstrapDns, true
}

// HasBootstrapDns returns a boolean if a field has been set.
func (o *DNSConfig) HasBootstrapDns() bool {
	if o != nil && !IsNil(o.BootstrapDns) {
		return true
	}

	return false
}

// SetBootstrapDns gets a reference to the given []string and assigns it to the BootstrapDns field.
func (o *DNSConfig) SetBootstrapDns(v []string) {
	o.BootstrapDns = v
}

// GetUpstreamDns returns the UpstreamDns field value if set, zero value otherwise.
func (o *DNSConfig) GetUpstreamDns() []string {
	if o == nil || IsNil(o.UpstreamDns) {
		var ret []string
		return ret
	}
	return o.UpstreamDns
}

// GetUpstreamDnsOk returns a tuple with the UpstreamDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetUpstreamDnsOk() ([]string, bool) {
	if o == nil || IsNil(o.UpstreamDns) {
		return nil, false
	}
	return o.UpstreamDns, true
}

// HasUpstreamDns returns a boolean if a field has been set.
func (o *DNSConfig) HasUpstreamDns() bool {
	if o != nil && !IsNil(o.UpstreamDns) {
		return true
	}

	return false
}

// SetUpstreamDns gets a reference to the given []string and assigns it to the UpstreamDns field.
func (o *DNSConfig) SetUpstreamDns(v []string) {
	o.UpstreamDns = v
}

// GetFallbackDns returns the FallbackDns field value if set, zero value otherwise.
func (o *DNSConfig) GetFallbackDns() []string {
	if o == nil || IsNil(o.FallbackDns) {
		var ret []string
		return ret
	}
	return o.FallbackDns
}

// GetFallbackDnsOk returns a tuple with the FallbackDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetFallbackDnsOk() ([]string, bool) {
	if o == nil || IsNil(o.FallbackDns) {
		return nil, false
	}
	return o.FallbackDns, true
}

// HasFallbackDns returns a boolean if a field has been set.
func (o *DNSConfig) HasFallbackDns() bool {
	if o != nil && !IsNil(o.FallbackDns) {
		return true
	}

	return false
}

// SetFallbackDns gets a reference to the given []string and assigns it to the FallbackDns field.
func (o *DNSConfig) SetFallbackDns(v []string) {
	o.FallbackDns = v
}

// GetUpstreamDnsFile returns the UpstreamDnsFile field value if set, zero value otherwise.
func (o *DNSConfig) GetUpstreamDnsFile() string {
	if o == nil || IsNil(o.UpstreamDnsFile) {
		var ret string
		return ret
	}
	return *o.UpstreamDnsFile
}

// GetUpstreamDnsFileOk returns a tuple with the UpstreamDnsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetUpstreamDnsFileOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamDnsFile) {
		return nil, false
	}
	return o.UpstreamDnsFile, true
}

// HasUpstreamDnsFile returns a boolean if a field has been set.
func (o *DNSConfig) HasUpstreamDnsFile() bool {
	if o != nil && !IsNil(o.UpstreamDnsFile) {
		return true
	}

	return false
}

// SetUpstreamDnsFile gets a reference to the given string and assigns it to the UpstreamDnsFile field.
func (o *DNSConfig) SetUpstreamDnsFile(v string) {
	o.UpstreamDnsFile = &v
}

// GetProtectionEnabled returns the ProtectionEnabled field value if set, zero value otherwise.
func (o *DNSConfig) GetProtectionEnabled() bool {
	if o == nil || IsNil(o.ProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.ProtectionEnabled
}

// GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectionEnabled) {
		return nil, false
	}
	return o.ProtectionEnabled, true
}

// HasProtectionEnabled returns a boolean if a field has been set.
func (o *DNSConfig) HasProtectionEnabled() bool {
	if o != nil && !IsNil(o.ProtectionEnabled) {
		return true
	}

	return false
}

// SetProtectionEnabled gets a reference to the given bool and assigns it to the ProtectionEnabled field.
func (o *DNSConfig) SetProtectionEnabled(v bool) {
	o.ProtectionEnabled = &v
}

// GetRatelimit returns the Ratelimit field value if set, zero value otherwise.
func (o *DNSConfig) GetRatelimit() int32 {
	if o == nil || IsNil(o.Ratelimit) {
		var ret int32
		return ret
	}
	return *o.Ratelimit
}

// GetRatelimitOk returns a tuple with the Ratelimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetRatelimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Ratelimit) {
		return nil, false
	}
	return o.Ratelimit, true
}

// HasRatelimit returns a boolean if a field has been set.
func (o *DNSConfig) HasRatelimit() bool {
	if o != nil && !IsNil(o.Ratelimit) {
		return true
	}

	return false
}

// SetRatelimit gets a reference to the given int32 and assigns it to the Ratelimit field.
func (o *DNSConfig) SetRatelimit(v int32) {
	o.Ratelimit = &v
}

// GetRatelimitSubnetSubnetLenIpv4 returns the RatelimitSubnetSubnetLenIpv4 field value if set, zero value otherwise.
func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv4() int32 {
	if o == nil || IsNil(o.RatelimitSubnetSubnetLenIpv4) {
		var ret int32
		return ret
	}
	return *o.RatelimitSubnetSubnetLenIpv4
}

// GetRatelimitSubnetSubnetLenIpv4Ok returns a tuple with the RatelimitSubnetSubnetLenIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv4Ok() (*int32, bool) {
	if o == nil || IsNil(o.RatelimitSubnetSubnetLenIpv4) {
		return nil, false
	}
	return o.RatelimitSubnetSubnetLenIpv4, true
}

// HasRatelimitSubnetSubnetLenIpv4 returns a boolean if a field has been set.
func (o *DNSConfig) HasRatelimitSubnetSubnetLenIpv4() bool {
	if o != nil && !IsNil(o.RatelimitSubnetSubnetLenIpv4) {
		return true
	}

	return false
}

// SetRatelimitSubnetSubnetLenIpv4 gets a reference to the given int32 and assigns it to the RatelimitSubnetSubnetLenIpv4 field.
func (o *DNSConfig) SetRatelimitSubnetSubnetLenIpv4(v int32) {
	o.RatelimitSubnetSubnetLenIpv4 = &v
}

// GetRatelimitSubnetSubnetLenIpv6 returns the RatelimitSubnetSubnetLenIpv6 field value if set, zero value otherwise.
func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv6() int32 {
	if o == nil || IsNil(o.RatelimitSubnetSubnetLenIpv6) {
		var ret int32
		return ret
	}
	return *o.RatelimitSubnetSubnetLenIpv6
}

// GetRatelimitSubnetSubnetLenIpv6Ok returns a tuple with the RatelimitSubnetSubnetLenIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetRatelimitSubnetSubnetLenIpv6Ok() (*int32, bool) {
	if o == nil || IsNil(o.RatelimitSubnetSubnetLenIpv6) {
		return nil, false
	}
	return o.RatelimitSubnetSubnetLenIpv6, true
}

// HasRatelimitSubnetSubnetLenIpv6 returns a boolean if a field has been set.
func (o *DNSConfig) HasRatelimitSubnetSubnetLenIpv6() bool {
	if o != nil && !IsNil(o.RatelimitSubnetSubnetLenIpv6) {
		return true
	}

	return false
}

// SetRatelimitSubnetSubnetLenIpv6 gets a reference to the given int32 and assigns it to the RatelimitSubnetSubnetLenIpv6 field.
func (o *DNSConfig) SetRatelimitSubnetSubnetLenIpv6(v int32) {
	o.RatelimitSubnetSubnetLenIpv6 = &v
}

// GetRatelimitWhitelist returns the RatelimitWhitelist field value if set, zero value otherwise.
func (o *DNSConfig) GetRatelimitWhitelist() []string {
	if o == nil || IsNil(o.RatelimitWhitelist) {
		var ret []string
		return ret
	}
	return o.RatelimitWhitelist
}

// GetRatelimitWhitelistOk returns a tuple with the RatelimitWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetRatelimitWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.RatelimitWhitelist) {
		return nil, false
	}
	return o.RatelimitWhitelist, true
}

// HasRatelimitWhitelist returns a boolean if a field has been set.
func (o *DNSConfig) HasRatelimitWhitelist() bool {
	if o != nil && !IsNil(o.RatelimitWhitelist) {
		return true
	}

	return false
}

// SetRatelimitWhitelist gets a reference to the given []string and assigns it to the RatelimitWhitelist field.
func (o *DNSConfig) SetRatelimitWhitelist(v []string) {
	o.RatelimitWhitelist = v
}

// GetBlockingMode returns the BlockingMode field value if set, zero value otherwise.
func (o *DNSConfig) GetBlockingMode() string {
	if o == nil || IsNil(o.BlockingMode) {
		var ret string
		return ret
	}
	return *o.BlockingMode
}

// GetBlockingModeOk returns a tuple with the BlockingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetBlockingModeOk() (*string, bool) {
	if o == nil || IsNil(o.BlockingMode) {
		return nil, false
	}
	return o.BlockingMode, true
}

// HasBlockingMode returns a boolean if a field has been set.
func (o *DNSConfig) HasBlockingMode() bool {
	if o != nil && !IsNil(o.BlockingMode) {
		return true
	}

	return false
}

// SetBlockingMode gets a reference to the given string and assigns it to the BlockingMode field.
func (o *DNSConfig) SetBlockingMode(v string) {
	o.BlockingMode = &v
}

// GetBlockingIpv4 returns the BlockingIpv4 field value if set, zero value otherwise.
func (o *DNSConfig) GetBlockingIpv4() string {
	if o == nil || IsNil(o.BlockingIpv4) {
		var ret string
		return ret
	}
	return *o.BlockingIpv4
}

// GetBlockingIpv4Ok returns a tuple with the BlockingIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetBlockingIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.BlockingIpv4) {
		return nil, false
	}
	return o.BlockingIpv4, true
}

// HasBlockingIpv4 returns a boolean if a field has been set.
func (o *DNSConfig) HasBlockingIpv4() bool {
	if o != nil && !IsNil(o.BlockingIpv4) {
		return true
	}

	return false
}

// SetBlockingIpv4 gets a reference to the given string and assigns it to the BlockingIpv4 field.
func (o *DNSConfig) SetBlockingIpv4(v string) {
	o.BlockingIpv4 = &v
}

// GetBlockingIpv6 returns the BlockingIpv6 field value if set, zero value otherwise.
func (o *DNSConfig) GetBlockingIpv6() string {
	if o == nil || IsNil(o.BlockingIpv6) {
		var ret string
		return ret
	}
	return *o.BlockingIpv6
}

// GetBlockingIpv6Ok returns a tuple with the BlockingIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetBlockingIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.BlockingIpv6) {
		return nil, false
	}
	return o.BlockingIpv6, true
}

// HasBlockingIpv6 returns a boolean if a field has been set.
func (o *DNSConfig) HasBlockingIpv6() bool {
	if o != nil && !IsNil(o.BlockingIpv6) {
		return true
	}

	return false
}

// SetBlockingIpv6 gets a reference to the given string and assigns it to the BlockingIpv6 field.
func (o *DNSConfig) SetBlockingIpv6(v string) {
	o.BlockingIpv6 = &v
}

// GetBlockedResponseTtl returns the BlockedResponseTtl field value if set, zero value otherwise.
func (o *DNSConfig) GetBlockedResponseTtl() int32 {
	if o == nil || IsNil(o.BlockedResponseTtl) {
		var ret int32
		return ret
	}
	return *o.BlockedResponseTtl
}

// GetBlockedResponseTtlOk returns a tuple with the BlockedResponseTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetBlockedResponseTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockedResponseTtl) {
		return nil, false
	}
	return o.BlockedResponseTtl, true
}

// HasBlockedResponseTtl returns a boolean if a field has been set.
func (o *DNSConfig) HasBlockedResponseTtl() bool {
	if o != nil && !IsNil(o.BlockedResponseTtl) {
		return true
	}

	return false
}

// SetBlockedResponseTtl gets a reference to the given int32 and assigns it to the BlockedResponseTtl field.
func (o *DNSConfig) SetBlockedResponseTtl(v int32) {
	o.BlockedResponseTtl = &v
}

// GetProtectionDisabledUntil returns the ProtectionDisabledUntil field value if set, zero value otherwise.
func (o *DNSConfig) GetProtectionDisabledUntil() string {
	if o == nil || IsNil(o.ProtectionDisabledUntil) {
		var ret string
		return ret
	}
	return *o.ProtectionDisabledUntil
}

// GetProtectionDisabledUntilOk returns a tuple with the ProtectionDisabledUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetProtectionDisabledUntilOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectionDisabledUntil) {
		return nil, false
	}
	return o.ProtectionDisabledUntil, true
}

// HasProtectionDisabledUntil returns a boolean if a field has been set.
func (o *DNSConfig) HasProtectionDisabledUntil() bool {
	if o != nil && !IsNil(o.ProtectionDisabledUntil) {
		return true
	}

	return false
}

// SetProtectionDisabledUntil gets a reference to the given string and assigns it to the ProtectionDisabledUntil field.
func (o *DNSConfig) SetProtectionDisabledUntil(v string) {
	o.ProtectionDisabledUntil = &v
}

// GetEdnsCsEnabled returns the EdnsCsEnabled field value if set, zero value otherwise.
func (o *DNSConfig) GetEdnsCsEnabled() bool {
	if o == nil || IsNil(o.EdnsCsEnabled) {
		var ret bool
		return ret
	}
	return *o.EdnsCsEnabled
}

// GetEdnsCsEnabledOk returns a tuple with the EdnsCsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetEdnsCsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EdnsCsEnabled) {
		return nil, false
	}
	return o.EdnsCsEnabled, true
}

// HasEdnsCsEnabled returns a boolean if a field has been set.
func (o *DNSConfig) HasEdnsCsEnabled() bool {
	if o != nil && !IsNil(o.EdnsCsEnabled) {
		return true
	}

	return false
}

// SetEdnsCsEnabled gets a reference to the given bool and assigns it to the EdnsCsEnabled field.
func (o *DNSConfig) SetEdnsCsEnabled(v bool) {
	o.EdnsCsEnabled = &v
}

// GetEdnsCsUseCustom returns the EdnsCsUseCustom field value if set, zero value otherwise.
func (o *DNSConfig) GetEdnsCsUseCustom() bool {
	if o == nil || IsNil(o.EdnsCsUseCustom) {
		var ret bool
		return ret
	}
	return *o.EdnsCsUseCustom
}

// GetEdnsCsUseCustomOk returns a tuple with the EdnsCsUseCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetEdnsCsUseCustomOk() (*bool, bool) {
	if o == nil || IsNil(o.EdnsCsUseCustom) {
		return nil, false
	}
	return o.EdnsCsUseCustom, true
}

// HasEdnsCsUseCustom returns a boolean if a field has been set.
func (o *DNSConfig) HasEdnsCsUseCustom() bool {
	if o != nil && !IsNil(o.EdnsCsUseCustom) {
		return true
	}

	return false
}

// SetEdnsCsUseCustom gets a reference to the given bool and assigns it to the EdnsCsUseCustom field.
func (o *DNSConfig) SetEdnsCsUseCustom(v bool) {
	o.EdnsCsUseCustom = &v
}

// GetEdnsCsCustomIp returns the EdnsCsCustomIp field value if set, zero value otherwise.
func (o *DNSConfig) GetEdnsCsCustomIp() string {
	if o == nil || IsNil(o.EdnsCsCustomIp) {
		var ret string
		return ret
	}
	return *o.EdnsCsCustomIp
}

// GetEdnsCsCustomIpOk returns a tuple with the EdnsCsCustomIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetEdnsCsCustomIpOk() (*string, bool) {
	if o == nil || IsNil(o.EdnsCsCustomIp) {
		return nil, false
	}
	return o.EdnsCsCustomIp, true
}

// HasEdnsCsCustomIp returns a boolean if a field has been set.
func (o *DNSConfig) HasEdnsCsCustomIp() bool {
	if o != nil && !IsNil(o.EdnsCsCustomIp) {
		return true
	}

	return false
}

// SetEdnsCsCustomIp gets a reference to the given string and assigns it to the EdnsCsCustomIp field.
func (o *DNSConfig) SetEdnsCsCustomIp(v string) {
	o.EdnsCsCustomIp = &v
}

// GetDisableIpv6 returns the DisableIpv6 field value if set, zero value otherwise.
func (o *DNSConfig) GetDisableIpv6() bool {
	if o == nil || IsNil(o.DisableIpv6) {
		var ret bool
		return ret
	}
	return *o.DisableIpv6
}

// GetDisableIpv6Ok returns a tuple with the DisableIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetDisableIpv6Ok() (*bool, bool) {
	if o == nil || IsNil(o.DisableIpv6) {
		return nil, false
	}
	return o.DisableIpv6, true
}

// HasDisableIpv6 returns a boolean if a field has been set.
func (o *DNSConfig) HasDisableIpv6() bool {
	if o != nil && !IsNil(o.DisableIpv6) {
		return true
	}

	return false
}

// SetDisableIpv6 gets a reference to the given bool and assigns it to the DisableIpv6 field.
func (o *DNSConfig) SetDisableIpv6(v bool) {
	o.DisableIpv6 = &v
}

// GetDnssecEnabled returns the DnssecEnabled field value if set, zero value otherwise.
func (o *DNSConfig) GetDnssecEnabled() bool {
	if o == nil || IsNil(o.DnssecEnabled) {
		var ret bool
		return ret
	}
	return *o.DnssecEnabled
}

// GetDnssecEnabledOk returns a tuple with the DnssecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetDnssecEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnssecEnabled) {
		return nil, false
	}
	return o.DnssecEnabled, true
}

// HasDnssecEnabled returns a boolean if a field has been set.
func (o *DNSConfig) HasDnssecEnabled() bool {
	if o != nil && !IsNil(o.DnssecEnabled) {
		return true
	}

	return false
}

// SetDnssecEnabled gets a reference to the given bool and assigns it to the DnssecEnabled field.
func (o *DNSConfig) SetDnssecEnabled(v bool) {
	o.DnssecEnabled = &v
}

// GetCacheSize returns the CacheSize field value if set, zero value otherwise.
func (o *DNSConfig) GetCacheSize() int32 {
	if o == nil || IsNil(o.CacheSize) {
		var ret int32
		return ret
	}
	return *o.CacheSize
}

// GetCacheSizeOk returns a tuple with the CacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetCacheSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheSize) {
		return nil, false
	}
	return o.CacheSize, true
}

// HasCacheSize returns a boolean if a field has been set.
func (o *DNSConfig) HasCacheSize() bool {
	if o != nil && !IsNil(o.CacheSize) {
		return true
	}

	return false
}

// SetCacheSize gets a reference to the given int32 and assigns it to the CacheSize field.
func (o *DNSConfig) SetCacheSize(v int32) {
	o.CacheSize = &v
}

// GetCacheTtlMin returns the CacheTtlMin field value if set, zero value otherwise.
func (o *DNSConfig) GetCacheTtlMin() int32 {
	if o == nil || IsNil(o.CacheTtlMin) {
		var ret int32
		return ret
	}
	return *o.CacheTtlMin
}

// GetCacheTtlMinOk returns a tuple with the CacheTtlMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetCacheTtlMinOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTtlMin) {
		return nil, false
	}
	return o.CacheTtlMin, true
}

// HasCacheTtlMin returns a boolean if a field has been set.
func (o *DNSConfig) HasCacheTtlMin() bool {
	if o != nil && !IsNil(o.CacheTtlMin) {
		return true
	}

	return false
}

// SetCacheTtlMin gets a reference to the given int32 and assigns it to the CacheTtlMin field.
func (o *DNSConfig) SetCacheTtlMin(v int32) {
	o.CacheTtlMin = &v
}

// GetCacheTtlMax returns the CacheTtlMax field value if set, zero value otherwise.
func (o *DNSConfig) GetCacheTtlMax() int32 {
	if o == nil || IsNil(o.CacheTtlMax) {
		var ret int32
		return ret
	}
	return *o.CacheTtlMax
}

// GetCacheTtlMaxOk returns a tuple with the CacheTtlMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetCacheTtlMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTtlMax) {
		return nil, false
	}
	return o.CacheTtlMax, true
}

// HasCacheTtlMax returns a boolean if a field has been set.
func (o *DNSConfig) HasCacheTtlMax() bool {
	if o != nil && !IsNil(o.CacheTtlMax) {
		return true
	}

	return false
}

// SetCacheTtlMax gets a reference to the given int32 and assigns it to the CacheTtlMax field.
func (o *DNSConfig) SetCacheTtlMax(v int32) {
	o.CacheTtlMax = &v
}

// GetCacheOptimistic returns the CacheOptimistic field value if set, zero value otherwise.
func (o *DNSConfig) GetCacheOptimistic() bool {
	if o == nil || IsNil(o.CacheOptimistic) {
		var ret bool
		return ret
	}
	return *o.CacheOptimistic
}

// GetCacheOptimisticOk returns a tuple with the CacheOptimistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetCacheOptimisticOk() (*bool, bool) {
	if o == nil || IsNil(o.CacheOptimistic) {
		return nil, false
	}
	return o.CacheOptimistic, true
}

// HasCacheOptimistic returns a boolean if a field has been set.
func (o *DNSConfig) HasCacheOptimistic() bool {
	if o != nil && !IsNil(o.CacheOptimistic) {
		return true
	}

	return false
}

// SetCacheOptimistic gets a reference to the given bool and assigns it to the CacheOptimistic field.
func (o *DNSConfig) SetCacheOptimistic(v bool) {
	o.CacheOptimistic = &v
}

// GetUpstreamMode returns the UpstreamMode field value if set, zero value otherwise.
func (o *DNSConfig) GetUpstreamMode() string {
	if o == nil || IsNil(o.UpstreamMode) {
		var ret string
		return ret
	}
	return *o.UpstreamMode
}

// GetUpstreamModeOk returns a tuple with the UpstreamMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetUpstreamModeOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamMode) {
		return nil, false
	}
	return o.UpstreamMode, true
}

// HasUpstreamMode returns a boolean if a field has been set.
func (o *DNSConfig) HasUpstreamMode() bool {
	if o != nil && !IsNil(o.UpstreamMode) {
		return true
	}

	return false
}

// SetUpstreamMode gets a reference to the given string and assigns it to the UpstreamMode field.
func (o *DNSConfig) SetUpstreamMode(v string) {
	o.UpstreamMode = &v
}

// GetUsePrivatePtrResolvers returns the UsePrivatePtrResolvers field value if set, zero value otherwise.
func (o *DNSConfig) GetUsePrivatePtrResolvers() bool {
	if o == nil || IsNil(o.UsePrivatePtrResolvers) {
		var ret bool
		return ret
	}
	return *o.UsePrivatePtrResolvers
}

// GetUsePrivatePtrResolversOk returns a tuple with the UsePrivatePtrResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetUsePrivatePtrResolversOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePrivatePtrResolvers) {
		return nil, false
	}
	return o.UsePrivatePtrResolvers, true
}

// HasUsePrivatePtrResolvers returns a boolean if a field has been set.
func (o *DNSConfig) HasUsePrivatePtrResolvers() bool {
	if o != nil && !IsNil(o.UsePrivatePtrResolvers) {
		return true
	}

	return false
}

// SetUsePrivatePtrResolvers gets a reference to the given bool and assigns it to the UsePrivatePtrResolvers field.
func (o *DNSConfig) SetUsePrivatePtrResolvers(v bool) {
	o.UsePrivatePtrResolvers = &v
}

// GetResolveClients returns the ResolveClients field value if set, zero value otherwise.
func (o *DNSConfig) GetResolveClients() bool {
	if o == nil || IsNil(o.ResolveClients) {
		var ret bool
		return ret
	}
	return *o.ResolveClients
}

// GetResolveClientsOk returns a tuple with the ResolveClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetResolveClientsOk() (*bool, bool) {
	if o == nil || IsNil(o.ResolveClients) {
		return nil, false
	}
	return o.ResolveClients, true
}

// HasResolveClients returns a boolean if a field has been set.
func (o *DNSConfig) HasResolveClients() bool {
	if o != nil && !IsNil(o.ResolveClients) {
		return true
	}

	return false
}

// SetResolveClients gets a reference to the given bool and assigns it to the ResolveClients field.
func (o *DNSConfig) SetResolveClients(v bool) {
	o.ResolveClients = &v
}

// GetLocalPtrUpstreams returns the LocalPtrUpstreams field value if set, zero value otherwise.
func (o *DNSConfig) GetLocalPtrUpstreams() []string {
	if o == nil || IsNil(o.LocalPtrUpstreams) {
		var ret []string
		return ret
	}
	return o.LocalPtrUpstreams
}

// GetLocalPtrUpstreamsOk returns a tuple with the LocalPtrUpstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSConfig) GetLocalPtrUpstreamsOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalPtrUpstreams) {
		return nil, false
	}
	return o.LocalPtrUpstreams, true
}

// HasLocalPtrUpstreams returns a boolean if a field has been set.
func (o *DNSConfig) HasLocalPtrUpstreams() bool {
	if o != nil && !IsNil(o.LocalPtrUpstreams) {
		return true
	}

	return false
}

// SetLocalPtrUpstreams gets a reference to the given []string and assigns it to the LocalPtrUpstreams field.
func (o *DNSConfig) SetLocalPtrUpstreams(v []string) {
	o.LocalPtrUpstreams = v
}

func (o DNSConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BootstrapDns) {
		toSerialize["bootstrap_dns"] = o.BootstrapDns
	}
	if !IsNil(o.UpstreamDns) {
		toSerialize["upstream_dns"] = o.UpstreamDns
	}
	if !IsNil(o.FallbackDns) {
		toSerialize["fallback_dns"] = o.FallbackDns
	}
	if !IsNil(o.UpstreamDnsFile) {
		toSerialize["upstream_dns_file"] = o.UpstreamDnsFile
	}
	if !IsNil(o.ProtectionEnabled) {
		toSerialize["protection_enabled"] = o.ProtectionEnabled
	}
	if !IsNil(o.Ratelimit) {
		toSerialize["ratelimit"] = o.Ratelimit
	}
	if !IsNil(o.RatelimitSubnetSubnetLenIpv4) {
		toSerialize["ratelimit_subnet_subnet_len_ipv4"] = o.RatelimitSubnetSubnetLenIpv4
	}
	if !IsNil(o.RatelimitSubnetSubnetLenIpv6) {
		toSerialize["ratelimit_subnet_subnet_len_ipv6"] = o.RatelimitSubnetSubnetLenIpv6
	}
	if !IsNil(o.RatelimitWhitelist) {
		toSerialize["ratelimit_whitelist"] = o.RatelimitWhitelist
	}
	if !IsNil(o.BlockingMode) {
		toSerialize["blocking_mode"] = o.BlockingMode
	}
	if !IsNil(o.BlockingIpv4) {
		toSerialize["blocking_ipv4"] = o.BlockingIpv4
	}
	if !IsNil(o.BlockingIpv6) {
		toSerialize["blocking_ipv6"] = o.BlockingIpv6
	}
	if !IsNil(o.BlockedResponseTtl) {
		toSerialize["blocked_response_ttl"] = o.BlockedResponseTtl
	}
	if !IsNil(o.ProtectionDisabledUntil) {
		toSerialize["protection_disabled_until"] = o.ProtectionDisabledUntil
	}
	if !IsNil(o.EdnsCsEnabled) {
		toSerialize["edns_cs_enabled"] = o.EdnsCsEnabled
	}
	if !IsNil(o.EdnsCsUseCustom) {
		toSerialize["edns_cs_use_custom"] = o.EdnsCsUseCustom
	}
	if !IsNil(o.EdnsCsCustomIp) {
		toSerialize["edns_cs_custom_ip"] = o.EdnsCsCustomIp
	}
	if !IsNil(o.DisableIpv6) {
		toSerialize["disable_ipv6"] = o.DisableIpv6
	}
	if !IsNil(o.DnssecEnabled) {
		toSerialize["dnssec_enabled"] = o.DnssecEnabled
	}
	if !IsNil(o.CacheSize) {
		toSerialize["cache_size"] = o.CacheSize
	}
	if !IsNil(o.CacheTtlMin) {
		toSerialize["cache_ttl_min"] = o.CacheTtlMin
	}
	if !IsNil(o.CacheTtlMax) {
		toSerialize["cache_ttl_max"] = o.CacheTtlMax
	}
	if !IsNil(o.CacheOptimistic) {
		toSerialize["cache_optimistic"] = o.CacheOptimistic
	}
	if !IsNil(o.UpstreamMode) {
		toSerialize["upstream_mode"] = o.UpstreamMode
	}
	if !IsNil(o.UsePrivatePtrResolvers) {
		toSerialize["use_private_ptr_resolvers"] = o.UsePrivatePtrResolvers
	}
	if !IsNil(o.ResolveClients) {
		toSerialize["resolve_clients"] = o.ResolveClients
	}
	if !IsNil(o.LocalPtrUpstreams) {
		toSerialize["local_ptr_upstreams"] = o.LocalPtrUpstreams
	}
	return toSerialize, nil
}

type NullableDNSConfig struct {
	value *DNSConfig
	isSet bool
}

func (v NullableDNSConfig) Get() *DNSConfig {
	return v.value
}

func (v *NullableDNSConfig) Set(val *DNSConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSConfig(val *DNSConfig) *NullableDNSConfig {
	return &NullableDNSConfig{value: val, isSet: true}
}

func (v NullableDNSConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


