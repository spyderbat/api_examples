# \SourceApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationSoarSrcList**](SourceApi.md#IntegrationSoarSrcList) | **Get** /api/v1/integration/soar/org/{orgUID}/source/ | List sources for integration with SOARs
[**SrcCreate**](SourceApi.md#SrcCreate) | **Post** /api/v1/org/{orgUID}/source/ | Create a source
[**SrcDelete**](SourceApi.md#SrcDelete) | **Delete** /api/v1/org/{orgUID}/source/{sourceUID} | Delete a source
[**SrcList**](SourceApi.md#SrcList) | **Get** /api/v1/org/{orgUID}/source/ | List sources
[**SrcLoad**](SourceApi.md#SrcLoad) | **Get** /api/v1/org/{orgUID}/source/{sourceUID} | Load a source
[**SrcUpdate**](SourceApi.md#SrcUpdate) | **Put** /api/v1/org/{orgUID}/source/{sourceUID} | Update a source



## IntegrationSoarSrcList

> []ApiSOARListHandlerOutput IntegrationSoarSrcList(ctx, orgUID).Et(et).Hostname(hostname).IpAddress(ipAddress).MacAddress(macAddress).Page(page).PageSize(pageSize).St(st).Execute()

List sources for integration with SOARs



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
    orgUID := "orgUID_example" // string | 
    et := int64(789) // int64 | optional end time of the query (optional)
    hostname := "hostname_example" // string | A single hostname to match (optional)
    ipAddress := "ipAddress_example" // string | A single IP address to match (optional)
    macAddress := "macAddress_example" // string | A single mac address to match (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    st := int64(789) // int64 | optional start time of the query, if only a start time is provided, end time will be start+10m (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.IntegrationSoarSrcList(context.Background(), orgUID).Et(et).Hostname(hostname).IpAddress(ipAddress).MacAddress(macAddress).Page(page).PageSize(pageSize).St(st).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.IntegrationSoarSrcList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationSoarSrcList`: []ApiSOARListHandlerOutput
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.IntegrationSoarSrcList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationSoarSrcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **et** | **int64** | optional end time of the query | 
 **hostname** | **string** | A single hostname to match | 
 **ipAddress** | **string** | A single IP address to match | 
 **macAddress** | **string** | A single mac address to match | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **st** | **int64** | optional start time of the query, if only a start time is provided, end time will be start+10m | 

### Return type

[**[]ApiSOARListHandlerOutput**](ApiSOARListHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrcCreate

> ApiSourceCreateHandlerOutput SrcCreate(ctx, orgUID).SrcCreateInput(srcCreateInput).Execute()

Create a source



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
    orgUID := "orgUID_example" // string | The org this source is associated with
    srcCreateInput := *openapiclient.NewSrcCreateInput() // SrcCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.SrcCreate(context.Background(), orgUID).SrcCreateInput(srcCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.SrcCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SrcCreate`: ApiSourceCreateHandlerOutput
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.SrcCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | The org this source is associated with | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **srcCreateInput** | [**SrcCreateInput**](SrcCreateInput.md) |  | 

### Return type

[**ApiSourceCreateHandlerOutput**](ApiSourceCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrcDelete

> SrcDelete(ctx, orgUID, sourceUID).Execute()

Delete a source



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
    orgUID := "orgUID_example" // string | 
    sourceUID := "sourceUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.SrcDelete(context.Background(), orgUID, sourceUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.SrcDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 
**sourceUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcDeleteRequest struct via the builder pattern


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


## SrcList

> []Source SrcList(ctx, orgUID).AgentUidEquals(agentUidEquals).DescriptionContains(descriptionContains).HasTags(hasTags).OriginalAssociation(originalAssociation).Page(page).PageSize(pageSize).Execute()

List sources



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
    orgUID := "orgUID_example" // string | 
    agentUidEquals := "agentUidEquals_example" // string |  (optional)
    descriptionContains := "descriptionContains_example" // string |  (optional)
    hasTags := []string{"Inner_example"} // []string |  (optional)
    originalAssociation := true // bool |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.SrcList(context.Background(), orgUID).AgentUidEquals(agentUidEquals).DescriptionContains(descriptionContains).HasTags(hasTags).OriginalAssociation(originalAssociation).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.SrcList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SrcList`: []Source
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.SrcList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentUidEquals** | **string** |  | 
 **descriptionContains** | **string** |  | 
 **hasTags** | **[]string** |  | 
 **originalAssociation** | **bool** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**[]Source**](Source.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrcLoad

> Source SrcLoad(ctx, orgUID, sourceUID).Execute()

Load a source



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
    orgUID := "orgUID_example" // string | 
    sourceUID := "sourceUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.SrcLoad(context.Background(), orgUID, sourceUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.SrcLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SrcLoad`: Source
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.SrcLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 
**sourceUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Source**](Source.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrcUpdate

> SrcUpdate(ctx, orgUID, sourceUID).SrcUpdateInput(srcUpdateInput).Execute()

Update a source



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
    orgUID := "orgUID_example" // string | The org this source is associated with
    sourceUID := "sourceUID_example" // string | The UID of the source
    srcUpdateInput := *openapiclient.NewSrcUpdateInput() // SrcUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.SrcUpdate(context.Background(), orgUID, sourceUID).SrcUpdateInput(srcUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.SrcUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | The org this source is associated with | 
**sourceUID** | **string** | The UID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **srcUpdateInput** | [**SrcUpdateInput**](SrcUpdateInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

