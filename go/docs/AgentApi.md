# \AgentApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentList**](AgentApi.md#AgentList) | **Get** /api/v1/org/{orgUID}/agent/ | List agents
[**AgentLoad**](AgentApi.md#AgentLoad) | **Get** /api/v1/org/{orgUID}/agent/{agentUID} | Load an agent



## AgentList

> []Agent AgentList(ctx, orgUID).AgentRegistrationUidEquals(agentRegistrationUidEquals).OriginalAssociation(originalAssociation).Page(page).PageSize(pageSize).SourceUidEquals(sourceUidEquals).Execute()

List agents



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
    agentRegistrationUidEquals := "agentRegistrationUidEquals_example" // string |  (optional)
    originalAssociation := true // bool |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    sourceUidEquals := "sourceUidEquals_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentApi.AgentList(context.Background(), orgUID).AgentRegistrationUidEquals(agentRegistrationUidEquals).OriginalAssociation(originalAssociation).Page(page).PageSize(pageSize).SourceUidEquals(sourceUidEquals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentApi.AgentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentList`: []Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentApi.AgentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentRegistrationUidEquals** | **string** |  | 
 **originalAssociation** | **bool** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **sourceUidEquals** | **string** |  | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentLoad

> Agent AgentLoad(ctx, agentUID, orgUID).Execute()

Load an agent



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
    agentUID := "agentUID_example" // string | 
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentApi.AgentLoad(context.Background(), agentUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentApi.AgentLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentLoad`: Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentApi.AgentLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** |  | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agent**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

