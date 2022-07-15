# spyderbat_api.RBACApi

All URIs are relative to *https://api.spyderbat.com*

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
import spyderbat_api
from spyderbat_api.api import rbac_api
from spyderbat_api.model.can_user_perform_input import CanUserPerformInput
from spyderbat_api.model.api_rbac_actions import ApiRBACActions
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
    except spyderbat_api.ApiException as e:
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

