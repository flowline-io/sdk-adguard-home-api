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

// checks if the SafebrowsingStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SafebrowsingStatus200Response{}

// SafebrowsingStatus200Response struct for SafebrowsingStatus200Response
type SafebrowsingStatus200Response struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSafebrowsingStatus200Response instantiates a new SafebrowsingStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafebrowsingStatus200Response() *SafebrowsingStatus200Response {
	this := SafebrowsingStatus200Response{}
	return &this
}

// NewSafebrowsingStatus200ResponseWithDefaults instantiates a new SafebrowsingStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafebrowsingStatus200ResponseWithDefaults() *SafebrowsingStatus200Response {
	this := SafebrowsingStatus200Response{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SafebrowsingStatus200Response) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafebrowsingStatus200Response) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SafebrowsingStatus200Response) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SafebrowsingStatus200Response) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SafebrowsingStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SafebrowsingStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableSafebrowsingStatus200Response struct {
	value *SafebrowsingStatus200Response
	isSet bool
}

func (v NullableSafebrowsingStatus200Response) Get() *SafebrowsingStatus200Response {
	return v.value
}

func (v *NullableSafebrowsingStatus200Response) Set(val *SafebrowsingStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSafebrowsingStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSafebrowsingStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafebrowsingStatus200Response(val *SafebrowsingStatus200Response) *NullableSafebrowsingStatus200Response {
	return &NullableSafebrowsingStatus200Response{value: val, isSet: true}
}

func (v NullableSafebrowsingStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafebrowsingStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


