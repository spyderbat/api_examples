/*
Spyderbat API UI & Public APIs

Restful APIs for use by UI & customers.

API version: 1.0.0
Contact: apisupport@spyderbat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spyderbat_api

import (
	"encoding/json"
)

// CanUserPerformInput struct for CanUserPerformInput
type CanUserPerformInput struct {
	Actions []RBACAction `json:"actions,omitempty"`
}

// NewCanUserPerformInput instantiates a new CanUserPerformInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanUserPerformInput() *CanUserPerformInput {
	this := CanUserPerformInput{}
	return &this
}

// NewCanUserPerformInputWithDefaults instantiates a new CanUserPerformInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanUserPerformInputWithDefaults() *CanUserPerformInput {
	this := CanUserPerformInput{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CanUserPerformInput) GetActions() []RBACAction {
	if o == nil || o.Actions == nil {
		var ret []RBACAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanUserPerformInput) GetActionsOk() ([]RBACAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CanUserPerformInput) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []RBACAction and assigns it to the Actions field.
func (o *CanUserPerformInput) SetActions(v []RBACAction) {
	o.Actions = v
}

func (o CanUserPerformInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableCanUserPerformInput struct {
	value *CanUserPerformInput
	isSet bool
}

func (v NullableCanUserPerformInput) Get() *CanUserPerformInput {
	return v.value
}

func (v *NullableCanUserPerformInput) Set(val *CanUserPerformInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCanUserPerformInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCanUserPerformInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanUserPerformInput(val *CanUserPerformInput) *NullableCanUserPerformInput {
	return &NullableCanUserPerformInput{value: val, isSet: true}
}

func (v NullableCanUserPerformInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanUserPerformInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


