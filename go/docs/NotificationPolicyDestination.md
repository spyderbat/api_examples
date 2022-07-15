# NotificationPolicyDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** | UI-supplied data | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **[]string** |  | [optional] 
**OrgUid** | Pointer to **string** |  | [optional] 
**Slack** | Pointer to [**NotificationPolicyDestinationSlack**](NotificationPolicyDestinationSlack.md) |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 
**Webhook** | Pointer to [**NotificationPolicyDestinationWebhook**](NotificationPolicyDestinationWebhook.md) |  | [optional] 

## Methods

### NewNotificationPolicyDestination

`func NewNotificationPolicyDestination() *NotificationPolicyDestination`

NewNotificationPolicyDestination instantiates a new NotificationPolicyDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPolicyDestinationWithDefaults

`func NewNotificationPolicyDestinationWithDefaults() *NotificationPolicyDestination`

NewNotificationPolicyDestinationWithDefaults instantiates a new NotificationPolicyDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NotificationPolicyDestination) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationPolicyDestination) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationPolicyDestination) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NotificationPolicyDestination) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *NotificationPolicyDestination) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *NotificationPolicyDestination) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDescription

`func (o *NotificationPolicyDestination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationPolicyDestination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationPolicyDestination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationPolicyDestination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *NotificationPolicyDestination) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NotificationPolicyDestination) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NotificationPolicyDestination) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NotificationPolicyDestination) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgUid

`func (o *NotificationPolicyDestination) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *NotificationPolicyDestination) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *NotificationPolicyDestination) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *NotificationPolicyDestination) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetSlack

`func (o *NotificationPolicyDestination) GetSlack() NotificationPolicyDestinationSlack`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *NotificationPolicyDestination) GetSlackOk() (*NotificationPolicyDestinationSlack, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *NotificationPolicyDestination) SetSlack(v NotificationPolicyDestinationSlack)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *NotificationPolicyDestination) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetUsers

`func (o *NotificationPolicyDestination) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NotificationPolicyDestination) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NotificationPolicyDestination) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NotificationPolicyDestination) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetWebhook

`func (o *NotificationPolicyDestination) GetWebhook() NotificationPolicyDestinationWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *NotificationPolicyDestination) GetWebhookOk() (*NotificationPolicyDestinationWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *NotificationPolicyDestination) SetWebhook(v NotificationPolicyDestinationWebhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *NotificationPolicyDestination) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


