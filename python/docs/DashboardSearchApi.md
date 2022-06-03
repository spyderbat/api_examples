# sbapi.DashboardSearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard_search_create**](DashboardSearchApi.md#dashboard_search_create) | **POST** /api/v1/org/{orgUID}/dashboard_search/ | Create a dashboard search
[**dashboard_search_delete**](DashboardSearchApi.md#dashboard_search_delete) | **DELETE** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**dashboard_search_get**](DashboardSearchApi.md#dashboard_search_get) | **GET** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**dashboard_search_list**](DashboardSearchApi.md#dashboard_search_list) | **GET** /api/v1/org/{orgUID}/dashboard_search/ | List dashboard searches
[**dashboard_search_update**](DashboardSearchApi.md#dashboard_search_update) | **PUT** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Update a dashboard search


# **dashboard_search_create**
> str dashboard_search_create(dashboard_search_uid, org_uid)

Create a dashboard search

 Create a dashboard search in an org.   * Requires action dashboard_search:Create

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import dashboard_search_api
from sbapi.model.dashboard_search_create_input import DashboardSearchCreateInput
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "http://localhost"
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
    api_instance = dashboard_search_api.DashboardSearchApi(api_client)
    dashboard_search_uid = "dashboardSearchUID_example" # str | UID for the DashboardSearch
    org_uid = "orgUID_example" # str | Org UID
    dashboard_search_create_input = DashboardSearchCreateInput(
        =dateutil_parser('1970-01-01T00:00:00.00Z'),
        data={
            "key": None,
        },
        description="description_example",
        notify=True,
        notify_frequency=1,
        search="search_example",
        tags=[
            "tags_example",
        ],
    ) # DashboardSearchCreateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a dashboard search
        api_response = api_instance.dashboard_search_create(dashboard_search_uid, org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_create: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a dashboard search
        api_response = api_instance.dashboard_search_create(dashboard_search_uid, org_uid, dashboard_search_create_input=dashboard_search_create_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_create: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard_search_uid** | **str**| UID for the DashboardSearch |
 **org_uid** | **str**| Org UID |
 **dashboard_search_create_input** | [**DashboardSearchCreateInput**](DashboardSearchCreateInput.md)|  | [optional]

### Return type

**str**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **dashboard_search_delete**
> dashboard_search_delete(dashboard_search_uid, org_uid)

Get a dashboard search

 Get a dashboard search in an org.   * Requires action dashboard_search:Delete

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import dashboard_search_api
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "http://localhost"
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
    api_instance = dashboard_search_api.DashboardSearchApi(api_client)
    dashboard_search_uid = "dashboardSearchUID_example" # str | 
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get a dashboard search
        api_instance.dashboard_search_delete(dashboard_search_uid, org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard_search_uid** | **str**|  |
 **org_uid** | **str**|  |

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

# **dashboard_search_get**
> DashboardSearch dashboard_search_get(dashboard_search_uid, org_uid)

Get a dashboard search

 Get a dashboard search in an org.   * Requires action dashboard_search:Get

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import dashboard_search_api
from sbapi.model.dashboard_search import DashboardSearch
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "http://localhost"
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
    api_instance = dashboard_search_api.DashboardSearchApi(api_client)
    dashboard_search_uid = "dashboardSearchUID_example" # str | 
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get a dashboard search
        api_response = api_instance.dashboard_search_get(dashboard_search_uid, org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard_search_uid** | **str**|  |
 **org_uid** | **str**|  |

### Return type

[**DashboardSearch**](DashboardSearch.md)

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

# **dashboard_search_list**
> [DashboardSearch] dashboard_search_list(org_uid)

List dashboard searches

 Lists dashboard searches by org.   * Requires action dashboard_search:Query

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import dashboard_search_api
from sbapi.model.dashboard_search import DashboardSearch
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "http://localhost"
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
    api_instance = dashboard_search_api.DashboardSearchApi(api_client)
    org_uid = "orgUID_example" # str | Org UID

    # example passing only required values which don't have defaults set
    try:
        # List dashboard searches
        api_response = api_instance.dashboard_search_list(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| Org UID |

### Return type

[**[DashboardSearch]**](DashboardSearch.md)

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

# **dashboard_search_update**
> dashboard_search_update(dashboard_search_uid, org_uid)

Update a dashboard search

 Update a dashboard search in an org.   * Requires action dashboard_search:Update

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import dashboard_search_api
from sbapi.model.dashboard_search_update_input import DashboardSearchUpdateInput
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "http://localhost"
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
    api_instance = dashboard_search_api.DashboardSearchApi(api_client)
    dashboard_search_uid = "dashboardSearchUID_example" # str | UID for the DashboardSearch
    org_uid = "orgUID_example" # str | Org UID
    dashboard_search_update_input = DashboardSearchUpdateInput(
        =dateutil_parser('1970-01-01T00:00:00.00Z'),
        data={
            "key": None,
        },
        description="description_example",
        notify=True,
        notify_frequency=1,
        search="search_example",
        tags=[
            "tags_example",
        ],
    ) # DashboardSearchUpdateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a dashboard search
        api_instance.dashboard_search_update(dashboard_search_uid, org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_update: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a dashboard search
        api_instance.dashboard_search_update(dashboard_search_uid, org_uid, dashboard_search_update_input=dashboard_search_update_input)
    except sbapi.ApiException as e:
        print("Exception when calling DashboardSearchApi->dashboard_search_update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard_search_uid** | **str**| UID for the DashboardSearch |
 **org_uid** | **str**| Org UID |
 **dashboard_search_update_input** | [**DashboardSearchUpdateInput**](DashboardSearchUpdateInput.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

