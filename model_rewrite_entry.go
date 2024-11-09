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

// checks if the RewriteEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewriteEntry{}

// RewriteEntry Rewrite rule
type RewriteEntry struct {
	// Domain name
	Domain *string `json:"domain,omitempty"`
	// value of A, AAAA or CNAME DNS record
	Answer *string `json:"answer,omitempty"`
}

// NewRewriteEntry instantiates a new RewriteEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewriteEntry() *RewriteEntry {
	this := RewriteEntry{}
	return &this
}

// NewRewriteEntryWithDefaults instantiates a new RewriteEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewriteEntryWithDefaults() *RewriteEntry {
	this := RewriteEntry{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *RewriteEntry) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewriteEntry) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RewriteEntry) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *RewriteEntry) SetDomain(v string) {
	o.Domain = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *RewriteEntry) GetAnswer() string {
	if o == nil || IsNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewriteEntry) GetAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *RewriteEntry) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *RewriteEntry) SetAnswer(v string) {
	o.Answer = &v
}

func (o RewriteEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewriteEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	return toSerialize, nil
}

type NullableRewriteEntry struct {
	value *RewriteEntry
	isSet bool
}

func (v NullableRewriteEntry) Get() *RewriteEntry {
	return v.value
}

func (v *NullableRewriteEntry) Set(val *RewriteEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRewriteEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRewriteEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewriteEntry(val *RewriteEntry) *NullableRewriteEntry {
	return &NullableRewriteEntry{value: val, isSet: true}
}

func (v NullableRewriteEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewriteEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

