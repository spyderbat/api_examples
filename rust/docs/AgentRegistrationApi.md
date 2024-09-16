# \AgentRegistrationApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agent_registration_create**](AgentRegistrationApi.md#agent_registration_create) | **POST** /api/v1/org/{orgUID}/agent_registration/ | Create an agent registration
[**agent_registration_download_link**](AgentRegistrationApi.md#agent_registration_download_link) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid}/download_link | Get a download link for this registration
[**agent_registration_get_log**](AgentRegistrationApi.md#agent_registration_get_log) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid}/log | Get log of recent agent registration activity
[**agent_registration_list**](AgentRegistrationApi.md#agent_registration_list) | **GET** /api/v1/org/{orgUID}/agent_registration/ | List agent registrations
[**agent_registration_load**](AgentRegistrationApi.md#agent_registration_load) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid} | Load an agent registration
[**agent_registration_update**](AgentRegistrationApi.md#agent_registration_update) | **PUT** /api/v1/org/{orgUID}/agent_registration/{uid} | Update an agent registration



## agent_registration_create

> crate::models::ApiAgentCreateHandlerOutput agent_registration_create(org_uid, agent_registration_create_input)
Create an agent registration

 Creates a new agent registration  * Requires the action  *agent_registration:Create* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | The OrgUID the registration is associated with | [required] |
**agent_registration_create_input** | Option<[**AgentRegistrationCreateInput**](AgentRegistrationCreateInput.md)> |  |  |

### Return type

[**crate::models::ApiAgentCreateHandlerOutput**](ApiAgentCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_registration_download_link

> crate::models::ApiAgentRegistrationDownloadLinkHandlerOutput agent_registration_download_link(org_uid, uid)
Get a download link for this registration

 Create a download link for the registration  * Requires the action  *agent_registration:Load* on the agent registration  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**uid** | **String** |  | [required] |

### Return type

[**crate::models::ApiAgentRegistrationDownloadLinkHandlerOutput**](ApiAgentRegistrationDownloadLinkHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_registration_get_log

> Vec<crate::models::DaoAgentLog> agent_registration_get_log(org_uid, uid, page, page_size)
Get log of recent agent registration activity

 Get lots relating to recent agent registration activity  * Requires the action  *agent_registration:ListLog* on the agent registration  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**uid** | **String** |  | [required] |
**page** | Option<**i32**> |  |  |
**page_size** | Option<**i32**> |  |  |

### Return type

[**Vec<crate::models::DaoAgentLog>**](DaoAgentLog.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_registration_list

> Vec<crate::models::AgentRegistration> agent_registration_list(org_uid, page, page_size)
List agent registrations

 Lists the agent registrations associated with an organization  * Requires the action  *agent_registration:List* on the organization  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**page** | Option<**i32**> |  |  |
**page_size** | Option<**i32**> |  |  |

### Return type

[**Vec<crate::models::AgentRegistration>**](AgentRegistration.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_registration_load

> crate::models::AgentRegistration agent_registration_load(org_uid, uid)
Load an agent registration

 Load a specified agent registration  * Requires the action  *agent_registration:Load* on the agent registration  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**uid** | **String** |  | [required] |

### Return type

[**crate::models::AgentRegistration**](AgentRegistration.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## agent_registration_update

> agent_registration_update(org_uid, uid, agent_registration_update_input)
Update an agent registration

 Updates an existing registration  * Requires the action  *agent_registration:Update* on the organization and the registration 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | The OrgUID the registration is associated with | [required] |
**uid** | **String** | Agent Registration UID | [required] |
**agent_registration_update_input** | Option<[**AgentRegistrationUpdateInput**](AgentRegistrationUpdateInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

