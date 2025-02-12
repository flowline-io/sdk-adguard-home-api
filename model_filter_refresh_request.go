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

// checks if the FilterRefreshRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterRefreshRequest{}

// FilterRefreshRequest Refresh Filters request data
type FilterRefreshRequest struct {
	Whitelist *bool `json:"whitelist,omitempty"`
}

// NewFilterRefreshRequest instantiates a new FilterRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterRefreshRequest() *FilterRefreshRequest {
	this := FilterRefreshRequest{}
	return &this
}

// NewFilterRefreshRequestWithDefaults instantiates a new FilterRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterRefreshRequestWithDefaults() *FilterRefreshRequest {
	this := FilterRefreshRequest{}
	return &this
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise.
func (o *FilterRefreshRequest) GetWhitelist() bool {
	if o == nil || IsNil(o.Whitelist) {
		var ret bool
		return ret
	}
	return *o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterRefreshRequest) GetWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *FilterRefreshRequest) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given bool and assigns it to the Whitelist field.
func (o *FilterRefreshRequest) SetWhitelist(v bool) {
	o.Whitelist = &v
}

func (o FilterRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterRefreshRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Whitelist) {
		toSerialize["whitelist"] = o.Whitelist
	}
	return toSerialize, nil
}

type NullableFilterRefreshRequest struct {
	value *FilterRefreshRequest
	isSet bool
}

func (v NullableFilterRefreshRequest) Get() *FilterRefreshRequest {
	return v.value
}

func (v *NullableFilterRefreshRequest) Set(val *FilterRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterRefreshRequest(val *FilterRefreshRequest) *NullableFilterRefreshRequest {
	return &NullableFilterRefreshRequest{value: val, isSet: true}
}

func (v NullableFilterRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


