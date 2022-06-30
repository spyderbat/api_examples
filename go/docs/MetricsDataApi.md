# \MetricsDataApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsDataQuery**](MetricsDataApi.md#MetricsDataQuery) | **Post** /api/v1/metrics/query/ | Query metrics data



## MetricsDataQuery

> string MetricsDataQuery(ctx).MetricsDataQueryInput(metricsDataQueryInput).Execute()

Query metrics data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    metricsDataQueryInput := *openapiclient.NewMetricsDataQueryInput("DataType_example", "OrgUid_example") // MetricsDataQueryInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsDataApi.MetricsDataQuery(context.Background()).MetricsDataQueryInput(metricsDataQueryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsDataApi.MetricsDataQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetricsDataQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `MetricsDataApi.MetricsDataQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricsDataQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsDataQueryInput** | [**MetricsDataQueryInput**](MetricsDataQueryInput.md) |  | 

### Return type

**string**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

