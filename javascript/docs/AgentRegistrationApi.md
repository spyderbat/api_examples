# SpyderbatApi.AgentRegistrationApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agentRegistrationCreate**](AgentRegistrationApi.md#agentRegistrationCreate) | **POST** /api/v1/org/{orgUID}/agent_registration/ | Create an agent registration
[**agentRegistrationDownloadLink**](AgentRegistrationApi.md#agentRegistrationDownloadLink) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid}/download_link | Get a download link for this registration
[**agentRegistrationGetLog**](AgentRegistrationApi.md#agentRegistrationGetLog) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid}/log | Get log of recent agent registration activity
[**agentRegistrationList**](AgentRegistrationApi.md#agentRegistrationList) | **GET** /api/v1/org/{orgUID}/agent_registration/ | List agent registrations
[**agentRegistrationLoad**](AgentRegistrationApi.md#agentRegistrationLoad) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid} | Load an agent registration
[**agentRegistrationUpdate**](AgentRegistrationApi.md#agentRegistrationUpdate) | **PUT** /api/v1/org/{orgUID}/agent_registration/{uid} | Update an agent registration



## agentRegistrationCreate

> ApiAgentCreateHandlerOutput agentRegistrationCreate(orgUID, opts)

Create an agent registration

 Creates a new agent registration  * Requires the action  *agent_registration:Create* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentRegistrationApi();
let orgUID = "orgUID_example"; // String | The OrgUID the registration is associated with
let opts = {
  'agentRegistrationCreateInput': new SpyderbatApi.AgentRegistrationCreateInput() // AgentRegistrationCreateInput | 
};
apiInstance.agentRegistrationCreate(orgUID, opts, (error, data, response) => {
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
 **orgUID** | **String**| The OrgUID the registration is associated with | 
 **agentRegistrationCreateInput** | [**AgentRegistrationCreateInput**](AgentRegistrationCreateInput.md)|  | [optional] 

### Return type

[**ApiAgentCreateHandlerOutput**](ApiAgentCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## agentRegistrationDownloadLink

> ApiAgentRegistrationDownloadLinkHandlerOutput agentRegistrationDownloadLink(orgUID, uid)

Get a download link for this registration

 Create a download link for the registration  * Requires the action  *agent_registration:Load* on the agent registration  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentRegistrationApi();
let orgUID = "orgUID_example"; // String | 
let uid = "uid_example"; // String | 
apiInstance.agentRegistrationDownloadLink(orgUID, uid, (error, data, response) => {
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
 **uid** | **String**|  | 

### Return type

[**ApiAgentRegistrationDownloadLinkHandlerOutput**](ApiAgentRegistrationDownloadLinkHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentRegistrationGetLog

> [DaoAgentLog] agentRegistrationGetLog(orgUID, uid, opts)

Get log of recent agent registration activity

 Get lots relating to recent agent registration activity  * Requires the action  *agent_registration:ListLog* on the agent registration  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentRegistrationApi();
let orgUID = "orgUID_example"; // String | 
let uid = "uid_example"; // String | 
let opts = {
  'page': 56, // Number | 
  'pageSize': 56 // Number | 
};
apiInstance.agentRegistrationGetLog(orgUID, uid, opts, (error, data, response) => {
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
 **uid** | **String**|  | 
 **page** | **Number**|  | [optional] 
 **pageSize** | **Number**|  | [optional] 

### Return type

[**[DaoAgentLog]**](DaoAgentLog.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentRegistrationList

> [AgentRegistration] agentRegistrationList(orgUID, opts)

List agent registrations

 Lists the agent registrations associated with an organization  * Requires the action  *agent_registration:List* on the organization  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentRegistrationApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'page': 56, // Number | 
  'pageSize': 56 // Number | 
};
apiInstance.agentRegistrationList(orgUID, opts, (error, data, response) => {
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
 **page** | **Number**|  | [optional] 
 **pageSize** | **Number**|  | [optional] 

### Return type

[**[AgentRegistration]**](AgentRegistration.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentRegistrationLoad

> AgentRegistration agentRegistrationLoad(orgUID, uid)

Load an agent registration

 Load a specified agent registration  * Requires the action  *agent_registration:Load* on the agent registration  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentRegistrationApi();
let orgUID = "orgUID_example"; // String | 
let uid = "uid_example"; // String | 
apiInstance.agentRegistrationLoad(orgUID, uid, (error, data, response) => {
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
 **uid** | **String**|  | 

### Return type

[**AgentRegistration**](AgentRegistration.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentRegistrationUpdate

> agentRegistrationUpdate(orgUID, uid, opts)

Update an agent registration

 Updates an existing registration  * Requires the action  *agent_registration:Update* on the organization and the registration 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentRegistrationApi();
let orgUID = "orgUID_example"; // String | The OrgUID the registration is associated with
let uid = "uid_example"; // String | Agent Registration UID
let opts = {
  'agentRegistrationUpdateInput': new SpyderbatApi.AgentRegistrationUpdateInput() // AgentRegistrationUpdateInput | 
};
apiInstance.agentRegistrationUpdate(orgUID, uid, opts, (error, data, response) => {
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
 **orgUID** | **String**| The OrgUID the registration is associated with | 
 **uid** | **String**| Agent Registration UID | 
 **agentRegistrationUpdateInput** | [**AgentRegistrationUpdateInput**](AgentRegistrationUpdateInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

