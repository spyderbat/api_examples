# sbapi.InvestigationApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**investigation_create**](InvestigationApi.md#investigation_create) | **POST** /api/v1/org/{orgUID}/investigation/ | Create an investigation
[**investigation_delete**](InvestigationApi.md#investigation_delete) | **DELETE** /api/v1/org/{orgUID}/investigation/{investigationUID} | Delete an investigation
[**investigation_list**](InvestigationApi.md#investigation_list) | **GET** /api/v1/org/{orgUID}/investigation/ | List investigations
[**investigation_list_versions**](InvestigationApi.md#investigation_list_versions) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/ | List Investigation Versions
[**investigation_load**](InvestigationApi.md#investigation_load) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID} | Load an investigation
[**investigation_load_version**](InvestigationApi.md#investigation_load_version) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/{version} | Load Investigation Version
[**investigation_update**](InvestigationApi.md#investigation_update) | **PUT** /api/v1/org/{orgUID}/investigation/{investigationUID} | Update an investigation


# **investigation_create**
> ApiInvestigationCreateOutput investigation_create(investigation_uid, org_uid)

Create an investigation

 Create an investigationan   * Requires the user have the action *investigation:Create* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
from sbapi.model.api_investigation_create_output import ApiInvestigationCreateOutput
from sbapi.model.validation_error import ValidationError
from sbapi.model.investigation_create_input import InvestigationCreateInput
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
    api_instance = investigation_api.InvestigationApi(api_client)
    investigation_uid = "investigationUID_example" # str | Investigation UID
    org_uid = "orgUID_example" # str | Investigation OrgUID
    investigation_create_input = InvestigationCreateInput(
        created_by="created_by_example",
        data={
            "key": None,
        },
        modified_by="modified_by_example",
        modified_on=dateutil_parser('1970-01-01T00:00:00.00Z'),
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
        tags=[
            "tags_example",
        ],
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
        version=1,
    ) # InvestigationCreateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create an investigation
        api_response = api_instance.investigation_create(investigation_uid, org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_create: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an investigation
        api_response = api_instance.investigation_create(investigation_uid, org_uid, investigation_create_input=investigation_create_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_create: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigation_uid** | **str**| Investigation UID |
 **org_uid** | **str**| Investigation OrgUID |
 **investigation_create_input** | [**InvestigationCreateInput**](InvestigationCreateInput.md)|  | [optional]

### Return type

[**ApiInvestigationCreateOutput**](ApiInvestigationCreateOutput.md)

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

# **investigation_delete**
> investigation_delete(investigation_uid, org_uid)

Delete an investigation

 Deletes an investigation, by setting valid_to=now so that the investigation is virtually deleted.   * Requires the user have the action *investigation:Delete* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
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
    api_instance = investigation_api.InvestigationApi(api_client)
    investigation_uid = "investigationUID_example" # str | Investigation UID
    org_uid = "orgUID_example" # str | Investigation OrgUID

    # example passing only required values which don't have defaults set
    try:
        # Delete an investigation
        api_instance.investigation_delete(investigation_uid, org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigation_uid** | **str**| Investigation UID |
 **org_uid** | **str**| Investigation OrgUID |

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

# **investigation_list**
> [DaoInvestigation] investigation_list(org_uid)

List investigations

 Lists investigations   * Will list investigations which the user has the action *investigation:Load* or *investigation:LoadExpired* on 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
from sbapi.model.dao_investigation import DaoInvestigation
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
    api_instance = investigation_api.InvestigationApi(api_client)
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # List investigations
        api_response = api_instance.investigation_list(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |

### Return type

[**[DaoInvestigation]**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | invalid query parameters |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **investigation_list_versions**
> [DaoInvestigation] investigation_list_versions(investigation_uid, org_uid)

List Investigation Versions

 Lists prior version of this investigation   * Requires the user have the action *investigation:ListVersions* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
from sbapi.model.dao_investigation import DaoInvestigation
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
    api_instance = investigation_api.InvestigationApi(api_client)
    investigation_uid = "investigationUID_example" # str | 
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # List Investigation Versions
        api_response = api_instance.investigation_list_versions(investigation_uid, org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_list_versions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigation_uid** | **str**|  |
 **org_uid** | **str**|  |

### Return type

[**[DaoInvestigation]**](DaoInvestigation.md)

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

# **investigation_load**
> DaoInvestigation investigation_load(investigation_uid, org_uid)

Load an investigation

 Loads an investigation by UID.    * Requires action  *investigation:Load* to load an active investigation  * Requires action *investigation:LoadExpired* to load expired investigations  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
from sbapi.model.dao_investigation import DaoInvestigation
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
    api_instance = investigation_api.InvestigationApi(api_client)
    investigation_uid = "investigationUID_example" # str | 
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load an investigation
        api_response = api_instance.investigation_load(investigation_uid, org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigation_uid** | **str**|  |
 **org_uid** | **str**|  |

### Return type

[**DaoInvestigation**](DaoInvestigation.md)

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

# **investigation_load_version**
> DaoInvestigation investigation_load_version(investigation_uid, org_uid, version)

Load Investigation Version

 Loads a specific version of an investigation   * Requires the user have the action *investigation:LoadVersion* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
from sbapi.model.dao_investigation import DaoInvestigation
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
    api_instance = investigation_api.InvestigationApi(api_client)
    investigation_uid = "investigationUID_example" # str | 
    org_uid = "orgUID_example" # str | 
    version = 1 # int | 

    # example passing only required values which don't have defaults set
    try:
        # Load Investigation Version
        api_response = api_instance.investigation_load_version(investigation_uid, org_uid, version)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_load_version: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigation_uid** | **str**|  |
 **org_uid** | **str**|  |
 **version** | **int**|  |

### Return type

[**DaoInvestigation**](DaoInvestigation.md)

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

# **investigation_update**
> investigation_update(investigation_uid, org_uid)

Update an investigation

 Updates the investigationan   * Requires the user have the action *investigation:Update* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import investigation_api
from sbapi.model.investigation_update_input import InvestigationUpdateInput
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
    api_instance = investigation_api.InvestigationApi(api_client)
    investigation_uid = "investigationUID_example" # str | Investigation UID
    org_uid = "orgUID_example" # str | Investigation OrgUID
    investigation_update_input = InvestigationUpdateInput(
        created_by="created_by_example",
        data={
            "key": None,
        },
        modified_by="modified_by_example",
        modified_on=dateutil_parser('1970-01-01T00:00:00.00Z'),
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
        tags=[
            "tags_example",
        ],
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
        version=1,
    ) # InvestigationUpdateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an investigation
        api_instance.investigation_update(investigation_uid, org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_update: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an investigation
        api_instance.investigation_update(investigation_uid, org_uid, investigation_update_input=investigation_update_input)
    except sbapi.ApiException as e:
        print("Exception when calling InvestigationApi->investigation_update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigation_uid** | **str**| Investigation UID |
 **org_uid** | **str**| Investigation OrgUID |
 **investigation_update_input** | [**InvestigationUpdateInput**](InvestigationUpdateInput.md)|  | [optional]

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

