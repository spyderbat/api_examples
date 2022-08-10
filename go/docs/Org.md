# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSources** | Pointer to **int32** | Total number of active sources within the last 5 minutes | [optional] 
**ActiveUsers** | Pointer to **int32** | Total number of active users within the last 7 days (which might be active on a different org) | [optional] 
**Name** | **string** | Name of the organization | 
**OrgTypeUid** | Pointer to **string** | Organization Type | [optional] 
**OwnerEmail** | **string** | The email address of the user who owns this org | 
**OwnerUid** | Pointer to **string** | The user UID who owns this organization | [optional] 
**ResourceName** | Pointer to **string** | Resource name utilized by RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**TotalSources** | Pointer to **int32** | Total number of sources | [optional] 
**TotalUsers** | Pointer to **int32** | Total number of users | [optional] 
**Uid** | Pointer to **string** | Org UID | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewOrg

`func NewOrg(name string, ownerEmail string, ) *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSources

`func (o *Org) GetActiveSources() int32`

GetActiveSources returns the ActiveSources field if non-nil, zero value otherwise.

### GetActiveSourcesOk

`func (o *Org) GetActiveSourcesOk() (*int32, bool)`

GetActiveSourcesOk returns a tuple with the ActiveSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSources

`func (o *Org) SetActiveSources(v int32)`

SetActiveSources sets ActiveSources field to given value.

### HasActiveSources

`func (o *Org) HasActiveSources() bool`

HasActiveSources returns a boolean if a field has been set.

### GetActiveUsers

`func (o *Org) GetActiveUsers() int32`

GetActiveUsers returns the ActiveUsers field if non-nil, zero value otherwise.

### GetActiveUsersOk

`func (o *Org) GetActiveUsersOk() (*int32, bool)`

GetActiveUsersOk returns a tuple with the ActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUsers

`func (o *Org) SetActiveUsers(v int32)`

SetActiveUsers sets ActiveUsers field to given value.

### HasActiveUsers

`func (o *Org) HasActiveUsers() bool`

HasActiveUsers returns a boolean if a field has been set.

### GetName

`func (o *Org) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Org) SetName(v string)`

SetName sets Name field to given value.


### GetOrgTypeUid

`func (o *Org) GetOrgTypeUid() string`

GetOrgTypeUid returns the OrgTypeUid field if non-nil, zero value otherwise.

### GetOrgTypeUidOk

`func (o *Org) GetOrgTypeUidOk() (*string, bool)`

GetOrgTypeUidOk returns a tuple with the OrgTypeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTypeUid

`func (o *Org) SetOrgTypeUid(v string)`

SetOrgTypeUid sets OrgTypeUid field to given value.

### HasOrgTypeUid

`func (o *Org) HasOrgTypeUid() bool`

HasOrgTypeUid returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Org) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Org) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Org) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.


### GetOwnerUid

`func (o *Org) GetOwnerUid() string`

GetOwnerUid returns the OwnerUid field if non-nil, zero value otherwise.

### GetOwnerUidOk

`func (o *Org) GetOwnerUidOk() (*string, bool)`

GetOwnerUidOk returns a tuple with the OwnerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUid

`func (o *Org) SetOwnerUid(v string)`

SetOwnerUid sets OwnerUid field to given value.

### HasOwnerUid

`func (o *Org) HasOwnerUid() bool`

HasOwnerUid returns a boolean if a field has been set.

### GetResourceName

`func (o *Org) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Org) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Org) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Org) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *Org) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *Org) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *Org) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *Org) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetTags

`func (o *Org) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Org) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Org) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Org) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTotalSources

`func (o *Org) GetTotalSources() int32`

GetTotalSources returns the TotalSources field if non-nil, zero value otherwise.

### GetTotalSourcesOk

`func (o *Org) GetTotalSourcesOk() (*int32, bool)`

GetTotalSourcesOk returns a tuple with the TotalSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSources

`func (o *Org) SetTotalSources(v int32)`

SetTotalSources sets TotalSources field to given value.

### HasTotalSources

`func (o *Org) HasTotalSources() bool`

HasTotalSources returns a boolean if a field has been set.

### GetTotalUsers

`func (o *Org) GetTotalUsers() int32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *Org) GetTotalUsersOk() (*int32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *Org) SetTotalUsers(v int32)`

SetTotalUsers sets TotalUsers field to given value.

### HasTotalUsers

`func (o *Org) HasTotalUsers() bool`

HasTotalUsers returns a boolean if a field has been set.

### GetUid

`func (o *Org) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Org) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Org) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Org) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *Org) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *Org) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *Org) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *Org) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *Org) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *Org) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *Org) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *Org) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


