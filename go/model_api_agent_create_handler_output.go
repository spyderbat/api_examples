/*
Spyderbat API UI & Public APIs

Restful APIs for use by UI & customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sbapi

import (
	"encoding/json"
)

// ApiAgentCreateHandlerOutput struct for ApiAgentCreateHandlerOutput
type ApiAgentCreateHandlerOutput struct {
	Uid *string `json:"uid,omitempty"`
}

// NewApiAgentCreateHandlerOutput instantiates a new ApiAgentCreateHandlerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAgentCreateHandlerOutput() *ApiAgentCreateHandlerOutput {
	this := ApiAgentCreateHandlerOutput{}
	return &this
}

// NewApiAgentCreateHandlerOutputWithDefaults instantiates a new ApiAgentCreateHandlerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAgentCreateHandlerOutputWithDefaults() *ApiAgentCreateHandlerOutput {
	this := ApiAgentCreateHandlerOutput{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ApiAgentCreateHandlerOutput) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentCreateHandlerOutput) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ApiAgentCreateHandlerOutput) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ApiAgentCreateHandlerOutput) SetUid(v string) {
	o.Uid = &v
}

func (o ApiAgentCreateHandlerOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableApiAgentCreateHandlerOutput struct {
	value *ApiAgentCreateHandlerOutput
	isSet bool
}

func (v NullableApiAgentCreateHandlerOutput) Get() *ApiAgentCreateHandlerOutput {
	return v.value
}

func (v *NullableApiAgentCreateHandlerOutput) Set(val *ApiAgentCreateHandlerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAgentCreateHandlerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAgentCreateHandlerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAgentCreateHandlerOutput(val *ApiAgentCreateHandlerOutput) *NullableApiAgentCreateHandlerOutput {
	return &NullableApiAgentCreateHandlerOutput{value: val, isSet: true}
}

func (v NullableApiAgentCreateHandlerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAgentCreateHandlerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

