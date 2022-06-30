/*
Spyderbat API UI & Public APIs

Restful APIs for use by UI & customers.

API version: 1.0.0
Contact: apisupport@spyderbat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sbapi

import (
	"encoding/json"
)

// NotificationPolicyDestination A notification policy destination, containing one and only one of the available types
type NotificationPolicyDestination struct {
	// UI-supplied data
	Data interface{} `json:"data,omitempty"`
	Description *string `json:"description,omitempty"`
	Email []string `json:"email,omitempty"`
	OrgUid *string `json:"org_uid,omitempty"`
	Slack *NotificationPolicyDestinationSlack `json:"slack,omitempty"`
	Users []string `json:"users,omitempty"`
	Webhook *NotificationPolicyDestinationWebhook `json:"webhook,omitempty"`
}

// NewNotificationPolicyDestination instantiates a new NotificationPolicyDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationPolicyDestination() *NotificationPolicyDestination {
	this := NotificationPolicyDestination{}
	return &this
}

// NewNotificationPolicyDestinationWithDefaults instantiates a new NotificationPolicyDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationPolicyDestinationWithDefaults() *NotificationPolicyDestination {
	this := NotificationPolicyDestination{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationPolicyDestination) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationPolicyDestination) GetDataOk() (*interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *NotificationPolicyDestination) SetData(v interface{}) {
	o.Data = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationPolicyDestination) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestination) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationPolicyDestination) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NotificationPolicyDestination) GetEmail() []string {
	if o == nil || o.Email == nil {
		var ret []string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestination) GetEmailOk() ([]string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given []string and assigns it to the Email field.
func (o *NotificationPolicyDestination) SetEmail(v []string) {
	o.Email = v
}

// GetOrgUid returns the OrgUid field value if set, zero value otherwise.
func (o *NotificationPolicyDestination) GetOrgUid() string {
	if o == nil || o.OrgUid == nil {
		var ret string
		return ret
	}
	return *o.OrgUid
}

// GetOrgUidOk returns a tuple with the OrgUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestination) GetOrgUidOk() (*string, bool) {
	if o == nil || o.OrgUid == nil {
		return nil, false
	}
	return o.OrgUid, true
}

// HasOrgUid returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasOrgUid() bool {
	if o != nil && o.OrgUid != nil {
		return true
	}

	return false
}

// SetOrgUid gets a reference to the given string and assigns it to the OrgUid field.
func (o *NotificationPolicyDestination) SetOrgUid(v string) {
	o.OrgUid = &v
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *NotificationPolicyDestination) GetSlack() NotificationPolicyDestinationSlack {
	if o == nil || o.Slack == nil {
		var ret NotificationPolicyDestinationSlack
		return ret
	}
	return *o.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestination) GetSlackOk() (*NotificationPolicyDestinationSlack, bool) {
	if o == nil || o.Slack == nil {
		return nil, false
	}
	return o.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasSlack() bool {
	if o != nil && o.Slack != nil {
		return true
	}

	return false
}

// SetSlack gets a reference to the given NotificationPolicyDestinationSlack and assigns it to the Slack field.
func (o *NotificationPolicyDestination) SetSlack(v NotificationPolicyDestinationSlack) {
	o.Slack = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *NotificationPolicyDestination) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestination) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *NotificationPolicyDestination) SetUsers(v []string) {
	o.Users = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *NotificationPolicyDestination) GetWebhook() NotificationPolicyDestinationWebhook {
	if o == nil || o.Webhook == nil {
		var ret NotificationPolicyDestinationWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestination) GetWebhookOk() (*NotificationPolicyDestinationWebhook, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *NotificationPolicyDestination) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NotificationPolicyDestinationWebhook and assigns it to the Webhook field.
func (o *NotificationPolicyDestination) SetWebhook(v NotificationPolicyDestinationWebhook) {
	o.Webhook = &v
}

func (o NotificationPolicyDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.OrgUid != nil {
		toSerialize["org_uid"] = o.OrgUid
	}
	if o.Slack != nil {
		toSerialize["slack"] = o.Slack
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationPolicyDestination struct {
	value *NotificationPolicyDestination
	isSet bool
}

func (v NullableNotificationPolicyDestination) Get() *NotificationPolicyDestination {
	return v.value
}

func (v *NullableNotificationPolicyDestination) Set(val *NotificationPolicyDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationPolicyDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationPolicyDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationPolicyDestination(val *NotificationPolicyDestination) *NullableNotificationPolicyDestination {
	return &NullableNotificationPolicyDestination{value: val, isSet: true}
}

func (v NullableNotificationPolicyDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationPolicyDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


