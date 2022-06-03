# Sbapi.UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**userAuth**](UserApi.md#userAuth) | **POST** /api/v1/app/user/auth | Authenticate a user, returns token required for authentication
[**userCurrent**](UserApi.md#userCurrent) | **GET** /api/v1/app/user/current | Returns the current user
[**userLoad**](UserApi.md#userLoad) | **GET** /api/v1/user/{userUID} | Load a user by ID
[**userSignup**](UserApi.md#userSignup) | **GET** /api/v1/app/signup | Sign up a new user



## userAuth

> User userAuth(opts)

Authenticate a user, returns token required for authentication

 Users may be authenticate by providing their email address and password as a JSON object, in response an authentication token will be returned in the header &#39;Authorization-Token&#39;. The returned token must be used in a request header with the name &#39;Authorization&#39; when making other authenticated API calls, for example &#39;Authorization: Bearer TOKEN&#39;. 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.UserApi();
let opts = {
  'userAuthInput': new Sbapi.UserAuthInput() // UserAuthInput | 
};
apiInstance.userAuth(opts, (error, data, response) => {
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
 **userAuthInput** | [**UserAuthInput**](UserAuthInput.md)|  | [optional] 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## userCurrent

> User userCurrent()

Returns the current user

 This api call will return the current user as a JSON object, or 404 if there is no user associated with the current authentication token or JWT. When a 403 is received the message will indicate if the user is expired or is pending approval, i.e. \&quot;user is expired\&quot;, or if the user is in a pending state \&quot;user is in a pending state\&quot;. 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.UserApi();
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

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## userLoad

> User userLoad(userUID)

Load a user by ID

 Loads a user given the user&#39;s user ID.    * By default users may only load their own user ID  * The action *user:Load* may be placed upon the users role to allowed loading other users 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.UserApi();
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


## userSignup

> userSignup(opts)

Sign up a new user

 Users can sign up by providing an email address. If the email is invalid or a user with the same existing email address has already registered an error is returned  ### Default configuration  By default the user will have an organization created for them, with the all actions allowed on the organization. The user can discover their organization id by  list all organizations.     

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.UserApi();
let opts = {
  'auth': "auth_example", // String | 
  'confirm': "confirm_example", // String | 
  'email': "email_example", // String | 
  'enable': true // Boolean | 
};
apiInstance.userSignup(opts, (error, data, response) => {
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
 **auth** | **String**|  | [optional] 
 **confirm** | **String**|  | [optional] 
 **email** | **String**|  | [optional] 
 **enable** | **Boolean**|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

