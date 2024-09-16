# DaoOrgUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the user who belongs to this org | 
**Roles** | **[]string** | The roles of the user who belongs to this org | 

## Methods

### NewDaoOrgUser

`func NewDaoOrgUser(email string, roles []string, ) *DaoOrgUser`

NewDaoOrgUser instantiates a new DaoOrgUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaoOrgUserWithDefaults

`func NewDaoOrgUserWithDefaults() *DaoOrgUser`

NewDaoOrgUserWithDefaults instantiates a new DaoOrgUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DaoOrgUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DaoOrgUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DaoOrgUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoles

`func (o *DaoOrgUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DaoOrgUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DaoOrgUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


