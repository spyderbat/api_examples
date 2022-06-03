# InvestigationCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | UID of user who created the investigation | [optional] 
**Data** | Pointer to **map[string]interface{}** | JSON Object associated with the investigation | [optional] 
**ModifiedBy** | Pointer to **string** | UID of the user who last modified the investigation | [optional] 
**ModifiedOn** | Pointer to **time.Time** | Date the investigation was last modified | [optional] 
**Name** | Pointer to **string** | Name of the investigation | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 
**Version** | Pointer to **int32** | Version of the investigation | [optional] 

## Methods

### NewInvestigationCreateInput

`func NewInvestigationCreateInput() *InvestigationCreateInput`

NewInvestigationCreateInput instantiates a new InvestigationCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationCreateInputWithDefaults

`func NewInvestigationCreateInputWithDefaults() *InvestigationCreateInput`

NewInvestigationCreateInputWithDefaults instantiates a new InvestigationCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *InvestigationCreateInput) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InvestigationCreateInput) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InvestigationCreateInput) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InvestigationCreateInput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetData

`func (o *InvestigationCreateInput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvestigationCreateInput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvestigationCreateInput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *InvestigationCreateInput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModifiedBy

`func (o *InvestigationCreateInput) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *InvestigationCreateInput) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *InvestigationCreateInput) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *InvestigationCreateInput) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *InvestigationCreateInput) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *InvestigationCreateInput) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *InvestigationCreateInput) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *InvestigationCreateInput) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *InvestigationCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvestigationCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceName

`func (o *InvestigationCreateInput) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *InvestigationCreateInput) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *InvestigationCreateInput) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *InvestigationCreateInput) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *InvestigationCreateInput) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *InvestigationCreateInput) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *InvestigationCreateInput) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *InvestigationCreateInput) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetTags

`func (o *InvestigationCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InvestigationCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InvestigationCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InvestigationCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValidFrom

`func (o *InvestigationCreateInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *InvestigationCreateInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *InvestigationCreateInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *InvestigationCreateInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *InvestigationCreateInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *InvestigationCreateInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *InvestigationCreateInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *InvestigationCreateInput) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetVersion

`func (o *InvestigationCreateInput) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InvestigationCreateInput) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InvestigationCreateInput) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InvestigationCreateInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


