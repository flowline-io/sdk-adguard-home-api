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

// checks if the QueryLogItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryLogItem{}

// QueryLogItem Query log item
type QueryLogItem struct {
	Answer []DnsAnswer `json:"answer,omitempty"`
	// Answer from upstream server (optional)
	OriginalAnswer []DnsAnswer `json:"original_answer,omitempty"`
	// Defines if the response has been served from cache. 
	Cached *bool `json:"cached,omitempty"`
	// Upstream URL starting with tcp://, tls://, https://, or with an IP address. 
	Upstream *string `json:"upstream,omitempty"`
	// If true, the response had the Authenticated Data (AD) flag set. 
	AnswerDnssec *bool `json:"answer_dnssec,omitempty"`
	// The client's IP address. 
	Client *string `json:"client,omitempty"`
	// The ClientID, if provided in DoH, DoQ, or DoT. 
	ClientId *string `json:"client_id,omitempty"`
	ClientInfo *QueryLogItemClient `json:"client_info,omitempty"`
	ClientProto *string `json:"client_proto,omitempty"`
	// The IP network defined by an EDNS Client-Subnet option in the request message if any. 
	Ecs *string `json:"ecs,omitempty"`
	ElapsedMs *string `json:"elapsedMs,omitempty"`
	Question *DnsQuestion `json:"question,omitempty"`
	// In case if there's a rule applied to this DNS request, this is ID of the filter list that the rule belongs to. Deprecated: use `rules[*].filter_list_id` instead. 
	// Deprecated
	FilterId *int32 `json:"filterId,omitempty"`
	// Filtering rule applied to the request (if any). Deprecated: use `rules[*].text` instead. 
	// Deprecated
	Rule *string `json:"rule,omitempty"`
	// Applied rules.
	Rules []ResultRule `json:"rules,omitempty"`
	// Request filtering status.
	Reason *string `json:"reason,omitempty"`
	// Set if reason=FilteredBlockedService
	ServiceName *string `json:"service_name,omitempty"`
	// DNS response status
	Status *string `json:"status,omitempty"`
	// DNS request processing start time
	Time *string `json:"time,omitempty"`
}

// NewQueryLogItem instantiates a new QueryLogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLogItem() *QueryLogItem {
	this := QueryLogItem{}
	return &this
}

// NewQueryLogItemWithDefaults instantiates a new QueryLogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLogItemWithDefaults() *QueryLogItem {
	this := QueryLogItem{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *QueryLogItem) GetAnswer() []DnsAnswer {
	if o == nil || IsNil(o.Answer) {
		var ret []DnsAnswer
		return ret
	}
	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetAnswerOk() ([]DnsAnswer, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *QueryLogItem) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given []DnsAnswer and assigns it to the Answer field.
func (o *QueryLogItem) SetAnswer(v []DnsAnswer) {
	o.Answer = v
}

// GetOriginalAnswer returns the OriginalAnswer field value if set, zero value otherwise.
func (o *QueryLogItem) GetOriginalAnswer() []DnsAnswer {
	if o == nil || IsNil(o.OriginalAnswer) {
		var ret []DnsAnswer
		return ret
	}
	return o.OriginalAnswer
}

// GetOriginalAnswerOk returns a tuple with the OriginalAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetOriginalAnswerOk() ([]DnsAnswer, bool) {
	if o == nil || IsNil(o.OriginalAnswer) {
		return nil, false
	}
	return o.OriginalAnswer, true
}

// HasOriginalAnswer returns a boolean if a field has been set.
func (o *QueryLogItem) HasOriginalAnswer() bool {
	if o != nil && !IsNil(o.OriginalAnswer) {
		return true
	}

	return false
}

// SetOriginalAnswer gets a reference to the given []DnsAnswer and assigns it to the OriginalAnswer field.
func (o *QueryLogItem) SetOriginalAnswer(v []DnsAnswer) {
	o.OriginalAnswer = v
}

// GetCached returns the Cached field value if set, zero value otherwise.
func (o *QueryLogItem) GetCached() bool {
	if o == nil || IsNil(o.Cached) {
		var ret bool
		return ret
	}
	return *o.Cached
}

// GetCachedOk returns a tuple with the Cached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetCachedOk() (*bool, bool) {
	if o == nil || IsNil(o.Cached) {
		return nil, false
	}
	return o.Cached, true
}

// HasCached returns a boolean if a field has been set.
func (o *QueryLogItem) HasCached() bool {
	if o != nil && !IsNil(o.Cached) {
		return true
	}

	return false
}

// SetCached gets a reference to the given bool and assigns it to the Cached field.
func (o *QueryLogItem) SetCached(v bool) {
	o.Cached = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *QueryLogItem) GetUpstream() string {
	if o == nil || IsNil(o.Upstream) {
		var ret string
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetUpstreamOk() (*string, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *QueryLogItem) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given string and assigns it to the Upstream field.
func (o *QueryLogItem) SetUpstream(v string) {
	o.Upstream = &v
}

// GetAnswerDnssec returns the AnswerDnssec field value if set, zero value otherwise.
func (o *QueryLogItem) GetAnswerDnssec() bool {
	if o == nil || IsNil(o.AnswerDnssec) {
		var ret bool
		return ret
	}
	return *o.AnswerDnssec
}

// GetAnswerDnssecOk returns a tuple with the AnswerDnssec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetAnswerDnssecOk() (*bool, bool) {
	if o == nil || IsNil(o.AnswerDnssec) {
		return nil, false
	}
	return o.AnswerDnssec, true
}

// HasAnswerDnssec returns a boolean if a field has been set.
func (o *QueryLogItem) HasAnswerDnssec() bool {
	if o != nil && !IsNil(o.AnswerDnssec) {
		return true
	}

	return false
}

// SetAnswerDnssec gets a reference to the given bool and assigns it to the AnswerDnssec field.
func (o *QueryLogItem) SetAnswerDnssec(v bool) {
	o.AnswerDnssec = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *QueryLogItem) GetClient() string {
	if o == nil || IsNil(o.Client) {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetClientOk() (*string, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *QueryLogItem) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *QueryLogItem) SetClient(v string) {
	o.Client = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *QueryLogItem) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *QueryLogItem) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *QueryLogItem) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientInfo returns the ClientInfo field value if set, zero value otherwise.
func (o *QueryLogItem) GetClientInfo() QueryLogItemClient {
	if o == nil || IsNil(o.ClientInfo) {
		var ret QueryLogItemClient
		return ret
	}
	return *o.ClientInfo
}

// GetClientInfoOk returns a tuple with the ClientInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetClientInfoOk() (*QueryLogItemClient, bool) {
	if o == nil || IsNil(o.ClientInfo) {
		return nil, false
	}
	return o.ClientInfo, true
}

// HasClientInfo returns a boolean if a field has been set.
func (o *QueryLogItem) HasClientInfo() bool {
	if o != nil && !IsNil(o.ClientInfo) {
		return true
	}

	return false
}

// SetClientInfo gets a reference to the given QueryLogItemClient and assigns it to the ClientInfo field.
func (o *QueryLogItem) SetClientInfo(v QueryLogItemClient) {
	o.ClientInfo = &v
}

// GetClientProto returns the ClientProto field value if set, zero value otherwise.
func (o *QueryLogItem) GetClientProto() string {
	if o == nil || IsNil(o.ClientProto) {
		var ret string
		return ret
	}
	return *o.ClientProto
}

// GetClientProtoOk returns a tuple with the ClientProto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetClientProtoOk() (*string, bool) {
	if o == nil || IsNil(o.ClientProto) {
		return nil, false
	}
	return o.ClientProto, true
}

// HasClientProto returns a boolean if a field has been set.
func (o *QueryLogItem) HasClientProto() bool {
	if o != nil && !IsNil(o.ClientProto) {
		return true
	}

	return false
}

// SetClientProto gets a reference to the given string and assigns it to the ClientProto field.
func (o *QueryLogItem) SetClientProto(v string) {
	o.ClientProto = &v
}

// GetEcs returns the Ecs field value if set, zero value otherwise.
func (o *QueryLogItem) GetEcs() string {
	if o == nil || IsNil(o.Ecs) {
		var ret string
		return ret
	}
	return *o.Ecs
}

// GetEcsOk returns a tuple with the Ecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetEcsOk() (*string, bool) {
	if o == nil || IsNil(o.Ecs) {
		return nil, false
	}
	return o.Ecs, true
}

// HasEcs returns a boolean if a field has been set.
func (o *QueryLogItem) HasEcs() bool {
	if o != nil && !IsNil(o.Ecs) {
		return true
	}

	return false
}

// SetEcs gets a reference to the given string and assigns it to the Ecs field.
func (o *QueryLogItem) SetEcs(v string) {
	o.Ecs = &v
}

// GetElapsedMs returns the ElapsedMs field value if set, zero value otherwise.
func (o *QueryLogItem) GetElapsedMs() string {
	if o == nil || IsNil(o.ElapsedMs) {
		var ret string
		return ret
	}
	return *o.ElapsedMs
}

// GetElapsedMsOk returns a tuple with the ElapsedMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetElapsedMsOk() (*string, bool) {
	if o == nil || IsNil(o.ElapsedMs) {
		return nil, false
	}
	return o.ElapsedMs, true
}

// HasElapsedMs returns a boolean if a field has been set.
func (o *QueryLogItem) HasElapsedMs() bool {
	if o != nil && !IsNil(o.ElapsedMs) {
		return true
	}

	return false
}

// SetElapsedMs gets a reference to the given string and assigns it to the ElapsedMs field.
func (o *QueryLogItem) SetElapsedMs(v string) {
	o.ElapsedMs = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *QueryLogItem) GetQuestion() DnsQuestion {
	if o == nil || IsNil(o.Question) {
		var ret DnsQuestion
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetQuestionOk() (*DnsQuestion, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *QueryLogItem) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given DnsQuestion and assigns it to the Question field.
func (o *QueryLogItem) SetQuestion(v DnsQuestion) {
	o.Question = &v
}

// GetFilterId returns the FilterId field value if set, zero value otherwise.
// Deprecated
func (o *QueryLogItem) GetFilterId() int32 {
	if o == nil || IsNil(o.FilterId) {
		var ret int32
		return ret
	}
	return *o.FilterId
}

// GetFilterIdOk returns a tuple with the FilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryLogItem) GetFilterIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FilterId) {
		return nil, false
	}
	return o.FilterId, true
}

// HasFilterId returns a boolean if a field has been set.
func (o *QueryLogItem) HasFilterId() bool {
	if o != nil && !IsNil(o.FilterId) {
		return true
	}

	return false
}

// SetFilterId gets a reference to the given int32 and assigns it to the FilterId field.
// Deprecated
func (o *QueryLogItem) SetFilterId(v int32) {
	o.FilterId = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
// Deprecated
func (o *QueryLogItem) GetRule() string {
	if o == nil || IsNil(o.Rule) {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QueryLogItem) GetRuleOk() (*string, bool) {
	if o == nil || IsNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *QueryLogItem) HasRule() bool {
	if o != nil && !IsNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
// Deprecated
func (o *QueryLogItem) SetRule(v string) {
	o.Rule = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *QueryLogItem) GetRules() []ResultRule {
	if o == nil || IsNil(o.Rules) {
		var ret []ResultRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetRulesOk() ([]ResultRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *QueryLogItem) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ResultRule and assigns it to the Rules field.
func (o *QueryLogItem) SetRules(v []ResultRule) {
	o.Rules = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *QueryLogItem) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *QueryLogItem) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *QueryLogItem) SetReason(v string) {
	o.Reason = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *QueryLogItem) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *QueryLogItem) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *QueryLogItem) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryLogItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryLogItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryLogItem) SetStatus(v string) {
	o.Status = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryLogItem) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLogItem) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryLogItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *QueryLogItem) SetTime(v string) {
	o.Time = &v
}

func (o QueryLogItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLogItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.OriginalAnswer) {
		toSerialize["original_answer"] = o.OriginalAnswer
	}
	if !IsNil(o.Cached) {
		toSerialize["cached"] = o.Cached
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !IsNil(o.AnswerDnssec) {
		toSerialize["answer_dnssec"] = o.AnswerDnssec
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientInfo) {
		toSerialize["client_info"] = o.ClientInfo
	}
	if !IsNil(o.ClientProto) {
		toSerialize["client_proto"] = o.ClientProto
	}
	if !IsNil(o.Ecs) {
		toSerialize["ecs"] = o.Ecs
	}
	if !IsNil(o.ElapsedMs) {
		toSerialize["elapsedMs"] = o.ElapsedMs
	}
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !IsNil(o.FilterId) {
		toSerialize["filterId"] = o.FilterId
	}
	if !IsNil(o.Rule) {
		toSerialize["rule"] = o.Rule
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableQueryLogItem struct {
	value *QueryLogItem
	isSet bool
}

func (v NullableQueryLogItem) Get() *QueryLogItem {
	return v.value
}

func (v *NullableQueryLogItem) Set(val *QueryLogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLogItem(val *QueryLogItem) *NullableQueryLogItem {
	return &NullableQueryLogItem{value: val, isSet: true}
}

func (v NullableQueryLogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

