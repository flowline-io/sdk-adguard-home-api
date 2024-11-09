# RewriteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to [**RewriteEntry**](RewriteEntry.md) |  | [optional] 
**Update** | Pointer to [**RewriteEntry**](RewriteEntry.md) |  | [optional] 

## Methods

### NewRewriteUpdate

`func NewRewriteUpdate() *RewriteUpdate`

NewRewriteUpdate instantiates a new RewriteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewriteUpdateWithDefaults

`func NewRewriteUpdateWithDefaults() *RewriteUpdate`

NewRewriteUpdateWithDefaults instantiates a new RewriteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *RewriteUpdate) GetTarget() RewriteEntry`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RewriteUpdate) GetTargetOk() (*RewriteEntry, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RewriteUpdate) SetTarget(v RewriteEntry)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RewriteUpdate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUpdate

`func (o *RewriteUpdate) GetUpdate() RewriteEntry`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *RewriteUpdate) GetUpdateOk() (*RewriteEntry, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *RewriteUpdate) SetUpdate(v RewriteEntry)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *RewriteUpdate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


