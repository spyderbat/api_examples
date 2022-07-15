# \SourceDataApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SrcDataQuery**](SourceDataApi.md#SrcDataQuery) | **Post** /api/v1/source/query/ | Query source data
[**SrcDataQueryV2**](SourceDataApi.md#SrcDataQueryV2) | **Get** /api/v1/org/{orgUID}/data/ | Query source data
[**SrcSendData**](SourceDataApi.md#SrcSendData) | **Post** /api/v1/org/{orgUID}/source/{sourceUID}/data/{dataType} | Send data to a source, this is expected to be gzip compressed nd-json. The &#39;Content-Encoding&#39; header should be specified with a value of &#39;gzip&#39;



## SrcDataQuery

> string SrcDataQuery(ctx).SrcDataQueryInput(srcDataQueryInput).Execute()

Query source data



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
    srcDataQueryInput := *openapiclient.NewSrcDataQueryInput("DataType_example", "OrgUid_example") // SrcDataQueryInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDataApi.SrcDataQuery(context.Background()).SrcDataQueryInput(srcDataQueryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDataApi.SrcDataQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SrcDataQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `SourceDataApi.SrcDataQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSrcDataQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcDataQueryInput** | [**SrcDataQueryInput**](SrcDataQueryInput.md) |  | 

### Return type

**string**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrcDataQueryV2

> string SrcDataQueryV2(ctx, orgUID).Dt(dt).E(e).Et(et).Id(id).Pj(pj).Q(q).Qf(qf).Qs(qs).Rr(rr).Src(src).St(st).Execute()

Query source data



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
    orgUID := "orgUID_example" // string | Organization UID to query
    dt := "dt_example" // string | DataType to query
    e := "e_example" // string | Data which matches this expression are returned, as a json object (optional)
    et := float64(1.2) // float64 | Time in unix epoch time, records < time are returned (optional)
    id := []string{"Inner_example"} // []string | List of IDs to resolve (optional)
    pj := []string{"Inner_example"} // []string | Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj=id&pj=version (optional)
    q := "q_example" // string | Lucene based search query (optional)
    qf := int32(56) // int32 | Where to start the query in the result set from (optional)
    qs := int32(56) // int32 | Size of the query result set (optional)
    rr := true // bool | Return rptrs mixed with the data (optional)
    src := "src_example" // string | Source UID to query (optional)
    st := float64(1.2) // float64 | Time in unix epoch time, records >= time are returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDataApi.SrcDataQueryV2(context.Background(), orgUID).Dt(dt).E(e).Et(et).Id(id).Pj(pj).Q(q).Qf(qf).Qs(qs).Rr(rr).Src(src).St(st).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDataApi.SrcDataQueryV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SrcDataQueryV2`: string
    fmt.Fprintf(os.Stdout, "Response from `SourceDataApi.SrcDataQueryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | Organization UID to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcDataQueryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dt** | **string** | DataType to query | 
 **e** | **string** | Data which matches this expression are returned, as a json object | 
 **et** | **float64** | Time in unix epoch time, records &lt; time are returned | 
 **id** | **[]string** | List of IDs to resolve | 
 **pj** | **[]string** | Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj&#x3D;id&amp;pj&#x3D;version | 
 **q** | **string** | Lucene based search query | 
 **qf** | **int32** | Where to start the query in the result set from | 
 **qs** | **int32** | Size of the query result set | 
 **rr** | **bool** | Return rptrs mixed with the data | 
 **src** | **string** | Source UID to query | 
 **st** | **float64** | Time in unix epoch time, records &gt;&#x3D; time are returned | 

### Return type

**string**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrcSendData

> SrcSendData(ctx, dataType, orgUID, sourceUID).Execute()

Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'



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
    dataType := "dataType_example" // string | 
    orgUID := "orgUID_example" // string | 
    sourceUID := "sourceUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDataApi.SrcSendData(context.Background(), dataType, orgUID, sourceUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDataApi.SrcSendData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataType** | **string** |  | 
**orgUID** | **string** |  | 
**sourceUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSrcSendDataRequest struct via the builder pattern


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

