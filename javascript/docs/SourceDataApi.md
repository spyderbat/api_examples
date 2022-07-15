# SpyderbatApi.SourceDataApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**srcDataQuery**](SourceDataApi.md#srcDataQuery) | **POST** /api/v1/source/query/ | Query source data
[**srcDataQueryV2**](SourceDataApi.md#srcDataQueryV2) | **GET** /api/v1/org/{orgUID}/data/ | Query source data
[**srcSendData**](SourceDataApi.md#srcSendData) | **POST** /api/v1/org/{orgUID}/source/{sourceUID}/data/{dataType} | Send data to a source, this is expected to be gzip compressed nd-json. The &#39;Content-Encoding&#39; header should be specified with a value of &#39;gzip&#39;



## srcDataQuery

> String srcDataQuery(opts)

Query source data

 Allows querying of the source data, data is stored as &#39;records&#39; which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 24 hour period of time, it is the callers responsibility to construct an appropriate 24 hour range. * Documentation for the returned spydergraph datatype can be found at https://app.spyderbat.com/schema/spydergraph/index.html  * The user must have both the *org.ListSourceData* action on the org and *source_data.Query* action on the source  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.SourceDataApi();
let opts = {
  'srcDataQueryInput': new SpyderbatApi.SrcDataQueryInput() // SrcDataQueryInput | 
};
apiInstance.srcDataQuery(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcDataQueryInput** | [**SrcDataQueryInput**](SrcDataQueryInput.md)|  | [optional] 

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json


## srcDataQueryV2

> String srcDataQueryV2(orgUID, dt, opts)

Query source data

Same as the post query above except results are cached

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.SourceDataApi();
let orgUID = "orgUID_example"; // String | Organization UID to query
let dt = "dt_example"; // String | DataType to query
let opts = {
  'e': "e_example", // String | Data which matches this expression are returned, as a json object
  'et': 3.4, // Number | Time in unix epoch time, records < time are returned
  'id': ["null"], // [String] | List of IDs to resolve
  'pj': ["null"], // [String] | Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj=id&pj=version
  'q': "q_example", // String | Lucene based search query
  'qf': 56, // Number | Where to start the query in the result set from
  'qs': 56, // Number | Size of the query result set
  'rr': true, // Boolean | Return rptrs mixed with the data
  'src': "src_example", // String | Source UID to query
  'st': 3.4 // Number | Time in unix epoch time, records >= time are returned
};
apiInstance.srcDataQueryV2(orgUID, dt, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**| Organization UID to query | 
 **dt** | **String**| DataType to query | 
 **e** | **String**| Data which matches this expression are returned, as a json object | [optional] 
 **et** | **Number**| Time in unix epoch time, records &lt; time are returned | [optional] 
 **id** | [**[String]**](String.md)| List of IDs to resolve | [optional] 
 **pj** | [**[String]**](String.md)| Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj&#x3D;id&amp;pj&#x3D;version | [optional] 
 **q** | **String**| Lucene based search query | [optional] 
 **qf** | **Number**| Where to start the query in the result set from | [optional] 
 **qs** | **Number**| Size of the query result set | [optional] 
 **rr** | **Boolean**| Return rptrs mixed with the data | [optional] 
 **src** | **String**| Source UID to query | [optional] 
 **st** | **Number**| Time in unix epoch time, records &gt;&#x3D; time are returned | [optional] 

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json


## srcSendData

> srcSendData(dataType, orgUID, sourceUID)

Send data to a source, this is expected to be gzip compressed nd-json. The &#39;Content-Encoding&#39; header should be specified with a value of &#39;gzip&#39;

Sends data to a source

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.SourceDataApi();
let dataType = "dataType_example"; // String | 
let orgUID = "orgUID_example"; // String | 
let sourceUID = "sourceUID_example"; // String | 
apiInstance.srcSendData(dataType, orgUID, sourceUID, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataType** | **String**|  | 
 **orgUID** | **String**|  | 
 **sourceUID** | **String**|  | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

