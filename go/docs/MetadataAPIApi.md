# \MetadataAPIApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataSearchSchema**](MetadataAPIApi.md#MetadataSearchSchema) | **Get** /api/v1/meta/search/schema | Returns the schema used for search



## MetadataSearchSchema

> map[string]ElasticRecordSchema MetadataSearchSchema(ctx).Execute()

Returns the schema used for search



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataAPIApi.MetadataSearchSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPIApi.MetadataSearchSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataSearchSchema`: map[string]ElasticRecordSchema
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPIApi.MetadataSearchSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataSearchSchemaRequest struct via the builder pattern


### Return type

[**map[string]ElasticRecordSchema**](ElasticRecordSchema.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

