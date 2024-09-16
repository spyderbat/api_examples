# \AgentWorkApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agent_delete_agent_work**](AgentWorkApi.md#agent_delete_agent_work) | **DELETE** /api/v1/org/{orgUID}/agent/{agentUID}/work | Delete agent work data for a specific agent
[**agent_delete_org_work**](AgentWorkApi.md#agent_delete_org_work) | **DELETE** /api/v1/org/{orgUID}/agent_work | Delete agent work for an org
[**agent_get_agent_work**](AgentWorkApi.md#agent_get_agent_work) | **GET** /api/v1/org/{orgUID}/agent/{agentUID}/work | Get agent work data for a specific agent
[**agent_get_org_work**](AgentWorkApi.md#agent_get_org_work) | **GET** /api/v1/org/{orgUID}/agent_work | Get agent work data for the organization
[**agent_set_agent_work**](AgentWorkApi.md#agent_set_agent_work) | **POST** /api/v1/org/{orgUID}/agent/{agentUID}/work | Set agent work data for a specific agent
[**agent_set_org_work**](AgentWorkApi.md#agent_set_org_work) | **POST** /api/v1/org/{orgUID}/agent_work | Set agent work data for a specific agent



## agent_delete_agent_work

> agent_delete_agent_work(agent_uid, org_uid)
Delete agent work data for a specific agent

 Delet the work data for a specified agent.   * Requires *agent_data:Delete*  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**agent_uid** | **String** | Agent UID | [required] |
**org_uid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_delete_org_work

> agent_delete_org_work(org_uid)
Delete agent work for an org

 Delete the work data for all agents for an organization.   * Requires *agent_data:Delete*  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_get_agent_work

> crate::models::ApiAgentWorkOutput agent_get_agent_work(agent_uid, org_uid)
Get agent work data for a specific agent

 Get the work data for a specified agent.   * Requires *agent_data:GetAgentData* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**agent_uid** | **String** | Agent UID | [required] |
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_get_org_work

> crate::models::ApiAgentWorkOutput agent_get_org_work(org_uid)
Get agent work data for the organization

 Get the work data for all agents associated with the organization.   * Requires *agent_data:GetOrgData* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_set_agent_work

> agent_set_agent_work(agent_uid, org_uid, agent_set_agent_work_input)
Set agent work data for a specific agent

 Set the work data for a specified agent.   * Requires *agent_data:Set*  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**agent_uid** | **String** | Agent UID | [required] |
**org_uid** | **String** |  | [required] |
**agent_set_agent_work_input** | Option<[**AgentSetAgentWorkInput**](AgentSetAgentWorkInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_set_org_work

> agent_set_org_work(org_uid, agent_set_org_work_input)
Set agent work data for a specific agent

 Set the work data for a specified agent.   * Requires *agent_data:SetData* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**agent_set_org_work_input** | Option<[**AgentSetOrgWorkInput**](AgentSetOrgWorkInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

