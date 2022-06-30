# sbapi.AgentApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agent_list**](AgentApi.md#agent_list) | **GET** /api/v1/org/{orgUID}/agent/ | List agents
[**agent_load**](AgentApi.md#agent_load) | **GET** /api/v1/org/{orgUID}/agent/{agentUID} | Load an agent


# **agent_list**
> [Agent] agent_list(org_uid)

List agents

 Lists the agents associated with an organization  * Requires the action  *agent:List* on the organization  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import agent_api
from sbapi.model.agent import Agent
from pprint import pprint
# Defining the host is optional and defaults to https://api.prod.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "https://api.prod.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = sbapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with sbapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_api.AgentApi(api_client)
    org_uid = "orgUID_example" # str | 
    agent_registration_uid_equals = "agent_registration_uid_equals_example" # str |  (optional)
    original_association = True # bool |  (optional)
    page = 1 # int |  (optional)
    page_size = 1 # int |  (optional)
    source_uid_equals = "source_uid_equals_example" # str |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # List agents
        api_response = api_instance.agent_list(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling AgentApi->agent_list: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List agents
        api_response = api_instance.agent_list(org_uid, agent_registration_uid_equals=agent_registration_uid_equals, original_association=original_association, page=page, page_size=page_size, source_uid_equals=source_uid_equals)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling AgentApi->agent_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **agent_registration_uid_equals** | **str**|  | [optional]
 **original_association** | **bool**|  | [optional]
 **page** | **int**|  | [optional]
 **page_size** | **int**|  | [optional]
 **source_uid_equals** | **str**|  | [optional]

### Return type

[**[Agent]**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_load**
> Agent agent_load(agent_uid, org_uid)

Load an agent

 Load a specified agent  * Requires the action  *agent:Load* on the organization  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import agent_api
from sbapi.model.agent import Agent
from pprint import pprint
# Defining the host is optional and defaults to https://api.prod.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "https://api.prod.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = sbapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with sbapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_api.AgentApi(api_client)
    agent_uid = "agentUID_example" # str | 
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load an agent
        api_response = api_instance.agent_load(agent_uid, org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling AgentApi->agent_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**|  |
 **org_uid** | **str**|  |

### Return type

[**Agent**](Agent.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | permission denied |  -  |
**404** | not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

