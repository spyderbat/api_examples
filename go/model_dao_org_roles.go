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

// DaoOrgRoles struct for DaoOrgRoles
type DaoOrgRoles struct {
	// Default roles for the user
	DefaultRoles []string `json:"default_roles,omitempty"`
	// Org UID
	OrgUid *string `json:"org_uid,omitempty"`
}

// NewDaoOrgRoles instantiates a new DaoOrgRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaoOrgRoles() *DaoOrgRoles {
	this := DaoOrgRoles{}
	return &this
}

// NewDaoOrgRolesWithDefaults instantiates a new DaoOrgRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaoOrgRolesWithDefaults() *DaoOrgRoles {
	this := DaoOrgRoles{}
	return &this
}

// GetDefaultRoles returns the DefaultRoles field value if set, zero value otherwise.
func (o *DaoOrgRoles) GetDefaultRoles() []string {
	if o == nil || o.DefaultRoles == nil {
		var ret []string
		return ret
	}
	return o.DefaultRoles
}

// GetDefaultRolesOk returns a tuple with the DefaultRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoOrgRoles) GetDefaultRolesOk() ([]string, bool) {
	if o == nil || o.DefaultRoles == nil {
		return nil, false
	}
	return o.DefaultRoles, true
}

// HasDefaultRoles returns a boolean if a field has been set.
func (o *DaoOrgRoles) HasDefaultRoles() bool {
	if o != nil && o.DefaultRoles != nil {
		return true
	}

	return false
}

// SetDefaultRoles gets a reference to the given []string and assigns it to the DefaultRoles field.
func (o *DaoOrgRoles) SetDefaultRoles(v []string) {
	o.DefaultRoles = v
}

// GetOrgUid returns the OrgUid field value if set, zero value otherwise.
func (o *DaoOrgRoles) GetOrgUid() string {
	if o == nil || o.OrgUid == nil {
		var ret string
		return ret
	}
	return *o.OrgUid
}

// GetOrgUidOk returns a tuple with the OrgUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoOrgRoles) GetOrgUidOk() (*string, bool) {
	if o == nil || o.OrgUid == nil {
		return nil, false
	}
	return o.OrgUid, true
}

// HasOrgUid returns a boolean if a field has been set.
func (o *DaoOrgRoles) HasOrgUid() bool {
	if o != nil && o.OrgUid != nil {
		return true
	}

	return false
}

// SetOrgUid gets a reference to the given string and assigns it to the OrgUid field.
func (o *DaoOrgRoles) SetOrgUid(v string) {
	o.OrgUid = &v
}

func (o DaoOrgRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultRoles != nil {
		toSerialize["default_roles"] = o.DefaultRoles
	}
	if o.OrgUid != nil {
		toSerialize["org_uid"] = o.OrgUid
	}
	return json.Marshal(toSerialize)
}

type NullableDaoOrgRoles struct {
	value *DaoOrgRoles
	isSet bool
}

func (v NullableDaoOrgRoles) Get() *DaoOrgRoles {
	return v.value
}

func (v *NullableDaoOrgRoles) Set(val *DaoOrgRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableDaoOrgRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableDaoOrgRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaoOrgRoles(val *DaoOrgRoles) *NullableDaoOrgRoles {
	return &NullableDaoOrgRoles{value: val, isSet: true}
}

func (v NullableDaoOrgRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaoOrgRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


