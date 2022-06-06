# UIData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | UI supplied JSON object | [optional] 
**DataKey** | Pointer to **string** | Key for the data | [optional] 
**LastUsed** | Pointer to **time.Time** | The time of las use of the key; only updated every 5 minutes | [optional] 
**OrgUid** | Pointer to **string** | Org UID | [optional] 
**OwnerUid** | Pointer to **string** | Owner UID | [optional] 
**SourceUid** | Pointer to **string** | Source UID | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Uid** | Pointer to **string** | UID for the UIData | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from is the creation date or first date when the API key became valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to is the expiration date or the last date this API key will be valid | [optional] 

## Methods

### NewUIData

`func NewUIData() *UIData`

NewUIData instantiates a new UIData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIDataWithDefaults

`func NewUIDataWithDefaults() *UIData`

NewUIDataWithDefaults instantiates a new UIData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UIData) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UIData) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UIData) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UIData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDataKey

`func (o *UIData) GetDataKey() string`

GetDataKey returns the DataKey field if non-nil, zero value otherwise.

### GetDataKeyOk

`func (o *UIData) GetDataKeyOk() (*string, bool)`

GetDataKeyOk returns a tuple with the DataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKey

`func (o *UIData) SetDataKey(v string)`

SetDataKey sets DataKey field to given value.

### HasDataKey

`func (o *UIData) HasDataKey() bool`

HasDataKey returns a boolean if a field has been set.

### GetLastUsed

`func (o *UIData) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UIData) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UIData) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UIData) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetOrgUid

`func (o *UIData) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *UIData) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *UIData) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *UIData) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetOwnerUid

`func (o *UIData) GetOwnerUid() string`

GetOwnerUid returns the OwnerUid field if non-nil, zero value otherwise.

### GetOwnerUidOk

`func (o *UIData) GetOwnerUidOk() (*string, bool)`

GetOwnerUidOk returns a tuple with the OwnerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUid

`func (o *UIData) SetOwnerUid(v string)`

SetOwnerUid sets OwnerUid field to given value.

### HasOwnerUid

`func (o *UIData) HasOwnerUid() bool`

HasOwnerUid returns a boolean if a field has been set.

### GetSourceUid

`func (o *UIData) GetSourceUid() string`

GetSourceUid returns the SourceUid field if non-nil, zero value otherwise.

### GetSourceUidOk

`func (o *UIData) GetSourceUidOk() (*string, bool)`

GetSourceUidOk returns a tuple with the SourceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUid

`func (o *UIData) SetSourceUid(v string)`

SetSourceUid sets SourceUid field to given value.

### HasSourceUid

`func (o *UIData) HasSourceUid() bool`

HasSourceUid returns a boolean if a field has been set.

### GetTags

`func (o *UIData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UIData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UIData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UIData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUid

`func (o *UIData) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UIData) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UIData) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UIData) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *UIData) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *UIData) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *UIData) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *UIData) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *UIData) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *UIData) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *UIData) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *UIData) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

