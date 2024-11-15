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

// checks if the QueryLogConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryLogConfig{}

// QueryLogConfig Query log configuration
type QueryLogConfig struct {
	// Is query log enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Time period for query log rotation. 
	Interval *float32 `json:"interval,omitempty"`
	// Anonymize clients' IP addresses
	AnonymizeClientIp *bool `json:"anonymize_client_ip,omitempty"`
}

// NewQueryLogConfig instantiates a new QueryLogConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLogConfig() *QueryLogConfig {
	this := QueryLogConfig{}
	return &this
}

// NewQueryLogConfigWithDefaults instantiates a new QueryLogConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLogConfigWithDefaults() *QueryLogConfig {
	this := QueryLogConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *QueryLogConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *QueryLogConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *QueryLogConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *QueryLogConfig) GetInterval() float32 {
	if o == nil || IsNil(o.Interval) {
		var ret float32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogConfig) GetIntervalOk() (*float32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *QueryLogConfig) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given float32 and assigns it to the Interval field.
func (o *QueryLogConfig) SetInterval(v float32) {
	o.Interval = &v
}

// GetAnonymizeClientIp returns the AnonymizeClientIp field value if set, zero value otherwise.
func (o *QueryLogConfig) GetAnonymizeClientIp() bool {
	if o == nil || IsNil(o.AnonymizeClientIp) {
		var ret bool
		return ret
	}
	return *o.AnonymizeClientIp
}

// GetAnonymizeClientIpOk returns a tuple with the AnonymizeClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogConfig) GetAnonymizeClientIpOk() (*bool, bool) {
	if o == nil || IsNil(o.AnonymizeClientIp) {
		return nil, false
	}
	return o.AnonymizeClientIp, true
}

// HasAnonymizeClientIp returns a boolean if a field has been set.
func (o *QueryLogConfig) HasAnonymizeClientIp() bool {
	if o != nil && !IsNil(o.AnonymizeClientIp) {
		return true
	}

	return false
}

// SetAnonymizeClientIp gets a reference to the given bool and assigns it to the AnonymizeClientIp field.
func (o *QueryLogConfig) SetAnonymizeClientIp(v bool) {
	o.AnonymizeClientIp = &v
}

func (o QueryLogConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLogConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.AnonymizeClientIp) {
		toSerialize["anonymize_client_ip"] = o.AnonymizeClientIp
	}
	return toSerialize, nil
}

type NullableQueryLogConfig struct {
	value *QueryLogConfig
	isSet bool
}

func (v NullableQueryLogConfig) Get() *QueryLogConfig {
	return v.value
}

func (v *NullableQueryLogConfig) Set(val *QueryLogConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLogConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLogConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLogConfig(val *QueryLogConfig) *NullableQueryLogConfig {
	return &NullableQueryLogConfig{value: val, isSet: true}
}

func (v NullableQueryLogConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLogConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


