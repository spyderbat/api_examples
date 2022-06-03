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

// ApiAgentWorkOutput struct for ApiAgentWorkOutput
type ApiAgentWorkOutput struct {
	// User defined tags
	Tags []string `json:"tags,omitempty"`
	Work *OrcApiAgentWork `json:"work,omitempty"`
}

// NewApiAgentWorkOutput instantiates a new ApiAgentWorkOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAgentWorkOutput() *ApiAgentWorkOutput {
	this := ApiAgentWorkOutput{}
	return &this
}

// NewApiAgentWorkOutputWithDefaults instantiates a new ApiAgentWorkOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAgentWorkOutputWithDefaults() *ApiAgentWorkOutput {
	this := ApiAgentWorkOutput{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApiAgentWorkOutput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentWorkOutput) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApiAgentWorkOutput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApiAgentWorkOutput) SetTags(v []string) {
	o.Tags = v
}

// GetWork returns the Work field value if set, zero value otherwise.
func (o *ApiAgentWorkOutput) GetWork() OrcApiAgentWork {
	if o == nil || o.Work == nil {
		var ret OrcApiAgentWork
		return ret
	}
	return *o.Work
}

// GetWorkOk returns a tuple with the Work field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentWorkOutput) GetWorkOk() (*OrcApiAgentWork, bool) {
	if o == nil || o.Work == nil {
		return nil, false
	}
	return o.Work, true
}

// HasWork returns a boolean if a field has been set.
func (o *ApiAgentWorkOutput) HasWork() bool {
	if o != nil && o.Work != nil {
		return true
	}

	return false
}

// SetWork gets a reference to the given OrcApiAgentWork and assigns it to the Work field.
func (o *ApiAgentWorkOutput) SetWork(v OrcApiAgentWork) {
	o.Work = &v
}

func (o ApiAgentWorkOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Work != nil {
		toSerialize["work"] = o.Work
	}
	return json.Marshal(toSerialize)
}

type NullableApiAgentWorkOutput struct {
	value *ApiAgentWorkOutput
	isSet bool
}

func (v NullableApiAgentWorkOutput) Get() *ApiAgentWorkOutput {
	return v.value
}

func (v *NullableApiAgentWorkOutput) Set(val *ApiAgentWorkOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAgentWorkOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAgentWorkOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAgentWorkOutput(val *ApiAgentWorkOutput) *NullableApiAgentWorkOutput {
	return &NullableApiAgentWorkOutput{value: val, isSet: true}
}

func (v NullableApiAgentWorkOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAgentWorkOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


