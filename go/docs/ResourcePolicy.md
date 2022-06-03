# ResourcePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the resource policy | [optional] 
**Statements** | Pointer to [**[]RbacStatement**](RbacStatement.md) | List of statements to be examined by the resource policy | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewResourcePolicy

`func NewResourcePolicy() *ResourcePolicy`

NewResourcePolicy instantiates a new ResourcePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePolicyWithDefaults

`func NewResourcePolicyWithDefaults() *ResourcePolicy`

NewResourcePolicyWithDefaults instantiates a new ResourcePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourcePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourcePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatements

`func (o *ResourcePolicy) GetStatements() []RbacStatement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *ResourcePolicy) GetStatementsOk() (*[]RbacStatement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *ResourcePolicy) SetStatements(v []RbacStatement)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *ResourcePolicy) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetVersion

`func (o *ResourcePolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResourcePolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResourcePolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResourcePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


