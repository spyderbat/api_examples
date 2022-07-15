# \InvestigationApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvestigationCreate**](InvestigationApi.md#InvestigationCreate) | **Post** /api/v1/org/{orgUID}/investigation/ | Create an investigation
[**InvestigationDelete**](InvestigationApi.md#InvestigationDelete) | **Delete** /api/v1/org/{orgUID}/investigation/{investigationUID} | Delete an investigation
[**InvestigationList**](InvestigationApi.md#InvestigationList) | **Get** /api/v1/org/{orgUID}/investigation/ | List investigations
[**InvestigationListVersions**](InvestigationApi.md#InvestigationListVersions) | **Get** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/ | List Investigation Versions
[**InvestigationLoad**](InvestigationApi.md#InvestigationLoad) | **Get** /api/v1/org/{orgUID}/investigation/{investigationUID} | Load an investigation
[**InvestigationLoadVersion**](InvestigationApi.md#InvestigationLoadVersion) | **Get** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/{version} | Load Investigation Version
[**InvestigationUpdate**](InvestigationApi.md#InvestigationUpdate) | **Put** /api/v1/org/{orgUID}/investigation/{investigationUID} | Update an investigation



## InvestigationCreate

> ApiInvestigationCreateOutput InvestigationCreate(ctx, investigationUID, orgUID).InvestigationCreateInput(investigationCreateInput).Execute()

Create an investigation



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
    investigationUID := "investigationUID_example" // string | Investigation UID
    orgUID := "orgUID_example" // string | Investigation OrgUID
    investigationCreateInput := *openapiclient.NewInvestigationCreateInput() // InvestigationCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationCreate(context.Background(), investigationUID, orgUID).InvestigationCreateInput(investigationCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationCreate`: ApiInvestigationCreateOutput
    fmt.Fprintf(os.Stdout, "Response from `InvestigationApi.InvestigationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationUID** | **string** | Investigation UID | 
**orgUID** | **string** | Investigation OrgUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **investigationCreateInput** | [**InvestigationCreateInput**](InvestigationCreateInput.md) |  | 

### Return type

[**ApiInvestigationCreateOutput**](ApiInvestigationCreateOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationDelete

> InvestigationDelete(ctx, investigationUID, orgUID).Execute()

Delete an investigation



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
    investigationUID := "investigationUID_example" // string | Investigation UID
    orgUID := "orgUID_example" // string | Investigation OrgUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationDelete(context.Background(), investigationUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationUID** | **string** | Investigation UID | 
**orgUID** | **string** | Investigation OrgUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationDeleteRequest struct via the builder pattern


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


## InvestigationList

> []DaoInvestigation InvestigationList(ctx, orgUID).Execute()

List investigations



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationList(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationList`: []DaoInvestigation
    fmt.Fprintf(os.Stdout, "Response from `InvestigationApi.InvestigationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationListVersions

> []DaoInvestigation InvestigationListVersions(ctx, investigationUID, orgUID).Execute()

List Investigation Versions



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
    investigationUID := "investigationUID_example" // string | 
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationListVersions(context.Background(), investigationUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationListVersions`: []DaoInvestigation
    fmt.Fprintf(os.Stdout, "Response from `InvestigationApi.InvestigationListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationUID** | **string** |  | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationLoad

> DaoInvestigation InvestigationLoad(ctx, investigationUID, orgUID).Execute()

Load an investigation



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
    investigationUID := "investigationUID_example" // string | 
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationLoad(context.Background(), investigationUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationLoad`: DaoInvestigation
    fmt.Fprintf(os.Stdout, "Response from `InvestigationApi.InvestigationLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationUID** | **string** |  | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationLoadVersion

> DaoInvestigation InvestigationLoadVersion(ctx, investigationUID, orgUID, version).Execute()

Load Investigation Version



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
    investigationUID := "investigationUID_example" // string | 
    orgUID := "orgUID_example" // string | 
    version := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationLoadVersion(context.Background(), investigationUID, orgUID, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationLoadVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationLoadVersion`: DaoInvestigation
    fmt.Fprintf(os.Stdout, "Response from `InvestigationApi.InvestigationLoadVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationUID** | **string** |  | 
**orgUID** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationLoadVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationUpdate

> InvestigationUpdate(ctx, investigationUID, orgUID).InvestigationUpdateInput(investigationUpdateInput).Execute()

Update an investigation



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
    investigationUID := "investigationUID_example" // string | Investigation UID
    orgUID := "orgUID_example" // string | Investigation OrgUID
    investigationUpdateInput := *openapiclient.NewInvestigationUpdateInput() // InvestigationUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestigationApi.InvestigationUpdate(context.Background(), investigationUID, orgUID).InvestigationUpdateInput(investigationUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestigationApi.InvestigationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationUID** | **string** | Investigation UID | 
**orgUID** | **string** | Investigation OrgUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **investigationUpdateInput** | [**InvestigationUpdateInput**](InvestigationUpdateInput.md) |  | 

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

