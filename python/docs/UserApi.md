# sbapi.UserApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_current**](UserApi.md#user_current) | **GET** /api/v1/app/user/current | Returns the current user
[**user_load**](UserApi.md#user_load) | **GET** /api/v1/user/{userUID} | Load a user by ID


# **user_current**
> User user_current()

Returns the current user

 This api call will return the current user as a JSON object, or 404 if there is no user associated with the current authentication token or JWT. When a 403 is received the message will indicate if the user is expired or is pending approval, i.e. \"user is expired\", or if the user is in a pending state \"user is in a pending state\". 

### Example


```python
import time
import sbapi
from sbapi.api import user_api
from sbapi.model.user import User
from pprint import pprint
# Defining the host is optional and defaults to https://api.prod.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "https://api.prod.spyderbat.com"
)


# Enter a context with an instance of the API client
with sbapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Returns the current user
        api_response = api_instance.user_current()
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UserApi->user_current: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | User is expired or in a pending state |  -  |
**404** | No user is associated |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_load**
> User user_load(user_uid)

Load a user by ID

 Loads a user given the user's user ID.    * By default users may only load their own user ID  * The action *user:Load* may be placed upon the users role to allowed loading other users 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import user_api
from sbapi.model.user import User
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
    api_instance = user_api.UserApi(api_client)
    user_uid = "userUID_example" # str | User UID

    # example passing only required values which don't have defaults set
    try:
        # Load a user by ID
        api_response = api_instance.user_load(user_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UserApi->user_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| User UID |

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**404** | No user with specified ID was found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

