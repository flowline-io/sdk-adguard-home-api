# BlockedServicesSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**Ids** | Pointer to **[]string** | The names of the blocked services.  | [optional] 

## Methods

### NewBlockedServicesSchedule

`func NewBlockedServicesSchedule() *BlockedServicesSchedule`

NewBlockedServicesSchedule instantiates a new BlockedServicesSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockedServicesScheduleWithDefaults

`func NewBlockedServicesScheduleWithDefaults() *BlockedServicesSchedule`

NewBlockedServicesScheduleWithDefaults instantiates a new BlockedServicesSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *BlockedServicesSchedule) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BlockedServicesSchedule) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BlockedServicesSchedule) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *BlockedServicesSchedule) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetIds

`func (o *BlockedServicesSchedule) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BlockedServicesSchedule) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BlockedServicesSchedule) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BlockedServicesSchedule) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


