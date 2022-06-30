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

// InvestigationCreateInput struct for InvestigationCreateInput
type InvestigationCreateInput struct {
	// UID of user who created the investigation
	CreatedBy *string `json:"created_by,omitempty"`
	// JSON Object associated with the investigation
	Data map[string]interface{} `json:"data,omitempty"`
	// UID of the user who last modified the investigation
	ModifiedBy *string `json:"modified_by,omitempty"`
	// Date the investigation was last modified
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	// Name of the investigation
	Name *string `json:"name,omitempty"`
	// Resource name used for RBAC
	ResourceName *string `json:"resource_name,omitempty"`
	ResourcePolicy *ResourcePolicy `json:"resource_policy,omitempty"`
	// User supplied tags
	Tags []string `json:"tags,omitempty"`
	// Valid from date, the first date this object was valid
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	// Valid to date, the date this object is valid to
	ValidTo *time.Time `json:"valid_to,omitempty"`
	// Version of the investigation
	Version *int32 `json:"version,omitempty"`
}

// NewInvestigationCreateInput instantiates a new InvestigationCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestigationCreateInput() *InvestigationCreateInput {
	this := InvestigationCreateInput{}
	return &this
}

// NewInvestigationCreateInputWithDefaults instantiates a new InvestigationCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestigationCreateInputWithDefaults() *InvestigationCreateInput {
	this := InvestigationCreateInput{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *InvestigationCreateInput) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *InvestigationCreateInput) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *InvestigationCreateInput) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *InvestigationCreateInput) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InvestigationCreateInput) SetName(v string) {
	o.Name = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *InvestigationCreateInput) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourcePolicy returns the ResourcePolicy field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetResourcePolicy() ResourcePolicy {
	if o == nil || o.ResourcePolicy == nil {
		var ret ResourcePolicy
		return ret
	}
	return *o.ResourcePolicy
}

// GetResourcePolicyOk returns a tuple with the ResourcePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetResourcePolicyOk() (*ResourcePolicy, bool) {
	if o == nil || o.ResourcePolicy == nil {
		return nil, false
	}
	return o.ResourcePolicy, true
}

// HasResourcePolicy returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasResourcePolicy() bool {
	if o != nil && o.ResourcePolicy != nil {
		return true
	}

	return false
}

// SetResourcePolicy gets a reference to the given ResourcePolicy and assigns it to the ResourcePolicy field.
func (o *InvestigationCreateInput) SetResourcePolicy(v ResourcePolicy) {
	o.ResourcePolicy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InvestigationCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *InvestigationCreateInput) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *InvestigationCreateInput) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InvestigationCreateInput) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestigationCreateInput) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InvestigationCreateInput) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *InvestigationCreateInput) SetVersion(v int32) {
	o.Version = &v
}

func (o InvestigationCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.ModifiedOn != nil {
		toSerialize["modified_on"] = o.ModifiedOn
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.ValidFrom != nil {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["valid_to"] = o.ValidTo
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInvestigationCreateInput struct {
	value *InvestigationCreateInput
	isSet bool
}

func (v NullableInvestigationCreateInput) Get() *InvestigationCreateInput {
	return v.value
}

func (v *NullableInvestigationCreateInput) Set(val *InvestigationCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestigationCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestigationCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestigationCreateInput(val *InvestigationCreateInput) *NullableInvestigationCreateInput {
	return &NullableInvestigationCreateInput{value: val, isSet: true}
}

func (v NullableInvestigationCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestigationCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


