# SrcDataQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**aggs** | Option<[**::std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> | Aggregations | [optional]
**causal** | Option<[**crate::models::RstoreCausalQuery**](RstoreCausalQuery.md)> |  | [optional]
**context_uid** | Option<**String**> | Context UID for this query, it's used to track the query as it flows through the system, and shouldn't be exposed to customers | [optional]
**data_type** | **String** | DataType to query | 
**end_time** | Option<**f64**> | Time in unix epoch time, records < time are returned | [optional]
**expr** | Option<[**crate::models::Expr**](Expr.md)> |  | [optional]
**ids** | Option<**Vec<String>**> | Array of IDs to resolve into records | [optional]
**org_uid** | **String** | Organization UID to query | 
**projection** | Option<**Vec<String>**> | Array of top level object properties which will be emitted, if none are specified all will be emitted | [optional]
**query** | Option<**String**> | Lucene based search query | [optional]
**query_from** | Option<**i32**> | Where to start the query in the result set from | [optional]
**query_size** | Option<**i32**> | Size of the query result set | [optional]
**return_rptrs** | Option<**bool**> |  | [optional]
**rptrs** | Option<**Vec<String>**> |  | [optional]
**sort_col** | Option<**String**> | Sort column | [optional]
**sort_order** | Option<**String**> | Sort order | [optional]
**src_uid** | Option<**String**> | Source UID to query | [optional]
**start_time** | Option<**f64**> | Time in unix epoch time, records >= time are returned | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


