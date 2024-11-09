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

// checks if the QueryLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryLog{}

// QueryLog Query log
type QueryLog struct {
	Oldest *string `json:"oldest,omitempty"`
	Data []QueryLogItem `json:"data,omitempty"`
}

// NewQueryLog instantiates a new QueryLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLog() *QueryLog {
	this := QueryLog{}
	return &this
}

// NewQueryLogWithDefaults instantiates a new QueryLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLogWithDefaults() *QueryLog {
	this := QueryLog{}
	return &this
}

// GetOldest returns the Oldest field value if set, zero value otherwise.
func (o *QueryLog) GetOldest() string {
	if o == nil || IsNil(o.Oldest) {
		var ret string
		return ret
	}
	return *o.Oldest
}

// GetOldestOk returns a tuple with the Oldest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLog) GetOldestOk() (*string, bool) {
	if o == nil || IsNil(o.Oldest) {
		return nil, false
	}
	return o.Oldest, true
}

// HasOldest returns a boolean if a field has been set.
func (o *QueryLog) HasOldest() bool {
	if o != nil && !IsNil(o.Oldest) {
		return true
	}

	return false
}

// SetOldest gets a reference to the given string and assigns it to the Oldest field.
func (o *QueryLog) SetOldest(v string) {
	o.Oldest = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryLog) GetData() []QueryLogItem {
	if o == nil || IsNil(o.Data) {
		var ret []QueryLogItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLog) GetDataOk() ([]QueryLogItem, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryLog) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []QueryLogItem and assigns it to the Data field.
func (o *QueryLog) SetData(v []QueryLogItem) {
	o.Data = v
}

func (o QueryLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oldest) {
		toSerialize["oldest"] = o.Oldest
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableQueryLog struct {
	value *QueryLog
	isSet bool
}

func (v NullableQueryLog) Get() *QueryLog {
	return v.value
}

func (v *NullableQueryLog) Set(val *QueryLog) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLog) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLog(val *QueryLog) *NullableQueryLog {
	return &NullableQueryLog{value: val, isSet: true}
}

func (v NullableQueryLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

