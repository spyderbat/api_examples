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

// DaoOrgType struct for DaoOrgType
type DaoOrgType struct {
	// Type of organization
	Description *string `json:"description,omitempty"`
	OrgUid *string `json:"org_uid,omitempty"`
	Policy *DaoOrgTypePolicy `json:"policy,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewDaoOrgType instantiates a new DaoOrgType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaoOrgType() *DaoOrgType {
	this := DaoOrgType{}
	return &this
}

// NewDaoOrgTypeWithDefaults instantiates a new DaoOrgType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaoOrgTypeWithDefaults() *DaoOrgType {
	this := DaoOrgType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DaoOrgType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoOrgType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DaoOrgType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DaoOrgType) SetDescription(v string) {
	o.Description = &v
}

// GetOrgUid returns the OrgUid field value if set, zero value otherwise.
func (o *DaoOrgType) GetOrgUid() string {
	if o == nil || o.OrgUid == nil {
		var ret string
		return ret
	}
	return *o.OrgUid
}

// GetOrgUidOk returns a tuple with the OrgUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoOrgType) GetOrgUidOk() (*string, bool) {
	if o == nil || o.OrgUid == nil {
		return nil, false
	}
	return o.OrgUid, true
}

// HasOrgUid returns a boolean if a field has been set.
func (o *DaoOrgType) HasOrgUid() bool {
	if o != nil && o.OrgUid != nil {
		return true
	}

	return false
}

// SetOrgUid gets a reference to the given string and assigns it to the OrgUid field.
func (o *DaoOrgType) SetOrgUid(v string) {
	o.OrgUid = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *DaoOrgType) GetPolicy() DaoOrgTypePolicy {
	if o == nil || o.Policy == nil {
		var ret DaoOrgTypePolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoOrgType) GetPolicyOk() (*DaoOrgTypePolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *DaoOrgType) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given DaoOrgTypePolicy and assigns it to the Policy field.
func (o *DaoOrgType) SetPolicy(v DaoOrgTypePolicy) {
	o.Policy = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *DaoOrgType) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoOrgType) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *DaoOrgType) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *DaoOrgType) SetUid(v string) {
	o.Uid = &v
}

func (o DaoOrgType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OrgUid != nil {
		toSerialize["org_uid"] = o.OrgUid
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableDaoOrgType struct {
	value *DaoOrgType
	isSet bool
}

func (v NullableDaoOrgType) Get() *DaoOrgType {
	return v.value
}

func (v *NullableDaoOrgType) Set(val *DaoOrgType) {
	v.value = val
	v.isSet = true
}

func (v NullableDaoOrgType) IsSet() bool {
	return v.isSet
}

func (v *NullableDaoOrgType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaoOrgType(val *DaoOrgType) *NullableDaoOrgType {
	return &NullableDaoOrgType{value: val, isSet: true}
}

func (v NullableDaoOrgType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaoOrgType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

