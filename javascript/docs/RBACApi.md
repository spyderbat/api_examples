# SpyderbatApi.RBACApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**canUserPerform**](RBACApi.md#canUserPerform) | **POST** /api/v1/rbac/capabilities/ | Query allows actions on objects



## canUserPerform

> ApiRBACActions canUserPerform(opts)

Query allows actions on objects

Allows for querying of what actions a user can perform; results may be cached for a short period of time.

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.RBACApi();
let opts = {
  'canUserPerformInput': new SpyderbatApi.CanUserPerformInput() // CanUserPerformInput | 
};
apiInstance.canUserPerform(opts, (error, data, response) => {
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
 **canUserPerformInput** | [**CanUserPerformInput**](CanUserPerformInput.md)|  | [optional] 

### Return type

[**ApiRBACActions**](ApiRBACActions.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

