# SpyderbatApi.AgentWorkApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agentDeleteAgentWork**](AgentWorkApi.md#agentDeleteAgentWork) | **DELETE** /api/v1/org/{orgUID}/agent/{agentUID}/work | Delete agent work data for a specific agent
[**agentDeleteOrgWork**](AgentWorkApi.md#agentDeleteOrgWork) | **DELETE** /api/v1/org/{orgUID}/agent_work | Delete agent work for an org
[**agentGetAgentWork**](AgentWorkApi.md#agentGetAgentWork) | **GET** /api/v1/org/{orgUID}/agent/{agentUID}/work | Get agent work data for a specific agent
[**agentGetOrgWork**](AgentWorkApi.md#agentGetOrgWork) | **GET** /api/v1/org/{orgUID}/agent_work | Get agent work data for the organization
[**agentSetAgentWork**](AgentWorkApi.md#agentSetAgentWork) | **POST** /api/v1/org/{orgUID}/agent/{agentUID}/work | Set agent work data for a specific agent
[**agentSetOrgWork**](AgentWorkApi.md#agentSetOrgWork) | **POST** /api/v1/org/{orgUID}/agent_work | Set agent work data for a specific agent



## agentDeleteAgentWork

> agentDeleteAgentWork(agentUID, orgUID)

Delete agent work data for a specific agent

 Delet the work data for a specified agent.   * Requires *agent_data:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentWorkApi();
let agentUID = "agentUID_example"; // String | Agent UID
let orgUID = "orgUID_example"; // String | 
apiInstance.agentDeleteAgentWork(agentUID, orgUID, (error, data, response) => {
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
 **agentUID** | **String**| Agent UID | 
 **orgUID** | **String**|  | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentDeleteOrgWork

> agentDeleteOrgWork(agentUID, orgUID)

Delete agent work for an org

 Delete the work data for all agents for an organization.   * Requires *agent_data:Delete*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentWorkApi();
let agentUID = "agentUID_example"; // String | Agent UID
let orgUID = "orgUID_example"; // String | 
apiInstance.agentDeleteOrgWork(agentUID, orgUID, (error, data, response) => {
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
 **agentUID** | **String**| Agent UID | 
 **orgUID** | **String**|  | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentGetAgentWork

> ApiAgentWorkOutput agentGetAgentWork(agentUID, orgUID)

Get agent work data for a specific agent

 Get the work data for a specified agent.   * Requires *agent_data:GetAgentData* 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentWorkApi();
let agentUID = "agentUID_example"; // String | Agent UID
let orgUID = "orgUID_example"; // String | 
apiInstance.agentGetAgentWork(agentUID, orgUID, (error, data, response) => {
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
 **agentUID** | **String**| Agent UID | 
 **orgUID** | **String**|  | 

### Return type

[**ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentGetOrgWork

> ApiAgentWorkOutput agentGetOrgWork(agentUID, orgUID)

Get agent work data for the organization

 Get the work data for all agents associated with the organization.   * Requires *agent_data:GetOrgData* 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentWorkApi();
let agentUID = "agentUID_example"; // String | Agent UID
let orgUID = "orgUID_example"; // String | 
apiInstance.agentGetOrgWork(agentUID, orgUID, (error, data, response) => {
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
 **agentUID** | **String**| Agent UID | 
 **orgUID** | **String**|  | 

### Return type

[**ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentSetAgentWork

> agentSetAgentWork(agentUID, orgUID, opts)

Set agent work data for a specific agent

 Set the work data for a specified agent.   * Requires *agent_data:Set*  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentWorkApi();
let agentUID = "agentUID_example"; // String | Agent UID
let orgUID = "orgUID_example"; // String | 
let opts = {
  'agentSetAgentWorkInput': new SpyderbatApi.AgentSetAgentWorkInput() // AgentSetAgentWorkInput | 
};
apiInstance.agentSetAgentWork(agentUID, orgUID, opts, (error, data, response) => {
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
 **agentUID** | **String**| Agent UID | 
 **orgUID** | **String**|  | 
 **agentSetAgentWorkInput** | [**AgentSetAgentWorkInput**](AgentSetAgentWorkInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## agentSetOrgWork

> agentSetOrgWork(agentUID, orgUID, opts)

Set agent work data for a specific agent

 Set the work data for a specified agent.   * Requires *agent_data:SetData* 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentWorkApi();
let agentUID = "agentUID_example"; // String | Agent UID
let orgUID = "orgUID_example"; // String | 
let opts = {
  'agentSetOrgWorkInput': new SpyderbatApi.AgentSetOrgWorkInput() // AgentSetOrgWorkInput | 
};
apiInstance.agentSetOrgWork(agentUID, orgUID, opts, (error, data, response) => {
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
 **agentUID** | **String**| Agent UID | 
 **orgUID** | **String**|  | 
 **agentSetOrgWorkInput** | [**AgentSetOrgWorkInput**](AgentSetOrgWorkInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

