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

// checks if the FilterConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterConfig{}

// FilterConfig Filtering settings
type FilterConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
	Interval *int32 `json:"interval,omitempty"`
}

// NewFilterConfig instantiates a new FilterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterConfig() *FilterConfig {
	this := FilterConfig{}
	return &this
}

// NewFilterConfigWithDefaults instantiates a new FilterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterConfigWithDefaults() *FilterConfig {
	this := FilterConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *FilterConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *FilterConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *FilterConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *FilterConfig) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterConfig) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *FilterConfig) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *FilterConfig) SetInterval(v int32) {
	o.Interval = &v
}

func (o FilterConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableFilterConfig struct {
	value *FilterConfig
	isSet bool
}

func (v NullableFilterConfig) Get() *FilterConfig {
	return v.value
}

func (v *NullableFilterConfig) Set(val *FilterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterConfig(val *FilterConfig) *NullableFilterConfig {
	return &NullableFilterConfig{value: val, isSet: true}
}

func (v NullableFilterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


