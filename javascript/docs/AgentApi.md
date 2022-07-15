# SpyderbatApi.AgentApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agentList**](AgentApi.md#agentList) | **GET** /api/v1/org/{orgUID}/agent/ | List agents
[**agentLoad**](AgentApi.md#agentLoad) | **GET** /api/v1/org/{orgUID}/agent/{agentUID} | Load an agent



## agentList

> [Agent] agentList(orgUID, opts)

List agents

 Lists the agents associated with an organization  * Requires the action  *agent:List* on the organization  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'agentRegistrationUidEquals': "agentRegistrationUidEquals_example", // String | 
  'originalAssociation': true, // Boolean | 
  'page': 56, // Number | 
  'pageSize': 56, // Number | 
  'sourceUidEquals': "sourceUidEquals_example" // String | 
};
apiInstance.agentList(orgUID, opts, (error, data, response) => {
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
 **agentRegistrationUidEquals** | **String**|  | [optional] 
 **originalAssociation** | **Boolean**|  | [optional] 
 **page** | **Number**|  | [optional] 
 **pageSize** | **Number**|  | [optional] 
 **sourceUidEquals** | **String**|  | [optional] 

### Return type

[**[Agent]**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## agentLoad

> Agent agentLoad(agentUID, orgUID)

Load an agent

 Load a specified agent  * Requires the action  *agent:Load* on the organization  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.AgentApi();
let agentUID = "agentUID_example"; // String | 
let orgUID = "orgUID_example"; // String | 
apiInstance.agentLoad(agentUID, orgUID, (error, data, response) => {
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
 **agentUID** | **String**|  | 
 **orgUID** | **String**|  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

