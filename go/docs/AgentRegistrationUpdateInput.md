# AgentRegistrationUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**DaoAgentConfig**](DaoAgentConfig.md) |  | [optional] 
**CreatedBy** | Pointer to **string** | The user UID of the user who created the agent registration | [optional] 
**Description** | Pointer to **string** | Description of the agent registration | [optional] 
**Name** | Pointer to **string** | Name of the agent registration | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewAgentRegistrationUpdateInput

`func NewAgentRegistrationUpdateInput() *AgentRegistrationUpdateInput`

NewAgentRegistrationUpdateInput instantiates a new AgentRegistrationUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRegistrationUpdateInputWithDefaults

`func NewAgentRegistrationUpdateInputWithDefaults() *AgentRegistrationUpdateInput`

NewAgentRegistrationUpdateInputWithDefaults instantiates a new AgentRegistrationUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AgentRegistrationUpdateInput) GetConfig() DaoAgentConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AgentRegistrationUpdateInput) GetConfigOk() (*DaoAgentConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AgentRegistrationUpdateInput) SetConfig(v DaoAgentConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AgentRegistrationUpdateInput) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AgentRegistrationUpdateInput) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AgentRegistrationUpdateInput) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AgentRegistrationUpdateInput) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AgentRegistrationUpdateInput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *AgentRegistrationUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentRegistrationUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentRegistrationUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentRegistrationUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AgentRegistrationUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentRegistrationUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentRegistrationUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentRegistrationUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceName

`func (o *AgentRegistrationUpdateInput) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AgentRegistrationUpdateInput) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AgentRegistrationUpdateInput) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AgentRegistrationUpdateInput) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetValidFrom

`func (o *AgentRegistrationUpdateInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *AgentRegistrationUpdateInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *AgentRegistrationUpdateInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *AgentRegistrationUpdateInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *AgentRegistrationUpdateInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *AgentRegistrationUpdateInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *AgentRegistrationUpdateInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *AgentRegistrationUpdateInput) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


