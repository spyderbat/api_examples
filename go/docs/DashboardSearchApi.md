# \DashboardSearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardSearchCreate**](DashboardSearchApi.md#DashboardSearchCreate) | **Post** /api/v1/org/{orgUID}/dashboard_search/ | Create a dashboard search
[**DashboardSearchDelete**](DashboardSearchApi.md#DashboardSearchDelete) | **Delete** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**DashboardSearchGet**](DashboardSearchApi.md#DashboardSearchGet) | **Get** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**DashboardSearchList**](DashboardSearchApi.md#DashboardSearchList) | **Get** /api/v1/org/{orgUID}/dashboard_search/ | List dashboard searches
[**DashboardSearchUpdate**](DashboardSearchApi.md#DashboardSearchUpdate) | **Put** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Update a dashboard search



## DashboardSearchCreate

> string DashboardSearchCreate(ctx, dashboardSearchUID, orgUID).DashboardSearchCreateInput(dashboardSearchCreateInput).Execute()

Create a dashboard search



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
    dashboardSearchUID := "dashboardSearchUID_example" // string | UID for the DashboardSearch
    orgUID := "orgUID_example" // string | Org UID
    dashboardSearchCreateInput := *openapiclient.NewDashboardSearchCreateInput() // DashboardSearchCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardSearchApi.DashboardSearchCreate(context.Background(), dashboardSearchUID, orgUID).DashboardSearchCreateInput(dashboardSearchCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardSearchApi.DashboardSearchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardSearchCreate`: string
    fmt.Fprintf(os.Stdout, "Response from `DashboardSearchApi.DashboardSearchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardSearchUID** | **string** | UID for the DashboardSearch | 
**orgUID** | **string** | Org UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardSearchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dashboardSearchCreateInput** | [**DashboardSearchCreateInput**](DashboardSearchCreateInput.md) |  | 

### Return type

**string**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardSearchDelete

> DashboardSearchDelete(ctx, dashboardSearchUID, orgUID).Execute()

Get a dashboard search



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
    dashboardSearchUID := "dashboardSearchUID_example" // string | 
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardSearchApi.DashboardSearchDelete(context.Background(), dashboardSearchUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardSearchApi.DashboardSearchDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardSearchUID** | **string** |  | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardSearchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardSearchGet

> DashboardSearch DashboardSearchGet(ctx, dashboardSearchUID, orgUID).Execute()

Get a dashboard search



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
    dashboardSearchUID := "dashboardSearchUID_example" // string | 
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardSearchApi.DashboardSearchGet(context.Background(), dashboardSearchUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardSearchApi.DashboardSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardSearchGet`: DashboardSearch
    fmt.Fprintf(os.Stdout, "Response from `DashboardSearchApi.DashboardSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardSearchUID** | **string** |  | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DashboardSearch**](DashboardSearch.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardSearchList

> []DashboardSearch DashboardSearchList(ctx, orgUID).Execute()

List dashboard searches



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
    orgUID := "orgUID_example" // string | Org UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardSearchApi.DashboardSearchList(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardSearchApi.DashboardSearchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardSearchList`: []DashboardSearch
    fmt.Fprintf(os.Stdout, "Response from `DashboardSearchApi.DashboardSearchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | Org UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardSearchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DashboardSearch**](DashboardSearch.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardSearchUpdate

> DashboardSearchUpdate(ctx, dashboardSearchUID, orgUID).DashboardSearchUpdateInput(dashboardSearchUpdateInput).Execute()

Update a dashboard search



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
    dashboardSearchUID := "dashboardSearchUID_example" // string | UID for the DashboardSearch
    orgUID := "orgUID_example" // string | Org UID
    dashboardSearchUpdateInput := *openapiclient.NewDashboardSearchUpdateInput() // DashboardSearchUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardSearchApi.DashboardSearchUpdate(context.Background(), dashboardSearchUID, orgUID).DashboardSearchUpdateInput(dashboardSearchUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardSearchApi.DashboardSearchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardSearchUID** | **string** | UID for the DashboardSearch | 
**orgUID** | **string** | Org UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardSearchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dashboardSearchUpdateInput** | [**DashboardSearchUpdateInput**](DashboardSearchUpdateInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

