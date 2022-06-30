# sbapi.UIDataApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ui_data_delete_org_data**](UIDataApi.md#ui_data_delete_org_data) | **DELETE** /api/v1/org/{orgUID}/uidata/{dataKey} | Delete Org UI Data
[**ui_data_delete_source_data**](UIDataApi.md#ui_data_delete_source_data) | **DELETE** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Delete Source UI Data
[**ui_data_delete_user_data**](UIDataApi.md#ui_data_delete_user_data) | **DELETE** /api/v1/user/{userUID}/uidata/{dataKey} | Delete User UI Data
[**ui_data_delete_user_org_data**](UIDataApi.md#ui_data_delete_user_org_data) | **DELETE** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Delete UserOrg UI Data
[**ui_data_delete_user_source_data**](UIDataApi.md#ui_data_delete_user_source_data) | **DELETE** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Delete UserSource UI Data
[**ui_data_get_org_data**](UIDataApi.md#ui_data_get_org_data) | **GET** /api/v1/org/{orgUID}/uidata/{dataKey} | Get Org UI Data
[**ui_data_get_source_data**](UIDataApi.md#ui_data_get_source_data) | **GET** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Get Source UI Data
[**ui_data_get_user_data**](UIDataApi.md#ui_data_get_user_data) | **GET** /api/v1/user/{userUID}/uidata/{dataKey} | Get User UI Data
[**ui_data_get_user_org_data**](UIDataApi.md#ui_data_get_user_org_data) | **GET** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Get UserOrg UI Data
[**ui_data_get_user_source_data**](UIDataApi.md#ui_data_get_user_source_data) | **GET** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Get UserSource UI Data
[**ui_data_query_org_data**](UIDataApi.md#ui_data_query_org_data) | **GET** /api/v1/org/{orgUID}/uidata/ | Query Org UI Data
[**ui_data_query_source_data**](UIDataApi.md#ui_data_query_source_data) | **GET** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/ | Query Source UI Data
[**ui_data_query_user_data**](UIDataApi.md#ui_data_query_user_data) | **GET** /api/v1/user/{userUID}/uidata/ | Query User UI Data
[**ui_data_query_user_org_data**](UIDataApi.md#ui_data_query_user_org_data) | **GET** /api/v1/user/{userUID}/org/{orgUID}/uidata/ | Query UserOrg UI Data
[**ui_data_query_user_source_data**](UIDataApi.md#ui_data_query_user_source_data) | **GET** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/ | Query UserSource UI Data
[**ui_data_set_org_data**](UIDataApi.md#ui_data_set_org_data) | **PUT** /api/v1/org/{orgUID}/uidata/{dataKey} | Set Org UI Data
[**ui_data_set_source_data**](UIDataApi.md#ui_data_set_source_data) | **PUT** /api/v1/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Set Source UI Data
[**ui_data_set_user_data**](UIDataApi.md#ui_data_set_user_data) | **PUT** /api/v1/user/{userUID}/uidata/{dataKey} | Set User UI Data
[**ui_data_set_user_org_data**](UIDataApi.md#ui_data_set_user_org_data) | **PUT** /api/v1/user/{userUID}/org/{orgUID}/uidata/{dataKey} | Set UserOrg UI Data
[**ui_data_set_user_source_data**](UIDataApi.md#ui_data_set_user_source_data) | **PUT** /api/v1/user/{userUID}/org/{orgUID}/source/{sourceUID}/uidata/{dataKey} | Set UserSource UI Data


# **ui_data_delete_org_data**
> ui_data_delete_org_data(data_key, org_uid, source_uid, user_uid)

Delete Org UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Delete Org UI Data
        api_instance.ui_data_delete_org_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_delete_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_delete_source_data**
> ui_data_delete_source_data(data_key, org_uid, source_uid, user_uid)

Delete Source UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Delete Source UI Data
        api_instance.ui_data_delete_source_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_delete_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_delete_user_data**
> ui_data_delete_user_data(data_key, org_uid, source_uid, user_uid)

Delete User UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Delete User UI Data
        api_instance.ui_data_delete_user_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_delete_user_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_delete_user_org_data**
> ui_data_delete_user_org_data(data_key, org_uid, source_uid, user_uid)

Delete UserOrg UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Delete UserOrg UI Data
        api_instance.ui_data_delete_user_org_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_delete_user_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_delete_user_source_data**
> ui_data_delete_user_source_data(data_key, org_uid, source_uid, user_uid)

Delete UserSource UI Data

 Sets UI Data   * Requires *uidata:Delete*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Delete UserSource UI Data
        api_instance.ui_data_delete_user_source_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_delete_user_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_get_org_data**
> UIData ui_data_get_org_data(data_key, org_uid, source_uid, user_uid)

Get Org UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Get Org UI Data
        api_response = api_instance.ui_data_get_org_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_get_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_get_source_data**
> UIData ui_data_get_source_data(data_key, org_uid, source_uid, user_uid)

Get Source UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Get Source UI Data
        api_response = api_instance.ui_data_get_source_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_get_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_get_user_data**
> UIData ui_data_get_user_data(data_key, org_uid, source_uid, user_uid)

Get User UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Get User UI Data
        api_response = api_instance.ui_data_get_user_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_get_user_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_get_user_org_data**
> UIData ui_data_get_user_org_data(data_key, org_uid, source_uid, user_uid)

Get UserOrg UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Get UserOrg UI Data
        api_response = api_instance.ui_data_get_user_org_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_get_user_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_get_user_source_data**
> UIData ui_data_get_user_source_data(data_key, org_uid, source_uid, user_uid)

Get UserSource UI Data

 Gets UI Data   * Updates lastused  * Requires *uidata:Get*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID

    # example passing only required values which don't have defaults set
    try:
        # Get UserSource UI Data
        api_response = api_instance.ui_data_get_user_source_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_get_user_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |

### Return type

[**UIData**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_query_org_data**
> [UIData] ui_data_query_org_data(data_key, org_uid, source_uid, user_uid)

Query Org UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID
    tags = [
        "tags_example",
    ] # [str] | Tags to search for (optional)

    # example passing only required values which don't have defaults set
    try:
        # Query Org UI Data
        api_response = api_instance.ui_data_query_org_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_org_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query Org UI Data
        api_response = api_instance.ui_data_query_org_data(data_key, org_uid, source_uid, user_uid, tags=tags)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |
 **tags** | **[str]**| Tags to search for | [optional]

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_query_source_data**
> [UIData] ui_data_query_source_data(data_key, org_uid, source_uid, user_uid)

Query Source UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID
    tags = [
        "tags_example",
    ] # [str] | Tags to search for (optional)

    # example passing only required values which don't have defaults set
    try:
        # Query Source UI Data
        api_response = api_instance.ui_data_query_source_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_source_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query Source UI Data
        api_response = api_instance.ui_data_query_source_data(data_key, org_uid, source_uid, user_uid, tags=tags)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |
 **tags** | **[str]**| Tags to search for | [optional]

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_query_user_data**
> [UIData] ui_data_query_user_data(data_key, org_uid, source_uid, user_uid)

Query User UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID
    tags = [
        "tags_example",
    ] # [str] | Tags to search for (optional)

    # example passing only required values which don't have defaults set
    try:
        # Query User UI Data
        api_response = api_instance.ui_data_query_user_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_user_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query User UI Data
        api_response = api_instance.ui_data_query_user_data(data_key, org_uid, source_uid, user_uid, tags=tags)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_user_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |
 **tags** | **[str]**| Tags to search for | [optional]

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_query_user_org_data**
> [UIData] ui_data_query_user_org_data(data_key, org_uid, source_uid, user_uid)

Query UserOrg UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID
    tags = [
        "tags_example",
    ] # [str] | Tags to search for (optional)

    # example passing only required values which don't have defaults set
    try:
        # Query UserOrg UI Data
        api_response = api_instance.ui_data_query_user_org_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_user_org_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query UserOrg UI Data
        api_response = api_instance.ui_data_query_user_org_data(data_key, org_uid, source_uid, user_uid, tags=tags)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_user_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |
 **tags** | **[str]**| Tags to search for | [optional]

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_query_user_source_data**
> [UIData] ui_data_query_user_source_data(data_key, org_uid, source_uid, user_uid)

Query UserSource UI Data

 Query UI Data   * Requires *uidata:Query*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data import UIData
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Src UID
    user_uid = "userUID_example" # str | Owner UID
    tags = [
        "tags_example",
    ] # [str] | Tags to search for (optional)

    # example passing only required values which don't have defaults set
    try:
        # Query UserSource UI Data
        api_response = api_instance.ui_data_query_user_source_data(data_key, org_uid, source_uid, user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_user_source_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query UserSource UI Data
        api_response = api_instance.ui_data_query_user_source_data(data_key, org_uid, source_uid, user_uid, tags=tags)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_query_user_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Src UID |
 **user_uid** | **str**| Owner UID |
 **tags** | **[str]**| Tags to search for | [optional]

### Return type

[**[UIData]**](UIData.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_set_org_data**
> ui_data_set_org_data(data_key, org_uid, source_uid, user_uid)

Set Org UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data_set_org_data_input import UiDataSetOrgDataInput
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Source UID
    user_uid = "userUID_example" # str | Owner UID
    ui_data_set_org_data_input = UiDataSetOrgDataInput(
        ="_example",
        data={
            "key": None,
        },
        last_used=dateutil_parser('1970-01-01T00:00:00.00Z'),
        tags=[
            "tags_example",
        ],
        uid="uid_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # UiDataSetOrgDataInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set Org UI Data
        api_instance.ui_data_set_org_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_org_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set Org UI Data
        api_instance.ui_data_set_org_data(data_key, org_uid, source_uid, user_uid, ui_data_set_org_data_input=ui_data_set_org_data_input)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Source UID |
 **user_uid** | **str**| Owner UID |
 **ui_data_set_org_data_input** | [**UiDataSetOrgDataInput**](UiDataSetOrgDataInput.md)|  | [optional]

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_set_source_data**
> ui_data_set_source_data(data_key, org_uid, source_uid, user_uid)

Set Source UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data_set_source_data_input import UiDataSetSourceDataInput
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Source UID
    user_uid = "userUID_example" # str | Owner UID
    ui_data_set_source_data_input = UiDataSetSourceDataInput(
        ="_example",
        data={
            "key": None,
        },
        last_used=dateutil_parser('1970-01-01T00:00:00.00Z'),
        tags=[
            "tags_example",
        ],
        uid="uid_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # UiDataSetSourceDataInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set Source UI Data
        api_instance.ui_data_set_source_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_source_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set Source UI Data
        api_instance.ui_data_set_source_data(data_key, org_uid, source_uid, user_uid, ui_data_set_source_data_input=ui_data_set_source_data_input)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Source UID |
 **user_uid** | **str**| Owner UID |
 **ui_data_set_source_data_input** | [**UiDataSetSourceDataInput**](UiDataSetSourceDataInput.md)|  | [optional]

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_set_user_data**
> ui_data_set_user_data(data_key, org_uid, source_uid, user_uid)

Set User UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data_set_user_data_input import UiDataSetUserDataInput
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Source UID
    user_uid = "userUID_example" # str | Owner UID
    ui_data_set_user_data_input = UiDataSetUserDataInput(
        ="_example",
        data={
            "key": None,
        },
        last_used=dateutil_parser('1970-01-01T00:00:00.00Z'),
        tags=[
            "tags_example",
        ],
        uid="uid_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # UiDataSetUserDataInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set User UI Data
        api_instance.ui_data_set_user_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_user_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set User UI Data
        api_instance.ui_data_set_user_data(data_key, org_uid, source_uid, user_uid, ui_data_set_user_data_input=ui_data_set_user_data_input)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_user_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Source UID |
 **user_uid** | **str**| Owner UID |
 **ui_data_set_user_data_input** | [**UiDataSetUserDataInput**](UiDataSetUserDataInput.md)|  | [optional]

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_set_user_org_data**
> ui_data_set_user_org_data(data_key, org_uid, source_uid, user_uid)

Set UserOrg UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data_set_user_org_data_input import UiDataSetUserOrgDataInput
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Source UID
    user_uid = "userUID_example" # str | Owner UID
    ui_data_set_user_org_data_input = UiDataSetUserOrgDataInput(
        ="_example",
        data={
            "key": None,
        },
        last_used=dateutil_parser('1970-01-01T00:00:00.00Z'),
        tags=[
            "tags_example",
        ],
        uid="uid_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # UiDataSetUserOrgDataInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set UserOrg UI Data
        api_instance.ui_data_set_user_org_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_user_org_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set UserOrg UI Data
        api_instance.ui_data_set_user_org_data(data_key, org_uid, source_uid, user_uid, ui_data_set_user_org_data_input=ui_data_set_user_org_data_input)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_user_org_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Source UID |
 **user_uid** | **str**| Owner UID |
 **ui_data_set_user_org_data_input** | [**UiDataSetUserOrgDataInput**](UiDataSetUserOrgDataInput.md)|  | [optional]

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ui_data_set_user_source_data**
> ui_data_set_user_source_data(data_key, org_uid, source_uid, user_uid)

Set UserSource UI Data

 Sets UI Data   * Updates data, valid_to (if set), lastused  * Requires *uidata:Set*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import ui_data_api
from sbapi.model.ui_data_set_user_source_data_input import UiDataSetUserSourceDataInput
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
    api_instance = ui_data_api.UIDataApi(api_client)
    data_key = "dataKey_example" # str | Key for the data
    org_uid = "orgUID_example" # str | Org UID
    source_uid = "sourceUID_example" # str | Source UID
    user_uid = "userUID_example" # str | Owner UID
    ui_data_set_user_source_data_input = UiDataSetUserSourceDataInput(
        ="_example",
        data={
            "key": None,
        },
        last_used=dateutil_parser('1970-01-01T00:00:00.00Z'),
        tags=[
            "tags_example",
        ],
        uid="uid_example",
        valid_from=dateutil_parser('1970-01-01T00:00:00.00Z'),
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # UiDataSetUserSourceDataInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set UserSource UI Data
        api_instance.ui_data_set_user_source_data(data_key, org_uid, source_uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_user_source_data: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set UserSource UI Data
        api_instance.ui_data_set_user_source_data(data_key, org_uid, source_uid, user_uid, ui_data_set_user_source_data_input=ui_data_set_user_source_data_input)
    except sbapi.ApiException as e:
        print("Exception when calling UIDataApi->ui_data_set_user_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_key** | **str**| Key for the data |
 **org_uid** | **str**| Org UID |
 **source_uid** | **str**| Source UID |
 **user_uid** | **str**| Owner UID |
 **ui_data_set_user_source_data_input** | [**UiDataSetUserSourceDataInput**](UiDataSetUserSourceDataInput.md)|  | [optional]

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

