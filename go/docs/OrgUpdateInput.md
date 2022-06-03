# OrgUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the organization | 
**OrgTypeUid** | Pointer to **string** | Organization Type | [optional] 
**OwnerEmail** | Pointer to **string** | The email address of the user who owns this org | [optional] 
**OwnerUid** | Pointer to **string** | The user UID who owns this organization | [optional] 
**ResourceName** | Pointer to **string** | Resource name utilized by RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewOrgUpdateInput

`func NewOrgUpdateInput(name string, ) *OrgUpdateInput`

NewOrgUpdateInput instantiates a new OrgUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUpdateInputWithDefaults

`func NewOrgUpdateInputWithDefaults() *OrgUpdateInput`

NewOrgUpdateInputWithDefaults instantiates a new OrgUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrgUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgUpdateInput) SetName(v string)`

SetName sets Name field to given value.


### GetOrgTypeUid

`func (o *OrgUpdateInput) GetOrgTypeUid() string`

GetOrgTypeUid returns the OrgTypeUid field if non-nil, zero value otherwise.

### GetOrgTypeUidOk

`func (o *OrgUpdateInput) GetOrgTypeUidOk() (*string, bool)`

GetOrgTypeUidOk returns a tuple with the OrgTypeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTypeUid

`func (o *OrgUpdateInput) SetOrgTypeUid(v string)`

SetOrgTypeUid sets OrgTypeUid field to given value.

### HasOrgTypeUid

`func (o *OrgUpdateInput) HasOrgTypeUid() bool`

HasOrgTypeUid returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *OrgUpdateInput) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *OrgUpdateInput) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *OrgUpdateInput) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *OrgUpdateInput) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetOwnerUid

`func (o *OrgUpdateInput) GetOwnerUid() string`

GetOwnerUid returns the OwnerUid field if non-nil, zero value otherwise.

### GetOwnerUidOk

`func (o *OrgUpdateInput) GetOwnerUidOk() (*string, bool)`

GetOwnerUidOk returns a tuple with the OwnerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUid

`func (o *OrgUpdateInput) SetOwnerUid(v string)`

SetOwnerUid sets OwnerUid field to given value.

### HasOwnerUid

`func (o *OrgUpdateInput) HasOwnerUid() bool`

HasOwnerUid returns a boolean if a field has been set.

### GetResourceName

`func (o *OrgUpdateInput) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *OrgUpdateInput) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *OrgUpdateInput) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *OrgUpdateInput) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *OrgUpdateInput) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *OrgUpdateInput) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *OrgUpdateInput) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *OrgUpdateInput) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetTags

`func (o *OrgUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrgUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrgUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrgUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValidFrom

`func (o *OrgUpdateInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *OrgUpdateInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *OrgUpdateInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *OrgUpdateInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *OrgUpdateInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *OrgUpdateInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *OrgUpdateInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *OrgUpdateInput) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


