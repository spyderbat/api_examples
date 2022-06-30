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

// DaoAgentClass struct for DaoAgentClass
type DaoAgentClass struct {
	Name string `json:"name"`
	RbacRoles []string `json:"rbac_roles,omitempty"`
}

// NewDaoAgentClass instantiates a new DaoAgentClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaoAgentClass(name string) *DaoAgentClass {
	this := DaoAgentClass{}
	this.Name = name
	return &this
}

// NewDaoAgentClassWithDefaults instantiates a new DaoAgentClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaoAgentClassWithDefaults() *DaoAgentClass {
	this := DaoAgentClass{}
	return &this
}

// GetName returns the Name field value
func (o *DaoAgentClass) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DaoAgentClass) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DaoAgentClass) SetName(v string) {
	o.Name = v
}

// GetRbacRoles returns the RbacRoles field value if set, zero value otherwise.
func (o *DaoAgentClass) GetRbacRoles() []string {
	if o == nil || o.RbacRoles == nil {
		var ret []string
		return ret
	}
	return o.RbacRoles
}

// GetRbacRolesOk returns a tuple with the RbacRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentClass) GetRbacRolesOk() ([]string, bool) {
	if o == nil || o.RbacRoles == nil {
		return nil, false
	}
	return o.RbacRoles, true
}

// HasRbacRoles returns a boolean if a field has been set.
func (o *DaoAgentClass) HasRbacRoles() bool {
	if o != nil && o.RbacRoles != nil {
		return true
	}

	return false
}

// SetRbacRoles gets a reference to the given []string and assigns it to the RbacRoles field.
func (o *DaoAgentClass) SetRbacRoles(v []string) {
	o.RbacRoles = v
}

func (o DaoAgentClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RbacRoles != nil {
		toSerialize["rbac_roles"] = o.RbacRoles
	}
	return json.Marshal(toSerialize)
}

type NullableDaoAgentClass struct {
	value *DaoAgentClass
	isSet bool
}

func (v NullableDaoAgentClass) Get() *DaoAgentClass {
	return v.value
}

func (v *NullableDaoAgentClass) Set(val *DaoAgentClass) {
	v.value = val
	v.isSet = true
}

func (v NullableDaoAgentClass) IsSet() bool {
	return v.isSet
}

func (v *NullableDaoAgentClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaoAgentClass(val *DaoAgentClass) *NullableDaoAgentClass {
	return &NullableDaoAgentClass{value: val, isSet: true}
}

func (v NullableDaoAgentClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaoAgentClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


