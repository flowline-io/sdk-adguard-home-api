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

// checks if the DhcpConfigV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpConfigV4{}

// DhcpConfigV4 struct for DhcpConfigV4
type DhcpConfigV4 struct {
	GatewayIp *string `json:"gateway_ip,omitempty"`
	SubnetMask *string `json:"subnet_mask,omitempty"`
	RangeStart *string `json:"range_start,omitempty"`
	RangeEnd *string `json:"range_end,omitempty"`
	LeaseDuration *int32 `json:"lease_duration,omitempty"`
}

// NewDhcpConfigV4 instantiates a new DhcpConfigV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpConfigV4() *DhcpConfigV4 {
	this := DhcpConfigV4{}
	return &this
}

// NewDhcpConfigV4WithDefaults instantiates a new DhcpConfigV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpConfigV4WithDefaults() *DhcpConfigV4 {
	this := DhcpConfigV4{}
	return &this
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *DhcpConfigV4) GetGatewayIp() string {
	if o == nil || IsNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpConfigV4) GetGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIp) {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *DhcpConfigV4) HasGatewayIp() bool {
	if o != nil && !IsNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *DhcpConfigV4) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *DhcpConfigV4) GetSubnetMask() string {
	if o == nil || IsNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpConfigV4) GetSubnetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetMask) {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *DhcpConfigV4) HasSubnetMask() bool {
	if o != nil && !IsNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *DhcpConfigV4) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetRangeStart returns the RangeStart field value if set, zero value otherwise.
func (o *DhcpConfigV4) GetRangeStart() string {
	if o == nil || IsNil(o.RangeStart) {
		var ret string
		return ret
	}
	return *o.RangeStart
}

// GetRangeStartOk returns a tuple with the RangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpConfigV4) GetRangeStartOk() (*string, bool) {
	if o == nil || IsNil(o.RangeStart) {
		return nil, false
	}
	return o.RangeStart, true
}

// HasRangeStart returns a boolean if a field has been set.
func (o *DhcpConfigV4) HasRangeStart() bool {
	if o != nil && !IsNil(o.RangeStart) {
		return true
	}

	return false
}

// SetRangeStart gets a reference to the given string and assigns it to the RangeStart field.
func (o *DhcpConfigV4) SetRangeStart(v string) {
	o.RangeStart = &v
}

// GetRangeEnd returns the RangeEnd field value if set, zero value otherwise.
func (o *DhcpConfigV4) GetRangeEnd() string {
	if o == nil || IsNil(o.RangeEnd) {
		var ret string
		return ret
	}
	return *o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpConfigV4) GetRangeEndOk() (*string, bool) {
	if o == nil || IsNil(o.RangeEnd) {
		return nil, false
	}
	return o.RangeEnd, true
}

// HasRangeEnd returns a boolean if a field has been set.
func (o *DhcpConfigV4) HasRangeEnd() bool {
	if o != nil && !IsNil(o.RangeEnd) {
		return true
	}

	return false
}

// SetRangeEnd gets a reference to the given string and assigns it to the RangeEnd field.
func (o *DhcpConfigV4) SetRangeEnd(v string) {
	o.RangeEnd = &v
}

// GetLeaseDuration returns the LeaseDuration field value if set, zero value otherwise.
func (o *DhcpConfigV4) GetLeaseDuration() int32 {
	if o == nil || IsNil(o.LeaseDuration) {
		var ret int32
		return ret
	}
	return *o.LeaseDuration
}

// GetLeaseDurationOk returns a tuple with the LeaseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpConfigV4) GetLeaseDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.LeaseDuration) {
		return nil, false
	}
	return o.LeaseDuration, true
}

// HasLeaseDuration returns a boolean if a field has been set.
func (o *DhcpConfigV4) HasLeaseDuration() bool {
	if o != nil && !IsNil(o.LeaseDuration) {
		return true
	}

	return false
}

// SetLeaseDuration gets a reference to the given int32 and assigns it to the LeaseDuration field.
func (o *DhcpConfigV4) SetLeaseDuration(v int32) {
	o.LeaseDuration = &v
}

func (o DhcpConfigV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpConfigV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayIp) {
		toSerialize["gateway_ip"] = o.GatewayIp
	}
	if !IsNil(o.SubnetMask) {
		toSerialize["subnet_mask"] = o.SubnetMask
	}
	if !IsNil(o.RangeStart) {
		toSerialize["range_start"] = o.RangeStart
	}
	if !IsNil(o.RangeEnd) {
		toSerialize["range_end"] = o.RangeEnd
	}
	if !IsNil(o.LeaseDuration) {
		toSerialize["lease_duration"] = o.LeaseDuration
	}
	return toSerialize, nil
}

type NullableDhcpConfigV4 struct {
	value *DhcpConfigV4
	isSet bool
}

func (v NullableDhcpConfigV4) Get() *DhcpConfigV4 {
	return v.value
}

func (v *NullableDhcpConfigV4) Set(val *DhcpConfigV4) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpConfigV4) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpConfigV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpConfigV4(val *DhcpConfigV4) *NullableDhcpConfigV4 {
	return &NullableDhcpConfigV4{value: val, isSet: true}
}

func (v NullableDhcpConfigV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpConfigV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

