# AgentSetOrgWorkInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | User defined tags | [optional] 
**Work** | Pointer to [**OrcApiAgentWork**](OrcApiAgentWork.md) |  | [optional] 

## Methods

### NewAgentSetOrgWorkInput

`func NewAgentSetOrgWorkInput() *AgentSetOrgWorkInput`

NewAgentSetOrgWorkInput instantiates a new AgentSetOrgWorkInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSetOrgWorkInputWithDefaults

`func NewAgentSetOrgWorkInputWithDefaults() *AgentSetOrgWorkInput`

NewAgentSetOrgWorkInputWithDefaults instantiates a new AgentSetOrgWorkInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AgentSetOrgWorkInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentSetOrgWorkInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentSetOrgWorkInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentSetOrgWorkInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWork

`func (o *AgentSetOrgWorkInput) GetWork() OrcApiAgentWork`

GetWork returns the Work field if non-nil, zero value otherwise.

### GetWorkOk

`func (o *AgentSetOrgWorkInput) GetWorkOk() (*OrcApiAgentWork, bool)`

GetWorkOk returns a tuple with the Work field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWork

`func (o *AgentSetOrgWorkInput) SetWork(v OrcApiAgentWork)`

SetWork sets Work field to given value.

### HasWork

`func (o *AgentSetOrgWorkInput) HasWork() bool`

HasWork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


