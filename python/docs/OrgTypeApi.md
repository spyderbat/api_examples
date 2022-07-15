# spyderbat_api.OrgTypeApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**org_type_limit_active_sources**](OrgTypeApi.md#org_type_limit_active_sources) | **GET** /api/v1/org/{orgUID}/type/limit/active_sources | Loads limits regarding active sources
[**org_type_limit_org_roles**](OrgTypeApi.md#org_type_limit_org_roles) | **GET** /api/v1/org/{orgUID}/type/limit/org_roles | Loads limits regarding org roles
[**org_type_load**](OrgTypeApi.md#org_type_load) | **GET** /api/v1/org/{orgUID}/type | Load the org type for the organization


# **org_type_limit_active_sources**
> SessionOrgTypeMaxLimit org_type_limit_active_sources(org_uid)

Loads limits regarding active sources

 Loads the limits regarding active sources allowed on the organization, the active sources in an org are calculated on a 5m0s basis.    * Requires action *org:Load* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import org_type_api
from spyderbat_api.model.session_org_type_max_limit import SessionOrgTypeMaxLimit
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
    api_instance = org_type_api.OrgTypeApi(api_client)
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Loads limits regarding active sources
        api_response = api_instance.org_type_limit_active_sources(org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling OrgTypeApi->org_type_limit_active_sources: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |

### Return type

[**SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

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

# **org_type_limit_org_roles**
> SessionOrgTypeMaxLimit org_type_limit_org_roles(org_uid)

Loads limits regarding org roles

 Loads the limit information regarding the number of associated roles allowed per an organization   * Requires action *org:Load* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import org_type_api
from spyderbat_api.model.session_org_type_max_limit import SessionOrgTypeMaxLimit
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
    api_instance = org_type_api.OrgTypeApi(api_client)
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Loads limits regarding org roles
        api_response = api_instance.org_type_limit_org_roles(org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling OrgTypeApi->org_type_limit_org_roles: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |

### Return type

[**SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

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

# **org_type_load**
> DaoOrgType org_type_load(org_uid)

Load the org type for the organization

 Loads the org type for the organiation   * Requires action *org:Load* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import org_type_api
from spyderbat_api.model.dao_org_type import DaoOrgType
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
    api_instance = org_type_api.OrgTypeApi(api_client)
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load the org type for the organization
        api_response = api_instance.org_type_load(org_uid)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling OrgTypeApi->org_type_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |

### Return type

[**DaoOrgType**](DaoOrgType.md)

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

