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

// Agent struct for Agent
type Agent struct {
	// Agent registration associated with the agent
	AgentRegistrationUid *string `json:"agent_registration_uid,omitempty"`
	// Agent type
	AgentType *string `json:"agent_type,omitempty"`
	// Description of the the purpose of the agent
	Description *string `json:"description,omitempty"`
	// Agent OrgUID
	OrgUid *string `json:"org_uid,omitempty"`
	// Resource name used for RBAC
	ResourceName *string `json:"resource_name,omitempty"`
	ResourcePolicy *ResourcePolicy `json:"resource_policy,omitempty"`
	// Description of the runtime of the agent
	RuntimeDescription *string `json:"runtime_description,omitempty"`
	RuntimeDetails *OrcApiRuntimeDetails `json:"runtime_details,omitempty"`
	// Agent UID
	Uid *string `json:"uid,omitempty"`
	// Valid from date, the first date this object was valid
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	// Valid to date, the date this object is valid to
	ValidTo *time.Time `json:"valid_to,omitempty"`
}

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent() *Agent {
	this := Agent{}
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetAgentRegistrationUid returns the AgentRegistrationUid field value if set, zero value otherwise.
func (o *Agent) GetAgentRegistrationUid() string {
	if o == nil || o.AgentRegistrationUid == nil {
		var ret string
		return ret
	}
	return *o.AgentRegistrationUid
}

// GetAgentRegistrationUidOk returns a tuple with the AgentRegistrationUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAgentRegistrationUidOk() (*string, bool) {
	if o == nil || o.AgentRegistrationUid == nil {
		return nil, false
	}
	return o.AgentRegistrationUid, true
}

// HasAgentRegistrationUid returns a boolean if a field has been set.
func (o *Agent) HasAgentRegistrationUid() bool {
	if o != nil && o.AgentRegistrationUid != nil {
		return true
	}

	return false
}

// SetAgentRegistrationUid gets a reference to the given string and assigns it to the AgentRegistrationUid field.
func (o *Agent) SetAgentRegistrationUid(v string) {
	o.AgentRegistrationUid = &v
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *Agent) GetAgentType() string {
	if o == nil || o.AgentType == nil {
		var ret string
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetAgentTypeOk() (*string, bool) {
	if o == nil || o.AgentType == nil {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *Agent) HasAgentType() bool {
	if o != nil && o.AgentType != nil {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given string and assigns it to the AgentType field.
func (o *Agent) SetAgentType(v string) {
	o.AgentType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Agent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Agent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Agent) SetDescription(v string) {
	o.Description = &v
}

// GetOrgUid returns the OrgUid field value if set, zero value otherwise.
func (o *Agent) GetOrgUid() string {
	if o == nil || o.OrgUid == nil {
		var ret string
		return ret
	}
	return *o.OrgUid
}

// GetOrgUidOk returns a tuple with the OrgUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetOrgUidOk() (*string, bool) {
	if o == nil || o.OrgUid == nil {
		return nil, false
	}
	return o.OrgUid, true
}

// HasOrgUid returns a boolean if a field has been set.
func (o *Agent) HasOrgUid() bool {
	if o != nil && o.OrgUid != nil {
		return true
	}

	return false
}

// SetOrgUid gets a reference to the given string and assigns it to the OrgUid field.
func (o *Agent) SetOrgUid(v string) {
	o.OrgUid = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *Agent) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *Agent) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *Agent) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourcePolicy returns the ResourcePolicy field value if set, zero value otherwise.
func (o *Agent) GetResourcePolicy() ResourcePolicy {
	if o == nil || o.ResourcePolicy == nil {
		var ret ResourcePolicy
		return ret
	}
	return *o.ResourcePolicy
}

// GetResourcePolicyOk returns a tuple with the ResourcePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetResourcePolicyOk() (*ResourcePolicy, bool) {
	if o == nil || o.ResourcePolicy == nil {
		return nil, false
	}
	return o.ResourcePolicy, true
}

// HasResourcePolicy returns a boolean if a field has been set.
func (o *Agent) HasResourcePolicy() bool {
	if o != nil && o.ResourcePolicy != nil {
		return true
	}

	return false
}

// SetResourcePolicy gets a reference to the given ResourcePolicy and assigns it to the ResourcePolicy field.
func (o *Agent) SetResourcePolicy(v ResourcePolicy) {
	o.ResourcePolicy = &v
}

// GetRuntimeDescription returns the RuntimeDescription field value if set, zero value otherwise.
func (o *Agent) GetRuntimeDescription() string {
	if o == nil || o.RuntimeDescription == nil {
		var ret string
		return ret
	}
	return *o.RuntimeDescription
}

// GetRuntimeDescriptionOk returns a tuple with the RuntimeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetRuntimeDescriptionOk() (*string, bool) {
	if o == nil || o.RuntimeDescription == nil {
		return nil, false
	}
	return o.RuntimeDescription, true
}

// HasRuntimeDescription returns a boolean if a field has been set.
func (o *Agent) HasRuntimeDescription() bool {
	if o != nil && o.RuntimeDescription != nil {
		return true
	}

	return false
}

// SetRuntimeDescription gets a reference to the given string and assigns it to the RuntimeDescription field.
func (o *Agent) SetRuntimeDescription(v string) {
	o.RuntimeDescription = &v
}

// GetRuntimeDetails returns the RuntimeDetails field value if set, zero value otherwise.
func (o *Agent) GetRuntimeDetails() OrcApiRuntimeDetails {
	if o == nil || o.RuntimeDetails == nil {
		var ret OrcApiRuntimeDetails
		return ret
	}
	return *o.RuntimeDetails
}

// GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool) {
	if o == nil || o.RuntimeDetails == nil {
		return nil, false
	}
	return o.RuntimeDetails, true
}

// HasRuntimeDetails returns a boolean if a field has been set.
func (o *Agent) HasRuntimeDetails() bool {
	if o != nil && o.RuntimeDetails != nil {
		return true
	}

	return false
}

// SetRuntimeDetails gets a reference to the given OrcApiRuntimeDetails and assigns it to the RuntimeDetails field.
func (o *Agent) SetRuntimeDetails(v OrcApiRuntimeDetails) {
	o.RuntimeDetails = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *Agent) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *Agent) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *Agent) SetUid(v string) {
	o.Uid = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *Agent) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *Agent) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *Agent) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *Agent) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *Agent) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *Agent) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentRegistrationUid != nil {
		toSerialize["agent_registration_uid"] = o.AgentRegistrationUid
	}
	if o.AgentType != nil {
		toSerialize["agent_type"] = o.AgentType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OrgUid != nil {
		toSerialize["org_uid"] = o.OrgUid
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourcePolicy != nil {
		toSerialize["resource_policy"] = o.ResourcePolicy
	}
	if o.RuntimeDescription != nil {
		toSerialize["runtime_description"] = o.RuntimeDescription
	}
	if o.RuntimeDetails != nil {
		toSerialize["runtime_details"] = o.RuntimeDetails
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

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


