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

// ValidationError struct for ValidationError
type ValidationError struct {
	// Message regarding the validation failure
	ErrMsg *string `json:"err_msg,omitempty"`
	// Field name which failed validation
	Field *string `json:"field,omitempty"`
	// JSON property name of the field which failed validation
	Property *string `json:"property,omitempty"`
	// Validation tag which failed
	Tags *string `json:"tags,omitempty"`
}

// NewValidationError instantiates a new ValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationError() *ValidationError {
	this := ValidationError{}
	return &this
}

// NewValidationErrorWithDefaults instantiates a new ValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorWithDefaults() *ValidationError {
	this := ValidationError{}
	return &this
}

// GetErrMsg returns the ErrMsg field value if set, zero value otherwise.
func (o *ValidationError) GetErrMsg() string {
	if o == nil || o.ErrMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrMsg
}

// GetErrMsgOk returns a tuple with the ErrMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetErrMsgOk() (*string, bool) {
	if o == nil || o.ErrMsg == nil {
		return nil, false
	}
	return o.ErrMsg, true
}

// HasErrMsg returns a boolean if a field has been set.
func (o *ValidationError) HasErrMsg() bool {
	if o != nil && o.ErrMsg != nil {
		return true
	}

	return false
}

// SetErrMsg gets a reference to the given string and assigns it to the ErrMsg field.
func (o *ValidationError) SetErrMsg(v string) {
	o.ErrMsg = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ValidationError) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ValidationError) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ValidationError) SetField(v string) {
	o.Field = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ValidationError) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ValidationError) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *ValidationError) SetProperty(v string) {
	o.Property = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ValidationError) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ValidationError) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ValidationError) SetTags(v string) {
	o.Tags = &v
}

func (o ValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrMsg != nil {
		toSerialize["err_msg"] = o.ErrMsg
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableValidationError struct {
	value *ValidationError
	isSet bool
}

func (v NullableValidationError) Get() *ValidationError {
	return v.value
}

func (v *NullableValidationError) Set(val *ValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationError(val *ValidationError) *NullableValidationError {
	return &NullableValidationError{value: val, isSet: true}
}

func (v NullableValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


