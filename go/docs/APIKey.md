# APIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User supplied description of the API key | [optional] 
**Enabled** | Pointer to **bool** | Enables or disables the API key | [optional] 
**Jwt** | Pointer to **string** | JWT to use for authentication | [optional] 
**LastUsed** | Pointer to **time.Time** | The time of las use of the key; only updated every 5 minutes | [optional] 
**OwnerUid** | Pointer to **string** | Owner UID | [optional] 
**Uid** | Pointer to **string** | APIKey UID | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from is the creation date or first date when the API key became valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to is the expiration date or the last date this API key will be valid | [optional] 

## Methods

### NewAPIKey

`func NewAPIKey() *APIKey`

NewAPIKey instantiates a new APIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyWithDefaults

`func NewAPIKeyWithDefaults() *APIKey`

NewAPIKeyWithDefaults instantiates a new APIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *APIKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *APIKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *APIKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *APIKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *APIKey) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *APIKey) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *APIKey) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *APIKey) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetJwt

`func (o *APIKey) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *APIKey) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *APIKey) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *APIKey) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetLastUsed

`func (o *APIKey) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *APIKey) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *APIKey) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *APIKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetOwnerUid

`func (o *APIKey) GetOwnerUid() string`

GetOwnerUid returns the OwnerUid field if non-nil, zero value otherwise.

### GetOwnerUidOk

`func (o *APIKey) GetOwnerUidOk() (*string, bool)`

GetOwnerUidOk returns a tuple with the OwnerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUid

`func (o *APIKey) SetOwnerUid(v string)`

SetOwnerUid sets OwnerUid field to given value.

### HasOwnerUid

`func (o *APIKey) HasOwnerUid() bool`

HasOwnerUid returns a boolean if a field has been set.

### GetUid

`func (o *APIKey) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *APIKey) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *APIKey) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *APIKey) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *APIKey) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *APIKey) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *APIKey) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *APIKey) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *APIKey) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *APIKey) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *APIKey) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *APIKey) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


