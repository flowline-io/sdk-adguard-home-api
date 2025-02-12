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

// checks if the RemoveUrlRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveUrlRequest{}

// RemoveUrlRequest /remove_url request data
type RemoveUrlRequest struct {
	// Previously added URL containing filtering rules
	Url *string `json:"url,omitempty"`
	Whitelist *bool `json:"whitelist,omitempty"`
}

// NewRemoveUrlRequest instantiates a new RemoveUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveUrlRequest() *RemoveUrlRequest {
	this := RemoveUrlRequest{}
	return &this
}

// NewRemoveUrlRequestWithDefaults instantiates a new RemoveUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveUrlRequestWithDefaults() *RemoveUrlRequest {
	this := RemoveUrlRequest{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RemoveUrlRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveUrlRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RemoveUrlRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RemoveUrlRequest) SetUrl(v string) {
	o.Url = &v
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise.
func (o *RemoveUrlRequest) GetWhitelist() bool {
	if o == nil || IsNil(o.Whitelist) {
		var ret bool
		return ret
	}
	return *o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveUrlRequest) GetWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *RemoveUrlRequest) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given bool and assigns it to the Whitelist field.
func (o *RemoveUrlRequest) SetWhitelist(v bool) {
	o.Whitelist = &v
}

func (o RemoveUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Whitelist) {
		toSerialize["whitelist"] = o.Whitelist
	}
	return toSerialize, nil
}

type NullableRemoveUrlRequest struct {
	value *RemoveUrlRequest
	isSet bool
}

func (v NullableRemoveUrlRequest) Get() *RemoveUrlRequest {
	return v.value
}

func (v *NullableRemoveUrlRequest) Set(val *RemoveUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveUrlRequest(val *RemoveUrlRequest) *NullableRemoveUrlRequest {
	return &NullableRemoveUrlRequest{value: val, isSet: true}
}

func (v NullableRemoveUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


