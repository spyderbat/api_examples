# \AgentRegistrationApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRegistrationCreate**](AgentRegistrationApi.md#AgentRegistrationCreate) | **Post** /api/v1/org/{orgUID}/agent_registration/ | Create an agent registration
[**AgentRegistrationDownloadLink**](AgentRegistrationApi.md#AgentRegistrationDownloadLink) | **Get** /api/v1/org/{orgUID}/agent_registration/{uid}/download_link | Get a download link for this registration
[**AgentRegistrationGetLog**](AgentRegistrationApi.md#AgentRegistrationGetLog) | **Get** /api/v1/org/{orgUID}/agent_registration/{uid}/log | Get log of recent agent registration activity
[**AgentRegistrationList**](AgentRegistrationApi.md#AgentRegistrationList) | **Get** /api/v1/org/{orgUID}/agent_registration/ | List agent registrations
[**AgentRegistrationLoad**](AgentRegistrationApi.md#AgentRegistrationLoad) | **Get** /api/v1/org/{orgUID}/agent_registration/{uid} | Load an agent registration
[**AgentRegistrationUpdate**](AgentRegistrationApi.md#AgentRegistrationUpdate) | **Put** /api/v1/org/{orgUID}/agent_registration/{uid} | Update an agent registration



## AgentRegistrationCreate

> ApiAgentCreateHandlerOutput AgentRegistrationCreate(ctx, orgUID).AgentRegistrationCreateInput(agentRegistrationCreateInput).Execute()

Create an agent registration



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
    orgUID := "orgUID_example" // string | The OrgUID the registration is associated with
    agentRegistrationCreateInput := *openapiclient.NewAgentRegistrationCreateInput() // AgentRegistrationCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentRegistrationApi.AgentRegistrationCreate(context.Background(), orgUID).AgentRegistrationCreateInput(agentRegistrationCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationApi.AgentRegistrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentRegistrationCreate`: ApiAgentCreateHandlerOutput
    fmt.Fprintf(os.Stdout, "Response from `AgentRegistrationApi.AgentRegistrationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | The OrgUID the registration is associated with | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRegistrationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentRegistrationCreateInput** | [**AgentRegistrationCreateInput**](AgentRegistrationCreateInput.md) |  | 

### Return type

[**ApiAgentCreateHandlerOutput**](ApiAgentCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRegistrationDownloadLink

> ApiAgentRegistrationDownloadLinkHandlerOutput AgentRegistrationDownloadLink(ctx, orgUID, uid).Execute()

Get a download link for this registration



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
    uid := "uid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentRegistrationApi.AgentRegistrationDownloadLink(context.Background(), orgUID, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationApi.AgentRegistrationDownloadLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentRegistrationDownloadLink`: ApiAgentRegistrationDownloadLinkHandlerOutput
    fmt.Fprintf(os.Stdout, "Response from `AgentRegistrationApi.AgentRegistrationDownloadLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRegistrationDownloadLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiAgentRegistrationDownloadLinkHandlerOutput**](ApiAgentRegistrationDownloadLinkHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRegistrationGetLog

> []DaoAgentLog AgentRegistrationGetLog(ctx, orgUID, uid).Page(page).PageSize(pageSize).Execute()

Get log of recent agent registration activity



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
    uid := "uid_example" // string | 
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentRegistrationApi.AgentRegistrationGetLog(context.Background(), orgUID, uid).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationApi.AgentRegistrationGetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentRegistrationGetLog`: []DaoAgentLog
    fmt.Fprintf(os.Stdout, "Response from `AgentRegistrationApi.AgentRegistrationGetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRegistrationGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**[]DaoAgentLog**](DaoAgentLog.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRegistrationList

> []AgentRegistration AgentRegistrationList(ctx, orgUID).Page(page).PageSize(pageSize).Execute()

List agent registrations



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
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentRegistrationApi.AgentRegistrationList(context.Background(), orgUID).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationApi.AgentRegistrationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentRegistrationList`: []AgentRegistration
    fmt.Fprintf(os.Stdout, "Response from `AgentRegistrationApi.AgentRegistrationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRegistrationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**[]AgentRegistration**](AgentRegistration.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRegistrationLoad

> AgentRegistration AgentRegistrationLoad(ctx, orgUID, uid).Execute()

Load an agent registration



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
    uid := "uid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentRegistrationApi.AgentRegistrationLoad(context.Background(), orgUID, uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationApi.AgentRegistrationLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentRegistrationLoad`: AgentRegistration
    fmt.Fprintf(os.Stdout, "Response from `AgentRegistrationApi.AgentRegistrationLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRegistrationLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentRegistration**](AgentRegistration.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRegistrationUpdate

> AgentRegistrationUpdate(ctx, orgUID, uid).AgentRegistrationUpdateInput(agentRegistrationUpdateInput).Execute()

Update an agent registration



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
    orgUID := "orgUID_example" // string | The OrgUID the registration is associated with
    uid := "uid_example" // string | Agent Registration UID
    agentRegistrationUpdateInput := *openapiclient.NewAgentRegistrationUpdateInput() // AgentRegistrationUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentRegistrationApi.AgentRegistrationUpdate(context.Background(), orgUID, uid).AgentRegistrationUpdateInput(agentRegistrationUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentRegistrationApi.AgentRegistrationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | The OrgUID the registration is associated with | 
**uid** | **string** | Agent Registration UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRegistrationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentRegistrationUpdateInput** | [**AgentRegistrationUpdateInput**](AgentRegistrationUpdateInput.md) |  | 

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

