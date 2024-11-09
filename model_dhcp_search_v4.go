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

// checks if the DhcpSearchV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpSearchV4{}

// DhcpSearchV4 struct for DhcpSearchV4
type DhcpSearchV4 struct {
	OtherServer *DhcpSearchResultOtherServer `json:"other_server,omitempty"`
	StaticIp *DhcpSearchResultStaticIP `json:"static_ip,omitempty"`
}

// NewDhcpSearchV4 instantiates a new DhcpSearchV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSearchV4() *DhcpSearchV4 {
	this := DhcpSearchV4{}
	return &this
}

// NewDhcpSearchV4WithDefaults instantiates a new DhcpSearchV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSearchV4WithDefaults() *DhcpSearchV4 {
	this := DhcpSearchV4{}
	return &this
}

// GetOtherServer returns the OtherServer field value if set, zero value otherwise.
func (o *DhcpSearchV4) GetOtherServer() DhcpSearchResultOtherServer {
	if o == nil || IsNil(o.OtherServer) {
		var ret DhcpSearchResultOtherServer
		return ret
	}
	return *o.OtherServer
}

// GetOtherServerOk returns a tuple with the OtherServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSearchV4) GetOtherServerOk() (*DhcpSearchResultOtherServer, bool) {
	if o == nil || IsNil(o.OtherServer) {
		return nil, false
	}
	return o.OtherServer, true
}

// HasOtherServer returns a boolean if a field has been set.
func (o *DhcpSearchV4) HasOtherServer() bool {
	if o != nil && !IsNil(o.OtherServer) {
		return true
	}

	return false
}

// SetOtherServer gets a reference to the given DhcpSearchResultOtherServer and assigns it to the OtherServer field.
func (o *DhcpSearchV4) SetOtherServer(v DhcpSearchResultOtherServer) {
	o.OtherServer = &v
}

// GetStaticIp returns the StaticIp field value if set, zero value otherwise.
func (o *DhcpSearchV4) GetStaticIp() DhcpSearchResultStaticIP {
	if o == nil || IsNil(o.StaticIp) {
		var ret DhcpSearchResultStaticIP
		return ret
	}
	return *o.StaticIp
}

// GetStaticIpOk returns a tuple with the StaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSearchV4) GetStaticIpOk() (*DhcpSearchResultStaticIP, bool) {
	if o == nil || IsNil(o.StaticIp) {
		return nil, false
	}
	return o.StaticIp, true
}

// HasStaticIp returns a boolean if a field has been set.
func (o *DhcpSearchV4) HasStaticIp() bool {
	if o != nil && !IsNil(o.StaticIp) {
		return true
	}

	return false
}

// SetStaticIp gets a reference to the given DhcpSearchResultStaticIP and assigns it to the StaticIp field.
func (o *DhcpSearchV4) SetStaticIp(v DhcpSearchResultStaticIP) {
	o.StaticIp = &v
}

func (o DhcpSearchV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpSearchV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OtherServer) {
		toSerialize["other_server"] = o.OtherServer
	}
	if !IsNil(o.StaticIp) {
		toSerialize["static_ip"] = o.StaticIp
	}
	return toSerialize, nil
}

type NullableDhcpSearchV4 struct {
	value *DhcpSearchV4
	isSet bool
}

func (v NullableDhcpSearchV4) Get() *DhcpSearchV4 {
	return v.value
}

func (v *NullableDhcpSearchV4) Set(val *DhcpSearchV4) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSearchV4) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSearchV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSearchV4(val *DhcpSearchV4) *NullableDhcpSearchV4 {
	return &NullableDhcpSearchV4{value: val, isSet: true}
}

func (v NullableDhcpSearchV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSearchV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

