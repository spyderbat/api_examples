# NotificationPolicyRoutesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** | UI-supplied data | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to [**NotificationPolicyDestination**](NotificationPolicyDestination.md) |  | [optional] 
**Expr** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Target** | Pointer to **string** | One of the targets specified in &#x60;targets&#x60; | [optional] 

## Methods

### NewNotificationPolicyRoutesInner

`func NewNotificationPolicyRoutesInner() *NotificationPolicyRoutesInner`

NewNotificationPolicyRoutesInner instantiates a new NotificationPolicyRoutesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPolicyRoutesInnerWithDefaults

`func NewNotificationPolicyRoutesInnerWithDefaults() *NotificationPolicyRoutesInner`

NewNotificationPolicyRoutesInnerWithDefaults instantiates a new NotificationPolicyRoutesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NotificationPolicyRoutesInner) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationPolicyRoutesInner) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationPolicyRoutesInner) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NotificationPolicyRoutesInner) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *NotificationPolicyRoutesInner) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *NotificationPolicyRoutesInner) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDescription

`func (o *NotificationPolicyRoutesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationPolicyRoutesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationPolicyRoutesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationPolicyRoutesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestination

`func (o *NotificationPolicyRoutesInner) GetDestination() NotificationPolicyDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NotificationPolicyRoutesInner) GetDestinationOk() (*NotificationPolicyDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NotificationPolicyRoutesInner) SetDestination(v NotificationPolicyDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *NotificationPolicyRoutesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetExpr

`func (o *NotificationPolicyRoutesInner) GetExpr() Expr`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *NotificationPolicyRoutesInner) GetExprOk() (*Expr, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *NotificationPolicyRoutesInner) SetExpr(v Expr)`

SetExpr sets Expr field to given value.

### HasExpr

`func (o *NotificationPolicyRoutesInner) HasExpr() bool`

HasExpr returns a boolean if a field has been set.

### GetTarget

`func (o *NotificationPolicyRoutesInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NotificationPolicyRoutesInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NotificationPolicyRoutesInner) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *NotificationPolicyRoutesInner) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


