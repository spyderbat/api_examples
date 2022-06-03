# RBACAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action which meets the requirements of RBAC action naming | [optional] 
**CanPerform** | Pointer to **bool** | Return result of querying the users RBAC capabilities | [optional] 
**Error** | Pointer to **string** | Error returned from permission checking | [optional] 
**ResourceName** | Pointer to **string** | ResourceName which meets the requirements of RBAC resource naming | [optional] 

## Methods

### NewRBACAction

`func NewRBACAction() *RBACAction`

NewRBACAction instantiates a new RBACAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBACActionWithDefaults

`func NewRBACActionWithDefaults() *RBACAction`

NewRBACActionWithDefaults instantiates a new RBACAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RBACAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RBACAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RBACAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RBACAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCanPerform

`func (o *RBACAction) GetCanPerform() bool`

GetCanPerform returns the CanPerform field if non-nil, zero value otherwise.

### GetCanPerformOk

`func (o *RBACAction) GetCanPerformOk() (*bool, bool)`

GetCanPerformOk returns a tuple with the CanPerform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPerform

`func (o *RBACAction) SetCanPerform(v bool)`

SetCanPerform sets CanPerform field to given value.

### HasCanPerform

`func (o *RBACAction) HasCanPerform() bool`

HasCanPerform returns a boolean if a field has been set.

### GetError

`func (o *RBACAction) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RBACAction) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RBACAction) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RBACAction) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResourceName

`func (o *RBACAction) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *RBACAction) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *RBACAction) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *RBACAction) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


