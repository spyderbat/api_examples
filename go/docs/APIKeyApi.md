# \APIKeyApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyCreate**](APIKeyApi.md#ApiKeyCreate) | **Post** /api/v1/user/{userUID}/apikey/ | Creates a new API key
[**ApiKeyDelete**](APIKeyApi.md#ApiKeyDelete) | **Delete** /api/v1/user/{userUID}/apikey/{uid} | Delete an API key
[**ApiKeyList**](APIKeyApi.md#ApiKeyList) | **Get** /api/v1/user/{userUID}/apikey/ | Lists an API key
[**ApiKeyUpdate**](APIKeyApi.md#ApiKeyUpdate) | **Put** /api/v1/user/{userUID}/apikey/{uid} | Updates an API key



## ApiKeyCreate

> string ApiKeyCreate(ctx, userUID).ApiKeyCreateInput(apiKeyCreateInput).Execute()

Creates a new API key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    userUID := "userUID_example" // string | User UID
    apiKeyCreateInput := *openapiclient.NewApiKeyCreateInput(time.Now()) // ApiKeyCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeyApi.ApiKeyCreate(context.Background(), userUID).ApiKeyCreateInput(apiKeyCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeyApi.ApiKeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiKeyCreate`: string
    fmt.Fprintf(os.Stdout, "Response from `APIKeyApi.ApiKeyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUID** | **string** | User UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKeyCreateInput** | [**ApiKeyCreateInput**](ApiKeyCreateInput.md) |  | 

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


## ApiKeyDelete

> ApiKeyDelete(ctx, uid, userUID).Execute()

Delete an API key



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
    uid := "uid_example" // string | API Key UID
    userUID := "userUID_example" // string | User UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeyApi.ApiKeyDelete(context.Background(), uid, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeyApi.ApiKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | API Key UID | 
**userUID** | **string** | User UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyDeleteRequest struct via the builder pattern


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


## ApiKeyList

> []APIKey ApiKeyList(ctx, userUID).Execute()

Lists an API key



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
    userUID := "userUID_example" // string | User UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeyApi.ApiKeyList(context.Background(), userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeyApi.ApiKeyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiKeyList`: []APIKey
    fmt.Fprintf(os.Stdout, "Response from `APIKeyApi.ApiKeyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUID** | **string** | User UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]APIKey**](APIKey.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyUpdate

> ApiKeyUpdate(ctx, uid, userUID).ApiKeyUpdateInput(apiKeyUpdateInput).Execute()

Updates an API key



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
    uid := "uid_example" // string | API Key UID
    userUID := "userUID_example" // string | User UID
    apiKeyUpdateInput := *openapiclient.NewApiKeyUpdateInput() // ApiKeyUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeyApi.ApiKeyUpdate(context.Background(), uid, userUID).ApiKeyUpdateInput(apiKeyUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeyApi.ApiKeyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | API Key UID | 
**userUID** | **string** | User UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKeyUpdateInput** | [**ApiKeyUpdateInput**](ApiKeyUpdateInput.md) |  | 

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

