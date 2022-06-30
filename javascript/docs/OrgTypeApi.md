# Sbapi.OrgTypeApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**orgTypeLimitActiveSources**](OrgTypeApi.md#orgTypeLimitActiveSources) | **GET** /api/v1/org/{orgUID}/type/limit/active_sources | Loads limits regarding active sources
[**orgTypeLimitOrgRoles**](OrgTypeApi.md#orgTypeLimitOrgRoles) | **GET** /api/v1/org/{orgUID}/type/limit/org_roles | Loads limits regarding org roles
[**orgTypeLoad**](OrgTypeApi.md#orgTypeLoad) | **GET** /api/v1/org/{orgUID}/type | Load the org type for the organization



## orgTypeLimitActiveSources

> SessionOrgTypeMaxLimit orgTypeLimitActiveSources(orgUID)

Loads limits regarding active sources

 Loads the limits regarding active sources allowed on the organization, the active sources in an org are calculated on a 5m0s basis.    * Requires action *org:Load* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.OrgTypeApi();
let orgUID = "orgUID_example"; // String | 
apiInstance.orgTypeLimitActiveSources(orgUID, (error, data, response) => {
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

### Return type

[**SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgTypeLimitOrgRoles

> SessionOrgTypeMaxLimit orgTypeLimitOrgRoles(orgUID)

Loads limits regarding org roles

 Loads the limit information regarding the number of associated roles allowed per an organization   * Requires action *org:Load* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.OrgTypeApi();
let orgUID = "orgUID_example"; // String | 
apiInstance.orgTypeLimitOrgRoles(orgUID, (error, data, response) => {
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

### Return type

[**SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgTypeLoad

> DaoOrgType orgTypeLoad(orgUID)

Load the org type for the organization

 Loads the org type for the organiation   * Requires action *org:Load* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.OrgTypeApi();
let orgUID = "orgUID_example"; // String | 
apiInstance.orgTypeLoad(orgUID, (error, data, response) => {
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

### Return type

[**DaoOrgType**](DaoOrgType.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

