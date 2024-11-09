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

// checks if the Schedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule{}

// Schedule Sets periods of inactivity for filtering blocked services.  The schedule contains 7 days (Sunday to Saturday) and a time zone. 
type Schedule struct {
	// Time zone name according to IANA time zone database.  For example `Europe/Brussels`.  `Local` represents the system's local time zone. 
	TimeZone *string `json:"time_zone,omitempty"`
	Sun *DayRange `json:"sun,omitempty"`
	Mon *DayRange `json:"mon,omitempty"`
	Tue *DayRange `json:"tue,omitempty"`
	Wed *DayRange `json:"wed,omitempty"`
	Thu *DayRange `json:"thu,omitempty"`
	Fri *DayRange `json:"fri,omitempty"`
	Sat *DayRange `json:"sat,omitempty"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *Schedule) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Schedule) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *Schedule) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetSun returns the Sun field value if set, zero value otherwise.
func (o *Schedule) GetSun() DayRange {
	if o == nil || IsNil(o.Sun) {
		var ret DayRange
		return ret
	}
	return *o.Sun
}

// GetSunOk returns a tuple with the Sun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetSunOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Sun) {
		return nil, false
	}
	return o.Sun, true
}

// HasSun returns a boolean if a field has been set.
func (o *Schedule) HasSun() bool {
	if o != nil && !IsNil(o.Sun) {
		return true
	}

	return false
}

// SetSun gets a reference to the given DayRange and assigns it to the Sun field.
func (o *Schedule) SetSun(v DayRange) {
	o.Sun = &v
}

// GetMon returns the Mon field value if set, zero value otherwise.
func (o *Schedule) GetMon() DayRange {
	if o == nil || IsNil(o.Mon) {
		var ret DayRange
		return ret
	}
	return *o.Mon
}

// GetMonOk returns a tuple with the Mon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetMonOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Mon) {
		return nil, false
	}
	return o.Mon, true
}

// HasMon returns a boolean if a field has been set.
func (o *Schedule) HasMon() bool {
	if o != nil && !IsNil(o.Mon) {
		return true
	}

	return false
}

// SetMon gets a reference to the given DayRange and assigns it to the Mon field.
func (o *Schedule) SetMon(v DayRange) {
	o.Mon = &v
}

// GetTue returns the Tue field value if set, zero value otherwise.
func (o *Schedule) GetTue() DayRange {
	if o == nil || IsNil(o.Tue) {
		var ret DayRange
		return ret
	}
	return *o.Tue
}

// GetTueOk returns a tuple with the Tue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTueOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Tue) {
		return nil, false
	}
	return o.Tue, true
}

// HasTue returns a boolean if a field has been set.
func (o *Schedule) HasTue() bool {
	if o != nil && !IsNil(o.Tue) {
		return true
	}

	return false
}

// SetTue gets a reference to the given DayRange and assigns it to the Tue field.
func (o *Schedule) SetTue(v DayRange) {
	o.Tue = &v
}

// GetWed returns the Wed field value if set, zero value otherwise.
func (o *Schedule) GetWed() DayRange {
	if o == nil || IsNil(o.Wed) {
		var ret DayRange
		return ret
	}
	return *o.Wed
}

// GetWedOk returns a tuple with the Wed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetWedOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Wed) {
		return nil, false
	}
	return o.Wed, true
}

// HasWed returns a boolean if a field has been set.
func (o *Schedule) HasWed() bool {
	if o != nil && !IsNil(o.Wed) {
		return true
	}

	return false
}

// SetWed gets a reference to the given DayRange and assigns it to the Wed field.
func (o *Schedule) SetWed(v DayRange) {
	o.Wed = &v
}

// GetThu returns the Thu field value if set, zero value otherwise.
func (o *Schedule) GetThu() DayRange {
	if o == nil || IsNil(o.Thu) {
		var ret DayRange
		return ret
	}
	return *o.Thu
}

// GetThuOk returns a tuple with the Thu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetThuOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Thu) {
		return nil, false
	}
	return o.Thu, true
}

// HasThu returns a boolean if a field has been set.
func (o *Schedule) HasThu() bool {
	if o != nil && !IsNil(o.Thu) {
		return true
	}

	return false
}

// SetThu gets a reference to the given DayRange and assigns it to the Thu field.
func (o *Schedule) SetThu(v DayRange) {
	o.Thu = &v
}

// GetFri returns the Fri field value if set, zero value otherwise.
func (o *Schedule) GetFri() DayRange {
	if o == nil || IsNil(o.Fri) {
		var ret DayRange
		return ret
	}
	return *o.Fri
}

// GetFriOk returns a tuple with the Fri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetFriOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Fri) {
		return nil, false
	}
	return o.Fri, true
}

// HasFri returns a boolean if a field has been set.
func (o *Schedule) HasFri() bool {
	if o != nil && !IsNil(o.Fri) {
		return true
	}

	return false
}

// SetFri gets a reference to the given DayRange and assigns it to the Fri field.
func (o *Schedule) SetFri(v DayRange) {
	o.Fri = &v
}

// GetSat returns the Sat field value if set, zero value otherwise.
func (o *Schedule) GetSat() DayRange {
	if o == nil || IsNil(o.Sat) {
		var ret DayRange
		return ret
	}
	return *o.Sat
}

// GetSatOk returns a tuple with the Sat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetSatOk() (*DayRange, bool) {
	if o == nil || IsNil(o.Sat) {
		return nil, false
	}
	return o.Sat, true
}

// HasSat returns a boolean if a field has been set.
func (o *Schedule) HasSat() bool {
	if o != nil && !IsNil(o.Sat) {
		return true
	}

	return false
}

// SetSat gets a reference to the given DayRange and assigns it to the Sat field.
func (o *Schedule) SetSat(v DayRange) {
	o.Sat = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeZone) {
		toSerialize["time_zone"] = o.TimeZone
	}
	if !IsNil(o.Sun) {
		toSerialize["sun"] = o.Sun
	}
	if !IsNil(o.Mon) {
		toSerialize["mon"] = o.Mon
	}
	if !IsNil(o.Tue) {
		toSerialize["tue"] = o.Tue
	}
	if !IsNil(o.Wed) {
		toSerialize["wed"] = o.Wed
	}
	if !IsNil(o.Thu) {
		toSerialize["thu"] = o.Thu
	}
	if !IsNil(o.Fri) {
		toSerialize["fri"] = o.Fri
	}
	if !IsNil(o.Sat) {
		toSerialize["sat"] = o.Sat
	}
	return toSerialize, nil
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

