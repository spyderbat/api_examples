# \AgentApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agent_list**](AgentApi.md#agent_list) | **GET** /api/v1/org/{orgUID}/agent/ | List agents
[**agent_load**](AgentApi.md#agent_load) | **GET** /api/v1/org/{orgUID}/agent/{agentUID} | Load an agent



## agent_list

> Vec<crate::models::Agent> agent_list(org_uid, agent_registration_uid_equals, original_association, page, page_size, source_uid_equals)
List agents

 Lists the agents associated with an organization  * Requires the action  *agent:List* on the organization  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**agent_registration_uid_equals** | Option<**String**> |  |  |
**original_association** | Option<**bool**> |  |  |
**page** | Option<**i32**> |  |  |
**page_size** | Option<**i32**> |  |  |
**source_uid_equals** | Option<**String**> |  |  |

### Return type

[**Vec<crate::models::Agent>**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_load

> crate::models::Agent agent_load(agent_uid, org_uid)
Load an agent

 Load a specified agent  * Requires the action  *agent:Load* on the organization  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**agent_uid** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::Agent**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

