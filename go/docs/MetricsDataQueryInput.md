# MetricsDataQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggs** | Pointer to **map[string]interface{}** | Aggregations | [optional] 
**Causal** | Pointer to [**RstoreCausalQuery**](RstoreCausalQuery.md) |  | [optional] 
**ContextUid** | Pointer to **string** | Context UID for this query, it&#39;s used to track the query as it flows through the system, and shouldn&#39;t be exposed to customers | [optional] 
**DataType** | **string** | DataType to query | 
**EndTime** | Pointer to **float64** | Time in unix epoch time, records &lt; time are returned | [optional] 
**Expr** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Ids** | Pointer to **[]string** | Array of IDs to resolve into records | [optional] 
**OrgUid** | **string** | Organization UID to query | 
**Projection** | Pointer to **[]string** | Array of top level object properties which will be emitted, if none are specified all will be emitted | [optional] 
**Query** | Pointer to **string** | Lucene based search query | [optional] 
**QueryFrom** | Pointer to **int32** | Where to start the query in the result set from | [optional] 
**QuerySize** | Pointer to **int32** | Size of the query result set | [optional] 
**ReturnRptrs** | Pointer to **bool** |  | [optional] 
**Rptrs** | Pointer to **[]string** |  | [optional] 
**SortCol** | Pointer to **string** | Sort column | [optional] 
**SortOrder** | Pointer to **string** | Sort order | [optional] 
**SrcUid** | Pointer to **string** | Source UID to query | [optional] 
**StartTime** | Pointer to **float64** | Time in unix epoch time, records &gt;&#x3D; time are returned | [optional] 

## Methods

### NewMetricsDataQueryInput

`func NewMetricsDataQueryInput(dataType string, orgUid string, ) *MetricsDataQueryInput`

NewMetricsDataQueryInput instantiates a new MetricsDataQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDataQueryInputWithDefaults

`func NewMetricsDataQueryInputWithDefaults() *MetricsDataQueryInput`

NewMetricsDataQueryInputWithDefaults instantiates a new MetricsDataQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggs

`func (o *MetricsDataQueryInput) GetAggs() map[string]interface{}`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *MetricsDataQueryInput) GetAggsOk() (*map[string]interface{}, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *MetricsDataQueryInput) SetAggs(v map[string]interface{})`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *MetricsDataQueryInput) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetCausal

`func (o *MetricsDataQueryInput) GetCausal() RstoreCausalQuery`

GetCausal returns the Causal field if non-nil, zero value otherwise.

### GetCausalOk

`func (o *MetricsDataQueryInput) GetCausalOk() (*RstoreCausalQuery, bool)`

GetCausalOk returns a tuple with the Causal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausal

`func (o *MetricsDataQueryInput) SetCausal(v RstoreCausalQuery)`

SetCausal sets Causal field to given value.

### HasCausal

`func (o *MetricsDataQueryInput) HasCausal() bool`

HasCausal returns a boolean if a field has been set.

### GetContextUid

`func (o *MetricsDataQueryInput) GetContextUid() string`

GetContextUid returns the ContextUid field if non-nil, zero value otherwise.

### GetContextUidOk

`func (o *MetricsDataQueryInput) GetContextUidOk() (*string, bool)`

GetContextUidOk returns a tuple with the ContextUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextUid

`func (o *MetricsDataQueryInput) SetContextUid(v string)`

SetContextUid sets ContextUid field to given value.

### HasContextUid

`func (o *MetricsDataQueryInput) HasContextUid() bool`

HasContextUid returns a boolean if a field has been set.

### GetDataType

`func (o *MetricsDataQueryInput) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MetricsDataQueryInput) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MetricsDataQueryInput) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetEndTime

`func (o *MetricsDataQueryInput) GetEndTime() float64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MetricsDataQueryInput) GetEndTimeOk() (*float64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MetricsDataQueryInput) SetEndTime(v float64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MetricsDataQueryInput) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpr

`func (o *MetricsDataQueryInput) GetExpr() Expr`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *MetricsDataQueryInput) GetExprOk() (*Expr, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *MetricsDataQueryInput) SetExpr(v Expr)`

SetExpr sets Expr field to given value.

### HasExpr

`func (o *MetricsDataQueryInput) HasExpr() bool`

HasExpr returns a boolean if a field has been set.

### GetIds

`func (o *MetricsDataQueryInput) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MetricsDataQueryInput) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MetricsDataQueryInput) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MetricsDataQueryInput) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetOrgUid

`func (o *MetricsDataQueryInput) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *MetricsDataQueryInput) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *MetricsDataQueryInput) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.


### GetProjection

`func (o *MetricsDataQueryInput) GetProjection() []string`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *MetricsDataQueryInput) GetProjectionOk() (*[]string, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *MetricsDataQueryInput) SetProjection(v []string)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *MetricsDataQueryInput) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### GetQuery

`func (o *MetricsDataQueryInput) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricsDataQueryInput) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricsDataQueryInput) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetricsDataQueryInput) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryFrom

`func (o *MetricsDataQueryInput) GetQueryFrom() int32`

GetQueryFrom returns the QueryFrom field if non-nil, zero value otherwise.

### GetQueryFromOk

`func (o *MetricsDataQueryInput) GetQueryFromOk() (*int32, bool)`

GetQueryFromOk returns a tuple with the QueryFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFrom

`func (o *MetricsDataQueryInput) SetQueryFrom(v int32)`

SetQueryFrom sets QueryFrom field to given value.

### HasQueryFrom

`func (o *MetricsDataQueryInput) HasQueryFrom() bool`

HasQueryFrom returns a boolean if a field has been set.

### GetQuerySize

`func (o *MetricsDataQueryInput) GetQuerySize() int32`

GetQuerySize returns the QuerySize field if non-nil, zero value otherwise.

### GetQuerySizeOk

`func (o *MetricsDataQueryInput) GetQuerySizeOk() (*int32, bool)`

GetQuerySizeOk returns a tuple with the QuerySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySize

`func (o *MetricsDataQueryInput) SetQuerySize(v int32)`

SetQuerySize sets QuerySize field to given value.

### HasQuerySize

`func (o *MetricsDataQueryInput) HasQuerySize() bool`

HasQuerySize returns a boolean if a field has been set.

### GetReturnRptrs

`func (o *MetricsDataQueryInput) GetReturnRptrs() bool`

GetReturnRptrs returns the ReturnRptrs field if non-nil, zero value otherwise.

### GetReturnRptrsOk

`func (o *MetricsDataQueryInput) GetReturnRptrsOk() (*bool, bool)`

GetReturnRptrsOk returns a tuple with the ReturnRptrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnRptrs

`func (o *MetricsDataQueryInput) SetReturnRptrs(v bool)`

SetReturnRptrs sets ReturnRptrs field to given value.

### HasReturnRptrs

`func (o *MetricsDataQueryInput) HasReturnRptrs() bool`

HasReturnRptrs returns a boolean if a field has been set.

### GetRptrs

`func (o *MetricsDataQueryInput) GetRptrs() []string`

GetRptrs returns the Rptrs field if non-nil, zero value otherwise.

### GetRptrsOk

`func (o *MetricsDataQueryInput) GetRptrsOk() (*[]string, bool)`

GetRptrsOk returns a tuple with the Rptrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRptrs

`func (o *MetricsDataQueryInput) SetRptrs(v []string)`

SetRptrs sets Rptrs field to given value.

### HasRptrs

`func (o *MetricsDataQueryInput) HasRptrs() bool`

HasRptrs returns a boolean if a field has been set.

### GetSortCol

`func (o *MetricsDataQueryInput) GetSortCol() string`

GetSortCol returns the SortCol field if non-nil, zero value otherwise.

### GetSortColOk

`func (o *MetricsDataQueryInput) GetSortColOk() (*string, bool)`

GetSortColOk returns a tuple with the SortCol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCol

`func (o *MetricsDataQueryInput) SetSortCol(v string)`

SetSortCol sets SortCol field to given value.

### HasSortCol

`func (o *MetricsDataQueryInput) HasSortCol() bool`

HasSortCol returns a boolean if a field has been set.

### GetSortOrder

`func (o *MetricsDataQueryInput) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *MetricsDataQueryInput) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *MetricsDataQueryInput) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *MetricsDataQueryInput) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSrcUid

`func (o *MetricsDataQueryInput) GetSrcUid() string`

GetSrcUid returns the SrcUid field if non-nil, zero value otherwise.

### GetSrcUidOk

`func (o *MetricsDataQueryInput) GetSrcUidOk() (*string, bool)`

GetSrcUidOk returns a tuple with the SrcUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUid

`func (o *MetricsDataQueryInput) SetSrcUid(v string)`

SetSrcUid sets SrcUid field to given value.

### HasSrcUid

`func (o *MetricsDataQueryInput) HasSrcUid() bool`

HasSrcUid returns a boolean if a field has been set.

### GetStartTime

`func (o *MetricsDataQueryInput) GetStartTime() float64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MetricsDataQueryInput) GetStartTimeOk() (*float64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MetricsDataQueryInput) SetStartTime(v float64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MetricsDataQueryInput) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


