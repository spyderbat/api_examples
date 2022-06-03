# SrcCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User supplied description of the source | [optional] 
**Name** | Pointer to **string** | User supplied name of the source | [optional] 
**RuntimeDescription** | Pointer to **string** | Description of the runtime of the source | [optional] 
**RuntimeDetails** | Pointer to [**OrcApiRuntimeDetails**](OrcApiRuntimeDetails.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Type** | Pointer to **string** | Type of source | [optional] 
**Uid** | Pointer to **string** | The UID of the source | [optional] 

## Methods

### NewSrcCreateInput

`func NewSrcCreateInput() *SrcCreateInput`

NewSrcCreateInput instantiates a new SrcCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcCreateInputWithDefaults

`func NewSrcCreateInputWithDefaults() *SrcCreateInput`

NewSrcCreateInputWithDefaults instantiates a new SrcCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SrcCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SrcCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SrcCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SrcCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SrcCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SrcCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SrcCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SrcCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuntimeDescription

`func (o *SrcCreateInput) GetRuntimeDescription() string`

GetRuntimeDescription returns the RuntimeDescription field if non-nil, zero value otherwise.

### GetRuntimeDescriptionOk

`func (o *SrcCreateInput) GetRuntimeDescriptionOk() (*string, bool)`

GetRuntimeDescriptionOk returns a tuple with the RuntimeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDescription

`func (o *SrcCreateInput) SetRuntimeDescription(v string)`

SetRuntimeDescription sets RuntimeDescription field to given value.

### HasRuntimeDescription

`func (o *SrcCreateInput) HasRuntimeDescription() bool`

HasRuntimeDescription returns a boolean if a field has been set.

### GetRuntimeDetails

`func (o *SrcCreateInput) GetRuntimeDetails() OrcApiRuntimeDetails`

GetRuntimeDetails returns the RuntimeDetails field if non-nil, zero value otherwise.

### GetRuntimeDetailsOk

`func (o *SrcCreateInput) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool)`

GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDetails

`func (o *SrcCreateInput) SetRuntimeDetails(v OrcApiRuntimeDetails)`

SetRuntimeDetails sets RuntimeDetails field to given value.

### HasRuntimeDetails

`func (o *SrcCreateInput) HasRuntimeDetails() bool`

HasRuntimeDetails returns a boolean if a field has been set.

### GetTags

`func (o *SrcCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SrcCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SrcCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SrcCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SrcCreateInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SrcCreateInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SrcCreateInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SrcCreateInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *SrcCreateInput) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SrcCreateInput) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SrcCreateInput) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SrcCreateInput) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


