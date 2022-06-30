# Sbapi.APIKeyApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiKeyCreate**](APIKeyApi.md#apiKeyCreate) | **POST** /api/v1/user/{userUID}/apikey/ | Creates a new API key
[**apiKeyDelete**](APIKeyApi.md#apiKeyDelete) | **DELETE** /api/v1/user/{userUID}/apikey/{uid} | Delete an API key
[**apiKeyList**](APIKeyApi.md#apiKeyList) | **GET** /api/v1/user/{userUID}/apikey/ | Lists an API key
[**apiKeyUpdate**](APIKeyApi.md#apiKeyUpdate) | **PUT** /api/v1/user/{userUID}/apikey/{uid} | Updates an API key



## apiKeyCreate

> String apiKeyCreate(userUID, opts)

Creates a new API key

 Creates a new API key which is associated with the user    * Requires global action *user:APIKeyCreate* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.APIKeyApi();
let userUID = "userUID_example"; // String | User UID
let opts = {
  'apiKeyCreateInput': new Sbapi.ApiKeyCreateInput() // ApiKeyCreateInput | 
};
apiInstance.apiKeyCreate(userUID, opts, (error, data, response) => {
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
 **userUID** | **String**| User UID | 
 **apiKeyCreateInput** | [**ApiKeyCreateInput**](ApiKeyCreateInput.md)|  | [optional] 

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## apiKeyDelete

> apiKeyDelete(uid, userUID)

Delete an API key

 Deletes a specific API key   * Requires global action *user:APIKeyDelete* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.APIKeyApi();
let uid = "uid_example"; // String | API Key UID
let userUID = "userUID_example"; // String | User UID
apiInstance.apiKeyDelete(uid, userUID, (error, data, response) => {
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
 **uid** | **String**| API Key UID | 
 **userUID** | **String**| User UID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## apiKeyList

> [APIKey] apiKeyList(userUID)

Lists an API key

 Lists API keys associated with a user   * Requires global action *user:APIKeyList*  

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.APIKeyApi();
let userUID = "userUID_example"; // String | User UID
apiInstance.apiKeyList(userUID, (error, data, response) => {
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
 **userUID** | **String**| User UID | 

### Return type

[**[APIKey]**](APIKey.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## apiKeyUpdate

> apiKeyUpdate(uid, userUID, opts)

Updates an API key

 Updates a specific API Key, the only fields which can be updated are description and enabled   * Requires global action *user:APIKeyUpdate* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.APIKeyApi();
let uid = "uid_example"; // String | API Key UID
let userUID = "userUID_example"; // String | User UID
let opts = {
  'apiKeyUpdateInput': new Sbapi.ApiKeyUpdateInput() // ApiKeyUpdateInput | 
};
apiInstance.apiKeyUpdate(uid, userUID, opts, (error, data, response) => {
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
 **uid** | **String**| API Key UID | 
 **userUID** | **String**| User UID | 
 **apiKeyUpdateInput** | [**ApiKeyUpdateInput**](ApiKeyUpdateInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

