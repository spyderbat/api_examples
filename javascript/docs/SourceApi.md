# Sbapi.SourceApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**integrationSoarSrcList**](SourceApi.md#integrationSoarSrcList) | **GET** /api/v1/integration/soar/org/{orgUID}/source/ | List sources for integration with SOARs
[**srcCreate**](SourceApi.md#srcCreate) | **POST** /api/v1/org/{orgUID}/source/ | Create a source
[**srcDelete**](SourceApi.md#srcDelete) | **DELETE** /api/v1/org/{orgUID}/source/{sourceUID} | Delete a source
[**srcList**](SourceApi.md#srcList) | **GET** /api/v1/org/{orgUID}/source/ | List sources
[**srcLoad**](SourceApi.md#srcLoad) | **GET** /api/v1/org/{orgUID}/source/{sourceUID} | Load a source
[**srcUpdate**](SourceApi.md#srcUpdate) | **PUT** /api/v1/org/{orgUID}/source/{sourceUID} | Update a source



## integrationSoarSrcList

> [ApiSOARListHandlerOutput] integrationSoarSrcList(orgUID, opts)

List sources for integration with SOARs

 Lists the sources of data that match the specified query parameters, and return  URL entry points into the UI for matching sources.   * Requires the action  *org:ListSources* on the organization 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.SourceApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'et': 789, // Number | optional end time of the query
  'hostname': "hostname_example", // String | A single hostname to match
  'ipAddress': "ipAddress_example", // String | A single IP address to match
  'macAddress': "macAddress_example", // String | A single mac address to match
  'page': 56, // Number | 
  'pageSize': 56, // Number | 
  'st': 789 // Number | optional start time of the query, if only a start time is provided, end time will be start+10m
};
apiInstance.integrationSoarSrcList(orgUID, opts, (error, data, response) => {
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
 **orgUID** | **String**|  | 
 **et** | **Number**| optional end time of the query | [optional] 
 **hostname** | **String**| A single hostname to match | [optional] 
 **ipAddress** | **String**| A single IP address to match | [optional] 
 **macAddress** | **String**| A single mac address to match | [optional] 
 **page** | **Number**|  | [optional] 
 **pageSize** | **Number**|  | [optional] 
 **st** | **Number**| optional start time of the query, if only a start time is provided, end time will be start+10m | [optional] 

### Return type

[**[ApiSOARListHandlerOutput]**](ApiSOARListHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## srcCreate

> ApiSourceCreateHandlerOutput srcCreate(orgUID, opts)

Create a source

 Creates a new source of data  * Requires the action  *org:CreateSource* on the organization 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.SourceApi();
let orgUID = "orgUID_example"; // String | The org this source is associated with
let opts = {
  'srcCreateInput': new Sbapi.SrcCreateInput() // SrcCreateInput | 
};
apiInstance.srcCreate(orgUID, opts, (error, data, response) => {
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
 **orgUID** | **String**| The org this source is associated with | 
 **srcCreateInput** | [**SrcCreateInput**](SrcCreateInput.md)|  | [optional] 

### Return type

[**ApiSourceCreateHandlerOutput**](ApiSourceCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## srcDelete

> srcDelete(orgUID, sourceUID)

Delete a source

 Delete a source  * Requires the action  *org:DeleteSource* on the organization 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.SourceApi();
let orgUID = "orgUID_example"; // String | 
let sourceUID = "sourceUID_example"; // String | 
apiInstance.srcDelete(orgUID, sourceUID, (error, data, response) => {
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
 **orgUID** | **String**|  | 
 **sourceUID** | **String**|  | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## srcList

> [Source] srcList(orgUID, opts)

List sources

 Lists the sources of data for an organization  * Requires the action  *org:ListSources* on the organization 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.SourceApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'agentUidEquals': "agentUidEquals_example", // String | 
  'descriptionContains': "descriptionContains_example", // String | 
  'hasTags': ["null"], // [String] | 
  'originalAssociation': true, // Boolean | 
  'page': 56, // Number | 
  'pageSize': 56 // Number | 
};
apiInstance.srcList(orgUID, opts, (error, data, response) => {
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
 **orgUID** | **String**|  | 
 **agentUidEquals** | **String**|  | [optional] 
 **descriptionContains** | **String**|  | [optional] 
 **hasTags** | [**[String]**](String.md)|  | [optional] 
 **originalAssociation** | **Boolean**|  | [optional] 
 **page** | **Number**|  | [optional] 
 **pageSize** | **Number**|  | [optional] 

### Return type

[**[Source]**](Source.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## srcLoad

> Source srcLoad(orgUID, sourceUID)

Load a source

 Loads the meta data for a source of data  * Requires the action  *org:LoadSource* on the organization 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.SourceApi();
let orgUID = "orgUID_example"; // String | 
let sourceUID = "sourceUID_example"; // String | 
apiInstance.srcLoad(orgUID, sourceUID, (error, data, response) => {
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
 **orgUID** | **String**|  | 
 **sourceUID** | **String**|  | 

### Return type

[**Source**](Source.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## srcUpdate

> srcUpdate(orgUID, sourceUID, opts)

Update a source

 Updates the meta data for a source of data  * Requires the action  *org:UpdateSource* on the organization 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.SourceApi();
let orgUID = "orgUID_example"; // String | The org this source is associated with
let sourceUID = "sourceUID_example"; // String | The UID of the source
let opts = {
  'srcUpdateInput': new Sbapi.SrcUpdateInput() // SrcUpdateInput | 
};
apiInstance.srcUpdate(orgUID, sourceUID, opts, (error, data, response) => {
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
 **orgUID** | **String**| The org this source is associated with | 
 **sourceUID** | **String**| The UID of the source | 
 **srcUpdateInput** | [**SrcUpdateInput**](SrcUpdateInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

