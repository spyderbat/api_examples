# Sbapi.DashboardSearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboardSearchCreate**](DashboardSearchApi.md#dashboardSearchCreate) | **POST** /api/v1/org/{orgUID}/dashboard_search/ | Create a dashboard search
[**dashboardSearchDelete**](DashboardSearchApi.md#dashboardSearchDelete) | **DELETE** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**dashboardSearchGet**](DashboardSearchApi.md#dashboardSearchGet) | **GET** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**dashboardSearchList**](DashboardSearchApi.md#dashboardSearchList) | **GET** /api/v1/org/{orgUID}/dashboard_search/ | List dashboard searches
[**dashboardSearchUpdate**](DashboardSearchApi.md#dashboardSearchUpdate) | **PUT** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Update a dashboard search



## dashboardSearchCreate

> String dashboardSearchCreate(dashboardSearchUID, orgUID, opts)

Create a dashboard search

 Create a dashboard search in an org.   * Requires action dashboard_search:Create

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.DashboardSearchApi();
let dashboardSearchUID = "dashboardSearchUID_example"; // String | UID for the DashboardSearch
let orgUID = "orgUID_example"; // String | Org UID
let opts = {
  'dashboardSearchCreateInput': new Sbapi.DashboardSearchCreateInput() // DashboardSearchCreateInput | 
};
apiInstance.dashboardSearchCreate(dashboardSearchUID, orgUID, opts, (error, data, response) => {
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
 **dashboardSearchUID** | **String**| UID for the DashboardSearch | 
 **orgUID** | **String**| Org UID | 
 **dashboardSearchCreateInput** | [**DashboardSearchCreateInput**](DashboardSearchCreateInput.md)|  | [optional] 

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## dashboardSearchDelete

> dashboardSearchDelete(dashboardSearchUID, orgUID)

Get a dashboard search

 Get a dashboard search in an org.   * Requires action dashboard_search:Delete

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.DashboardSearchApi();
let dashboardSearchUID = "dashboardSearchUID_example"; // String | 
let orgUID = "orgUID_example"; // String | 
apiInstance.dashboardSearchDelete(dashboardSearchUID, orgUID, (error, data, response) => {
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
 **dashboardSearchUID** | **String**|  | 
 **orgUID** | **String**|  | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## dashboardSearchGet

> DashboardSearch dashboardSearchGet(dashboardSearchUID, orgUID)

Get a dashboard search

 Get a dashboard search in an org.   * Requires action dashboard_search:Get

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.DashboardSearchApi();
let dashboardSearchUID = "dashboardSearchUID_example"; // String | 
let orgUID = "orgUID_example"; // String | 
apiInstance.dashboardSearchGet(dashboardSearchUID, orgUID, (error, data, response) => {
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
 **dashboardSearchUID** | **String**|  | 
 **orgUID** | **String**|  | 

### Return type

[**DashboardSearch**](DashboardSearch.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## dashboardSearchList

> [DashboardSearch] dashboardSearchList(orgUID)

List dashboard searches

 Lists dashboard searches by org.   * Requires action dashboard_search:Query

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.DashboardSearchApi();
let orgUID = "orgUID_example"; // String | Org UID
apiInstance.dashboardSearchList(orgUID, (error, data, response) => {
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
 **orgUID** | **String**| Org UID | 

### Return type

[**[DashboardSearch]**](DashboardSearch.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## dashboardSearchUpdate

> dashboardSearchUpdate(dashboardSearchUID, orgUID, opts)

Update a dashboard search

 Update a dashboard search in an org.   * Requires action dashboard_search:Update

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.DashboardSearchApi();
let dashboardSearchUID = "dashboardSearchUID_example"; // String | UID for the DashboardSearch
let orgUID = "orgUID_example"; // String | Org UID
let opts = {
  'dashboardSearchUpdateInput': new Sbapi.DashboardSearchUpdateInput() // DashboardSearchUpdateInput | 
};
apiInstance.dashboardSearchUpdate(dashboardSearchUID, orgUID, opts, (error, data, response) => {
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
 **dashboardSearchUID** | **String**| UID for the DashboardSearch | 
 **orgUID** | **String**| Org UID | 
 **dashboardSearchUpdateInput** | [**DashboardSearchUpdateInput**](DashboardSearchUpdateInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

