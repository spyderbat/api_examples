# \SourceDataApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**src_data_query**](SourceDataApi.md#src_data_query) | **POST** /api/v1/source/query/ | Query source data
[**src_data_query_v2**](SourceDataApi.md#src_data_query_v2) | **GET** /api/v1/org/{orgUID}/data/ | Query source data
[**src_send_data**](SourceDataApi.md#src_send_data) | **POST** /api/v1/org/{orgUID}/source/{sourceUID}/data/{dataType} | Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'. Alternatively, a multi-part form upload may be used with gzipped data up to a maximum size of 1MB.



## src_data_query

> String src_data_query(src_data_query_input)
Query source data

 Allows querying of the source data, data is stored as 'records' which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 24 hour period of time, it is the callers responsibility to construct an appropriate 24 hour range. * Documentation for the returned spydergraph datatype can be found at https://app.spyderbat.com/schema/spydergraph/index.html  * The user must have both the *org.ListSourceData* action on the org and *source_data.Query* action on the source * To get a count of results (up to 10K) but no data, use querySize: 0  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**src_data_query_input** | Option<[**SrcDataQueryInput**](SrcDataQueryInput.md)> |  |  |

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_data_query_v2

> String src_data_query_v2(org_uid, dt, e, et, id, pj, q, qf, qs, rr, src, st)
Query source data

Same as the post query above except results are cached

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | Organization UID to query | [required] |
**dt** | **String** | DataType to query | [required] |
**e** | Option<**String**> | Data which matches this expression are returned, as a json object |  |
**et** | Option<**f64**> | Time in unix epoch time, records < time are returned |  |
**id** | Option<[**Vec<String>**](String.md)> | List of IDs to resolve |  |
**pj** | Option<[**Vec<String>**](String.md)> | Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj=id&pj=version |  |
**q** | Option<**String**> | Lucene based search query |  |
**qf** | Option<**i32**> | Where to start the query in the result set from |  |
**qs** | Option<**i32**> | Size of the query result set |  |
**rr** | Option<**bool**> | Return rptrs mixed with the data |  |
**src** | Option<**String**> | Source UID to query |  |
**st** | Option<**f64**> | Time in unix epoch time, records >= time are returned |  |

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_send_data

> src_send_data(data_type, org_uid, source_uid, encoding, file)
Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'. Alternatively, a multi-part form upload may be used with gzipped data up to a maximum size of 1MB.

Sends data to a source

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**data_type** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |
**source_uid** | **String** |  | [required] |
**encoding** | **String** | must be gzip | [required] |
**file** | **std::path::PathBuf** | The file to upload. The file must be a valid gzip-ed JSON file. | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

