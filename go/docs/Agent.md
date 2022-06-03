# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentRegistrationUid** | Pointer to **string** | Agent registration associated with the agent | [optional] 
**AgentType** | Pointer to **string** | Agent type | [optional] 
**Description** | Pointer to **string** | Description of the the purpose of the agent | [optional] 
**OrgUid** | Pointer to **string** | Agent OrgUID | [optional] 
**ResourceName** | Pointer to **string** | Resource name used for RBAC | [optional] 
**ResourcePolicy** | Pointer to [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**RuntimeDescription** | Pointer to **string** | Description of the runtime of the agent | [optional] 
**RuntimeDetails** | Pointer to [**OrcApiRuntimeDetails**](OrcApiRuntimeDetails.md) |  | [optional] 
**Uid** | Pointer to **string** | Agent UID | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from date, the first date this object was valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to date, the date this object is valid to | [optional] 

## Methods

### NewAgent

`func NewAgent() *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentRegistrationUid

`func (o *Agent) GetAgentRegistrationUid() string`

GetAgentRegistrationUid returns the AgentRegistrationUid field if non-nil, zero value otherwise.

### GetAgentRegistrationUidOk

`func (o *Agent) GetAgentRegistrationUidOk() (*string, bool)`

GetAgentRegistrationUidOk returns a tuple with the AgentRegistrationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRegistrationUid

`func (o *Agent) SetAgentRegistrationUid(v string)`

SetAgentRegistrationUid sets AgentRegistrationUid field to given value.

### HasAgentRegistrationUid

`func (o *Agent) HasAgentRegistrationUid() bool`

HasAgentRegistrationUid returns a boolean if a field has been set.

### GetAgentType

`func (o *Agent) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *Agent) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *Agent) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *Agent) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrgUid

`func (o *Agent) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *Agent) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *Agent) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *Agent) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetResourceName

`func (o *Agent) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Agent) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Agent) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Agent) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourcePolicy

`func (o *Agent) GetResourcePolicy() ResourcePolicy`

GetResourcePolicy returns the ResourcePolicy field if non-nil, zero value otherwise.

### GetResourcePolicyOk

`func (o *Agent) GetResourcePolicyOk() (*ResourcePolicy, bool)`

GetResourcePolicyOk returns a tuple with the ResourcePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePolicy

`func (o *Agent) SetResourcePolicy(v ResourcePolicy)`

SetResourcePolicy sets ResourcePolicy field to given value.

### HasResourcePolicy

`func (o *Agent) HasResourcePolicy() bool`

HasResourcePolicy returns a boolean if a field has been set.

### GetRuntimeDescription

`func (o *Agent) GetRuntimeDescription() string`

GetRuntimeDescription returns the RuntimeDescription field if non-nil, zero value otherwise.

### GetRuntimeDescriptionOk

`func (o *Agent) GetRuntimeDescriptionOk() (*string, bool)`

GetRuntimeDescriptionOk returns a tuple with the RuntimeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDescription

`func (o *Agent) SetRuntimeDescription(v string)`

SetRuntimeDescription sets RuntimeDescription field to given value.

### HasRuntimeDescription

`func (o *Agent) HasRuntimeDescription() bool`

HasRuntimeDescription returns a boolean if a field has been set.

### GetRuntimeDetails

`func (o *Agent) GetRuntimeDetails() OrcApiRuntimeDetails`

GetRuntimeDetails returns the RuntimeDetails field if non-nil, zero value otherwise.

### GetRuntimeDetailsOk

`func (o *Agent) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool)`

GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDetails

`func (o *Agent) SetRuntimeDetails(v OrcApiRuntimeDetails)`

SetRuntimeDetails sets RuntimeDetails field to given value.

### HasRuntimeDetails

`func (o *Agent) HasRuntimeDetails() bool`

HasRuntimeDetails returns a boolean if a field has been set.

### GetUid

`func (o *Agent) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Agent) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Agent) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Agent) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *Agent) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *Agent) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *Agent) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *Agent) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *Agent) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *Agent) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *Agent) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *Agent) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


