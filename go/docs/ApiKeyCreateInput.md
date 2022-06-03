# ApiKeyCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the API key | [optional] 
**ValidTo** | **time.Time** | When the key should expire, in a RFC3339 formated string | 

## Methods

### NewApiKeyCreateInput

`func NewApiKeyCreateInput(validTo time.Time, ) *ApiKeyCreateInput`

NewApiKeyCreateInput instantiates a new ApiKeyCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreateInputWithDefaults

`func NewApiKeyCreateInputWithDefaults() *ApiKeyCreateInput`

NewApiKeyCreateInputWithDefaults instantiates a new ApiKeyCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiKeyCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValidTo

`func (o *ApiKeyCreateInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *ApiKeyCreateInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *ApiKeyCreateInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


