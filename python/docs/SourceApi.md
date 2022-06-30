# sbapi.SourceApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**integration_soar_src_list**](SourceApi.md#integration_soar_src_list) | **GET** /api/v1/integration/soar/org/{orgUID}/source/ | List sources for integration with SOARs
[**src_create**](SourceApi.md#src_create) | **POST** /api/v1/org/{orgUID}/source/ | Create a source
[**src_delete**](SourceApi.md#src_delete) | **DELETE** /api/v1/org/{orgUID}/source/{sourceUID} | Delete a source
[**src_list**](SourceApi.md#src_list) | **GET** /api/v1/org/{orgUID}/source/ | List sources
[**src_load**](SourceApi.md#src_load) | **GET** /api/v1/org/{orgUID}/source/{sourceUID} | Load a source
[**src_update**](SourceApi.md#src_update) | **PUT** /api/v1/org/{orgUID}/source/{sourceUID} | Update a source


# **integration_soar_src_list**
> [ApiSOARListHandlerOutput] integration_soar_src_list(org_uid)

List sources for integration with SOARs

 Lists the sources of data that match the specified query parameters, and return  URL entry points into the UI for matching sources.   * Requires the action  *org:ListSources* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import source_api
from sbapi.model.api_soar_list_handler_output import ApiSOARListHandlerOutput
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
    api_instance = source_api.SourceApi(api_client)
    org_uid = "orgUID_example" # str | 
    et = 1 # int | optional end time of the query (optional)
    hostname = "hostname_example" # str | A single hostname to match (optional)
    ip_address = "ip_address_example" # str | A single IP address to match (optional)
    mac_address = "mac_address_example" # str | A single mac address to match (optional)
    page = 1 # int |  (optional)
    page_size = 1 # int |  (optional)
    st = 1 # int | optional start time of the query, if only a start time is provided, end time will be start+10m (optional)

    # example passing only required values which don't have defaults set
    try:
        # List sources for integration with SOARs
        api_response = api_instance.integration_soar_src_list(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->integration_soar_src_list: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List sources for integration with SOARs
        api_response = api_instance.integration_soar_src_list(org_uid, et=et, hostname=hostname, ip_address=ip_address, mac_address=mac_address, page=page, page_size=page_size, st=st)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->integration_soar_src_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **et** | **int**| optional end time of the query | [optional]
 **hostname** | **str**| A single hostname to match | [optional]
 **ip_address** | **str**| A single IP address to match | [optional]
 **mac_address** | **str**| A single mac address to match | [optional]
 **page** | **int**|  | [optional]
 **page_size** | **int**|  | [optional]
 **st** | **int**| optional start time of the query, if only a start time is provided, end time will be start+10m | [optional]

### Return type

[**[ApiSOARListHandlerOutput]**](ApiSOARListHandlerOutput.md)

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

# **src_create**
> ApiSourceCreateHandlerOutput src_create(org_uid)

Create a source

 Creates a new source of data  * Requires the action  *org:CreateSource* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import source_api
from sbapi.model.src_create_input import SrcCreateInput
from sbapi.model.api_source_create_handler_output import ApiSourceCreateHandlerOutput
from sbapi.model.validation_error import ValidationError
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
    api_instance = source_api.SourceApi(api_client)
    org_uid = "orgUID_example" # str | The org this source is associated with
    src_create_input = SrcCreateInput(
        description="description_example",
        name="name_example",
        runtime_description="runtime_description_example",
        runtime_details=OrcApiRuntimeDetails(
            agent_arch="agent_arch_example",
            agent_version="agent_version_example",
            boot_time=3.14,
            cloud_image_id="cloud_image_id_example",
            cloud_instance_id="cloud_instance_id_example",
            cloud_region="cloud_region_example",
            cloud_tags=[
                "cloud_tags_example",
            ],
            cloud_type="cloud_type_example",
            cpu_cores=1,
            cpu_make="cpu_make_example",
            cpu_model="cpu_model_example",
            hostname="hostname_example",
            ip_addresses=[
                "ip_addresses_example",
            ],
            mac_addresses=[
                "mac_addresses_example",
            ],
            os_name="os_name_example",
            os_pretty_name="os_pretty_name_example",
            request_ip="request_ip_example",
            src_uid="src_uid_example",
            uname="uname_example",
        ),
        tags=[
            "tags_example",
        ],
        type="type_example",
        uid="uid_example",
    ) # SrcCreateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a source
        api_response = api_instance.src_create(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_create: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a source
        api_response = api_instance.src_create(org_uid, src_create_input=src_create_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_create: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| The org this source is associated with |
 **src_create_input** | [**SrcCreateInput**](SrcCreateInput.md)|  | [optional]

### Return type

[**ApiSourceCreateHandlerOutput**](ApiSourceCreateHandlerOutput.md)

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

# **src_delete**
> src_delete(org_uid, source_uid)

Delete a source

 Delete a source  * Requires the action  *org:DeleteSource* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import source_api
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
    api_instance = source_api.SourceApi(api_client)
    org_uid = "orgUID_example" # str | 
    source_uid = "sourceUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Delete a source
        api_instance.src_delete(org_uid, source_uid)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **source_uid** | **str**|  |

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **src_list**
> [Source] src_list(org_uid)

List sources

 Lists the sources of data for an organization  * Requires the action  *org:ListSources* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import source_api
from sbapi.model.source import Source
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
    api_instance = source_api.SourceApi(api_client)
    org_uid = "orgUID_example" # str | 
    agent_uid_equals = "agent_uid_equals_example" # str |  (optional)
    description_contains = "description_contains_example" # str |  (optional)
    has_tags = [
        "has_tags_example",
    ] # [str] |  (optional)
    original_association = True # bool |  (optional)
    page = 1 # int |  (optional)
    page_size = 1 # int |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # List sources
        api_response = api_instance.src_list(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_list: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List sources
        api_response = api_instance.src_list(org_uid, agent_uid_equals=agent_uid_equals, description_contains=description_contains, has_tags=has_tags, original_association=original_association, page=page, page_size=page_size)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **agent_uid_equals** | **str**|  | [optional]
 **description_contains** | **str**|  | [optional]
 **has_tags** | **[str]**|  | [optional]
 **original_association** | **bool**|  | [optional]
 **page** | **int**|  | [optional]
 **page_size** | **int**|  | [optional]

### Return type

[**[Source]**](Source.md)

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

# **src_load**
> Source src_load(org_uid, source_uid)

Load a source

 Loads the meta data for a source of data  * Requires the action  *org:LoadSource* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import source_api
from sbapi.model.source import Source
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
    api_instance = source_api.SourceApi(api_client)
    org_uid = "orgUID_example" # str | 
    source_uid = "sourceUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load a source
        api_response = api_instance.src_load(org_uid, source_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **source_uid** | **str**|  |

### Return type

[**Source**](Source.md)

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

# **src_update**
> src_update(org_uid, source_uid)

Update a source

 Updates the meta data for a source of data  * Requires the action  *org:UpdateSource* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import source_api
from sbapi.model.src_update_input import SrcUpdateInput
from sbapi.model.validation_error import ValidationError
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
    api_instance = source_api.SourceApi(api_client)
    org_uid = "orgUID_example" # str | The org this source is associated with
    source_uid = "sourceUID_example" # str | The UID of the source
    src_update_input = SrcUpdateInput(
        description="description_example",
        last_data=dateutil_parser('1970-01-01T00:00:00.00Z'),
        last_ingest_chunk_end_time=dateutil_parser('1970-01-01T00:00:00.00Z'),
        last_stored_chunk_end_time=dateutil_parser('1970-01-01T00:00:00.00Z'),
        name="name_example",
        resource_name="resource_name_example",
        resource_policy=ResourcePolicy(
            name="name_example",
            statements=[
                RbacStatement(
                    actions=[
                        "actions_example",
                    ],
                    condition={},
                    effect="effect_example",
                    resources=[
                        "resources_example",
                    ],
                    sid="sid_example",
                ),
            ],
            version="version_example",
        ),
        runtime_description="runtime_description_example",
        runtime_details=OrcApiRuntimeDetails(
            agent_arch="agent_arch_example",
            agent_version="agent_version_example",
            boot_time=3.14,
            cloud_image_id="cloud_image_id_example",
            cloud_instance_id="cloud_instance_id_example",
            cloud_region="cloud_region_example",
            cloud_tags=[
                "cloud_tags_example",
            ],
            cloud_type="cloud_type_example",
            cpu_cores=1,
            cpu_make="cpu_make_example",
            cpu_model="cpu_model_example",
            hostname="hostname_example",
            ip_addresses=[
                "ip_addresses_example",
            ],
            mac_addresses=[
                "mac_addresses_example",
            ],
            os_name="os_name_example",
            os_pretty_name="os_pretty_name_example",
            request_ip="request_ip_example",
            src_uid="src_uid_example",
            uname="uname_example",
        ),
        tags=[
            "tags_example",
        ],
        type="type_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # SrcUpdateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a source
        api_instance.src_update(org_uid, source_uid)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_update: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a source
        api_instance.src_update(org_uid, source_uid, src_update_input=src_update_input)
    except sbapi.ApiException as e:
        print("Exception when calling SourceApi->src_update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| The org this source is associated with |
 **source_uid** | **str**| The UID of the source |
 **src_update_input** | [**SrcUpdateInput**](SrcUpdateInput.md)|  | [optional]

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

