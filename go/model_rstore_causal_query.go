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

// RstoreCausalQuery struct for RstoreCausalQuery
type RstoreCausalQuery struct {
	// Causal tree matching rules for peer causal records
	Peer []map[string]interface{} `json:"peer,omitempty"`
	// Causal tree matching rules for post causal records
	Post []map[string]interface{} `json:"post,omitempty"`
	// Causal tree matching rules for pre causal records
	Pre []map[string]interface{} `json:"pre,omitempty"`
}

// NewRstoreCausalQuery instantiates a new RstoreCausalQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRstoreCausalQuery() *RstoreCausalQuery {
	this := RstoreCausalQuery{}
	return &this
}

// NewRstoreCausalQueryWithDefaults instantiates a new RstoreCausalQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRstoreCausalQueryWithDefaults() *RstoreCausalQuery {
	this := RstoreCausalQuery{}
	return &this
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *RstoreCausalQuery) GetPeer() []map[string]interface{} {
	if o == nil || o.Peer == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RstoreCausalQuery) GetPeerOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *RstoreCausalQuery) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given []map[string]interface{} and assigns it to the Peer field.
func (o *RstoreCausalQuery) SetPeer(v []map[string]interface{}) {
	o.Peer = v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *RstoreCausalQuery) GetPost() []map[string]interface{} {
	if o == nil || o.Post == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RstoreCausalQuery) GetPostOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *RstoreCausalQuery) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given []map[string]interface{} and assigns it to the Post field.
func (o *RstoreCausalQuery) SetPost(v []map[string]interface{}) {
	o.Post = v
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *RstoreCausalQuery) GetPre() []map[string]interface{} {
	if o == nil || o.Pre == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RstoreCausalQuery) GetPreOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Pre == nil {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *RstoreCausalQuery) HasPre() bool {
	if o != nil && o.Pre != nil {
		return true
	}

	return false
}

// SetPre gets a reference to the given []map[string]interface{} and assigns it to the Pre field.
func (o *RstoreCausalQuery) SetPre(v []map[string]interface{}) {
	o.Pre = v
}

func (o RstoreCausalQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Peer != nil {
		toSerialize["peer"] = o.Peer
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	if o.Pre != nil {
		toSerialize["pre"] = o.Pre
	}
	return json.Marshal(toSerialize)
}

type NullableRstoreCausalQuery struct {
	value *RstoreCausalQuery
	isSet bool
}

func (v NullableRstoreCausalQuery) Get() *RstoreCausalQuery {
	return v.value
}

func (v *NullableRstoreCausalQuery) Set(val *RstoreCausalQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRstoreCausalQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRstoreCausalQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRstoreCausalQuery(val *RstoreCausalQuery) *NullableRstoreCausalQuery {
	return &NullableRstoreCausalQuery{value: val, isSet: true}
}

func (v NullableRstoreCausalQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRstoreCausalQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


