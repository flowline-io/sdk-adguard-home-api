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

// checks if the DhcpSearchResultOtherServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpSearchResultOtherServer{}

// DhcpSearchResultOtherServer struct for DhcpSearchResultOtherServer
type DhcpSearchResultOtherServer struct {
	// The result of searching the other DHCP server. 
	Found *string `json:"found,omitempty"`
	// Set if found=error
	Error *string `json:"error,omitempty"`
}

// NewDhcpSearchResultOtherServer instantiates a new DhcpSearchResultOtherServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSearchResultOtherServer() *DhcpSearchResultOtherServer {
	this := DhcpSearchResultOtherServer{}
	return &this
}

// NewDhcpSearchResultOtherServerWithDefaults instantiates a new DhcpSearchResultOtherServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSearchResultOtherServerWithDefaults() *DhcpSearchResultOtherServer {
	this := DhcpSearchResultOtherServer{}
	return &this
}

// GetFound returns the Found field value if set, zero value otherwise.
func (o *DhcpSearchResultOtherServer) GetFound() string {
	if o == nil || IsNil(o.Found) {
		var ret string
		return ret
	}
	return *o.Found
}

// GetFoundOk returns a tuple with the Found field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSearchResultOtherServer) GetFoundOk() (*string, bool) {
	if o == nil || IsNil(o.Found) {
		return nil, false
	}
	return o.Found, true
}

// HasFound returns a boolean if a field has been set.
func (o *DhcpSearchResultOtherServer) HasFound() bool {
	if o != nil && !IsNil(o.Found) {
		return true
	}

	return false
}

// SetFound gets a reference to the given string and assigns it to the Found field.
func (o *DhcpSearchResultOtherServer) SetFound(v string) {
	o.Found = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DhcpSearchResultOtherServer) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSearchResultOtherServer) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DhcpSearchResultOtherServer) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DhcpSearchResultOtherServer) SetError(v string) {
	o.Error = &v
}

func (o DhcpSearchResultOtherServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpSearchResultOtherServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Found) {
		toSerialize["found"] = o.Found
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableDhcpSearchResultOtherServer struct {
	value *DhcpSearchResultOtherServer
	isSet bool
}

func (v NullableDhcpSearchResultOtherServer) Get() *DhcpSearchResultOtherServer {
	return v.value
}

func (v *NullableDhcpSearchResultOtherServer) Set(val *DhcpSearchResultOtherServer) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSearchResultOtherServer) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSearchResultOtherServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSearchResultOtherServer(val *DhcpSearchResultOtherServer) *NullableDhcpSearchResultOtherServer {
	return &NullableDhcpSearchResultOtherServer{value: val, isSet: true}
}

func (v NullableDhcpSearchResultOtherServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSearchResultOtherServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


