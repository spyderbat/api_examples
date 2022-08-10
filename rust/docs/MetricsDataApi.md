# \MetricsDataApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metrics_data_query**](MetricsDataApi.md#metrics_data_query) | **POST** /api/v1/metrics/query/ | Query metrics data



## metrics_data_query

> String metrics_data_query(metrics_data_query_input)
Query metrics data

 Allows querying of the metrics for metricss, data is stored as 'records' which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 1 week period of time, it is the callers responsibility to construct an appropriate 24 hour range.  * The user must have both the *metrics:QueryData* action on the org 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**metrics_data_query_input** | Option<[**MetricsDataQueryInput**](MetricsDataQueryInput.md)> |  |  |

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

