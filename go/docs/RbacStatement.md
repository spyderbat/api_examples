# RbacStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | The actions that may be performed | [optional] 
**Condition** | Pointer to **map[string]interface{}** | Additional conditions which may be applied | [optional] 
**Effect** | **string** | The effect of the statement, i.e. allow or deny | 
**Resources** | Pointer to **[]string** | The resource the statement applies to | [optional] 
**Sid** | **string** | Statement ID, used to identify the statement in audit logs | 

## Methods

### NewRbacStatement

`func NewRbacStatement(effect string, sid string, ) *RbacStatement`

NewRbacStatement instantiates a new RbacStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacStatementWithDefaults

`func NewRbacStatementWithDefaults() *RbacStatement`

NewRbacStatementWithDefaults instantiates a new RbacStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RbacStatement) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RbacStatement) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RbacStatement) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RbacStatement) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCondition

`func (o *RbacStatement) GetCondition() map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RbacStatement) GetConditionOk() (*map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RbacStatement) SetCondition(v map[string]interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RbacStatement) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEffect

`func (o *RbacStatement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *RbacStatement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *RbacStatement) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetResources

`func (o *RbacStatement) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RbacStatement) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RbacStatement) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RbacStatement) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSid

`func (o *RbacStatement) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *RbacStatement) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *RbacStatement) SetSid(v string)`

SetSid sets Sid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


