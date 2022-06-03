# SrcDataQueryInput


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_type** | **str** | DataType to query | 
**org_uid** | **str** | Organization UID to query | 
**aggs** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Aggregations | [optional] 
**causal** | [**RstoreCausalQuery**](RstoreCausalQuery.md) |  | [optional] 
**context_uid** | **str** | Context UID for this query, it&#39;s used to track the query as it flows through the system, and shouldn&#39;t be exposed to customers | [optional] 
**end_time** | **float** | Time in unix epoch time, records &lt; time are returned | [optional] 
**expr** | [**Expr**](Expr.md) |  | [optional] 
**ids** | **[str]** | Array of IDs to resolve into records | [optional] 
**projection** | **[str]** | Array of top level object properties which will be emitted, if none are specified all will be emitted | [optional] 
**query** | **str** | Lucene based search query | [optional] 
**query_from** | **int** | Where to start the query in the result set from | [optional] 
**query_size** | **int** | Size of the query result set | [optional] 
**return_rptrs** | **bool** |  | [optional] 
**rptrs** | **[str]** |  | [optional] 
**sort_col** | **str** | Sort column | [optional] 
**sort_order** | **str** | Sort order | [optional] 
**src_uid** | **str** | Source UID to query | [optional] 
**start_time** | **float** | Time in unix epoch time, records &gt;&#x3D; time are returned | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


