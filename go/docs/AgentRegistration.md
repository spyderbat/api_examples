# AgentRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**DaoAgentConfig**](DaoAgentConfig.md) |  | [optional] 
**CreatedBy** | Pointer to **string** | The user UID of the user who created the agent registration | [optional] 
**Description** | Pointer to **string** | Description of the agent registration | [optional] 
**Name** | Pointer to **string** | Name of the agent registration | [optional] 
**OrgUid** | Pointer to **string** | The OrgUID the registration is associated with | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**Uid** | Pointer to **string** | Agent Registration UID | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewAgentRegistration

`func NewAgentRegistration() *AgentRegistration`

NewAgentRegistration instantiates a new AgentRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRegistrationWithDefaults

`func NewAgentRegistrationWithDefaults() *AgentRegistration`

NewAgentRegistrationWithDefaults instantiates a new AgentRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AgentRegistration) GetConfig() DaoAgentConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AgentRegistration) GetConfigOk() (*DaoAgentConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AgentRegistration) SetConfig(v DaoAgentConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AgentRegistration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AgentRegistration) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AgentRegistration) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AgentRegistration) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AgentRegistration) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *AgentRegistration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentRegistration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentRegistration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentRegistration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AgentRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentRegistration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentRegistration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgUid

`func (o *AgentRegistration) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *AgentRegistration) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *AgentRegistration) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *AgentRegistration) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetResourceName

`func (o *AgentRegistration) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AgentRegistration) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AgentRegistration) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AgentRegistration) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetUid

`func (o *AgentRegistration) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AgentRegistration) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AgentRegistration) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AgentRegistration) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *AgentRegistration) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *AgentRegistration) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *AgentRegistration) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *AgentRegistration) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *AgentRegistration) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *AgentRegistration) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *AgentRegistration) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *AgentRegistration) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


