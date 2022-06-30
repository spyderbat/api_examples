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

// RstreamTimeRange returns true if the property is a time and is start<=time<end
type RstreamTimeRange struct {
	EndTime *float64 `json:"end_time,omitempty"`
	StartTime *float64 `json:"start_time,omitempty"`
}

// NewRstreamTimeRange instantiates a new RstreamTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRstreamTimeRange() *RstreamTimeRange {
	this := RstreamTimeRange{}
	return &this
}

// NewRstreamTimeRangeWithDefaults instantiates a new RstreamTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRstreamTimeRangeWithDefaults() *RstreamTimeRange {
	this := RstreamTimeRange{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RstreamTimeRange) GetEndTime() float64 {
	if o == nil || o.EndTime == nil {
		var ret float64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RstreamTimeRange) GetEndTimeOk() (*float64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RstreamTimeRange) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given float64 and assigns it to the EndTime field.
func (o *RstreamTimeRange) SetEndTime(v float64) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RstreamTimeRange) GetStartTime() float64 {
	if o == nil || o.StartTime == nil {
		var ret float64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RstreamTimeRange) GetStartTimeOk() (*float64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RstreamTimeRange) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given float64 and assigns it to the StartTime field.
func (o *RstreamTimeRange) SetStartTime(v float64) {
	o.StartTime = &v
}

func (o RstreamTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableRstreamTimeRange struct {
	value *RstreamTimeRange
	isSet bool
}

func (v NullableRstreamTimeRange) Get() *RstreamTimeRange {
	return v.value
}

func (v *NullableRstreamTimeRange) Set(val *RstreamTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableRstreamTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRstreamTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRstreamTimeRange(val *RstreamTimeRange) *NullableRstreamTimeRange {
	return &NullableRstreamTimeRange{value: val, isSet: true}
}

func (v NullableRstreamTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRstreamTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


