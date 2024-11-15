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

// checks if the BlockedServicesAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockedServicesAll{}

// BlockedServicesAll struct for BlockedServicesAll
type BlockedServicesAll struct {
	BlockedServices []BlockedService `json:"blocked_services"`
}

type _BlockedServicesAll BlockedServicesAll

// NewBlockedServicesAll instantiates a new BlockedServicesAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockedServicesAll(blockedServices []BlockedService) *BlockedServicesAll {
	this := BlockedServicesAll{}
	this.BlockedServices = blockedServices
	return &this
}

// NewBlockedServicesAllWithDefaults instantiates a new BlockedServicesAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockedServicesAllWithDefaults() *BlockedServicesAll {
	this := BlockedServicesAll{}
	return &this
}

// GetBlockedServices returns the BlockedServices field value
func (o *BlockedServicesAll) GetBlockedServices() []BlockedService {
	if o == nil {
		var ret []BlockedService
		return ret
	}

	return o.BlockedServices
}

// GetBlockedServicesOk returns a tuple with the BlockedServices field value
// and a boolean to check if the value has been set.
func (o *BlockedServicesAll) GetBlockedServicesOk() ([]BlockedService, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockedServices, true
}

// SetBlockedServices sets field value
func (o *BlockedServicesAll) SetBlockedServices(v []BlockedService) {
	o.BlockedServices = v
}

func (o BlockedServicesAll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockedServicesAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blocked_services"] = o.BlockedServices
	return toSerialize, nil
}

func (o *BlockedServicesAll) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"blocked_services",
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

	varBlockedServicesAll := _BlockedServicesAll{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockedServicesAll)

	if err != nil {
		return err
	}

	*o = BlockedServicesAll(varBlockedServicesAll)

	return err
}

type NullableBlockedServicesAll struct {
	value *BlockedServicesAll
	isSet bool
}

func (v NullableBlockedServicesAll) Get() *BlockedServicesAll {
	return v.value
}

func (v *NullableBlockedServicesAll) Set(val *BlockedServicesAll) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockedServicesAll) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockedServicesAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockedServicesAll(val *BlockedServicesAll) *NullableBlockedServicesAll {
	return &NullableBlockedServicesAll{value: val, isSet: true}
}

func (v NullableBlockedServicesAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockedServicesAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


