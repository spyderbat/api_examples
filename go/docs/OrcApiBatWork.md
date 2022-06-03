# OrcApiBatWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **[]string** | arguments to pass to the commandline | [optional] 
**BatUid** | Pointer to **string** | ID of this specific type of bat, specified by Spyderbat | [optional] 
**Enabled** | Pointer to **bool** | Execute this bat or not? | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | input parameters to the bat | [optional] 
**StartOrder** | Pointer to **int32** | Order in which to start this bat | [optional] 
**Uid** | Pointer to **string** | uid of the specific bat work for a specific agent | [optional] 
**Version** | Pointer to **map[string]interface{}** | Newest version of the bat from the repository | [optional] 

## Methods

### NewOrcApiBatWork

`func NewOrcApiBatWork() *OrcApiBatWork`

NewOrcApiBatWork instantiates a new OrcApiBatWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrcApiBatWorkWithDefaults

`func NewOrcApiBatWorkWithDefaults() *OrcApiBatWork`

NewOrcApiBatWorkWithDefaults instantiates a new OrcApiBatWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *OrcApiBatWork) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *OrcApiBatWork) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *OrcApiBatWork) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *OrcApiBatWork) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetBatUid

`func (o *OrcApiBatWork) GetBatUid() string`

GetBatUid returns the BatUid field if non-nil, zero value otherwise.

### GetBatUidOk

`func (o *OrcApiBatWork) GetBatUidOk() (*string, bool)`

GetBatUidOk returns a tuple with the BatUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatUid

`func (o *OrcApiBatWork) SetBatUid(v string)`

SetBatUid sets BatUid field to given value.

### HasBatUid

`func (o *OrcApiBatWork) HasBatUid() bool`

HasBatUid returns a boolean if a field has been set.

### GetEnabled

`func (o *OrcApiBatWork) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrcApiBatWork) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrcApiBatWork) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrcApiBatWork) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParameters

`func (o *OrcApiBatWork) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *OrcApiBatWork) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *OrcApiBatWork) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *OrcApiBatWork) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetStartOrder

`func (o *OrcApiBatWork) GetStartOrder() int32`

GetStartOrder returns the StartOrder field if non-nil, zero value otherwise.

### GetStartOrderOk

`func (o *OrcApiBatWork) GetStartOrderOk() (*int32, bool)`

GetStartOrderOk returns a tuple with the StartOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOrder

`func (o *OrcApiBatWork) SetStartOrder(v int32)`

SetStartOrder sets StartOrder field to given value.

### HasStartOrder

`func (o *OrcApiBatWork) HasStartOrder() bool`

HasStartOrder returns a boolean if a field has been set.

### GetUid

`func (o *OrcApiBatWork) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OrcApiBatWork) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OrcApiBatWork) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *OrcApiBatWork) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersion

`func (o *OrcApiBatWork) GetVersion() map[string]interface{}`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OrcApiBatWork) GetVersionOk() (*map[string]interface{}, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OrcApiBatWork) SetVersion(v map[string]interface{})`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OrcApiBatWork) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


