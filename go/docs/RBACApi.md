# \RBACApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CanUserPerform**](RBACApi.md#CanUserPerform) | **Post** /api/v1/rbac/capabilities/ | Query allows actions on objects



## CanUserPerform

> ApiRBACActions CanUserPerform(ctx).CanUserPerformInput(canUserPerformInput).Execute()

Query allows actions on objects



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
    canUserPerformInput := *openapiclient.NewCanUserPerformInput() // CanUserPerformInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RBACApi.CanUserPerform(context.Background()).CanUserPerformInput(canUserPerformInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACApi.CanUserPerform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CanUserPerform`: ApiRBACActions
    fmt.Fprintf(os.Stdout, "Response from `RBACApi.CanUserPerform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCanUserPerformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canUserPerformInput** | [**CanUserPerformInput**](CanUserPerformInput.md) |  | 

### Return type

[**ApiRBACActions**](ApiRBACActions.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

