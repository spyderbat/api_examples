# spyderbat_api.AgentRegistrationApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**agent_registration_create**](AgentRegistrationApi.md#agent_registration_create) | **POST** /api/v1/org/{orgUID}/agent_registration/ | Create an agent registration
[**agent_registration_download_link**](AgentRegistrationApi.md#agent_registration_download_link) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid}/download_link | Get a download link for this registration
[**agent_registration_get_log**](AgentRegistrationApi.md#agent_registration_get_log) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid}/log | Get log of recent agent registration activity
[**agent_registration_list**](AgentRegistrationApi.md#agent_registration_list) | **GET** /api/v1/org/{orgUID}/agent_registration/ | List agent registrations
[**agent_registration_load**](AgentRegistrationApi.md#agent_registration_load) | **GET** /api/v1/org/{orgUID}/agent_registration/{uid} | Load an agent registration
[**agent_registration_update**](AgentRegistrationApi.md#agent_registration_update) | **PUT** /api/v1/org/{orgUID}/agent_registration/{uid} | Update an agent registration


# **agent_registration_create**
> ApiAgentCreateHandlerOutput agent_registration_create(org_uid)

Create an agent registration

 Creates a new agent registration  * Requires the action  *agent_registration:Create* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_registration_api
from spyderbat_api.model.api_agent_create_handler_output import ApiAgentCreateHandlerOutput
from spyderbat_api.model.agent_registration_create_input import AgentRegistrationCreateInput
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
    api_instance = agent_registration_api.AgentRegistrationApi(api_client)
    org_uid = "orgUID_example" # str | The OrgUID the registration is associated with
    agent_registration_create_input = AgentRegistrationCreateInput(
        config=DaoAgentConfig(
            classes=[
                DaoAgentClass(
                    name="name_example",
                    rbac_roles=[
                        "rbac_roles_example",
                    ],
                ),
            ],
            description="description_example",
            name="name_example",
            source_tags=[
                "source_tags_example",
            ],
            type="type_example",
        ),
        created_by="created_by_example",
        description="description_example",
        name="name_example",
        resource_name="resource_name_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # AgentRegistrationCreateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create an agent registration
        api_response = api_instance.agent_registration_create(org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_create: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an agent registration
        api_response = api_instance.agent_registration_create(org_uid, agent_registration_create_input=agent_registration_create_input)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_create: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| The OrgUID the registration is associated with |
 **agent_registration_create_input** | [**AgentRegistrationCreateInput**](AgentRegistrationCreateInput.md)|  | [optional]

### Return type

[**ApiAgentCreateHandlerOutput**](ApiAgentCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | invalid input parameters |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_registration_download_link**
> ApiAgentRegistrationDownloadLinkHandlerOutput agent_registration_download_link(org_uid, uid)

Get a download link for this registration

 Create a download link for the registration  * Requires the action  *agent_registration:Load* on the agent registration  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_registration_api
from spyderbat_api.model.api_agent_registration_download_link_handler_output import ApiAgentRegistrationDownloadLinkHandlerOutput
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
    api_instance = agent_registration_api.AgentRegistrationApi(api_client)
    org_uid = "orgUID_example" # str | 
    uid = "uid_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get a download link for this registration
        api_response = api_instance.agent_registration_download_link(org_uid, uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_download_link: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **uid** | **str**|  |

### Return type

[**ApiAgentRegistrationDownloadLinkHandlerOutput**](ApiAgentRegistrationDownloadLinkHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | bad request |  -  |
**403** | permission denied |  -  |
**404** | not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_registration_get_log**
> [DaoAgentLog] agent_registration_get_log(org_uid, uid)

Get log of recent agent registration activity

 Get lots relating to recent agent registration activity  * Requires the action  *agent_registration:ListLog* on the agent registration  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_registration_api
from spyderbat_api.model.dao_agent_log import DaoAgentLog
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
    api_instance = agent_registration_api.AgentRegistrationApi(api_client)
    org_uid = "orgUID_example" # str | 
    uid = "uid_example" # str | 
    page = 1 # int |  (optional)
    page_size = 1 # int |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get log of recent agent registration activity
        api_response = api_instance.agent_registration_get_log(org_uid, uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_get_log: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get log of recent agent registration activity
        api_response = api_instance.agent_registration_get_log(org_uid, uid, page=page, page_size=page_size)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_get_log: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **uid** | **str**|  |
 **page** | **int**|  | [optional]
 **page_size** | **int**|  | [optional]

### Return type

[**[DaoAgentLog]**](DaoAgentLog.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | bad request |  -  |
**403** | permission denied |  -  |
**404** | not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **agent_registration_list**
> [AgentRegistration] agent_registration_list(org_uid)

List agent registrations

 Lists the agent registrations associated with an organization  * Requires the action  *agent_registration:List* on the organization  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_registration_api
from spyderbat_api.model.agent_registration import AgentRegistration
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
    api_instance = agent_registration_api.AgentRegistrationApi(api_client)
    org_uid = "orgUID_example" # str | 
    page = 1 # int |  (optional)
    page_size = 1 # int |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # List agent registrations
        api_response = api_instance.agent_registration_list(org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_list: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List agent registrations
        api_response = api_instance.agent_registration_list(org_uid, page=page, page_size=page_size)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **page** | **int**|  | [optional]
 **page_size** | **int**|  | [optional]

### Return type

[**[AgentRegistration]**](AgentRegistration.md)

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

# **agent_registration_load**
> AgentRegistration agent_registration_load(org_uid, uid)

Load an agent registration

 Load a specified agent registration  * Requires the action  *agent_registration:Load* on the agent registration  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_registration_api
from spyderbat_api.model.agent_registration import AgentRegistration
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
    api_instance = agent_registration_api.AgentRegistrationApi(api_client)
    org_uid = "orgUID_example" # str | 
    uid = "uid_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load an agent registration
        api_response = api_instance.agent_registration_load(org_uid, uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **uid** | **str**|  |

### Return type

[**AgentRegistration**](AgentRegistration.md)

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

# **agent_registration_update**
> agent_registration_update(org_uid, uid)

Update an agent registration

 Updates an existing registration  * Requires the action  *agent_registration:Update* on the organization and the registration 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import agent_registration_api
from spyderbat_api.model.agent_registration_update_input import AgentRegistrationUpdateInput
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
    api_instance = agent_registration_api.AgentRegistrationApi(api_client)
    org_uid = "orgUID_example" # str | The OrgUID the registration is associated with
    uid = "uid_example" # str | Agent Registration UID
    agent_registration_update_input = AgentRegistrationUpdateInput(
        config=DaoAgentConfig(
            classes=[
                DaoAgentClass(
                    name="name_example",
                    rbac_roles=[
                        "rbac_roles_example",
                    ],
                ),
            ],
            description="description_example",
            name="name_example",
            source_tags=[
                "source_tags_example",
            ],
            type="type_example",
        ),
        created_by="created_by_example",
        description="description_example",
        name="name_example",
        resource_name="resource_name_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # AgentRegistrationUpdateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an agent registration
        api_instance.agent_registration_update(org_uid, uid)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_update: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an agent registration
        api_instance.agent_registration_update(org_uid, uid, agent_registration_update_input=agent_registration_update_input)
    except spyderbat_api.ApiException as e:
        print("Exception when calling AgentRegistrationApi->agent_registration_update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| The OrgUID the registration is associated with |
 **uid** | **str**| Agent Registration UID |
 **agent_registration_update_input** | [**AgentRegistrationUpdateInput**](AgentRegistrationUpdateInput.md)|  | [optional]

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
**400** | invalid input parameters |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

