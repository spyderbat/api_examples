# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableEmail** | Pointer to **bool** | Disable sending email, this is set when either the user explicitly asks to be unsubscribed to email or there are errors sending emails to the user | [optional] 
**Email** | **string** | User email | 
**ResourceName** | Pointer to **string** | Resource name used by RBAC | [optional] 
**Uid** | Pointer to **string** | User UID | [optional] 
**ValidFrom** | Pointer to **time.Time** | Valid from is the creation date or first date when the object became valid | [optional] 
**ValidTo** | Pointer to **time.Time** | Valid to is the expiration date or the last date this object was valid | [optional] 

## Methods

### NewUser

`func NewUser(email string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableEmail

`func (o *User) GetDisableEmail() bool`

GetDisableEmail returns the DisableEmail field if non-nil, zero value otherwise.

### GetDisableEmailOk

`func (o *User) GetDisableEmailOk() (*bool, bool)`

GetDisableEmailOk returns a tuple with the DisableEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmail

`func (o *User) SetDisableEmail(v bool)`

SetDisableEmail sets DisableEmail field to given value.

### HasDisableEmail

`func (o *User) HasDisableEmail() bool`

HasDisableEmail returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetResourceName

`func (o *User) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *User) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *User) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *User) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetUid

`func (o *User) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *User) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *User) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *User) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetValidFrom

`func (o *User) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *User) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *User) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *User) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *User) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *User) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *User) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *User) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


