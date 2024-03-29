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

// AgentSetAgentWorkInput struct for AgentSetAgentWorkInput
type AgentSetAgentWorkInput struct {
	// User defined tags
	Tags []string `json:"tags,omitempty"`
	Work *OrcApiAgentWork `json:"work,omitempty"`
}

// NewAgentSetAgentWorkInput instantiates a new AgentSetAgentWorkInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentSetAgentWorkInput() *AgentSetAgentWorkInput {
	this := AgentSetAgentWorkInput{}
	return &this
}

// NewAgentSetAgentWorkInputWithDefaults instantiates a new AgentSetAgentWorkInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentSetAgentWorkInputWithDefaults() *AgentSetAgentWorkInput {
	this := AgentSetAgentWorkInput{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AgentSetAgentWorkInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSetAgentWorkInput) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AgentSetAgentWorkInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AgentSetAgentWorkInput) SetTags(v []string) {
	o.Tags = v
}

// GetWork returns the Work field value if set, zero value otherwise.
func (o *AgentSetAgentWorkInput) GetWork() OrcApiAgentWork {
	if o == nil || o.Work == nil {
		var ret OrcApiAgentWork
		return ret
	}
	return *o.Work
}

// GetWorkOk returns a tuple with the Work field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSetAgentWorkInput) GetWorkOk() (*OrcApiAgentWork, bool) {
	if o == nil || o.Work == nil {
		return nil, false
	}
	return o.Work, true
}

// HasWork returns a boolean if a field has been set.
func (o *AgentSetAgentWorkInput) HasWork() bool {
	if o != nil && o.Work != nil {
		return true
	}

	return false
}

// SetWork gets a reference to the given OrcApiAgentWork and assigns it to the Work field.
func (o *AgentSetAgentWorkInput) SetWork(v OrcApiAgentWork) {
	o.Work = &v
}

func (o AgentSetAgentWorkInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Work != nil {
		toSerialize["work"] = o.Work
	}
	return json.Marshal(toSerialize)
}

type NullableAgentSetAgentWorkInput struct {
	value *AgentSetAgentWorkInput
	isSet bool
}

func (v NullableAgentSetAgentWorkInput) Get() *AgentSetAgentWorkInput {
	return v.value
}

func (v *NullableAgentSetAgentWorkInput) Set(val *AgentSetAgentWorkInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSetAgentWorkInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSetAgentWorkInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSetAgentWorkInput(val *AgentSetAgentWorkInput) *NullableAgentSetAgentWorkInput {
	return &NullableAgentSetAgentWorkInput{value: val, isSet: true}
}

func (v NullableAgentSetAgentWorkInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSetAgentWorkInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


