# SessionOrgTypeMaxLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **float64** | What is the current value | [optional] 
**Description** | Pointer to **string** | Description of the limit | [optional] 
**LimitExceeded** | Pointer to **bool** | Has the limit been met or exceeded | [optional] 
**LimitValue** | Pointer to **float64** | What is the max limit value | [optional] 
**RemainingCapacity** | Pointer to **float64** | How many items can be added to the current value before it is exceeded | [optional] 
**TimeWindowInSeconds** | Pointer to **int32** | The time window in seconds that the limit is calcuated on | [optional] 

## Methods

### NewSessionOrgTypeMaxLimit

`func NewSessionOrgTypeMaxLimit() *SessionOrgTypeMaxLimit`

NewSessionOrgTypeMaxLimit instantiates a new SessionOrgTypeMaxLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionOrgTypeMaxLimitWithDefaults

`func NewSessionOrgTypeMaxLimitWithDefaults() *SessionOrgTypeMaxLimit`

NewSessionOrgTypeMaxLimitWithDefaults instantiates a new SessionOrgTypeMaxLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *SessionOrgTypeMaxLimit) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *SessionOrgTypeMaxLimit) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *SessionOrgTypeMaxLimit) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *SessionOrgTypeMaxLimit) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetDescription

`func (o *SessionOrgTypeMaxLimit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SessionOrgTypeMaxLimit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SessionOrgTypeMaxLimit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SessionOrgTypeMaxLimit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimitExceeded

`func (o *SessionOrgTypeMaxLimit) GetLimitExceeded() bool`

GetLimitExceeded returns the LimitExceeded field if non-nil, zero value otherwise.

### GetLimitExceededOk

`func (o *SessionOrgTypeMaxLimit) GetLimitExceededOk() (*bool, bool)`

GetLimitExceededOk returns a tuple with the LimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitExceeded

`func (o *SessionOrgTypeMaxLimit) SetLimitExceeded(v bool)`

SetLimitExceeded sets LimitExceeded field to given value.

### HasLimitExceeded

`func (o *SessionOrgTypeMaxLimit) HasLimitExceeded() bool`

HasLimitExceeded returns a boolean if a field has been set.

### GetLimitValue

`func (o *SessionOrgTypeMaxLimit) GetLimitValue() float64`

GetLimitValue returns the LimitValue field if non-nil, zero value otherwise.

### GetLimitValueOk

`func (o *SessionOrgTypeMaxLimit) GetLimitValueOk() (*float64, bool)`

GetLimitValueOk returns a tuple with the LimitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitValue

`func (o *SessionOrgTypeMaxLimit) SetLimitValue(v float64)`

SetLimitValue sets LimitValue field to given value.

### HasLimitValue

`func (o *SessionOrgTypeMaxLimit) HasLimitValue() bool`

HasLimitValue returns a boolean if a field has been set.

### GetRemainingCapacity

`func (o *SessionOrgTypeMaxLimit) GetRemainingCapacity() float64`

GetRemainingCapacity returns the RemainingCapacity field if non-nil, zero value otherwise.

### GetRemainingCapacityOk

`func (o *SessionOrgTypeMaxLimit) GetRemainingCapacityOk() (*float64, bool)`

GetRemainingCapacityOk returns a tuple with the RemainingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCapacity

`func (o *SessionOrgTypeMaxLimit) SetRemainingCapacity(v float64)`

SetRemainingCapacity sets RemainingCapacity field to given value.

### HasRemainingCapacity

`func (o *SessionOrgTypeMaxLimit) HasRemainingCapacity() bool`

HasRemainingCapacity returns a boolean if a field has been set.

### GetTimeWindowInSeconds

`func (o *SessionOrgTypeMaxLimit) GetTimeWindowInSeconds() int32`

GetTimeWindowInSeconds returns the TimeWindowInSeconds field if non-nil, zero value otherwise.

### GetTimeWindowInSecondsOk

`func (o *SessionOrgTypeMaxLimit) GetTimeWindowInSecondsOk() (*int32, bool)`

GetTimeWindowInSecondsOk returns a tuple with the TimeWindowInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowInSeconds

`func (o *SessionOrgTypeMaxLimit) SetTimeWindowInSeconds(v int32)`

SetTimeWindowInSeconds sets TimeWindowInSeconds field to given value.

### HasTimeWindowInSeconds

`func (o *SessionOrgTypeMaxLimit) HasTimeWindowInSeconds() bool`

HasTimeWindowInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


