# Sbapi.MetadataAPIApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metadataSearchSchema**](MetadataAPIApi.md#metadataSearchSchema) | **GET** /api/v1/meta/search/schema | Returns the schema used for search



## metadataSearchSchema

> {String: ElasticRecordSchema} metadataSearchSchema()

Returns the schema used for search

 Returns meta data around which model and event properties are indexed for search, i.e. the search schema 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.MetadataAPIApi();
apiInstance.metadataSearchSchema((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**{String: ElasticRecordSchema}**](ElasticRecordSchema.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

