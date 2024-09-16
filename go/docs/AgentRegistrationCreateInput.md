# AgentRegistrationCreateInput

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

### NewAgentRegistrationCreateInput

`func NewAgentRegistrationCreateInput() *AgentRegistrationCreateInput`

NewAgentRegistrationCreateInput instantiates a new AgentRegistrationCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRegistrationCreateInputWithDefaults

`func NewAgentRegistrationCreateInputWithDefaults() *AgentRegistrationCreateInput`

NewAgentRegistrationCreateInputWithDefaults instantiates a new AgentRegistrationCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AgentRegistrationCreateInput) GetConfig() DaoAgentConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AgentRegistrationCreateInput) GetConfigOk() (*DaoAgentConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AgentRegistrationCreateInput) SetConfig(v DaoAgentConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AgentRegistrationCreateInput) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AgentRegistrationCreateInput) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AgentRegistrationCreateInput) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AgentRegistrationCreateInput) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AgentRegistrationCreateInput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *AgentRegistrationCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentRegistrationCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentRegistrationCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentRegistrationCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AgentRegistrationCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentRegistrationCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentRegistrationCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentRegistrationCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceName

`func (o *AgentRegistrationCreateInput) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AgentRegistrationCreateInput) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AgentRegistrationCreateInput) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AgentRegistrationCreateInput) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetValidFrom

`func (o *AgentRegistrationCreateInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *AgentRegistrationCreateInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *AgentRegistrationCreateInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *AgentRegistrationCreateInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *AgentRegistrationCreateInput) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *AgentRegistrationCreateInput) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *AgentRegistrationCreateInput) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *AgentRegistrationCreateInput) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


