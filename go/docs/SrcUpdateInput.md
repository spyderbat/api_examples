# SrcUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User supplied description of the source | [optional] 
**LastData** | Pointer to **time.Time** |  | [optional] 
**LastIngestChunkEndTime** | Pointer to **time.Time** | The end of the last chunk ingested from the agent | [optional] 
**LastStoredChunkEndTime** | Pointer to **time.Time** | The end of the last chunk stored from the agent | [optional] 
**Name** | Pointer to **string** | User supplied name of the source | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**RuntimeDescription** | Pointer to **string** | Description of the runtime of the source | [optional] 
**RuntimeDetails** | Pointer to [**OrcApiRuntimeDetails**](OrcApiRuntimeDetails.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Type** | Pointer to **string** | Type of source | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewSrcUpdateInput

`func NewSrcUpdateInput() *SrcUpdateInput`

NewSrcUpdateInput instantiates a new SrcUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcUpdateInputWithDefaults

`func NewSrcUpdateInputWithDefaults() *SrcUpdateInput`

NewSrcUpdateInputWithDefaults instantiates a new SrcUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SrcUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SrcUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SrcUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SrcUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastData

`func (o *SrcUpdateInput) GetLastData() time.Time`

GetLastData returns the LastData field if non-nil, zero value otherwise.

### GetLastDataOk

`func (o *SrcUpdateInput) GetLastDataOk() (*time.Time, bool)`

GetLastDataOk returns a tuple with the LastData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastData

`func (o *SrcUpdateInput) SetLastData(v time.Time)`

SetLastData sets LastData field to given value.

### HasLastData

`func (o *SrcUpdateInput) HasLastData() bool`

HasLastData returns a boolean if a field has been set.

### GetLastIngestChunkEndTime

`func (o *SrcUpdateInput) GetLastIngestChunkEndTime() time.Time`

GetLastIngestChunkEndTime returns the LastIngestChunkEndTime field if non-nil, zero value otherwise.

### GetLastIngestChunkEndTimeOk

`func (o *SrcUpdateInput) GetLastIngestChunkEndTimeOk() (*time.Time, bool)`

GetLastIngestChunkEndTimeOk returns a tuple with the LastIngestChunkEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIngestChunkEndTime

`func (o *SrcUpdateInput) SetLastIngestChunkEndTime(v time.Time)`

SetLastIngestChunkEndTime sets LastIngestChunkEndTime field to given value.

### HasLastIngestChunkEndTime

`func (o *SrcUpdateInput) HasLastIngestChunkEndTime() bool`

HasLastIngestChunkEndTime returns a boolean if a field has been set.

### GetLastStoredChunkEndTime

`func (o *SrcUpdateInput) GetLastStoredChunkEndTime() time.Time`

GetLastStoredChunkEndTime returns the LastStoredChunkEndTime field if non-nil, zero value otherwise.

### GetLastStoredChunkEndTimeOk

`func (o *SrcUpdateInput) GetLastStoredChunkEndTimeOk() (*time.Time, bool)`

GetLastStoredChunkEndTimeOk returns a tuple with the LastStoredChunkEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStoredChunkEndTime

`func (o *SrcUpdateInput) SetLastStoredChunkEndTime(v time.Time)`

SetLastStoredChunkEndTime sets LastStoredChunkEndTime field to given value.

### HasLastStoredChunkEndTime

`func (o *SrcUpdateInput) HasLastStoredChunkEndTime() bool`

HasLastStoredChunkEndTime returns a boolean if a field has been set.

### GetName

`func (o *SrcUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SrcUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SrcUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SrcUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceName

`func (o *SrcUpdateInput) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SrcUpdateInput) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SrcUpdateInput) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *SrcUpdateInput) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *SrcUpdateInput) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *SrcUpdateInput) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *SrcUpdateInput) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *SrcUpdateInput) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetRuntimeDescription

`func (o *SrcUpdateInput) GetRuntimeDescription() string`

GetRuntimeDescription returns the RuntimeDescription field if non-nil, zero value otherwise.

### GetRuntimeDescriptionOk

`func (o *SrcUpdateInput) GetRuntimeDescriptionOk() (*string, bool)`

GetRuntimeDescriptionOk returns a tuple with the RuntimeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDescription

`func (o *SrcUpdateInput) SetRuntimeDescription(v string)`

SetRuntimeDescription sets RuntimeDescription field to given value.

### HasRuntimeDescription

`func (o *SrcUpdateInput) HasRuntimeDescription() bool`

HasRuntimeDescription returns a boolean if a field has been set.

### GetRuntimeDetails

`func (o *SrcUpdateInput) GetRuntimeDetails() OrcApiRuntimeDetails`

GetRuntimeDetails returns the RuntimeDetails field if non-nil, zero value otherwise.

### GetRuntimeDetailsOk

`func (o *SrcUpdateInput) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool)`

GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDetails

`func (o *SrcUpdateInput) SetRuntimeDetails(v OrcApiRuntimeDetails)`

SetRuntimeDetails sets RuntimeDetails field to given value.

### HasRuntimeDetails

`func (o *SrcUpdateInput) HasRuntimeDetails() bool`

HasRuntimeDetails returns a boolean if a field has been set.

### GetTags

`func (o *SrcUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SrcUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SrcUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SrcUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SrcUpdateInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SrcUpdateInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SrcUpdateInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SrcUpdateInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidFrom

`func (o *SrcUpdateInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SrcUpdateInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *SrcUpdateInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *SrcUpdateInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *SrcUpdateInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SrcUpdateInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *SrcUpdateInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *SrcUpdateInput) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


