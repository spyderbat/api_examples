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

// NotificationPolicyDestinationWebhook struct for NotificationPolicyDestinationWebhook
type NotificationPolicyDestinationWebhook struct {
	NoTlsValidation *bool `json:"no_tls_validation,omitempty"`
	Url string `json:"url"`
}

// NewNotificationPolicyDestinationWebhook instantiates a new NotificationPolicyDestinationWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationPolicyDestinationWebhook(url string) *NotificationPolicyDestinationWebhook {
	this := NotificationPolicyDestinationWebhook{}
	this.Url = url
	return &this
}

// NewNotificationPolicyDestinationWebhookWithDefaults instantiates a new NotificationPolicyDestinationWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationPolicyDestinationWebhookWithDefaults() *NotificationPolicyDestinationWebhook {
	this := NotificationPolicyDestinationWebhook{}
	return &this
}

// GetNoTlsValidation returns the NoTlsValidation field value if set, zero value otherwise.
func (o *NotificationPolicyDestinationWebhook) GetNoTlsValidation() bool {
	if o == nil || o.NoTlsValidation == nil {
		var ret bool
		return ret
	}
	return *o.NoTlsValidation
}

// GetNoTlsValidationOk returns a tuple with the NoTlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestinationWebhook) GetNoTlsValidationOk() (*bool, bool) {
	if o == nil || o.NoTlsValidation == nil {
		return nil, false
	}
	return o.NoTlsValidation, true
}

// HasNoTlsValidation returns a boolean if a field has been set.
func (o *NotificationPolicyDestinationWebhook) HasNoTlsValidation() bool {
	if o != nil && o.NoTlsValidation != nil {
		return true
	}

	return false
}

// SetNoTlsValidation gets a reference to the given bool and assigns it to the NoTlsValidation field.
func (o *NotificationPolicyDestinationWebhook) SetNoTlsValidation(v bool) {
	o.NoTlsValidation = &v
}

// GetUrl returns the Url field value
func (o *NotificationPolicyDestinationWebhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationPolicyDestinationWebhook) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationPolicyDestinationWebhook) SetUrl(v string) {
	o.Url = v
}

func (o NotificationPolicyDestinationWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NoTlsValidation != nil {
		toSerialize["no_tls_validation"] = o.NoTlsValidation
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationPolicyDestinationWebhook struct {
	value *NotificationPolicyDestinationWebhook
	isSet bool
}

func (v NullableNotificationPolicyDestinationWebhook) Get() *NotificationPolicyDestinationWebhook {
	return v.value
}

func (v *NullableNotificationPolicyDestinationWebhook) Set(val *NotificationPolicyDestinationWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationPolicyDestinationWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationPolicyDestinationWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationPolicyDestinationWebhook(val *NotificationPolicyDestinationWebhook) *NullableNotificationPolicyDestinationWebhook {
	return &NullableNotificationPolicyDestinationWebhook{value: val, isSet: true}
}

func (v NullableNotificationPolicyDestinationWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationPolicyDestinationWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


