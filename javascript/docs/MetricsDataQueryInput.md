# SpyderbatApi.MetricsDataQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**aggs** | **{String: Object}** | Aggregations | [optional] 
**causal** | [**RstoreCausalQuery**](RstoreCausalQuery.md) |  | [optional] 
**contextUid** | **String** | Context UID for this query, it&#39;s used to track the query as it flows through the system, and shouldn&#39;t be exposed to customers | [optional] 
**dataType** | **String** | DataType to query | 
**endTime** | **Number** | Time in unix epoch time, records &lt; time are returned | [optional] 
**expr** | [**Expr**](Expr.md) |  | [optional] 
**ids** | **[String]** | Array of IDs to resolve into records | [optional] 
**orgUid** | **String** | Organization UID to query | 
**projection** | **[String]** | Array of top level object properties which will be emitted, if none are specified all will be emitted | [optional] 
**query** | **String** | Lucene based search query | [optional] 
**queryFrom** | **Number** | Where to start the query in the result set from | [optional] 
**querySize** | **Number** | Size of the query result set | [optional] 
**returnRptrs** | **Boolean** |  | [optional] 
**rptrs** | **[String]** |  | [optional] 
**sortCol** | **String** | Sort column | [optional] 
**sortOrder** | **String** | Sort order | [optional] 
**srcUid** | **String** | Source UID to query | [optional] 
**startTime** | **Number** | Time in unix epoch time, records &gt;&#x3D; time are returned | [optional] 


