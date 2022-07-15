# SpyderbatApi.UIDataApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**uiDataDeleteOrgData**](UIDataApi.md#uiDataDeleteOrgData) | **DELETE** /api/v1/org/{orgUID}/uidata/{dataKey} | Delete Org UI Data
[**uiDataDeleteSourceData**](UIDataApi.md#uiDataDeleteSourceData) | **DELETE** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Delete Source UI Data
[**uiDataDeleteUserData**](UIDataApi.md#uiDataDeleteUserData) | **DELETE** /api/v1/user/{userUID}/uidata/{dataKey} | Delete User UI Data
[**uiDataDeleteUserOrgData**](UIDataApi.md#uiDataDeleteUserOrgData) | **DELETE** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Delete UserOrg UI Data
[**uiDataDeleteUserSourceData**](UIDataApi.md#uiDataDeleteUserSourceData) | **DELETE** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Delete UserSource UI Data
[**uiDataGetOrgData**](UIDataApi.md#uiDataGetOrgData) | **GET** /api/v1/org/{orgUID}/uidata/{dataKey} | Get Org UI Data
[**uiDataGetSourceData**](UIDataApi.md#uiDataGetSourceData) | **GET** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Get Source UI Data
[**uiDataGetUserData**](UIDataApi.md#uiDataGetUserData) | **GET** /api/v1/user/{userUID}/uidata/{dataKey} | Get User UI Data
[**uiDataGetUserOrgData**](UIDataApi.md#uiDataGetUserOrgData) | **GET** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Get UserOrg UI Data
[**uiDataGetUserSourceData**](UIDataApi.md#uiDataGetUserSourceData) | **GET** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Get UserSource UI Data
[**uiDataQueryOrgData**](UIDataApi.md#uiDataQueryOrgData) | **GET** /api/v1/org/{orgUID}/uidata/ | Query Org UI Data
[**uiDataQuerySourceData**](UIDataApi.md#uiDataQuerySourceData) | **GET** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/ | Query Source UI Data
[**uiDataQueryUserData**](UIDataApi.md#uiDataQueryUserData) | **GET** /api/v1/user/{userUID}/uidata/ | Query User UI Data
[**uiDataQueryUserOrgData**](UIDataApi.md#uiDataQueryUserOrgData) | **GET** /api/v1/user/{userUID}/org/{orgUID}/uidata/ | Query UserOrg UI Data
[**uiDataQueryUserSourceData**](UIDataApi.md#uiDataQueryUserSourceData) | **GET** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/ | Query UserSource UI Data
[**uiDataSetOrgData**](UIDataApi.md#uiDataSetOrgData) | **PUT** /api/v1/org/{orgUID}/uidata/{dataKey} | Set Org UI Data
[**uiDataSetSourceData**](UIDataApi.md#uiDataSetSourceData) | **PUT** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Set Source UI Data
[**uiDataSetUserData**](UIDataApi.md#uiDataSetUserData) | **PUT** /api/v1/user/{userUID}/uidata/{dataKey} | Set User UI Data
[**uiDataSetUserOrgData**](UIDataApi.md#uiDataSetUserOrgData) | **PUT** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Set UserOrg UI Data
[**uiDataSetUserSourceData**](UIDataApi.md#uiDataSetUserSourceData) | **PUT** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Set UserSource UI Data



## uiDataDeleteOrgData

> uiDataDeleteOrgData(dataKey, orgUID, sourceUID, userUID)

Delete Org UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataDeleteOrgData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataDeleteSourceData

> uiDataDeleteSourceData(dataKey, orgUID, sourceUID, userUID)

Delete Source UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataDeleteSourceData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataDeleteUserData

> uiDataDeleteUserData(dataKey, orgUID, sourceUID, userUID)

Delete User UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataDeleteUserData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataDeleteUserOrgData

> uiDataDeleteUserOrgData(dataKey, orgUID, sourceUID, userUID)

Delete UserOrg UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataDeleteUserOrgData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataDeleteUserSourceData

> uiDataDeleteUserSourceData(dataKey, orgUID, sourceUID, userUID)

Delete UserSource UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataDeleteUserSourceData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataGetOrgData

> UIData uiDataGetOrgData(dataKey, orgUID, sourceUID, userUID)

Get Org UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataGetOrgData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataGetSourceData

> UIData uiDataGetSourceData(dataKey, orgUID, sourceUID, userUID)

Get Source UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataGetSourceData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataGetUserData

> UIData uiDataGetUserData(dataKey, orgUID, sourceUID, userUID)

Get User UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataGetUserData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataGetUserOrgData

> UIData uiDataGetUserOrgData(dataKey, orgUID, sourceUID, userUID)

Get UserOrg UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataGetUserOrgData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataGetUserSourceData

> UIData uiDataGetUserSourceData(dataKey, orgUID, sourceUID, userUID)

Get UserSource UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
apiInstance.uiDataGetUserSourceData(dataKey, orgUID, sourceUID, userUID, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataQueryOrgData

> [UIData] uiDataQueryOrgData(dataKey, orgUID, sourceUID, userUID, opts)

Query Org UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'tags': ["null"] // [String] | Tags to search for
};
apiInstance.uiDataQueryOrgData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 
 **tags** | [**[String]**](String.md)| Tags to search for | [optional] 

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataQuerySourceData

> [UIData] uiDataQuerySourceData(dataKey, orgUID, sourceUID, userUID, opts)

Query Source UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'tags': ["null"] // [String] | Tags to search for
};
apiInstance.uiDataQuerySourceData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 
 **tags** | [**[String]**](String.md)| Tags to search for | [optional] 

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataQueryUserData

> [UIData] uiDataQueryUserData(dataKey, orgUID, sourceUID, userUID, opts)

Query User UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'tags': ["null"] // [String] | Tags to search for
};
apiInstance.uiDataQueryUserData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 
 **tags** | [**[String]**](String.md)| Tags to search for | [optional] 

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataQueryUserOrgData

> [UIData] uiDataQueryUserOrgData(dataKey, orgUID, sourceUID, userUID, opts)

Query UserOrg UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'tags': ["null"] // [String] | Tags to search for
};
apiInstance.uiDataQueryUserOrgData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 
 **tags** | [**[String]**](String.md)| Tags to search for | [optional] 

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataQueryUserSourceData

> [UIData] uiDataQueryUserSourceData(dataKey, orgUID, sourceUID, userUID, opts)

Query UserSource UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Src UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'tags': ["null"] // [String] | Tags to search for
};
apiInstance.uiDataQueryUserSourceData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Src UID | 
 **userUID** | **String**| Owner UID | 
 **tags** | [**[String]**](String.md)| Tags to search for | [optional] 

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## uiDataSetOrgData

> uiDataSetOrgData(dataKey, orgUID, sourceUID, userUID, opts)

Set Org UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Source UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'uiDataSetOrgDataInput': new SpyderbatApi.UiDataSetOrgDataInput() // UiDataSetOrgDataInput | 
};
apiInstance.uiDataSetOrgData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Source UID | 
 **userUID** | **String**| Owner UID | 
 **uiDataSetOrgDataInput** | [**UiDataSetOrgDataInput**](UiDataSetOrgDataInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## uiDataSetSourceData

> uiDataSetSourceData(dataKey, orgUID, sourceUID, userUID, opts)

Set Source UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Source UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'uiDataSetSourceDataInput': new SpyderbatApi.UiDataSetSourceDataInput() // UiDataSetSourceDataInput | 
};
apiInstance.uiDataSetSourceData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Source UID | 
 **userUID** | **String**| Owner UID | 
 **uiDataSetSourceDataInput** | [**UiDataSetSourceDataInput**](UiDataSetSourceDataInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## uiDataSetUserData

> uiDataSetUserData(dataKey, orgUID, sourceUID, userUID, opts)

Set User UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Source UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'uiDataSetUserDataInput': new SpyderbatApi.UiDataSetUserDataInput() // UiDataSetUserDataInput | 
};
apiInstance.uiDataSetUserData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Source UID | 
 **userUID** | **String**| Owner UID | 
 **uiDataSetUserDataInput** | [**UiDataSetUserDataInput**](UiDataSetUserDataInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## uiDataSetUserOrgData

> uiDataSetUserOrgData(dataKey, orgUID, sourceUID, userUID, opts)

Set UserOrg UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Source UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'uiDataSetUserOrgDataInput': new SpyderbatApi.UiDataSetUserOrgDataInput() // UiDataSetUserOrgDataInput | 
};
apiInstance.uiDataSetUserOrgData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Source UID | 
 **userUID** | **String**| Owner UID | 
 **uiDataSetUserOrgDataInput** | [**UiDataSetUserOrgDataInput**](UiDataSetUserOrgDataInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## uiDataSetUserSourceData

> uiDataSetUserSourceData(dataKey, orgUID, sourceUID, userUID, opts)

Set UserSource UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.UIDataApi();
let dataKey = "dataKey_example"; // String | Key for the data
let orgUID = "orgUID_example"; // String | Org UID
let sourceUID = "sourceUID_example"; // String | Source UID
let userUID = "userUID_example"; // String | Owner UID
let opts = {
  'uiDataSetUserSourceDataInput': new SpyderbatApi.UiDataSetUserSourceDataInput() // UiDataSetUserSourceDataInput | 
};
apiInstance.uiDataSetUserSourceData(dataKey, orgUID, sourceUID, userUID, opts, (error, data, response) => {
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
 **dataKey** | **String**| Key for the data | 
 **orgUID** | **String**| Org UID | 
 **sourceUID** | **String**| Source UID | 
 **userUID** | **String**| Owner UID | 
 **uiDataSetUserSourceDataInput** | [**UiDataSetUserSourceDataInput**](UiDataSetUserSourceDataInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

