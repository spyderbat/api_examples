# NotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**[]NotificationPolicyRoutesInner**](NotificationPolicyRoutesInner.md) |  | [optional] 
**Targets** | Pointer to [**map[string]NotificationPolicyDestination**](NotificationPolicyDestination.md) |  | [optional] 

## Methods

### NewNotificationPolicy

`func NewNotificationPolicy() *NotificationPolicy`

NewNotificationPolicy instantiates a new NotificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPolicyWithDefaults

`func NewNotificationPolicyWithDefaults() *NotificationPolicy`

NewNotificationPolicyWithDefaults instantiates a new NotificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *NotificationPolicy) GetRoutes() []NotificationPolicyRoutesInner`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NotificationPolicy) GetRoutesOk() (*[]NotificationPolicyRoutesInner, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NotificationPolicy) SetRoutes(v []NotificationPolicyRoutesInner)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NotificationPolicy) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetTargets

`func (o *NotificationPolicy) GetTargets() map[string]NotificationPolicyDestination`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *NotificationPolicy) GetTargetsOk() (*map[string]NotificationPolicyDestination, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *NotificationPolicy) SetTargets(v map[string]NotificationPolicyDestination)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *NotificationPolicy) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


