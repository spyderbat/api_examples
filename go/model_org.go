/*
Spyderbat API UI & Public APIs

Restful APIs for use by UI & customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sbapi

import (
	"encoding/json"
	"time"
)

// Org struct for Org
type Org struct {
	// Name of the organization
	Name string `json:"name"`
	// Organization Type
	OrgTypeUid *string `json:"org_type_uid,omitempty"`
	// The email address of the user who owns this org
	OwnerEmail *string `json:"owner_email,omitempty"`
	// The user UID who owns this organization
	OwnerUid *string `json:"owner_uid,omitempty"`
	// Resource name utilized by RBAC
	ResourceName *string `json:"resource_name,omitempty"`
	ResourcePolicy *ResourcePolicy `json:"resource_policy,omitempty"`
	// User supplied tags
	Tags []string `json:"tags,omitempty"`
	// Org UID
	Uid *string `json:"uid,omitempty"`
	// Valid from date, the first date this object was valid
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	// Valid to date, the date this object is valid to
	ValidTo *time.Time `json:"valid_to,omitempty"`
}

// NewOrg instantiates a new Org object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrg(name string) *Org {
	this := Org{}
	this.Name = name
	return &this
}

// NewOrgWithDefaults instantiates a new Org object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgWithDefaults() *Org {
	this := Org{}
	return &this
}

// GetName returns the Name field value
func (o *Org) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Org) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Org) SetName(v string) {
	o.Name = v
}

// GetOrgTypeUid returns the OrgTypeUid field value if set, zero value otherwise.
func (o *Org) GetOrgTypeUid() string {
	if o == nil || o.OrgTypeUid == nil {
		var ret string
		return ret
	}
	return *o.OrgTypeUid
}

// GetOrgTypeUidOk returns a tuple with the OrgTypeUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetOrgTypeUidOk() (*string, bool) {
	if o == nil || o.OrgTypeUid == nil {
		return nil, false
	}
	return o.OrgTypeUid, true
}

// HasOrgTypeUid returns a boolean if a field has been set.
func (o *Org) HasOrgTypeUid() bool {
	if o != nil && o.OrgTypeUid != nil {
		return true
	}

	return false
}

// SetOrgTypeUid gets a reference to the given string and assigns it to the OrgTypeUid field.
func (o *Org) SetOrgTypeUid(v string) {
	o.OrgTypeUid = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *Org) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *Org) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *Org) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetOwnerUid returns the OwnerUid field value if set, zero value otherwise.
func (o *Org) GetOwnerUid() string {
	if o == nil || o.OwnerUid == nil {
		var ret string
		return ret
	}
	return *o.OwnerUid
}

// GetOwnerUidOk returns a tuple with the OwnerUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetOwnerUidOk() (*string, bool) {
	if o == nil || o.OwnerUid == nil {
		return nil, false
	}
	return o.OwnerUid, true
}

// HasOwnerUid returns a boolean if a field has been set.
func (o *Org) HasOwnerUid() bool {
	if o != nil && o.OwnerUid != nil {
		return true
	}

	return false
}

// SetOwnerUid gets a reference to the given string and assigns it to the OwnerUid field.
func (o *Org) SetOwnerUid(v string) {
	o.OwnerUid = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *Org) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *Org) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *Org) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourcePolicy returns the ResourcePolicy field value if set, zero value otherwise.
func (o *Org) GetResourcePolicy() ResourcePolicy {
	if o == nil || o.ResourcePolicy == nil {
		var ret ResourcePolicy
		return ret
	}
	return *o.ResourcePolicy
}

// GetResourcePolicyOk returns a tuple with the ResourcePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetResourcePolicyOk() (*ResourcePolicy, bool) {
	if o == nil || o.ResourcePolicy == nil {
		return nil, false
	}
	return o.ResourcePolicy, true
}

// HasResourcePolicy returns a boolean if a field has been set.
func (o *Org) HasResourcePolicy() bool {
	if o != nil && o.ResourcePolicy != nil {
		return true
	}

	return false
}

// SetResourcePolicy gets a reference to the given ResourcePolicy and assigns it to the ResourcePolicy field.
func (o *Org) SetResourcePolicy(v ResourcePolicy) {
	o.ResourcePolicy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Org) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Org) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Org) SetTags(v []string) {
	o.Tags = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *Org) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *Org) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *Org) SetUid(v string) {
	o.Uid = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *Org) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *Org) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *Org) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *Org) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *Org) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *Org) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

func (o Org) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OrgTypeUid != nil {
		toSerialize["org_type_uid"] = o.OrgTypeUid
	}
	if o.OwnerEmail != nil {
		toSerialize["owner_email"] = o.OwnerEmail
	}
	if o.OwnerUid != nil {
		toSerialize["owner_uid"] = o.OwnerUid
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourcePolicy != nil {
		toSerialize["resource_policy"] = o.ResourcePolicy
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.ValidFrom != nil {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["valid_to"] = o.ValidTo
	}
	return json.Marshal(toSerialize)
}

type NullableOrg struct {
	value *Org
	isSet bool
}

func (v NullableOrg) Get() *Org {
	return v.value
}

func (v *NullableOrg) Set(val *Org) {
	v.value = val
	v.isSet = true
}

func (v NullableOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrg(val *Org) *NullableOrg {
	return &NullableOrg{value: val, isSet: true}
}

func (v NullableOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

