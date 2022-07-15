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

// MetricsDataQueryInput struct for MetricsDataQueryInput
type MetricsDataQueryInput struct {
	// Aggregations
	Aggs map[string]interface{} `json:"aggs,omitempty"`
	Causal *RstoreCausalQuery `json:"causal,omitempty"`
	// Context UID for this query, it's used to track the query as it flows through the system, and shouldn't be exposed to customers
	ContextUid *string `json:"context_uid,omitempty"`
	// DataType to query
	DataType string `json:"data_type"`
	// Time in unix epoch time, records < time are returned
	EndTime *float64 `json:"end_time,omitempty"`
	Expr *Expr `json:"expr,omitempty"`
	// Array of IDs to resolve into records
	Ids []string `json:"ids,omitempty"`
	// Organization UID to query
	OrgUid string `json:"org_uid"`
	// Array of top level object properties which will be emitted, if none are specified all will be emitted
	Projection []string `json:"projection,omitempty"`
	// Lucene based search query
	Query *string `json:"query,omitempty"`
	// Where to start the query in the result set from
	QueryFrom *int32 `json:"query_from,omitempty"`
	// Size of the query result set
	QuerySize *int32 `json:"query_size,omitempty"`
	ReturnRptrs *bool `json:"return_rptrs,omitempty"`
	Rptrs []string `json:"rptrs,omitempty"`
	// Sort column
	SortCol *string `json:"sort_col,omitempty"`
	// Sort order
	SortOrder *string `json:"sort_order,omitempty"`
	// Source UID to query
	SrcUid *string `json:"src_uid,omitempty"`
	// Time in unix epoch time, records >= time are returned
	StartTime *float64 `json:"start_time,omitempty"`
}

// NewMetricsDataQueryInput instantiates a new MetricsDataQueryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsDataQueryInput(dataType string, orgUid string) *MetricsDataQueryInput {
	this := MetricsDataQueryInput{}
	this.DataType = dataType
	this.OrgUid = orgUid
	return &this
}

// NewMetricsDataQueryInputWithDefaults instantiates a new MetricsDataQueryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsDataQueryInputWithDefaults() *MetricsDataQueryInput {
	this := MetricsDataQueryInput{}
	return &this
}

// GetAggs returns the Aggs field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetAggs() map[string]interface{} {
	if o == nil || o.Aggs == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Aggs
}

// GetAggsOk returns a tuple with the Aggs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetAggsOk() (map[string]interface{}, bool) {
	if o == nil || o.Aggs == nil {
		return nil, false
	}
	return o.Aggs, true
}

// HasAggs returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasAggs() bool {
	if o != nil && o.Aggs != nil {
		return true
	}

	return false
}

// SetAggs gets a reference to the given map[string]interface{} and assigns it to the Aggs field.
func (o *MetricsDataQueryInput) SetAggs(v map[string]interface{}) {
	o.Aggs = v
}

// GetCausal returns the Causal field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetCausal() RstoreCausalQuery {
	if o == nil || o.Causal == nil {
		var ret RstoreCausalQuery
		return ret
	}
	return *o.Causal
}

// GetCausalOk returns a tuple with the Causal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetCausalOk() (*RstoreCausalQuery, bool) {
	if o == nil || o.Causal == nil {
		return nil, false
	}
	return o.Causal, true
}

// HasCausal returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasCausal() bool {
	if o != nil && o.Causal != nil {
		return true
	}

	return false
}

// SetCausal gets a reference to the given RstoreCausalQuery and assigns it to the Causal field.
func (o *MetricsDataQueryInput) SetCausal(v RstoreCausalQuery) {
	o.Causal = &v
}

// GetContextUid returns the ContextUid field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetContextUid() string {
	if o == nil || o.ContextUid == nil {
		var ret string
		return ret
	}
	return *o.ContextUid
}

// GetContextUidOk returns a tuple with the ContextUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetContextUidOk() (*string, bool) {
	if o == nil || o.ContextUid == nil {
		return nil, false
	}
	return o.ContextUid, true
}

// HasContextUid returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasContextUid() bool {
	if o != nil && o.ContextUid != nil {
		return true
	}

	return false
}

// SetContextUid gets a reference to the given string and assigns it to the ContextUid field.
func (o *MetricsDataQueryInput) SetContextUid(v string) {
	o.ContextUid = &v
}

// GetDataType returns the DataType field value
func (o *MetricsDataQueryInput) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *MetricsDataQueryInput) SetDataType(v string) {
	o.DataType = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetEndTime() float64 {
	if o == nil || o.EndTime == nil {
		var ret float64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetEndTimeOk() (*float64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given float64 and assigns it to the EndTime field.
func (o *MetricsDataQueryInput) SetEndTime(v float64) {
	o.EndTime = &v
}

// GetExpr returns the Expr field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetExpr() Expr {
	if o == nil || o.Expr == nil {
		var ret Expr
		return ret
	}
	return *o.Expr
}

// GetExprOk returns a tuple with the Expr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetExprOk() (*Expr, bool) {
	if o == nil || o.Expr == nil {
		return nil, false
	}
	return o.Expr, true
}

// HasExpr returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasExpr() bool {
	if o != nil && o.Expr != nil {
		return true
	}

	return false
}

// SetExpr gets a reference to the given Expr and assigns it to the Expr field.
func (o *MetricsDataQueryInput) SetExpr(v Expr) {
	o.Expr = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *MetricsDataQueryInput) SetIds(v []string) {
	o.Ids = v
}

// GetOrgUid returns the OrgUid field value
func (o *MetricsDataQueryInput) GetOrgUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgUid
}

// GetOrgUidOk returns a tuple with the OrgUid field value
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetOrgUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgUid, true
}

// SetOrgUid sets field value
func (o *MetricsDataQueryInput) SetOrgUid(v string) {
	o.OrgUid = v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetProjection() []string {
	if o == nil || o.Projection == nil {
		var ret []string
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetProjectionOk() ([]string, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasProjection() bool {
	if o != nil && o.Projection != nil {
		return true
	}

	return false
}

// SetProjection gets a reference to the given []string and assigns it to the Projection field.
func (o *MetricsDataQueryInput) SetProjection(v []string) {
	o.Projection = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MetricsDataQueryInput) SetQuery(v string) {
	o.Query = &v
}

// GetQueryFrom returns the QueryFrom field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetQueryFrom() int32 {
	if o == nil || o.QueryFrom == nil {
		var ret int32
		return ret
	}
	return *o.QueryFrom
}

// GetQueryFromOk returns a tuple with the QueryFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetQueryFromOk() (*int32, bool) {
	if o == nil || o.QueryFrom == nil {
		return nil, false
	}
	return o.QueryFrom, true
}

// HasQueryFrom returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasQueryFrom() bool {
	if o != nil && o.QueryFrom != nil {
		return true
	}

	return false
}

// SetQueryFrom gets a reference to the given int32 and assigns it to the QueryFrom field.
func (o *MetricsDataQueryInput) SetQueryFrom(v int32) {
	o.QueryFrom = &v
}

// GetQuerySize returns the QuerySize field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetQuerySize() int32 {
	if o == nil || o.QuerySize == nil {
		var ret int32
		return ret
	}
	return *o.QuerySize
}

// GetQuerySizeOk returns a tuple with the QuerySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetQuerySizeOk() (*int32, bool) {
	if o == nil || o.QuerySize == nil {
		return nil, false
	}
	return o.QuerySize, true
}

// HasQuerySize returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasQuerySize() bool {
	if o != nil && o.QuerySize != nil {
		return true
	}

	return false
}

// SetQuerySize gets a reference to the given int32 and assigns it to the QuerySize field.
func (o *MetricsDataQueryInput) SetQuerySize(v int32) {
	o.QuerySize = &v
}

// GetReturnRptrs returns the ReturnRptrs field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetReturnRptrs() bool {
	if o == nil || o.ReturnRptrs == nil {
		var ret bool
		return ret
	}
	return *o.ReturnRptrs
}

// GetReturnRptrsOk returns a tuple with the ReturnRptrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetReturnRptrsOk() (*bool, bool) {
	if o == nil || o.ReturnRptrs == nil {
		return nil, false
	}
	return o.ReturnRptrs, true
}

// HasReturnRptrs returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasReturnRptrs() bool {
	if o != nil && o.ReturnRptrs != nil {
		return true
	}

	return false
}

// SetReturnRptrs gets a reference to the given bool and assigns it to the ReturnRptrs field.
func (o *MetricsDataQueryInput) SetReturnRptrs(v bool) {
	o.ReturnRptrs = &v
}

// GetRptrs returns the Rptrs field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetRptrs() []string {
	if o == nil || o.Rptrs == nil {
		var ret []string
		return ret
	}
	return o.Rptrs
}

// GetRptrsOk returns a tuple with the Rptrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetRptrsOk() ([]string, bool) {
	if o == nil || o.Rptrs == nil {
		return nil, false
	}
	return o.Rptrs, true
}

// HasRptrs returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasRptrs() bool {
	if o != nil && o.Rptrs != nil {
		return true
	}

	return false
}

// SetRptrs gets a reference to the given []string and assigns it to the Rptrs field.
func (o *MetricsDataQueryInput) SetRptrs(v []string) {
	o.Rptrs = v
}

// GetSortCol returns the SortCol field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetSortCol() string {
	if o == nil || o.SortCol == nil {
		var ret string
		return ret
	}
	return *o.SortCol
}

// GetSortColOk returns a tuple with the SortCol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetSortColOk() (*string, bool) {
	if o == nil || o.SortCol == nil {
		return nil, false
	}
	return o.SortCol, true
}

// HasSortCol returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasSortCol() bool {
	if o != nil && o.SortCol != nil {
		return true
	}

	return false
}

// SetSortCol gets a reference to the given string and assigns it to the SortCol field.
func (o *MetricsDataQueryInput) SetSortCol(v string) {
	o.SortCol = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *MetricsDataQueryInput) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetSrcUid returns the SrcUid field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetSrcUid() string {
	if o == nil || o.SrcUid == nil {
		var ret string
		return ret
	}
	return *o.SrcUid
}

// GetSrcUidOk returns a tuple with the SrcUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetSrcUidOk() (*string, bool) {
	if o == nil || o.SrcUid == nil {
		return nil, false
	}
	return o.SrcUid, true
}

// HasSrcUid returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasSrcUid() bool {
	if o != nil && o.SrcUid != nil {
		return true
	}

	return false
}

// SetSrcUid gets a reference to the given string and assigns it to the SrcUid field.
func (o *MetricsDataQueryInput) SetSrcUid(v string) {
	o.SrcUid = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MetricsDataQueryInput) GetStartTime() float64 {
	if o == nil || o.StartTime == nil {
		var ret float64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsDataQueryInput) GetStartTimeOk() (*float64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MetricsDataQueryInput) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given float64 and assigns it to the StartTime field.
func (o *MetricsDataQueryInput) SetStartTime(v float64) {
	o.StartTime = &v
}

func (o MetricsDataQueryInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggs != nil {
		toSerialize["aggs"] = o.Aggs
	}
	if o.Causal != nil {
		toSerialize["causal"] = o.Causal
	}
	if o.ContextUid != nil {
		toSerialize["context_uid"] = o.ContextUid
	}
	if true {
		toSerialize["data_type"] = o.DataType
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.Expr != nil {
		toSerialize["expr"] = o.Expr
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if true {
		toSerialize["org_uid"] = o.OrgUid
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryFrom != nil {
		toSerialize["query_from"] = o.QueryFrom
	}
	if o.QuerySize != nil {
		toSerialize["query_size"] = o.QuerySize
	}
	if o.ReturnRptrs != nil {
		toSerialize["return_rptrs"] = o.ReturnRptrs
	}
	if o.Rptrs != nil {
		toSerialize["rptrs"] = o.Rptrs
	}
	if o.SortCol != nil {
		toSerialize["sort_col"] = o.SortCol
	}
	if o.SortOrder != nil {
		toSerialize["sort_order"] = o.SortOrder
	}
	if o.SrcUid != nil {
		toSerialize["src_uid"] = o.SrcUid
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsDataQueryInput struct {
	value *MetricsDataQueryInput
	isSet bool
}

func (v NullableMetricsDataQueryInput) Get() *MetricsDataQueryInput {
	return v.value
}

func (v *NullableMetricsDataQueryInput) Set(val *MetricsDataQueryInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsDataQueryInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsDataQueryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsDataQueryInput(val *MetricsDataQueryInput) *NullableMetricsDataQueryInput {
	return &NullableMetricsDataQueryInput{value: val, isSet: true}
}

func (v NullableMetricsDataQueryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsDataQueryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


