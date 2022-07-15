# SpyderbatApi.UserApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**userCurrent**](UserApi.md#userCurrent) | **GET** /api/v1/app/user/current | Returns the current user
[**userLoad**](UserApi.md#userLoad) | **GET** /api/v1/user/{userUID} | Load a user by ID



## userCurrent

> User userCurrent()

Returns the current user

 This api call will return the current user as a JSON object, or 404 if there is no user associated with the current authentication token or JWT. When a 403 is received the message will indicate if the user is expired or is pending approval, i.e. \&quot;user is expired\&quot;, or if the user is in a pending state \&quot;user is in a pending state\&quot;. 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';

let apiInstance = new SpyderbatApi.UserApi();
apiInstance.userCurrent((error, data, response) => {
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

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## userLoad

> User userLoad(userUID)

Load a user by ID

 Loads a user given the user&#39;s user ID.    * By default users may only load their own user ID  * The action *user:Load* may be placed upon the users role to allowed loading other users 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UserApi();
let userUID = "userUID_example"; // String | User UID
apiInstance.userLoad(userUID, (error, data, response) => {
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

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

