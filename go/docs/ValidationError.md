# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrMsg** | Pointer to **string** | Message regarding the validation failure | [optional] 
**Field** | Pointer to **string** | Field name which failed validation | [optional] 
**Property** | Pointer to **string** | JSON property name of the field which failed validation | [optional] 
**Tags** | Pointer to **string** | Validation tag which failed | [optional] 

## Methods

### NewValidationError

`func NewValidationError() *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrMsg

`func (o *ValidationError) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *ValidationError) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *ValidationError) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.

### HasErrMsg

`func (o *ValidationError) HasErrMsg() bool`

HasErrMsg returns a boolean if a field has been set.

### GetField

`func (o *ValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ValidationError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetProperty

`func (o *ValidationError) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ValidationError) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ValidationError) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ValidationError) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetTags

`func (o *ValidationError) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ValidationError) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ValidationError) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ValidationError) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


