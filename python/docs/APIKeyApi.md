# sbapi.APIKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_key_create**](APIKeyApi.md#api_key_create) | **POST** /api/v1/user/{userUID}/apikey/ | Creates a new API key
[**api_key_delete**](APIKeyApi.md#api_key_delete) | **DELETE** /api/v1/user/{userUID}/apikey/{uid} | Delete an API key
[**api_key_list**](APIKeyApi.md#api_key_list) | **GET** /api/v1/user/{userUID}/apikey/ | Lists an API key
[**api_key_update**](APIKeyApi.md#api_key_update) | **PUT** /api/v1/user/{userUID}/apikey/{uid} | Updates an API key


# **api_key_create**
> str api_key_create(user_uid)

Creates a new API key

 Creates a new API key which is associated with the user    * Requires global action *user:APIKeyCreate* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import api_key_api
from sbapi.model.api_key_create_input import ApiKeyCreateInput
from sbapi.model.validation_error import ValidationError
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
    api_instance = api_key_api.APIKeyApi(api_client)
    user_uid = "userUID_example" # str | User UID
    api_key_create_input = ApiKeyCreateInput(
        description="description_example",
        valid_to=dateutil_parser('1970-01-01T00:00:00.00Z'),
    ) # ApiKeyCreateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates a new API key
        api_response = api_instance.api_key_create(user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling APIKeyApi->api_key_create: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a new API key
        api_response = api_instance.api_key_create(user_uid, api_key_create_input=api_key_create_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling APIKeyApi->api_key_create: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| User UID |
 **api_key_create_input** | [**ApiKeyCreateInput**](ApiKeyCreateInput.md)|  | [optional]

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
**400** | Invalid user |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_key_delete**
> api_key_delete(uid, user_uid)

Delete an API key

 Deletes a specific API key   * Requires global action *user:APIKeyDelete* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import api_key_api
from sbapi.model.validation_error import ValidationError
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
    api_instance = api_key_api.APIKeyApi(api_client)
    uid = "uid_example" # str | API Key UID
    user_uid = "userUID_example" # str | User UID

    # example passing only required values which don't have defaults set
    try:
        # Delete an API key
        api_instance.api_key_delete(uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling APIKeyApi->api_key_delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| API Key UID |
 **user_uid** | **str**| User UID |

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

# **api_key_list**
> [APIKey] api_key_list(user_uid)

Lists an API key

 Lists API keys associated with a user   * Requires global action *user:APIKeyList*  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import api_key_api
from sbapi.model.api_key import APIKey
from sbapi.model.validation_error import ValidationError
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
    api_instance = api_key_api.APIKeyApi(api_client)
    user_uid = "userUID_example" # str | User UID

    # example passing only required values which don't have defaults set
    try:
        # Lists an API key
        api_response = api_instance.api_key_list(user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling APIKeyApi->api_key_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| User UID |

### Return type

[**[APIKey]**](APIKey.md)

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

# **api_key_update**
> api_key_update(uid, user_uid)

Updates an API key

 Updates a specific API Key, the only fields which can be updated are description and enabled   * Requires global action *user:APIKeyUpdate* 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import api_key_api
from sbapi.model.api_key_update_input import ApiKeyUpdateInput
from sbapi.model.validation_error import ValidationError
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
    api_instance = api_key_api.APIKeyApi(api_client)
    uid = "uid_example" # str | API Key UID
    user_uid = "userUID_example" # str | User UID
    api_key_update_input = ApiKeyUpdateInput(
        description="description_example",
        enabled=True,
    ) # ApiKeyUpdateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates an API key
        api_instance.api_key_update(uid, user_uid)
    except sbapi.ApiException as e:
        print("Exception when calling APIKeyApi->api_key_update: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates an API key
        api_instance.api_key_update(uid, user_uid, api_key_update_input=api_key_update_input)
    except sbapi.ApiException as e:
        print("Exception when calling APIKeyApi->api_key_update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| API Key UID |
 **user_uid** | **str**| User UID |
 **api_key_update_input** | [**ApiKeyUpdateInput**](ApiKeyUpdateInput.md)|  | [optional]

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
