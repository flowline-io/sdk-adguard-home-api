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

// checks if the DhcpLease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpLease{}

// DhcpLease DHCP lease information
type DhcpLease struct {
	Mac string `json:"mac"`
	Ip string `json:"ip"`
	Hostname string `json:"hostname"`
	Expires string `json:"expires"`
}

type _DhcpLease DhcpLease

// NewDhcpLease instantiates a new DhcpLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpLease(mac string, ip string, hostname string, expires string) *DhcpLease {
	this := DhcpLease{}
	this.Mac = mac
	this.Ip = ip
	this.Hostname = hostname
	this.Expires = expires
	return &this
}

// NewDhcpLeaseWithDefaults instantiates a new DhcpLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpLeaseWithDefaults() *DhcpLease {
	this := DhcpLease{}
	return &this
}

// GetMac returns the Mac field value
func (o *DhcpLease) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *DhcpLease) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *DhcpLease) SetMac(v string) {
	o.Mac = v
}

// GetIp returns the Ip field value
func (o *DhcpLease) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *DhcpLease) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *DhcpLease) SetIp(v string) {
	o.Ip = v
}

// GetHostname returns the Hostname field value
func (o *DhcpLease) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *DhcpLease) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *DhcpLease) SetHostname(v string) {
	o.Hostname = v
}

// GetExpires returns the Expires field value
func (o *DhcpLease) GetExpires() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *DhcpLease) GetExpiresOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *DhcpLease) SetExpires(v string) {
	o.Expires = v
}

func (o DhcpLease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpLease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	toSerialize["ip"] = o.Ip
	toSerialize["hostname"] = o.Hostname
	toSerialize["expires"] = o.Expires
	return toSerialize, nil
}

func (o *DhcpLease) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mac",
		"ip",
		"hostname",
		"expires",
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

	varDhcpLease := _DhcpLease{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDhcpLease)

	if err != nil {
		return err
	}

	*o = DhcpLease(varDhcpLease)

	return err
}

type NullableDhcpLease struct {
	value *DhcpLease
	isSet bool
}

func (v NullableDhcpLease) Get() *DhcpLease {
	return v.value
}

func (v *NullableDhcpLease) Set(val *DhcpLease) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpLease) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpLease(val *DhcpLease) *NullableDhcpLease {
	return &NullableDhcpLease{value: val, isSet: true}
}

func (v NullableDhcpLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


