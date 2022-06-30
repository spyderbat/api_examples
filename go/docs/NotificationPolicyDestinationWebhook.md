# NotificationPolicyDestinationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoTlsValidation** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewNotificationPolicyDestinationWebhook

`func NewNotificationPolicyDestinationWebhook(url string, ) *NotificationPolicyDestinationWebhook`

NewNotificationPolicyDestinationWebhook instantiates a new NotificationPolicyDestinationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPolicyDestinationWebhookWithDefaults

`func NewNotificationPolicyDestinationWebhookWithDefaults() *NotificationPolicyDestinationWebhook`

NewNotificationPolicyDestinationWebhookWithDefaults instantiates a new NotificationPolicyDestinationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoTlsValidation

`func (o *NotificationPolicyDestinationWebhook) GetNoTlsValidation() bool`

GetNoTlsValidation returns the NoTlsValidation field if non-nil, zero value otherwise.

### GetNoTlsValidationOk

`func (o *NotificationPolicyDestinationWebhook) GetNoTlsValidationOk() (*bool, bool)`

GetNoTlsValidationOk returns a tuple with the NoTlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTlsValidation

`func (o *NotificationPolicyDestinationWebhook) SetNoTlsValidation(v bool)`

SetNoTlsValidation sets NoTlsValidation field to given value.

### HasNoTlsValidation

`func (o *NotificationPolicyDestinationWebhook) HasNoTlsValidation() bool`

HasNoTlsValidation returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationPolicyDestinationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationPolicyDestinationWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationPolicyDestinationWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


