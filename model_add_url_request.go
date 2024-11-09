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

// checks if the AddUrlRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUrlRequest{}

// AddUrlRequest /add_url request data
type AddUrlRequest struct {
	Name *string `json:"name,omitempty"`
	// URL or an absolute path to the file containing filtering rules. 
	Url *string `json:"url,omitempty"`
	Whitelist *bool `json:"whitelist,omitempty"`
}

// NewAddUrlRequest instantiates a new AddUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUrlRequest() *AddUrlRequest {
	this := AddUrlRequest{}
	return &this
}

// NewAddUrlRequestWithDefaults instantiates a new AddUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUrlRequestWithDefaults() *AddUrlRequest {
	this := AddUrlRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddUrlRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUrlRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddUrlRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddUrlRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AddUrlRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUrlRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AddUrlRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AddUrlRequest) SetUrl(v string) {
	o.Url = &v
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise.
func (o *AddUrlRequest) GetWhitelist() bool {
	if o == nil || IsNil(o.Whitelist) {
		var ret bool
		return ret
	}
	return *o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUrlRequest) GetWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *AddUrlRequest) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given bool and assigns it to the Whitelist field.
func (o *AddUrlRequest) SetWhitelist(v bool) {
	o.Whitelist = &v
}

func (o AddUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Whitelist) {
		toSerialize["whitelist"] = o.Whitelist
	}
	return toSerialize, nil
}

type NullableAddUrlRequest struct {
	value *AddUrlRequest
	isSet bool
}

func (v NullableAddUrlRequest) Get() *AddUrlRequest {
	return v.value
}

func (v *NullableAddUrlRequest) Set(val *AddUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUrlRequest(val *AddUrlRequest) *NullableAddUrlRequest {
	return &NullableAddUrlRequest{value: val, isSet: true}
}

func (v NullableAddUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


