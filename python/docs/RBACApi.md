# sbapi.RBACApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**can_user_perform**](RBACApi.md#can_user_perform) | **POST** /api/v1/rbac/capabilities/ | Query allows actions on objects


# **can_user_perform**
> ApiRBACActions can_user_perform()

Query allows actions on objects

Allows for querying of what actions a user can perform; results may be cached for a short period of time.

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import rbac_api
from sbapi.model.api_rbac_actions import ApiRBACActions
from sbapi.model.can_user_perform_input import CanUserPerformInput
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
    api_instance = rbac_api.RBACApi(api_client)
    can_user_perform_input = CanUserPerformInput(
        actions=[
            RBACAction(
                action="action_example",
                can_perform=True,
                error="error_example",
                resource_name="resource_name_example",
            ),
        ],
    ) # CanUserPerformInput |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query allows actions on objects
        api_response = api_instance.can_user_perform(can_user_perform_input=can_user_perform_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling RBACApi->can_user_perform: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **can_user_perform_input** | [**CanUserPerformInput**](CanUserPerformInput.md)|  | [optional]

### Return type

[**ApiRBACActions**](ApiRBACActions.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user signup |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

