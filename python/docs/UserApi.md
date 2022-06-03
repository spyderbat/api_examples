# sbapi.UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_auth**](UserApi.md#user_auth) | **POST** /api/v1/app/user/auth | Authenticate a user, returns token required for authentication
[**user_current**](UserApi.md#user_current) | **GET** /api/v1/app/user/current | Returns the current user
[**user_load**](UserApi.md#user_load) | **GET** /api/v1/user/{userUID} | Load a user by ID
[**user_signup**](UserApi.md#user_signup) | **GET** /api/v1/app/signup | Sign up a new user


# **user_auth**
> User user_auth()

Authenticate a user, returns token required for authentication

 Users may be authenticate by providing their email address and password as a JSON object, in response an authentication token will be returned in the header 'Authorization-Token'. The returned token must be used in a request header with the name 'Authorization' when making other authenticated API calls, for example 'Authorization: Bearer TOKEN'. 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import user_api
from sbapi.model.user_auth_input import UserAuthInput
from sbapi.model.user import User
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
    api_instance = user_api.UserApi(api_client)
    user_auth_input = UserAuthInput(
        email="email_example",
        password="password_example",
    ) # UserAuthInput |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Authenticate a user, returns token required for authentication
        api_response = api_instance.user_auth(user_auth_input=user_auth_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling UserApi->user_auth: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_auth_input** | [**UserAuthInput**](UserAuthInput.md)|  | [optional]

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  * Authorization-Token - Authorization token to be provided when making authenticated calls <br>  |
**400** | Invalid auth request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_current**
> User user_current()

Returns the current user

 This api call will return the current user as a JSON object, or 404 if there is no user associated with the current authentication token or JWT. When a 403 is received the message will indicate if the user is expired or is pending approval, i.e. \"user is expired\", or if the user is in a pending state \"user is in a pending state\". 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import user_api
from sbapi.model.user import User
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

[apiToken](../README.md#apiToken)

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

# **user_signup**
> user_signup()

Sign up a new user

 Users can sign up by providing an email address. If the email is invalid or a user with the same existing email address has already registered an error is returned  ### Default configuration  By default the user will have an organization created for them, with the all actions allowed on the organization. The user can discover their organization id by  list all organizations.     

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import user_api
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
    api_instance = user_api.UserApi(api_client)
    auth = "auth_example" # str |  (optional)
    confirm = "confirm_example" # str |  (optional)
    email = "email_example" # str |  (optional)
    enable = True # bool |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Sign up a new user
        api_instance.user_signup(auth=auth, confirm=confirm, email=email, enable=enable)
    except sbapi.ApiException as e:
        print("Exception when calling UserApi->user_signup: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth** | **str**|  | [optional]
 **confirm** | **str**|  | [optional]
 **email** | **str**|  | [optional]
 **enable** | **bool**|  | [optional]

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
**400** | Invalid user signup |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

