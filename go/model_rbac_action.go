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

// RBACAction struct for RBACAction
type RBACAction struct {
	// Action which meets the requirements of RBAC action naming
	Action *string `json:"action,omitempty"`
	// Return result of querying the users RBAC capabilities
	CanPerform *bool `json:"can_perform,omitempty"`
	// Error returned from permission checking
	Error *string `json:"error,omitempty"`
	// ResourceName which meets the requirements of RBAC resource naming
	ResourceName *string `json:"resource_name,omitempty"`
}

// NewRBACAction instantiates a new RBACAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRBACAction() *RBACAction {
	this := RBACAction{}
	return &this
}

// NewRBACActionWithDefaults instantiates a new RBACAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRBACActionWithDefaults() *RBACAction {
	this := RBACAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RBACAction) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBACAction) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RBACAction) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RBACAction) SetAction(v string) {
	o.Action = &v
}

// GetCanPerform returns the CanPerform field value if set, zero value otherwise.
func (o *RBACAction) GetCanPerform() bool {
	if o == nil || o.CanPerform == nil {
		var ret bool
		return ret
	}
	return *o.CanPerform
}

// GetCanPerformOk returns a tuple with the CanPerform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBACAction) GetCanPerformOk() (*bool, bool) {
	if o == nil || o.CanPerform == nil {
		return nil, false
	}
	return o.CanPerform, true
}

// HasCanPerform returns a boolean if a field has been set.
func (o *RBACAction) HasCanPerform() bool {
	if o != nil && o.CanPerform != nil {
		return true
	}

	return false
}

// SetCanPerform gets a reference to the given bool and assigns it to the CanPerform field.
func (o *RBACAction) SetCanPerform(v bool) {
	o.CanPerform = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RBACAction) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBACAction) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RBACAction) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RBACAction) SetError(v string) {
	o.Error = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *RBACAction) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBACAction) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *RBACAction) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *RBACAction) SetResourceName(v string) {
	o.ResourceName = &v
}

func (o RBACAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.CanPerform != nil {
		toSerialize["can_perform"] = o.CanPerform
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	return json.Marshal(toSerialize)
}

type NullableRBACAction struct {
	value *RBACAction
	isSet bool
}

func (v NullableRBACAction) Get() *RBACAction {
	return v.value
}

func (v *NullableRBACAction) Set(val *RBACAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRBACAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRBACAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRBACAction(val *RBACAction) *NullableRBACAction {
	return &NullableRBACAction{value: val, isSet: true}
}

func (v NullableRBACAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRBACAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

