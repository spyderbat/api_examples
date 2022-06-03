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

// DaoAgentLog struct for DaoAgentLog
type DaoAgentLog struct {
	AgentRegistrationUid *string `json:"agent_registration_uid,omitempty"`
	AgentUid *string `json:"agent_uid,omitempty"`
	BytesSent *int32 `json:"bytes_sent,omitempty"`
	ErrUid *string `json:"err_uid,omitempty"`
	Error *string `json:"error,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	Msg *string `json:"msg,omitempty"`
	OrgUid *string `json:"org_uid,omitempty"`
	RuntimeDetails *OrcApiRuntimeDetails `json:"runtime_details,omitempty"`
	SourceUid *string `json:"source_uid,omitempty"`
	Time *float64 `json:"time,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDaoAgentLog instantiates a new DaoAgentLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaoAgentLog() *DaoAgentLog {
	this := DaoAgentLog{}
	return &this
}

// NewDaoAgentLogWithDefaults instantiates a new DaoAgentLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaoAgentLogWithDefaults() *DaoAgentLog {
	this := DaoAgentLog{}
	return &this
}

// GetAgentRegistrationUid returns the AgentRegistrationUid field value if set, zero value otherwise.
func (o *DaoAgentLog) GetAgentRegistrationUid() string {
	if o == nil || o.AgentRegistrationUid == nil {
		var ret string
		return ret
	}
	return *o.AgentRegistrationUid
}

// GetAgentRegistrationUidOk returns a tuple with the AgentRegistrationUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetAgentRegistrationUidOk() (*string, bool) {
	if o == nil || o.AgentRegistrationUid == nil {
		return nil, false
	}
	return o.AgentRegistrationUid, true
}

// HasAgentRegistrationUid returns a boolean if a field has been set.
func (o *DaoAgentLog) HasAgentRegistrationUid() bool {
	if o != nil && o.AgentRegistrationUid != nil {
		return true
	}

	return false
}

// SetAgentRegistrationUid gets a reference to the given string and assigns it to the AgentRegistrationUid field.
func (o *DaoAgentLog) SetAgentRegistrationUid(v string) {
	o.AgentRegistrationUid = &v
}

// GetAgentUid returns the AgentUid field value if set, zero value otherwise.
func (o *DaoAgentLog) GetAgentUid() string {
	if o == nil || o.AgentUid == nil {
		var ret string
		return ret
	}
	return *o.AgentUid
}

// GetAgentUidOk returns a tuple with the AgentUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetAgentUidOk() (*string, bool) {
	if o == nil || o.AgentUid == nil {
		return nil, false
	}
	return o.AgentUid, true
}

// HasAgentUid returns a boolean if a field has been set.
func (o *DaoAgentLog) HasAgentUid() bool {
	if o != nil && o.AgentUid != nil {
		return true
	}

	return false
}

// SetAgentUid gets a reference to the given string and assigns it to the AgentUid field.
func (o *DaoAgentLog) SetAgentUid(v string) {
	o.AgentUid = &v
}

// GetBytesSent returns the BytesSent field value if set, zero value otherwise.
func (o *DaoAgentLog) GetBytesSent() int32 {
	if o == nil || o.BytesSent == nil {
		var ret int32
		return ret
	}
	return *o.BytesSent
}

// GetBytesSentOk returns a tuple with the BytesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetBytesSentOk() (*int32, bool) {
	if o == nil || o.BytesSent == nil {
		return nil, false
	}
	return o.BytesSent, true
}

// HasBytesSent returns a boolean if a field has been set.
func (o *DaoAgentLog) HasBytesSent() bool {
	if o != nil && o.BytesSent != nil {
		return true
	}

	return false
}

// SetBytesSent gets a reference to the given int32 and assigns it to the BytesSent field.
func (o *DaoAgentLog) SetBytesSent(v int32) {
	o.BytesSent = &v
}

// GetErrUid returns the ErrUid field value if set, zero value otherwise.
func (o *DaoAgentLog) GetErrUid() string {
	if o == nil || o.ErrUid == nil {
		var ret string
		return ret
	}
	return *o.ErrUid
}

// GetErrUidOk returns a tuple with the ErrUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetErrUidOk() (*string, bool) {
	if o == nil || o.ErrUid == nil {
		return nil, false
	}
	return o.ErrUid, true
}

// HasErrUid returns a boolean if a field has been set.
func (o *DaoAgentLog) HasErrUid() bool {
	if o != nil && o.ErrUid != nil {
		return true
	}

	return false
}

// SetErrUid gets a reference to the given string and assigns it to the ErrUid field.
func (o *DaoAgentLog) SetErrUid(v string) {
	o.ErrUid = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DaoAgentLog) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DaoAgentLog) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DaoAgentLog) SetError(v string) {
	o.Error = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *DaoAgentLog) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *DaoAgentLog) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *DaoAgentLog) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *DaoAgentLog) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *DaoAgentLog) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *DaoAgentLog) SetMsg(v string) {
	o.Msg = &v
}

// GetOrgUid returns the OrgUid field value if set, zero value otherwise.
func (o *DaoAgentLog) GetOrgUid() string {
	if o == nil || o.OrgUid == nil {
		var ret string
		return ret
	}
	return *o.OrgUid
}

// GetOrgUidOk returns a tuple with the OrgUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetOrgUidOk() (*string, bool) {
	if o == nil || o.OrgUid == nil {
		return nil, false
	}
	return o.OrgUid, true
}

// HasOrgUid returns a boolean if a field has been set.
func (o *DaoAgentLog) HasOrgUid() bool {
	if o != nil && o.OrgUid != nil {
		return true
	}

	return false
}

// SetOrgUid gets a reference to the given string and assigns it to the OrgUid field.
func (o *DaoAgentLog) SetOrgUid(v string) {
	o.OrgUid = &v
}

// GetRuntimeDetails returns the RuntimeDetails field value if set, zero value otherwise.
func (o *DaoAgentLog) GetRuntimeDetails() OrcApiRuntimeDetails {
	if o == nil || o.RuntimeDetails == nil {
		var ret OrcApiRuntimeDetails
		return ret
	}
	return *o.RuntimeDetails
}

// GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool) {
	if o == nil || o.RuntimeDetails == nil {
		return nil, false
	}
	return o.RuntimeDetails, true
}

// HasRuntimeDetails returns a boolean if a field has been set.
func (o *DaoAgentLog) HasRuntimeDetails() bool {
	if o != nil && o.RuntimeDetails != nil {
		return true
	}

	return false
}

// SetRuntimeDetails gets a reference to the given OrcApiRuntimeDetails and assigns it to the RuntimeDetails field.
func (o *DaoAgentLog) SetRuntimeDetails(v OrcApiRuntimeDetails) {
	o.RuntimeDetails = &v
}

// GetSourceUid returns the SourceUid field value if set, zero value otherwise.
func (o *DaoAgentLog) GetSourceUid() string {
	if o == nil || o.SourceUid == nil {
		var ret string
		return ret
	}
	return *o.SourceUid
}

// GetSourceUidOk returns a tuple with the SourceUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetSourceUidOk() (*string, bool) {
	if o == nil || o.SourceUid == nil {
		return nil, false
	}
	return o.SourceUid, true
}

// HasSourceUid returns a boolean if a field has been set.
func (o *DaoAgentLog) HasSourceUid() bool {
	if o != nil && o.SourceUid != nil {
		return true
	}

	return false
}

// SetSourceUid gets a reference to the given string and assigns it to the SourceUid field.
func (o *DaoAgentLog) SetSourceUid(v string) {
	o.SourceUid = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DaoAgentLog) GetTime() float64 {
	if o == nil || o.Time == nil {
		var ret float64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetTimeOk() (*float64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DaoAgentLog) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given float64 and assigns it to the Time field.
func (o *DaoAgentLog) SetTime(v float64) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DaoAgentLog) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaoAgentLog) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DaoAgentLog) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DaoAgentLog) SetType(v string) {
	o.Type = &v
}

func (o DaoAgentLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentRegistrationUid != nil {
		toSerialize["agent_registration_uid"] = o.AgentRegistrationUid
	}
	if o.AgentUid != nil {
		toSerialize["agent_uid"] = o.AgentUid
	}
	if o.BytesSent != nil {
		toSerialize["bytes_sent"] = o.BytesSent
	}
	if o.ErrUid != nil {
		toSerialize["err_uid"] = o.ErrUid
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.OrgUid != nil {
		toSerialize["org_uid"] = o.OrgUid
	}
	if o.RuntimeDetails != nil {
		toSerialize["runtime_details"] = o.RuntimeDetails
	}
	if o.SourceUid != nil {
		toSerialize["source_uid"] = o.SourceUid
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDaoAgentLog struct {
	value *DaoAgentLog
	isSet bool
}

func (v NullableDaoAgentLog) Get() *DaoAgentLog {
	return v.value
}

func (v *NullableDaoAgentLog) Set(val *DaoAgentLog) {
	v.value = val
	v.isSet = true
}

func (v NullableDaoAgentLog) IsSet() bool {
	return v.isSet
}

func (v *NullableDaoAgentLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaoAgentLog(val *DaoAgentLog) *NullableDaoAgentLog {
	return &NullableDaoAgentLog{value: val, isSet: true}
}

func (v NullableDaoAgentLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaoAgentLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


