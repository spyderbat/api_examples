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
	"time"
)

// UiDataSetUserSourceDataInput struct for UiDataSetUserSourceDataInput
type UiDataSetUserSourceDataInput struct {
	 *string `json:",omitempty"`
	// UI supplied JSON object
	Data map[string]interface{} `json:"data,omitempty"`
	// The time of las use of the key; only updated every 5 minutes
	LastUsed *time.Time `json:"last_used,omitempty"`
	// User supplied tags
	Tags []string `json:"tags,omitempty"`
	// UID for the UIData
	Uid *string `json:"uid,omitempty"`
	// Valid from is the creation date or first date when the API key became valid
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	// Valid to is the expiration date or the last date this API key will be valid
	ValidTo *time.Time `json:"valid_to,omitempty"`
}

// NewUiDataSetUserSourceDataInput instantiates a new UiDataSetUserSourceDataInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiDataSetUserSourceDataInput() *UiDataSetUserSourceDataInput {
	this := UiDataSetUserSourceDataInput{}
	return &this
}

// NewUiDataSetUserSourceDataInputWithDefaults instantiates a new UiDataSetUserSourceDataInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiDataSetUserSourceDataInputWithDefaults() *UiDataSetUserSourceDataInput {
	this := UiDataSetUserSourceDataInput{}
	return &this
}

// Get returns the  field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) Get() string {
	if o == nil || o. == nil {
		var ret string
		return ret
	}
	return *o.
}

// GetOk returns a tuple with the  field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetOk() (*string, bool) {
	if o == nil || o. == nil {
		return nil, false
	}
	return o., true
}

// Has returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) Has() bool {
	if o != nil && o. != nil {
		return true
	}

	return false
}

// Set gets a reference to the given string and assigns it to the  field.
func (o *UiDataSetUserSourceDataInput) Set(v string) {
	o. = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *UiDataSetUserSourceDataInput) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) GetLastUsed() time.Time {
	if o == nil || o.LastUsed == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetLastUsedOk() (*time.Time, bool) {
	if o == nil || o.LastUsed == nil {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) HasLastUsed() bool {
	if o != nil && o.LastUsed != nil {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given time.Time and assigns it to the LastUsed field.
func (o *UiDataSetUserSourceDataInput) SetLastUsed(v time.Time) {
	o.LastUsed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UiDataSetUserSourceDataInput) SetTags(v []string) {
	o.Tags = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *UiDataSetUserSourceDataInput) SetUid(v string) {
	o.Uid = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *UiDataSetUserSourceDataInput) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *UiDataSetUserSourceDataInput) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiDataSetUserSourceDataInput) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *UiDataSetUserSourceDataInput) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *UiDataSetUserSourceDataInput) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

func (o UiDataSetUserSourceDataInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o. != nil {
		toSerialize[""] = o.
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.LastUsed != nil {
		toSerialize["last_used"] = o.LastUsed
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

type NullableUiDataSetUserSourceDataInput struct {
	value *UiDataSetUserSourceDataInput
	isSet bool
}

func (v NullableUiDataSetUserSourceDataInput) Get() *UiDataSetUserSourceDataInput {
	return v.value
}

func (v *NullableUiDataSetUserSourceDataInput) Set(val *UiDataSetUserSourceDataInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUiDataSetUserSourceDataInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUiDataSetUserSourceDataInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiDataSetUserSourceDataInput(val *UiDataSetUserSourceDataInput) *NullableUiDataSetUserSourceDataInput {
	return &NullableUiDataSetUserSourceDataInput{value: val, isSet: true}
}

func (v NullableUiDataSetUserSourceDataInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiDataSetUserSourceDataInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


