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

// checks if the DhcpFindActiveReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpFindActiveReq{}

// DhcpFindActiveReq Request for checking for other DHCP servers in the network. 
type DhcpFindActiveReq struct {
	// The name of the network interface
	Interface *string `json:"interface,omitempty"`
}

// NewDhcpFindActiveReq instantiates a new DhcpFindActiveReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpFindActiveReq() *DhcpFindActiveReq {
	this := DhcpFindActiveReq{}
	return &this
}

// NewDhcpFindActiveReqWithDefaults instantiates a new DhcpFindActiveReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpFindActiveReqWithDefaults() *DhcpFindActiveReq {
	this := DhcpFindActiveReq{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *DhcpFindActiveReq) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpFindActiveReq) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *DhcpFindActiveReq) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *DhcpFindActiveReq) SetInterface(v string) {
	o.Interface = &v
}

func (o DhcpFindActiveReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpFindActiveReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return toSerialize, nil
}

type NullableDhcpFindActiveReq struct {
	value *DhcpFindActiveReq
	isSet bool
}

func (v NullableDhcpFindActiveReq) Get() *DhcpFindActiveReq {
	return v.value
}

func (v *NullableDhcpFindActiveReq) Set(val *DhcpFindActiveReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpFindActiveReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpFindActiveReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpFindActiveReq(val *DhcpFindActiveReq) *NullableDhcpFindActiveReq {
	return &NullableDhcpFindActiveReq{value: val, isSet: true}
}

func (v NullableDhcpFindActiveReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpFindActiveReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


