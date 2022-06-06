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

// OrcApiAgentWork Agent work, which is sent to the agents
type OrcApiAgentWork struct {
	// Array of bats to execute
	Work []OrcApiBatWork `json:"work,omitempty"`
}

// NewOrcApiAgentWork instantiates a new OrcApiAgentWork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrcApiAgentWork() *OrcApiAgentWork {
	this := OrcApiAgentWork{}
	return &this
}

// NewOrcApiAgentWorkWithDefaults instantiates a new OrcApiAgentWork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrcApiAgentWorkWithDefaults() *OrcApiAgentWork {
	this := OrcApiAgentWork{}
	return &this
}

// GetWork returns the Work field value if set, zero value otherwise.
func (o *OrcApiAgentWork) GetWork() []OrcApiBatWork {
	if o == nil || o.Work == nil {
		var ret []OrcApiBatWork
		return ret
	}
	return o.Work
}

// GetWorkOk returns a tuple with the Work field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrcApiAgentWork) GetWorkOk() ([]OrcApiBatWork, bool) {
	if o == nil || o.Work == nil {
		return nil, false
	}
	return o.Work, true
}

// HasWork returns a boolean if a field has been set.
func (o *OrcApiAgentWork) HasWork() bool {
	if o != nil && o.Work != nil {
		return true
	}

	return false
}

// SetWork gets a reference to the given []OrcApiBatWork and assigns it to the Work field.
func (o *OrcApiAgentWork) SetWork(v []OrcApiBatWork) {
	o.Work = v
}

func (o OrcApiAgentWork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Work != nil {
		toSerialize["work"] = o.Work
	}
	return json.Marshal(toSerialize)
}

type NullableOrcApiAgentWork struct {
	value *OrcApiAgentWork
	isSet bool
}

func (v NullableOrcApiAgentWork) Get() *OrcApiAgentWork {
	return v.value
}

func (v *NullableOrcApiAgentWork) Set(val *OrcApiAgentWork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrcApiAgentWork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrcApiAgentWork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrcApiAgentWork(val *OrcApiAgentWork) *NullableOrcApiAgentWork {
	return &NullableOrcApiAgentWork{value: val, isSet: true}
}

func (v NullableOrcApiAgentWork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrcApiAgentWork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

