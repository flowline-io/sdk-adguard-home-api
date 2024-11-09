/*
AdGuard Home

AdGuard Home REST-ish API.  Our admin web interface is built on top of this REST-ish API. 

API version: 0.107
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetStatsConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatsConfigResponse{}

// GetStatsConfigResponse Statistics configuration
type GetStatsConfigResponse struct {
	// Are statistics enabled
	Enabled bool `json:"enabled"`
	// Statistics rotation interval in milliseconds
	Interval float32 `json:"interval"`
	// List of host names, which should not be counted
	Ignored []string `json:"ignored"`
}

type _GetStatsConfigResponse GetStatsConfigResponse

// NewGetStatsConfigResponse instantiates a new GetStatsConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatsConfigResponse(enabled bool, interval float32, ignored []string) *GetStatsConfigResponse {
	this := GetStatsConfigResponse{}
	this.Enabled = enabled
	this.Interval = interval
	this.Ignored = ignored
	return &this
}

// NewGetStatsConfigResponseWithDefaults instantiates a new GetStatsConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatsConfigResponseWithDefaults() *GetStatsConfigResponse {
	this := GetStatsConfigResponse{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *GetStatsConfigResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetStatsConfigResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetStatsConfigResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInterval returns the Interval field value
func (o *GetStatsConfigResponse) GetInterval() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *GetStatsConfigResponse) GetIntervalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *GetStatsConfigResponse) SetInterval(v float32) {
	o.Interval = v
}

// GetIgnored returns the Ignored field value
func (o *GetStatsConfigResponse) GetIgnored() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value
// and a boolean to check if the value has been set.
func (o *GetStatsConfigResponse) GetIgnoredOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ignored, true
}

// SetIgnored sets field value
func (o *GetStatsConfigResponse) SetIgnored(v []string) {
	o.Ignored = v
}

func (o GetStatsConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatsConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["interval"] = o.Interval
	toSerialize["ignored"] = o.Ignored
	return toSerialize, nil
}

func (o *GetStatsConfigResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"interval",
		"ignored",
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

	varGetStatsConfigResponse := _GetStatsConfigResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatsConfigResponse)

	if err != nil {
		return err
	}

	*o = GetStatsConfigResponse(varGetStatsConfigResponse)

	return err
}

type NullableGetStatsConfigResponse struct {
	value *GetStatsConfigResponse
	isSet bool
}

func (v NullableGetStatsConfigResponse) Get() *GetStatsConfigResponse {
	return v.value
}

func (v *NullableGetStatsConfigResponse) Set(val *GetStatsConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatsConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatsConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatsConfigResponse(val *GetStatsConfigResponse) *NullableGetStatsConfigResponse {
	return &NullableGetStatsConfigResponse{value: val, isSet: true}
}

func (v NullableGetStatsConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatsConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

