# DaoInvestigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | UID of user who created the investigation | [optional] 
**Data** | Pointer to **map[string]interface{}** | JSON Object associated with the investigation | [optional] 
**ModifiedBy** | Pointer to **string** | UID of the user who last modified the investigation | [optional] 
**ModifiedOn** | Pointer to **time.Time** | Date the investigation was last modified | [optional] 
**Name** | Pointer to **string** | Name of the investigation | [optional] 
**OrgUid** | Pointer to **string** | Investigation OrgUID | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Uid** | Pointer to **string** | Investigation UID | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 
**Version** | Pointer to **int32** | Version of the investigation | [optional] 

## Methods

### NewDaoInvestigation

`func NewDaoInvestigation() *DaoInvestigation`

NewDaoInvestigation instantiates a new DaoInvestigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaoInvestigationWithDefaults

`func NewDaoInvestigationWithDefaults() *DaoInvestigation`

NewDaoInvestigationWithDefaults instantiates a new DaoInvestigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DaoInvestigation) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DaoInvestigation) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DaoInvestigation) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DaoInvestigation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetData

`func (o *DaoInvestigation) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DaoInvestigation) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DaoInvestigation) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DaoInvestigation) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModifiedBy

`func (o *DaoInvestigation) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *DaoInvestigation) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *DaoInvestigation) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *DaoInvestigation) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *DaoInvestigation) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *DaoInvestigation) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *DaoInvestigation) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *DaoInvestigation) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *DaoInvestigation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaoInvestigation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaoInvestigation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DaoInvestigation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgUid

`func (o *DaoInvestigation) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *DaoInvestigation) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *DaoInvestigation) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *DaoInvestigation) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetResourceName

`func (o *DaoInvestigation) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *DaoInvestigation) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *DaoInvestigation) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *DaoInvestigation) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *DaoInvestigation) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *DaoInvestigation) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *DaoInvestigation) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *DaoInvestigation) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetTags

`func (o *DaoInvestigation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DaoInvestigation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DaoInvestigation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DaoInvestigation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUid

`func (o *DaoInvestigation) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DaoInvestigation) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DaoInvestigation) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DaoInvestigation) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *DaoInvestigation) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *DaoInvestigation) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *DaoInvestigation) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *DaoInvestigation) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *DaoInvestigation) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *DaoInvestigation) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *DaoInvestigation) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *DaoInvestigation) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetVersion

`func (o *DaoInvestigation) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaoInvestigation) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaoInvestigation) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DaoInvestigation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


