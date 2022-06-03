# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User supplied description of the source | [optional] 
**LastData** | Pointer to **time.Time** |  | [optional] 
**LastIngestChunkEndTime** | Pointer to **time.Time** | The end of the last chunk ingested from the agent | [optional] 
**LastStoredChunkEndTime** | Pointer to **time.Time** | The end of the last chunk stored from the agent | [optional] 
**Name** | Pointer to **string** | User supplied name of the source | [optional] 
**OrgUid** | Pointer to **string** | The org this source is associated with | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**RuntimeDescription** | Pointer to **string** | Description of the runtime of the source | [optional] 
**RuntimeDetails** | Pointer to [**OrcApiRuntimeDetails**](OrcApiRuntimeDetails.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Type** | Pointer to **string** | Type of source | [optional] 
**Uid** | Pointer to **string** | The UID of the source | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewSource

`func NewSource() *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Source) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Source) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Source) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Source) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastData

`func (o *Source) GetLastData() time.Time`

GetLastData returns the LastData field if non-nil, zero value otherwise.

### GetLastDataOk

`func (o *Source) GetLastDataOk() (*time.Time, bool)`

GetLastDataOk returns a tuple with the LastData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastData

`func (o *Source) SetLastData(v time.Time)`

SetLastData sets LastData field to given value.

### HasLastData

`func (o *Source) HasLastData() bool`

HasLastData returns a boolean if a field has been set.

### GetLastIngestChunkEndTime

`func (o *Source) GetLastIngestChunkEndTime() time.Time`

GetLastIngestChunkEndTime returns the LastIngestChunkEndTime field if non-nil, zero value otherwise.

### GetLastIngestChunkEndTimeOk

`func (o *Source) GetLastIngestChunkEndTimeOk() (*time.Time, bool)`

GetLastIngestChunkEndTimeOk returns a tuple with the LastIngestChunkEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIngestChunkEndTime

`func (o *Source) SetLastIngestChunkEndTime(v time.Time)`

SetLastIngestChunkEndTime sets LastIngestChunkEndTime field to given value.

### HasLastIngestChunkEndTime

`func (o *Source) HasLastIngestChunkEndTime() bool`

HasLastIngestChunkEndTime returns a boolean if a field has been set.

### GetLastStoredChunkEndTime

`func (o *Source) GetLastStoredChunkEndTime() time.Time`

GetLastStoredChunkEndTime returns the LastStoredChunkEndTime field if non-nil, zero value otherwise.

### GetLastStoredChunkEndTimeOk

`func (o *Source) GetLastStoredChunkEndTimeOk() (*time.Time, bool)`

GetLastStoredChunkEndTimeOk returns a tuple with the LastStoredChunkEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStoredChunkEndTime

`func (o *Source) SetLastStoredChunkEndTime(v time.Time)`

SetLastStoredChunkEndTime sets LastStoredChunkEndTime field to given value.

### HasLastStoredChunkEndTime

`func (o *Source) HasLastStoredChunkEndTime() bool`

HasLastStoredChunkEndTime returns a boolean if a field has been set.

### GetName

`func (o *Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Source) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgUid

`func (o *Source) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *Source) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *Source) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *Source) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetResourceName

`func (o *Source) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Source) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Source) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Source) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *Source) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *Source) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *Source) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *Source) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetRuntimeDescription

`func (o *Source) GetRuntimeDescription() string`

GetRuntimeDescription returns the RuntimeDescription field if non-nil, zero value otherwise.

### GetRuntimeDescriptionOk

`func (o *Source) GetRuntimeDescriptionOk() (*string, bool)`

GetRuntimeDescriptionOk returns a tuple with the RuntimeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDescription

`func (o *Source) SetRuntimeDescription(v string)`

SetRuntimeDescription sets RuntimeDescription field to given value.

### HasRuntimeDescription

`func (o *Source) HasRuntimeDescription() bool`

HasRuntimeDescription returns a boolean if a field has been set.

### GetRuntimeDetails

`func (o *Source) GetRuntimeDetails() OrcApiRuntimeDetails`

GetRuntimeDetails returns the RuntimeDetails field if non-nil, zero value otherwise.

### GetRuntimeDetailsOk

`func (o *Source) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool)`

GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDetails

`func (o *Source) SetRuntimeDetails(v OrcApiRuntimeDetails)`

SetRuntimeDetails sets RuntimeDetails field to given value.

### HasRuntimeDetails

`func (o *Source) HasRuntimeDetails() bool`

HasRuntimeDetails returns a boolean if a field has been set.

### GetTags

`func (o *Source) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Source) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Source) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Source) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Source) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Source) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Source) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Source) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *Source) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Source) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Source) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Source) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *Source) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *Source) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *Source) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *Source) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *Source) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *Source) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *Source) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *Source) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


