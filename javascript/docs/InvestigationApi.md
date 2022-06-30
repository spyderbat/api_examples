# Sbapi.InvestigationApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**investigationCreate**](InvestigationApi.md#investigationCreate) | **POST** /api/v1/org/{orgUID}/investigation/ | Create an investigation
[**investigationDelete**](InvestigationApi.md#investigationDelete) | **DELETE** /api/v1/org/{orgUID}/investigation/{investigationUID} | Delete an investigation
[**investigationList**](InvestigationApi.md#investigationList) | **GET** /api/v1/org/{orgUID}/investigation/ | List investigations
[**investigationListVersions**](InvestigationApi.md#investigationListVersions) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/ | List Investigation Versions
[**investigationLoad**](InvestigationApi.md#investigationLoad) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID} | Load an investigation
[**investigationLoadVersion**](InvestigationApi.md#investigationLoadVersion) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/{version} | Load Investigation Version
[**investigationUpdate**](InvestigationApi.md#investigationUpdate) | **PUT** /api/v1/org/{orgUID}/investigation/{investigationUID} | Update an investigation



## investigationCreate

> ApiInvestigationCreateOutput investigationCreate(investigationUID, orgUID, opts)

Create an investigation

 Create an investigationan   * Requires the user have the action *investigation:Create* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let investigationUID = "investigationUID_example"; // String | Investigation UID
let orgUID = "orgUID_example"; // String | Investigation OrgUID
let opts = {
  'investigationCreateInput': new Sbapi.InvestigationCreateInput() // InvestigationCreateInput | 
};
apiInstance.investigationCreate(investigationUID, orgUID, opts, (error, data, response) => {
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
 **investigationUID** | **String**| Investigation UID | 
 **orgUID** | **String**| Investigation OrgUID | 
 **investigationCreateInput** | [**InvestigationCreateInput**](InvestigationCreateInput.md)|  | [optional] 

### Return type

[**ApiInvestigationCreateOutput**](ApiInvestigationCreateOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## investigationDelete

> investigationDelete(investigationUID, orgUID)

Delete an investigation

 Deletes an investigation, by setting valid_to&#x3D;now so that the investigation is virtually deleted.   * Requires the user have the action *investigation:Delete* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let investigationUID = "investigationUID_example"; // String | Investigation UID
let orgUID = "orgUID_example"; // String | Investigation OrgUID
apiInstance.investigationDelete(investigationUID, orgUID, (error, data, response) => {
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
 **investigationUID** | **String**| Investigation UID | 
 **orgUID** | **String**| Investigation OrgUID | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## investigationList

> [DaoInvestigation] investigationList(orgUID)

List investigations

 Lists investigations   * Will list investigations which the user has the action *investigation:Load* or *investigation:LoadExpired* on 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let orgUID = "orgUID_example"; // String | 
apiInstance.investigationList(orgUID, (error, data, response) => {
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

[**[DaoInvestigation]**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigationListVersions

> [DaoInvestigation] investigationListVersions(investigationUID, orgUID)

List Investigation Versions

 Lists prior version of this investigation   * Requires the user have the action *investigation:ListVersions* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let investigationUID = "investigationUID_example"; // String | 
let orgUID = "orgUID_example"; // String | 
apiInstance.investigationListVersions(investigationUID, orgUID, (error, data, response) => {
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
 **investigationUID** | **String**|  | 
 **orgUID** | **String**|  | 

### Return type

[**[DaoInvestigation]**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigationLoad

> DaoInvestigation investigationLoad(investigationUID, orgUID)

Load an investigation

 Loads an investigation by UID.    * Requires action  *investigation:Load* to load an active investigation  * Requires action *investigation:LoadExpired* to load expired investigations  

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let investigationUID = "investigationUID_example"; // String | 
let orgUID = "orgUID_example"; // String | 
apiInstance.investigationLoad(investigationUID, orgUID, (error, data, response) => {
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
 **investigationUID** | **String**|  | 
 **orgUID** | **String**|  | 

### Return type

[**DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigationLoadVersion

> DaoInvestigation investigationLoadVersion(investigationUID, orgUID, version)

Load Investigation Version

 Loads a specific version of an investigation   * Requires the user have the action *investigation:LoadVersion* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let investigationUID = "investigationUID_example"; // String | 
let orgUID = "orgUID_example"; // String | 
let version = 56; // Number | 
apiInstance.investigationLoadVersion(investigationUID, orgUID, version, (error, data, response) => {
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
 **investigationUID** | **String**|  | 
 **orgUID** | **String**|  | 
 **version** | **Number**|  | 

### Return type

[**DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## investigationUpdate

> investigationUpdate(investigationUID, orgUID, opts)

Update an investigation

 Updates the investigationan   * Requires the user have the action *investigation:Update* 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.InvestigationApi();
let investigationUID = "investigationUID_example"; // String | Investigation UID
let orgUID = "orgUID_example"; // String | Investigation OrgUID
let opts = {
  'investigationUpdateInput': new Sbapi.InvestigationUpdateInput() // InvestigationUpdateInput | 
};
apiInstance.investigationUpdate(investigationUID, orgUID, opts, (error, data, response) => {
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
 **investigationUID** | **String**| Investigation UID | 
 **orgUID** | **String**| Investigation OrgUID | 
 **investigationUpdateInput** | [**InvestigationUpdateInput**](InvestigationUpdateInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

