# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZone** | Pointer to **string** | Time zone name according to IANA time zone database.  For example &#x60;Europe/Brussels&#x60;.  &#x60;Local&#x60; represents the system&#39;s local time zone.  | [optional] 
**Sun** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 
**Mon** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 
**Tue** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 
**Wed** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 
**Thu** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 
**Fri** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 
**Sat** | Pointer to [**DayRange**](DayRange.md) |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZone

`func (o *Schedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Schedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Schedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Schedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetSun

`func (o *Schedule) GetSun() DayRange`

GetSun returns the Sun field if non-nil, zero value otherwise.

### GetSunOk

`func (o *Schedule) GetSunOk() (*DayRange, bool)`

GetSunOk returns a tuple with the Sun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSun

`func (o *Schedule) SetSun(v DayRange)`

SetSun sets Sun field to given value.

### HasSun

`func (o *Schedule) HasSun() bool`

HasSun returns a boolean if a field has been set.

### GetMon

`func (o *Schedule) GetMon() DayRange`

GetMon returns the Mon field if non-nil, zero value otherwise.

### GetMonOk

`func (o *Schedule) GetMonOk() (*DayRange, bool)`

GetMonOk returns a tuple with the Mon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMon

`func (o *Schedule) SetMon(v DayRange)`

SetMon sets Mon field to given value.

### HasMon

`func (o *Schedule) HasMon() bool`

HasMon returns a boolean if a field has been set.

### GetTue

`func (o *Schedule) GetTue() DayRange`

GetTue returns the Tue field if non-nil, zero value otherwise.

### GetTueOk

`func (o *Schedule) GetTueOk() (*DayRange, bool)`

GetTueOk returns a tuple with the Tue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTue

`func (o *Schedule) SetTue(v DayRange)`

SetTue sets Tue field to given value.

### HasTue

`func (o *Schedule) HasTue() bool`

HasTue returns a boolean if a field has been set.

### GetWed

`func (o *Schedule) GetWed() DayRange`

GetWed returns the Wed field if non-nil, zero value otherwise.

### GetWedOk

`func (o *Schedule) GetWedOk() (*DayRange, bool)`

GetWedOk returns a tuple with the Wed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWed

`func (o *Schedule) SetWed(v DayRange)`

SetWed sets Wed field to given value.

### HasWed

`func (o *Schedule) HasWed() bool`

HasWed returns a boolean if a field has been set.

### GetThu

`func (o *Schedule) GetThu() DayRange`

GetThu returns the Thu field if non-nil, zero value otherwise.

### GetThuOk

`func (o *Schedule) GetThuOk() (*DayRange, bool)`

GetThuOk returns a tuple with the Thu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThu

`func (o *Schedule) SetThu(v DayRange)`

SetThu sets Thu field to given value.

### HasThu

`func (o *Schedule) HasThu() bool`

HasThu returns a boolean if a field has been set.

### GetFri

`func (o *Schedule) GetFri() DayRange`

GetFri returns the Fri field if non-nil, zero value otherwise.

### GetFriOk

`func (o *Schedule) GetFriOk() (*DayRange, bool)`

GetFriOk returns a tuple with the Fri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFri

`func (o *Schedule) SetFri(v DayRange)`

SetFri sets Fri field to given value.

### HasFri

`func (o *Schedule) HasFri() bool`

HasFri returns a boolean if a field has been set.

### GetSat

`func (o *Schedule) GetSat() DayRange`

GetSat returns the Sat field if non-nil, zero value otherwise.

### GetSatOk

`func (o *Schedule) GetSatOk() (*DayRange, bool)`

GetSatOk returns a tuple with the Sat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSat

`func (o *Schedule) SetSat(v DayRange)`

SetSat sets Sat field to given value.

### HasSat

`func (o *Schedule) HasSat() bool`

HasSat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


