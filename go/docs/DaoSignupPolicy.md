# DaoSignupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatePersonalOrg** | Pointer to **bool** | Create a personal org for the user | [optional] 
**DefaultOrgRoles** | Pointer to [**[]DaoOrgRoles**](DaoOrgRoles.md) | Default org roles for new orgs | [optional] 
**RequiresApproval** | Pointer to **bool** | Require admin approval for new orgs | [optional] 

## Methods

### NewDaoSignupPolicy

`func NewDaoSignupPolicy() *DaoSignupPolicy`

NewDaoSignupPolicy instantiates a new DaoSignupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaoSignupPolicyWithDefaults

`func NewDaoSignupPolicyWithDefaults() *DaoSignupPolicy`

NewDaoSignupPolicyWithDefaults instantiates a new DaoSignupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatePersonalOrg

`func (o *DaoSignupPolicy) GetCreatePersonalOrg() bool`

GetCreatePersonalOrg returns the CreatePersonalOrg field if non-nil, zero value otherwise.

### GetCreatePersonalOrgOk

`func (o *DaoSignupPolicy) GetCreatePersonalOrgOk() (*bool, bool)`

GetCreatePersonalOrgOk returns a tuple with the CreatePersonalOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePersonalOrg

`func (o *DaoSignupPolicy) SetCreatePersonalOrg(v bool)`

SetCreatePersonalOrg sets CreatePersonalOrg field to given value.

### HasCreatePersonalOrg

`func (o *DaoSignupPolicy) HasCreatePersonalOrg() bool`

HasCreatePersonalOrg returns a boolean if a field has been set.

### GetDefaultOrgRoles

`func (o *DaoSignupPolicy) GetDefaultOrgRoles() []DaoOrgRoles`

GetDefaultOrgRoles returns the DefaultOrgRoles field if non-nil, zero value otherwise.

### GetDefaultOrgRolesOk

`func (o *DaoSignupPolicy) GetDefaultOrgRolesOk() (*[]DaoOrgRoles, bool)`

GetDefaultOrgRolesOk returns a tuple with the DefaultOrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrgRoles

`func (o *DaoSignupPolicy) SetDefaultOrgRoles(v []DaoOrgRoles)`

SetDefaultOrgRoles sets DefaultOrgRoles field to given value.

### HasDefaultOrgRoles

`func (o *DaoSignupPolicy) HasDefaultOrgRoles() bool`

HasDefaultOrgRoles returns a boolean if a field has been set.

### GetRequiresApproval

`func (o *DaoSignupPolicy) GetRequiresApproval() bool`

GetRequiresApproval returns the RequiresApproval field if non-nil, zero value otherwise.

### GetRequiresApprovalOk

`func (o *DaoSignupPolicy) GetRequiresApprovalOk() (*bool, bool)`

GetRequiresApprovalOk returns a tuple with the RequiresApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApproval

`func (o *DaoSignupPolicy) SetRequiresApproval(v bool)`

SetRequiresApproval sets RequiresApproval field to given value.

### HasRequiresApproval

`func (o *DaoSignupPolicy) HasRequiresApproval() bool`

HasRequiresApproval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


