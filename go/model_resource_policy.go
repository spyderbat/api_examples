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

// ResourcePolicy Resource policy for RBAC
type ResourcePolicy struct {
	// Name of the resource policy
	Name *string `json:"name,omitempty"`
	// List of statements to be examined by the resource policy
	Statements []RbacStatement `json:"statements,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewResourcePolicy instantiates a new ResourcePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePolicy() *ResourcePolicy {
	this := ResourcePolicy{}
	return &this
}

// NewResourcePolicyWithDefaults instantiates a new ResourcePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePolicyWithDefaults() *ResourcePolicy {
	this := ResourcePolicy{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourcePolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourcePolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourcePolicy) SetName(v string) {
	o.Name = &v
}

// GetStatements returns the Statements field value if set, zero value otherwise.
func (o *ResourcePolicy) GetStatements() []RbacStatement {
	if o == nil || o.Statements == nil {
		var ret []RbacStatement
		return ret
	}
	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicy) GetStatementsOk() ([]RbacStatement, bool) {
	if o == nil || o.Statements == nil {
		return nil, false
	}
	return o.Statements, true
}

// HasStatements returns a boolean if a field has been set.
func (o *ResourcePolicy) HasStatements() bool {
	if o != nil && o.Statements != nil {
		return true
	}

	return false
}

// SetStatements gets a reference to the given []RbacStatement and assigns it to the Statements field.
func (o *ResourcePolicy) SetStatements(v []RbacStatement) {
	o.Statements = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ResourcePolicy) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePolicy) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ResourcePolicy) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ResourcePolicy) SetVersion(v string) {
	o.Version = &v
}

func (o ResourcePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Statements != nil {
		toSerialize["statements"] = o.Statements
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableResourcePolicy struct {
	value *ResourcePolicy
	isSet bool
}

func (v NullableResourcePolicy) Get() *ResourcePolicy {
	return v.value
}

func (v *NullableResourcePolicy) Set(val *ResourcePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePolicy(val *ResourcePolicy) *NullableResourcePolicy {
	return &NullableResourcePolicy{value: val, isSet: true}
}

func (v NullableResourcePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


