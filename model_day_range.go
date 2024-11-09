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

// checks if the DayRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DayRange{}

// DayRange The single interval within a day.  It begins at the `start` and ends before the `end`. 
type DayRange struct {
	// The number of milliseconds elapsed from the start of a day.  It must be less than `end` and is expected to be rounded to minutes. So the maximum value is `86340000` (23 hours and 59 minutes). 
	Start *float32 `json:"start,omitempty"`
	// The number of milliseconds elapsed from the start of a day.  It is expected to be rounded to minutes.  The maximum value is `86400000` (24 hours). 
	End *float32 `json:"end,omitempty"`
}

// NewDayRange instantiates a new DayRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDayRange() *DayRange {
	this := DayRange{}
	return &this
}

// NewDayRangeWithDefaults instantiates a new DayRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDayRangeWithDefaults() *DayRange {
	this := DayRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *DayRange) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayRange) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *DayRange) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *DayRange) SetStart(v float32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *DayRange) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayRange) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *DayRange) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *DayRange) SetEnd(v float32) {
	o.End = &v
}

func (o DayRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DayRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableDayRange struct {
	value *DayRange
	isSet bool
}

func (v NullableDayRange) Get() *DayRange {
	return v.value
}

func (v *NullableDayRange) Set(val *DayRange) {
	v.value = val
	v.isSet = true
}

func (v NullableDayRange) IsSet() bool {
	return v.isSet
}

func (v *NullableDayRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayRange(val *DayRange) *NullableDayRange {
	return &NullableDayRange{value: val, isSet: true}
}

func (v NullableDayRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


