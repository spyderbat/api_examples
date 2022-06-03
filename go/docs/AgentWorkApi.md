# \AgentWorkApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentDeleteAgentWork**](AgentWorkApi.md#AgentDeleteAgentWork) | **Delete** /api/v1/org/{orgUID}/agent/{agentUID}/work | Delete agent work data for a specific agent
[**AgentDeleteOrgWork**](AgentWorkApi.md#AgentDeleteOrgWork) | **Delete** /api/v1/org/{orgUID}/agent_work | Delete agent work for an org
[**AgentGetAgentWork**](AgentWorkApi.md#AgentGetAgentWork) | **Get** /api/v1/org/{orgUID}/agent/{agentUID}/work | Get agent work data for a specific agent
[**AgentGetOrgWork**](AgentWorkApi.md#AgentGetOrgWork) | **Get** /api/v1/org/{orgUID}/agent_work | Get agent work data for the organization
[**AgentSetAgentWork**](AgentWorkApi.md#AgentSetAgentWork) | **Post** /api/v1/org/{orgUID}/agent/{agentUID}/work | Set agent work data for a specific agent
[**AgentSetOrgWork**](AgentWorkApi.md#AgentSetOrgWork) | **Post** /api/v1/org/{orgUID}/agent_work | Set agent work data for a specific agent



## AgentDeleteAgentWork

> AgentDeleteAgentWork(ctx, agentUID, orgUID).Execute()

Delete agent work data for a specific agent



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
    agentUID := "agentUID_example" // string | Agent UID
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentWorkApi.AgentDeleteAgentWork(context.Background(), agentUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentWorkApi.AgentDeleteAgentWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** | Agent UID | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentDeleteAgentWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDeleteOrgWork

> AgentDeleteOrgWork(ctx, agentUID, orgUID).Execute()

Delete agent work for an org



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
    agentUID := "agentUID_example" // string | Agent UID
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentWorkApi.AgentDeleteOrgWork(context.Background(), agentUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentWorkApi.AgentDeleteOrgWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** | Agent UID | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentDeleteOrgWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetAgentWork

> ApiAgentWorkOutput AgentGetAgentWork(ctx, agentUID, orgUID).Execute()

Get agent work data for a specific agent



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
    agentUID := "agentUID_example" // string | Agent UID
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentWorkApi.AgentGetAgentWork(context.Background(), agentUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentWorkApi.AgentGetAgentWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentGetAgentWork`: ApiAgentWorkOutput
    fmt.Fprintf(os.Stdout, "Response from `AgentWorkApi.AgentGetAgentWork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** | Agent UID | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetAgentWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetOrgWork

> ApiAgentWorkOutput AgentGetOrgWork(ctx, agentUID, orgUID).Execute()

Get agent work data for the organization



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
    agentUID := "agentUID_example" // string | Agent UID
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentWorkApi.AgentGetOrgWork(context.Background(), agentUID, orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentWorkApi.AgentGetOrgWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentGetOrgWork`: ApiAgentWorkOutput
    fmt.Fprintf(os.Stdout, "Response from `AgentWorkApi.AgentGetOrgWork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** | Agent UID | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetOrgWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentSetAgentWork

> AgentSetAgentWork(ctx, agentUID, orgUID).AgentSetAgentWorkInput(agentSetAgentWorkInput).Execute()

Set agent work data for a specific agent



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
    agentUID := "agentUID_example" // string | Agent UID
    orgUID := "orgUID_example" // string | 
    agentSetAgentWorkInput := *openapiclient.NewAgentSetAgentWorkInput() // AgentSetAgentWorkInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentWorkApi.AgentSetAgentWork(context.Background(), agentUID, orgUID).AgentSetAgentWorkInput(agentSetAgentWorkInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentWorkApi.AgentSetAgentWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** | Agent UID | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentSetAgentWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentSetAgentWorkInput** | [**AgentSetAgentWorkInput**](AgentSetAgentWorkInput.md) |  | 

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


## AgentSetOrgWork

> AgentSetOrgWork(ctx, agentUID, orgUID).AgentSetOrgWorkInput(agentSetOrgWorkInput).Execute()

Set agent work data for a specific agent



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
    agentUID := "agentUID_example" // string | Agent UID
    orgUID := "orgUID_example" // string | 
    agentSetOrgWorkInput := *openapiclient.NewAgentSetOrgWorkInput() // AgentSetOrgWorkInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentWorkApi.AgentSetOrgWork(context.Background(), agentUID, orgUID).AgentSetOrgWorkInput(agentSetOrgWorkInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentWorkApi.AgentSetOrgWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUID** | **string** | Agent UID | 
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentSetOrgWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentSetOrgWorkInput** | [**AgentSetOrgWorkInput**](AgentSetOrgWorkInput.md) |  | 

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

