# \OrgTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgTypeLimitActiveSources**](OrgTypeApi.md#OrgTypeLimitActiveSources) | **Get** /api/v1/org/{orgUID}/type/limit/active_sources | Loads limits regarding active sources
[**OrgTypeLimitOrgRoles**](OrgTypeApi.md#OrgTypeLimitOrgRoles) | **Get** /api/v1/org/{orgUID}/type/limit/org_roles | Loads limits regarding org roles
[**OrgTypeLoad**](OrgTypeApi.md#OrgTypeLoad) | **Get** /api/v1/org/{orgUID}/type | Load the org type for the organization



## OrgTypeLimitActiveSources

> SessionOrgTypeMaxLimit OrgTypeLimitActiveSources(ctx, orgUID).Execute()

Loads limits regarding active sources



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
    resp, r, err := apiClient.OrgTypeApi.OrgTypeLimitActiveSources(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgTypeApi.OrgTypeLimitActiveSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgTypeLimitActiveSources`: SessionOrgTypeMaxLimit
    fmt.Fprintf(os.Stdout, "Response from `OrgTypeApi.OrgTypeLimitActiveSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgTypeLimitActiveSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgTypeLimitOrgRoles

> SessionOrgTypeMaxLimit OrgTypeLimitOrgRoles(ctx, orgUID).Execute()

Loads limits regarding org roles



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
    resp, r, err := apiClient.OrgTypeApi.OrgTypeLimitOrgRoles(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgTypeApi.OrgTypeLimitOrgRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgTypeLimitOrgRoles`: SessionOrgTypeMaxLimit
    fmt.Fprintf(os.Stdout, "Response from `OrgTypeApi.OrgTypeLimitOrgRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgTypeLimitOrgRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgTypeLoad

> DaoOrgType OrgTypeLoad(ctx, orgUID).Execute()

Load the org type for the organization



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
    resp, r, err := apiClient.OrgTypeApi.OrgTypeLoad(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgTypeApi.OrgTypeLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgTypeLoad`: DaoOrgType
    fmt.Fprintf(os.Stdout, "Response from `OrgTypeApi.OrgTypeLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgTypeLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DaoOrgType**](DaoOrgType.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

