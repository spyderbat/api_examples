# DaoAgentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classes** | [**[]DaoAgentClass**](DaoAgentClass.md) | Agent classes | 
**Description** | Pointer to **string** | Description of the agent config | [optional] 
**Name** | Pointer to **string** | Name of the agent config | [optional] 
**SourceTags** | Pointer to **[]string** | Tags to be applied to a source upon agent registration | [optional] 
**Type** | Pointer to **string** | Type of agent config | [optional] 

## Methods

### NewDaoAgentConfig

`func NewDaoAgentConfig(classes []DaoAgentClass, ) *DaoAgentConfig`

NewDaoAgentConfig instantiates a new DaoAgentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaoAgentConfigWithDefaults

`func NewDaoAgentConfigWithDefaults() *DaoAgentConfig`

NewDaoAgentConfigWithDefaults instantiates a new DaoAgentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasses

`func (o *DaoAgentConfig) GetClasses() []DaoAgentClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *DaoAgentConfig) GetClassesOk() (*[]DaoAgentClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *DaoAgentConfig) SetClasses(v []DaoAgentClass)`

SetClasses sets Classes field to given value.


### GetDescription

`func (o *DaoAgentConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaoAgentConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaoAgentConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaoAgentConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DaoAgentConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaoAgentConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaoAgentConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DaoAgentConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceTags

`func (o *DaoAgentConfig) GetSourceTags() []string`

GetSourceTags returns the SourceTags field if non-nil, zero value otherwise.

### GetSourceTagsOk

`func (o *DaoAgentConfig) GetSourceTagsOk() (*[]string, bool)`

GetSourceTagsOk returns a tuple with the SourceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTags

`func (o *DaoAgentConfig) SetSourceTags(v []string)`

SetSourceTags sets SourceTags field to given value.

### HasSourceTags

`func (o *DaoAgentConfig) HasSourceTags() bool`

HasSourceTags returns a boolean if a field has been set.

### GetType

`func (o *DaoAgentConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DaoAgentConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DaoAgentConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DaoAgentConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


