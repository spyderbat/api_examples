# \UIDataApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UiDataDeleteOrgData**](UIDataApi.md#UiDataDeleteOrgData) | **Delete** /api/v1/org/{orgUID}/uidata/{dataKey} | Delete Org UI Data
[**UiDataDeleteSourceData**](UIDataApi.md#UiDataDeleteSourceData) | **Delete** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Delete Source UI Data
[**UiDataDeleteUserData**](UIDataApi.md#UiDataDeleteUserData) | **Delete** /api/v1/user/{userUID}/uidata/{dataKey} | Delete User UI Data
[**UiDataDeleteUserOrgData**](UIDataApi.md#UiDataDeleteUserOrgData) | **Delete** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Delete UserOrg UI Data
[**UiDataDeleteUserSourceData**](UIDataApi.md#UiDataDeleteUserSourceData) | **Delete** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Delete UserSource UI Data
[**UiDataGetOrgData**](UIDataApi.md#UiDataGetOrgData) | **Get** /api/v1/org/{orgUID}/uidata/{dataKey} | Get Org UI Data
[**UiDataGetSourceData**](UIDataApi.md#UiDataGetSourceData) | **Get** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Get Source UI Data
[**UiDataGetUserData**](UIDataApi.md#UiDataGetUserData) | **Get** /api/v1/user/{userUID}/uidata/{dataKey} | Get User UI Data
[**UiDataGetUserOrgData**](UIDataApi.md#UiDataGetUserOrgData) | **Get** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Get UserOrg UI Data
[**UiDataGetUserSourceData**](UIDataApi.md#UiDataGetUserSourceData) | **Get** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Get UserSource UI Data
[**UiDataQueryOrgData**](UIDataApi.md#UiDataQueryOrgData) | **Get** /api/v1/org/{orgUID}/uidata/ | Query Org UI Data
[**UiDataQuerySourceData**](UIDataApi.md#UiDataQuerySourceData) | **Get** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/ | Query Source UI Data
[**UiDataQueryUserData**](UIDataApi.md#UiDataQueryUserData) | **Get** /api/v1/user/{userUID}/uidata/ | Query User UI Data
[**UiDataQueryUserOrgData**](UIDataApi.md#UiDataQueryUserOrgData) | **Get** /api/v1/user/{userUID}/org/{orgUID}/uidata/ | Query UserOrg UI Data
[**UiDataQueryUserSourceData**](UIDataApi.md#UiDataQueryUserSourceData) | **Get** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/ | Query UserSource UI Data
[**UiDataSetOrgData**](UIDataApi.md#UiDataSetOrgData) | **Put** /api/v1/org/{orgUID}/uidata/{dataKey} | Set Org UI Data
[**UiDataSetSourceData**](UIDataApi.md#UiDataSetSourceData) | **Put** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Set Source UI Data
[**UiDataSetUserData**](UIDataApi.md#UiDataSetUserData) | **Put** /api/v1/user/{userUID}/uidata/{dataKey} | Set User UI Data
[**UiDataSetUserOrgData**](UIDataApi.md#UiDataSetUserOrgData) | **Put** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Set UserOrg UI Data
[**UiDataSetUserSourceData**](UIDataApi.md#UiDataSetUserSourceData) | **Put** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Set UserSource UI Data



## UiDataDeleteOrgData

> UiDataDeleteOrgData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Delete Org UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataDeleteOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataDeleteOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataDeleteOrgDataRequest struct via the builder pattern


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


## UiDataDeleteSourceData

> UiDataDeleteSourceData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Delete Source UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataDeleteSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataDeleteSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataDeleteSourceDataRequest struct via the builder pattern


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


## UiDataDeleteUserData

> UiDataDeleteUserData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Delete User UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataDeleteUserData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataDeleteUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataDeleteUserDataRequest struct via the builder pattern


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


## UiDataDeleteUserOrgData

> UiDataDeleteUserOrgData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Delete UserOrg UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataDeleteUserOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataDeleteUserOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataDeleteUserOrgDataRequest struct via the builder pattern


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


## UiDataDeleteUserSourceData

> UiDataDeleteUserSourceData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Delete UserSource UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataDeleteUserSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataDeleteUserSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataDeleteUserSourceDataRequest struct via the builder pattern


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


## UiDataGetOrgData

> UIData UiDataGetOrgData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Get Org UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataGetOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataGetOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataGetOrgData`: UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataGetOrgData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataGetOrgDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataGetSourceData

> UIData UiDataGetSourceData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Get Source UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataGetSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataGetSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataGetSourceData`: UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataGetSourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataGetSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataGetUserData

> UIData UiDataGetUserData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Get User UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataGetUserData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataGetUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataGetUserData`: UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataGetUserData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataGetUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataGetUserOrgData

> UIData UiDataGetUserOrgData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Get UserOrg UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataGetUserOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataGetUserOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataGetUserOrgData`: UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataGetUserOrgData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataGetUserOrgDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataGetUserSourceData

> UIData UiDataGetUserSourceData(ctx, dataKey, orgUID, sourceUID, userUID).Execute()

Get UserSource UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataGetUserSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataGetUserSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataGetUserSourceData`: UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataGetUserSourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataGetUserSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataQueryOrgData

> []UIData UiDataQueryOrgData(ctx, dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()

Query Org UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID
    tags := []string{"Inner_example"} // []string | Tags to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataQueryOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataQueryOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataQueryOrgData`: []UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataQueryOrgData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataQueryOrgDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tags** | **[]string** | Tags to search for | 

### Return type

[**[]UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataQuerySourceData

> []UIData UiDataQuerySourceData(ctx, dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()

Query Source UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID
    tags := []string{"Inner_example"} // []string | Tags to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataQuerySourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataQuerySourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataQuerySourceData`: []UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataQuerySourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataQuerySourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tags** | **[]string** | Tags to search for | 

### Return type

[**[]UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataQueryUserData

> []UIData UiDataQueryUserData(ctx, dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()

Query User UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID
    tags := []string{"Inner_example"} // []string | Tags to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataQueryUserData(context.Background(), dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataQueryUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataQueryUserData`: []UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataQueryUserData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataQueryUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tags** | **[]string** | Tags to search for | 

### Return type

[**[]UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataQueryUserOrgData

> []UIData UiDataQueryUserOrgData(ctx, dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()

Query UserOrg UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID
    tags := []string{"Inner_example"} // []string | Tags to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataQueryUserOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataQueryUserOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataQueryUserOrgData`: []UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataQueryUserOrgData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataQueryUserOrgDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tags** | **[]string** | Tags to search for | 

### Return type

[**[]UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataQueryUserSourceData

> []UIData UiDataQueryUserSourceData(ctx, dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()

Query UserSource UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Src UID
    userUID := "userUID_example" // string | Owner UID
    tags := []string{"Inner_example"} // []string | Tags to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataQueryUserSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataQueryUserSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UiDataQueryUserSourceData`: []UIData
    fmt.Fprintf(os.Stdout, "Response from `UIDataApi.UiDataQueryUserSourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Src UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataQueryUserSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tags** | **[]string** | Tags to search for | 

### Return type

[**[]UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UiDataSetOrgData

> UiDataSetOrgData(ctx, dataKey, orgUID, sourceUID, userUID).UiDataSetOrgDataInput(uiDataSetOrgDataInput).Execute()

Set Org UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Source UID
    userUID := "userUID_example" // string | Owner UID
    uiDataSetOrgDataInput := *openapiclient.NewUiDataSetOrgDataInput() // UiDataSetOrgDataInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataSetOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).UiDataSetOrgDataInput(uiDataSetOrgDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataSetOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Source UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataSetOrgDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **uiDataSetOrgDataInput** | [**UiDataSetOrgDataInput**](UiDataSetOrgDataInput.md) |  | 

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


## UiDataSetSourceData

> UiDataSetSourceData(ctx, dataKey, orgUID, sourceUID, userUID).UiDataSetSourceDataInput(uiDataSetSourceDataInput).Execute()

Set Source UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Source UID
    userUID := "userUID_example" // string | Owner UID
    uiDataSetSourceDataInput := *openapiclient.NewUiDataSetSourceDataInput() // UiDataSetSourceDataInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataSetSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).UiDataSetSourceDataInput(uiDataSetSourceDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataSetSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Source UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataSetSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **uiDataSetSourceDataInput** | [**UiDataSetSourceDataInput**](UiDataSetSourceDataInput.md) |  | 

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


## UiDataSetUserData

> UiDataSetUserData(ctx, dataKey, orgUID, sourceUID, userUID).UiDataSetUserDataInput(uiDataSetUserDataInput).Execute()

Set User UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Source UID
    userUID := "userUID_example" // string | Owner UID
    uiDataSetUserDataInput := *openapiclient.NewUiDataSetUserDataInput() // UiDataSetUserDataInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataSetUserData(context.Background(), dataKey, orgUID, sourceUID, userUID).UiDataSetUserDataInput(uiDataSetUserDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataSetUserData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Source UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataSetUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **uiDataSetUserDataInput** | [**UiDataSetUserDataInput**](UiDataSetUserDataInput.md) |  | 

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


## UiDataSetUserOrgData

> UiDataSetUserOrgData(ctx, dataKey, orgUID, sourceUID, userUID).UiDataSetUserOrgDataInput(uiDataSetUserOrgDataInput).Execute()

Set UserOrg UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Source UID
    userUID := "userUID_example" // string | Owner UID
    uiDataSetUserOrgDataInput := *openapiclient.NewUiDataSetUserOrgDataInput() // UiDataSetUserOrgDataInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataSetUserOrgData(context.Background(), dataKey, orgUID, sourceUID, userUID).UiDataSetUserOrgDataInput(uiDataSetUserOrgDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataSetUserOrgData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Source UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataSetUserOrgDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **uiDataSetUserOrgDataInput** | [**UiDataSetUserOrgDataInput**](UiDataSetUserOrgDataInput.md) |  | 

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


## UiDataSetUserSourceData

> UiDataSetUserSourceData(ctx, dataKey, orgUID, sourceUID, userUID).UiDataSetUserSourceDataInput(uiDataSetUserSourceDataInput).Execute()

Set UserSource UI Data



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
    dataKey := "dataKey_example" // string | Key for the data
    orgUID := "orgUID_example" // string | Org UID
    sourceUID := "sourceUID_example" // string | Source UID
    userUID := "userUID_example" // string | Owner UID
    uiDataSetUserSourceDataInput := *openapiclient.NewUiDataSetUserSourceDataInput() // UiDataSetUserSourceDataInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UIDataApi.UiDataSetUserSourceData(context.Background(), dataKey, orgUID, sourceUID, userUID).UiDataSetUserSourceDataInput(uiDataSetUserSourceDataInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UIDataApi.UiDataSetUserSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataKey** | **string** | Key for the data | 
**orgUID** | **string** | Org UID | 
**sourceUID** | **string** | Source UID | 
**userUID** | **string** | Owner UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUiDataSetUserSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **uiDataSetUserSourceDataInput** | [**UiDataSetUserSourceDataInput**](UiDataSetUserSourceDataInput.md) |  | 

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

