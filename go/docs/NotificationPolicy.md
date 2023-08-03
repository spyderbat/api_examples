# NotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | [**[]NotificationPolicyRoutesInner**](NotificationPolicyRoutesInner.md) |  | 
**Targets** | [**map[string]NotificationPolicyDestination**](NotificationPolicyDestination.md) |  | 

## Methods

### NewNotificationPolicy

`func NewNotificationPolicy(routes []NotificationPolicyRoutesInner, targets map[string]NotificationPolicyDestination, ) *NotificationPolicy`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


