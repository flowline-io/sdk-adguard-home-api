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

// checks if the CheckConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckConfigRequest{}

// CheckConfigRequest Configuration to be checked
type CheckConfigRequest struct {
	Dns *CheckConfigRequestInfo `json:"dns,omitempty"`
	Web *CheckConfigRequestInfo `json:"web,omitempty"`
	SetStaticIp *bool `json:"set_static_ip,omitempty"`
}

// NewCheckConfigRequest instantiates a new CheckConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckConfigRequest() *CheckConfigRequest {
	this := CheckConfigRequest{}
	return &this
}

// NewCheckConfigRequestWithDefaults instantiates a new CheckConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckConfigRequestWithDefaults() *CheckConfigRequest {
	this := CheckConfigRequest{}
	return &this
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *CheckConfigRequest) GetDns() CheckConfigRequestInfo {
	if o == nil || IsNil(o.Dns) {
		var ret CheckConfigRequestInfo
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConfigRequest) GetDnsOk() (*CheckConfigRequestInfo, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *CheckConfigRequest) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given CheckConfigRequestInfo and assigns it to the Dns field.
func (o *CheckConfigRequest) SetDns(v CheckConfigRequestInfo) {
	o.Dns = &v
}

// GetWeb returns the Web field value if set, zero value otherwise.
func (o *CheckConfigRequest) GetWeb() CheckConfigRequestInfo {
	if o == nil || IsNil(o.Web) {
		var ret CheckConfigRequestInfo
		return ret
	}
	return *o.Web
}

// GetWebOk returns a tuple with the Web field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConfigRequest) GetWebOk() (*CheckConfigRequestInfo, bool) {
	if o == nil || IsNil(o.Web) {
		return nil, false
	}
	return o.Web, true
}

// HasWeb returns a boolean if a field has been set.
func (o *CheckConfigRequest) HasWeb() bool {
	if o != nil && !IsNil(o.Web) {
		return true
	}

	return false
}

// SetWeb gets a reference to the given CheckConfigRequestInfo and assigns it to the Web field.
func (o *CheckConfigRequest) SetWeb(v CheckConfigRequestInfo) {
	o.Web = &v
}

// GetSetStaticIp returns the SetStaticIp field value if set, zero value otherwise.
func (o *CheckConfigRequest) GetSetStaticIp() bool {
	if o == nil || IsNil(o.SetStaticIp) {
		var ret bool
		return ret
	}
	return *o.SetStaticIp
}

// GetSetStaticIpOk returns a tuple with the SetStaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConfigRequest) GetSetStaticIpOk() (*bool, bool) {
	if o == nil || IsNil(o.SetStaticIp) {
		return nil, false
	}
	return o.SetStaticIp, true
}

// HasSetStaticIp returns a boolean if a field has been set.
func (o *CheckConfigRequest) HasSetStaticIp() bool {
	if o != nil && !IsNil(o.SetStaticIp) {
		return true
	}

	return false
}

// SetSetStaticIp gets a reference to the given bool and assigns it to the SetStaticIp field.
func (o *CheckConfigRequest) SetSetStaticIp(v bool) {
	o.SetStaticIp = &v
}

func (o CheckConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Web) {
		toSerialize["web"] = o.Web
	}
	if !IsNil(o.SetStaticIp) {
		toSerialize["set_static_ip"] = o.SetStaticIp
	}
	return toSerialize, nil
}

type NullableCheckConfigRequest struct {
	value *CheckConfigRequest
	isSet bool
}

func (v NullableCheckConfigRequest) Get() *CheckConfigRequest {
	return v.value
}

func (v *NullableCheckConfigRequest) Set(val *CheckConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckConfigRequest(val *CheckConfigRequest) *NullableCheckConfigRequest {
	return &NullableCheckConfigRequest{value: val, isSet: true}
}

func (v NullableCheckConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

