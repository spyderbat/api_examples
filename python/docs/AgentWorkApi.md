# spyderbat_api.AgentWorkApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agent_delete_agent_work**](AgentWorkApi.md#agent_delete_agent_work) | **DELETE** /api/v1/org/{orgUID}/agent/{agentUID}/work | Delete agent work data for a specific agent
[**agent_delete_org_work**](AgentWorkApi.md#agent_delete_org_work) | **DELETE** /api/v1/org/{orgUID}/agent_work | Delete agent work for an org
[**agent_get_agent_work**](AgentWorkApi.md#agent_get_agent_work) | **GET** /api/v1/org/{orgUID}/agent/{agentUID}/work | Get agent work data for a specific agent
[**agent_get_org_work**](AgentWorkApi.md#agent_get_org_work) | **GET** /api/v1/org/{orgUID}/agent_work | Get agent work data for the organization
[**agent_set_agent_work**](AgentWorkApi.md#agent_set_agent_work) | **POST** /api/v1/org/{orgUID}/agent/{agentUID}/work | Set agent work data for a specific agent
[**agent_set_org_work**](AgentWorkApi.md#agent_set_org_work) | **POST** /api/v1/org/{orgUID}/agent_work | Set agent work data for a specific agent


# **agent_delete_agent_work**
> agent_delete_agent_work(agent_uid, org_uid)

Delete agent work data for a specific agent

 Delet the work data for a specified agent.   * Requires *agent_data:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_work_api
from spyderbat_api.model.validation_error import ValidationError
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_work_api.AgentWorkApi(api_client)
    agent_uid = "agentUID_example" # str | Agent UID
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Delete agent work data for a specific agent
        api_instance.agent_delete_agent_work(agent_uid, org_uid)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_delete_agent_work: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**| Agent UID |
 **org_uid** | **str**|  |

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Validation error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_delete_org_work**
> agent_delete_org_work(agent_uid, org_uid)

Delete agent work for an org

 Delete the work data for all agents for an organization.   * Requires *agent_data:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_work_api
from spyderbat_api.model.validation_error import ValidationError
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_work_api.AgentWorkApi(api_client)
    agent_uid = "agentUID_example" # str | Agent UID
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Delete agent work for an org
        api_instance.agent_delete_org_work(agent_uid, org_uid)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_delete_org_work: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**| Agent UID |
 **org_uid** | **str**|  |

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Validation error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_get_agent_work**
> ApiAgentWorkOutput agent_get_agent_work(agent_uid, org_uid)

Get agent work data for a specific agent

 Get the work data for a specified agent.   * Requires *agent_data:GetAgentData* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_work_api
from spyderbat_api.model.api_agent_work_output import ApiAgentWorkOutput
from spyderbat_api.model.validation_error import ValidationError
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_work_api.AgentWorkApi(api_client)
    agent_uid = "agentUID_example" # str | Agent UID
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get agent work data for a specific agent
        api_response = api_instance.agent_get_agent_work(agent_uid, org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_get_agent_work: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**| Agent UID |
 **org_uid** | **str**|  |

### Return type

[**ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Validation error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_get_org_work**
> ApiAgentWorkOutput agent_get_org_work(agent_uid, org_uid)

Get agent work data for the organization

 Get the work data for all agents associated with the organization.   * Requires *agent_data:GetOrgData* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_work_api
from spyderbat_api.model.api_agent_work_output import ApiAgentWorkOutput
from spyderbat_api.model.validation_error import ValidationError
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_work_api.AgentWorkApi(api_client)
    agent_uid = "agentUID_example" # str | Agent UID
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get agent work data for the organization
        api_response = api_instance.agent_get_org_work(agent_uid, org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_get_org_work: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**| Agent UID |
 **org_uid** | **str**|  |

### Return type

[**ApiAgentWorkOutput**](ApiAgentWorkOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Validation error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_set_agent_work**
> agent_set_agent_work(agent_uid, org_uid)

Set agent work data for a specific agent

 Set the work data for a specified agent.   * Requires *agent_data:Set*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_work_api
from spyderbat_api.model.agent_set_agent_work_input import AgentSetAgentWorkInput
from spyderbat_api.model.validation_error import ValidationError
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_work_api.AgentWorkApi(api_client)
    agent_uid = "agentUID_example" # str | Agent UID
    org_uid = "orgUID_example" # str | 
    agent_set_agent_work_input = AgentSetAgentWorkInput(
        tags=[
            "tags_example",
        ],
        work=OrcApiAgentWork(
            work=[
                OrcApiBatWork(
                    arguments=[
                        "arguments_example",
                    ],
                    bat_uid="bat_uid_example",
                    enabled=True,
                    parameters={
                        "key": None,
                    },
                    start_order=1,
                    uid="uid_example",
                    version={
                        "key": None,
                    },
                ),
            ],
        ),
    ) # AgentSetAgentWorkInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set agent work data for a specific agent
        api_instance.agent_set_agent_work(agent_uid, org_uid)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_set_agent_work: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set agent work data for a specific agent
        api_instance.agent_set_agent_work(agent_uid, org_uid, agent_set_agent_work_input=agent_set_agent_work_input)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_set_agent_work: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**| Agent UID |
 **org_uid** | **str**|  |
 **agent_set_agent_work_input** | [**AgentSetAgentWorkInput**](AgentSetAgentWorkInput.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Validation error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_set_org_work**
> agent_set_org_work(agent_uid, org_uid)

Set agent work data for a specific agent

 Set the work data for a specified agent.   * Requires *agent_data:SetData* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_work_api
from spyderbat_api.model.validation_error import ValidationError
from spyderbat_api.model.agent_set_org_work_input import AgentSetOrgWorkInput
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = agent_work_api.AgentWorkApi(api_client)
    agent_uid = "agentUID_example" # str | Agent UID
    org_uid = "orgUID_example" # str | 
    agent_set_org_work_input = AgentSetOrgWorkInput(
        tags=[
            "tags_example",
        ],
        work=OrcApiAgentWork(
            work=[
                OrcApiBatWork(
                    arguments=[
                        "arguments_example",
                    ],
                    bat_uid="bat_uid_example",
                    enabled=True,
                    parameters={
                        "key": None,
                    },
                    start_order=1,
                    uid="uid_example",
                    version={
                        "key": None,
                    },
                ),
            ],
        ),
    ) # AgentSetOrgWorkInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set agent work data for a specific agent
        api_instance.agent_set_org_work(agent_uid, org_uid)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_set_org_work: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set agent work data for a specific agent
        api_instance.agent_set_org_work(agent_uid, org_uid, agent_set_org_work_input=agent_set_org_work_input)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentWorkApi->agent_set_org_work: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent_uid** | **str**| Agent UID |
 **org_uid** | **str**|  |
 **agent_set_org_work_input** | [**AgentSetOrgWorkInput**](AgentSetOrgWorkInput.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Validation error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

