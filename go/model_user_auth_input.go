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

// UserAuthInput struct for UserAuthInput
type UserAuthInput struct {
	// User email
	Email string `json:"email"`
	Password string `json:"password"`
}

// NewUserAuthInput instantiates a new UserAuthInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuthInput(email string, password string) *UserAuthInput {
	this := UserAuthInput{}
	this.Email = email
	this.Password = password
	return &this
}

// NewUserAuthInputWithDefaults instantiates a new UserAuthInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthInputWithDefaults() *UserAuthInput {
	this := UserAuthInput{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserAuthInput) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserAuthInput) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserAuthInput) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *UserAuthInput) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserAuthInput) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserAuthInput) SetPassword(v string) {
	o.Password = v
}

func (o UserAuthInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserAuthInput struct {
	value *UserAuthInput
	isSet bool
}

func (v NullableUserAuthInput) Get() *UserAuthInput {
	return v.value
}

func (v *NullableUserAuthInput) Set(val *UserAuthInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthInput(val *UserAuthInput) *NullableUserAuthInput {
	return &NullableUserAuthInput{value: val, isSet: true}
}

func (v NullableUserAuthInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

