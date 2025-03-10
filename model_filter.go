/*
AdGuard Home

AdGuard Home REST-ish API.  Our admin web interface is built on top of this REST-ish API. 

API version: 0.107
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter Filter subscription info
type Filter struct {
	Enabled bool `json:"enabled"`
	Id int64 `json:"id"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	Name string `json:"name"`
	RulesCount int32 `json:"rules_count"`
	Url string `json:"url"`
}

type _Filter Filter

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter(enabled bool, id int64, name string, rulesCount int32, url string) *Filter {
	this := Filter{}
	this.Enabled = enabled
	this.Id = id
	this.Name = name
	this.RulesCount = rulesCount
	this.Url = url
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *Filter) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Filter) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Filter) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value
func (o *Filter) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Filter) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Filter) SetId(v int64) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Filter) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Filter) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Filter) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *Filter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Filter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Filter) SetName(v string) {
	o.Name = v
}

// GetRulesCount returns the RulesCount field value
func (o *Filter) GetRulesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RulesCount
}

// GetRulesCountOk returns a tuple with the RulesCount field value
// and a boolean to check if the value has been set.
func (o *Filter) GetRulesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RulesCount, true
}

// SetRulesCount sets field value
func (o *Filter) SetRulesCount(v int32) {
	o.RulesCount = v
}

// GetUrl returns the Url field value
func (o *Filter) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Filter) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Filter) SetUrl(v string) {
	o.Url = v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	toSerialize["name"] = o.Name
	toSerialize["rules_count"] = o.RulesCount
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *Filter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"id",
		"name",
		"rules_count",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFilter := _Filter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilter)

	if err != nil {
		return err
	}

	*o = Filter(varFilter)

	return err
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


