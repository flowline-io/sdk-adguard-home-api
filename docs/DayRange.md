# DayRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **float32** | The number of milliseconds elapsed from the start of a day.  It must be less than &#x60;end&#x60; and is expected to be rounded to minutes. So the maximum value is &#x60;86340000&#x60; (23 hours and 59 minutes).  | [optional] 
**End** | Pointer to **float32** | The number of milliseconds elapsed from the start of a day.  It is expected to be rounded to minutes.  The maximum value is &#x60;86400000&#x60; (24 hours).  | [optional] 

## Methods

### NewDayRange

`func NewDayRange() *DayRange`

NewDayRange instantiates a new DayRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayRangeWithDefaults

`func NewDayRangeWithDefaults() *DayRange`

NewDayRangeWithDefaults instantiates a new DayRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *DayRange) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DayRange) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DayRange) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *DayRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *DayRange) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DayRange) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DayRange) SetEnd(v float32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DayRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


