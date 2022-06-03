# UiDataSetSourceDataInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | UI supplied JSON object | [optional] 
**LastUsed** | Pointer to **time.Time** | The time of las use of the key; only updated every 5 minutes | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Uid** | Pointer to **string** | UID for the UIData | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from is the creation date or first date when the API key became valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to is the expiration date or the last date this API key will be valid | [optional] 

## Methods

### NewUiDataSetSourceDataInput

`func NewUiDataSetSourceDataInput() *UiDataSetSourceDataInput`

NewUiDataSetSourceDataInput instantiates a new UiDataSetSourceDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiDataSetSourceDataInputWithDefaults

`func NewUiDataSetSourceDataInputWithDefaults() *UiDataSetSourceDataInput`

NewUiDataSetSourceDataInputWithDefaults instantiates a new UiDataSetSourceDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### Get

`func (o *UiDataSetSourceDataInput) Get() string`

Get returns the  field if non-nil, zero value otherwise.

### GetOk

`func (o *UiDataSetSourceDataInput) GetOk() (*string, bool)`

GetOk returns a tuple with the  field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### Set

`func (o *UiDataSetSourceDataInput) Set(v string)`

Set sets  field to given value.

### Has

`func (o *UiDataSetSourceDataInput) Has() bool`

Has returns a boolean if a field has been set.

### GetData

`func (o *UiDataSetSourceDataInput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UiDataSetSourceDataInput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UiDataSetSourceDataInput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UiDataSetSourceDataInput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLastUsed

`func (o *UiDataSetSourceDataInput) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UiDataSetSourceDataInput) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UiDataSetSourceDataInput) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UiDataSetSourceDataInput) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetTags

`func (o *UiDataSetSourceDataInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UiDataSetSourceDataInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UiDataSetSourceDataInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UiDataSetSourceDataInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUid

`func (o *UiDataSetSourceDataInput) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UiDataSetSourceDataInput) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UiDataSetSourceDataInput) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UiDataSetSourceDataInput) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *UiDataSetSourceDataInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *UiDataSetSourceDataInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *UiDataSetSourceDataInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *UiDataSetSourceDataInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *UiDataSetSourceDataInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *UiDataSetSourceDataInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *UiDataSetSourceDataInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *UiDataSetSourceDataInput) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


