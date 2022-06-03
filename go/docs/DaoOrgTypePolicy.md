# DaoOrgTypePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxActiveSources** | Pointer to **int32** | Maximum number of active sources | [optional] 
**MaxOrgRoles** | Pointer to **int32** | Maximum number of associated organizational roles | [optional] 
**ProcessingStack** | Pointer to **string** | Name of the processing stack | [optional] 
**RetentionPeriod** | Pointer to **int32** | Retention period for stored data, in days | [optional] 

## Methods

### NewDaoOrgTypePolicy

`func NewDaoOrgTypePolicy() *DaoOrgTypePolicy`

NewDaoOrgTypePolicy instantiates a new DaoOrgTypePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaoOrgTypePolicyWithDefaults

`func NewDaoOrgTypePolicyWithDefaults() *DaoOrgTypePolicy`

NewDaoOrgTypePolicyWithDefaults instantiates a new DaoOrgTypePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxActiveSources

`func (o *DaoOrgTypePolicy) GetMaxActiveSources() int32`

GetMaxActiveSources returns the MaxActiveSources field if non-nil, zero value otherwise.

### GetMaxActiveSourcesOk

`func (o *DaoOrgTypePolicy) GetMaxActiveSourcesOk() (*int32, bool)`

GetMaxActiveSourcesOk returns a tuple with the MaxActiveSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveSources

`func (o *DaoOrgTypePolicy) SetMaxActiveSources(v int32)`

SetMaxActiveSources sets MaxActiveSources field to given value.

### HasMaxActiveSources

`func (o *DaoOrgTypePolicy) HasMaxActiveSources() bool`

HasMaxActiveSources returns a boolean if a field has been set.

### GetMaxOrgRoles

`func (o *DaoOrgTypePolicy) GetMaxOrgRoles() int32`

GetMaxOrgRoles returns the MaxOrgRoles field if non-nil, zero value otherwise.

### GetMaxOrgRolesOk

`func (o *DaoOrgTypePolicy) GetMaxOrgRolesOk() (*int32, bool)`

GetMaxOrgRolesOk returns a tuple with the MaxOrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrgRoles

`func (o *DaoOrgTypePolicy) SetMaxOrgRoles(v int32)`

SetMaxOrgRoles sets MaxOrgRoles field to given value.

### HasMaxOrgRoles

`func (o *DaoOrgTypePolicy) HasMaxOrgRoles() bool`

HasMaxOrgRoles returns a boolean if a field has been set.

### GetProcessingStack

`func (o *DaoOrgTypePolicy) GetProcessingStack() string`

GetProcessingStack returns the ProcessingStack field if non-nil, zero value otherwise.

### GetProcessingStackOk

`func (o *DaoOrgTypePolicy) GetProcessingStackOk() (*string, bool)`

GetProcessingStackOk returns a tuple with the ProcessingStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStack

`func (o *DaoOrgTypePolicy) SetProcessingStack(v string)`

SetProcessingStack sets ProcessingStack field to given value.

### HasProcessingStack

`func (o *DaoOrgTypePolicy) HasProcessingStack() bool`

HasProcessingStack returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *DaoOrgTypePolicy) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *DaoOrgTypePolicy) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *DaoOrgTypePolicy) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *DaoOrgTypePolicy) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


