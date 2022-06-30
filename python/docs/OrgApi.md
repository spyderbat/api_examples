# sbapi.OrgApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**org_assign_role**](OrgApi.md#org_assign_role) | **POST** /api/v1/org/{orgUID}/assignedrole/add | Assign OrgRole
[**org_invite_users**](OrgApi.md#org_invite_users) | **POST** /api/v1/org/{orgUID}/invite | Invite users to an organization
[**org_list**](OrgApi.md#org_list) | **GET** /api/v1/org/ | List organizations
[**org_list_role**](OrgApi.md#org_list_role) | **GET** /api/v1/org/{orgUID}/assignedrole/ | List OrgRoles
[**org_load**](OrgApi.md#org_load) | **GET** /api/v1/org/{orgUID} | Load an organization
[**org_load_notification_policy**](OrgApi.md#org_load_notification_policy) | **GET** /api/v1/org/{orgUID}/notification_policy/ | Load Notification Policy
[**org_test_notification_target**](OrgApi.md#org_test_notification_target) | **POST** /api/v1/org/{orgUID}/notification_policy/test_target | Test Notification Target
[**org_unassign_role**](OrgApi.md#org_unassign_role) | **POST** /api/v1/org/{orgUID}/assignedrole/del | Unassign OrgRole
[**org_update**](OrgApi.md#org_update) | **PUT** /api/v1/org/{orgUID} | Update an organization
[**org_update_notification_policy**](OrgApi.md#org_update_notification_policy) | **PUT** /api/v1/org/{orgUID}/notification_policy | Update an organization&#39;s notification policy


# **org_assign_role**
> org_assign_role(org_uid)

Assign OrgRole

 Assigns a role to a particular user on an organization   * Requires the user have the action *user.AssignRole* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org_assign_role_input import OrgAssignRoleInput
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 
    org_assign_role_input = OrgAssignRoleInput(
        role_uid="role_uid_example",
        user_uid="user_uid_example",
    ) # OrgAssignRoleInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Assign OrgRole
        api_instance.org_assign_role(org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_assign_role: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Assign OrgRole
        api_instance.org_assign_role(org_uid, org_assign_role_input=org_assign_role_input)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_assign_role: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **org_assign_role_input** | [**OrgAssignRoleInput**](OrgAssignRoleInput.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **org_invite_users**
> org_invite_users(org_uid)

Invite users to an organization

 Invites users to an organization   * Requires action *org:InviteUsers* on the organization to invite users  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org_invite_users_input import OrgInviteUsersInput
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 
    org_invite_users_input = OrgInviteUsersInput(
        emails=[
            "emails_example",
        ],
        roles=[
            "roles_example",
        ],
    ) # OrgInviteUsersInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Invite users to an organization
        api_instance.org_invite_users(org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_invite_users: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Invite users to an organization
        api_instance.org_invite_users(org_uid, org_invite_users_input=org_invite_users_input)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_invite_users: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **org_invite_users_input** | [**OrgInviteUsersInput**](OrgInviteUsersInput.md)|  | [optional]

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

# **org_list**
> [Org] org_list()

List organizations

 Lists organizations   * Will list organizations which the user has the action *org:Load* or *org:LoadExpired* on 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org import Org
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
    api_instance = org_api.OrgApi(api_client)
    has_resource_policy = True # bool |  (optional)
    has_tags = [
        "has_tags_example",
    ] # [str] |  (optional)
    include_expired = True # bool |  (optional)
    name_contains = "name_contains_example" # str |  (optional)
    owner_uid_equals = "owner_uid_equals_example" # str |  (optional)
    page = 1 # int |  (optional)
    page_size = 10 # int |  (optional)
    uid_equals = "uid_equals_example" # str |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List organizations
        api_response = api_instance.org_list(has_resource_policy=has_resource_policy, has_tags=has_tags, include_expired=include_expired, name_contains=name_contains, owner_uid_equals=owner_uid_equals, page=page, page_size=page_size, uid_equals=uid_equals)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **has_resource_policy** | **bool**|  | [optional]
 **has_tags** | **[str]**|  | [optional]
 **include_expired** | **bool**|  | [optional]
 **name_contains** | **str**|  | [optional]
 **owner_uid_equals** | **str**|  | [optional]
 **page** | **int**|  | [optional]
 **page_size** | **int**|  | [optional]
 **uid_equals** | **str**|  | [optional]

### Return type

[**[Org]**](Org.md)

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

# **org_list_role**
> [DaoOrgRoleResponse] org_list_role(org_uid)

List OrgRoles

 Allows querying of roles on the organization   * Requires the user have the action *org:ListOrgRoles* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.dao_org_role_response import DaoOrgRoleResponse
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 
    role_uid_equals = "role_uid_equals_example" # str |  (optional)
    user_email_equals = "user_email_equals_example" # str |  (optional)
    user_uid_equals = "user_uid_equals_example" # str |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # List OrgRoles
        api_response = api_instance.org_list_role(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_list_role: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List OrgRoles
        api_response = api_instance.org_list_role(org_uid, role_uid_equals=role_uid_equals, user_email_equals=user_email_equals, user_uid_equals=user_uid_equals)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_list_role: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **role_uid_equals** | **str**|  | [optional]
 **user_email_equals** | **str**|  | [optional]
 **user_uid_equals** | **str**|  | [optional]

### Return type

[**[DaoOrgRoleResponse]**](DaoOrgRoleResponse.md)

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

# **org_load**
> Org org_load(org_uid)

Load an organization

 Loads an organization by UID.    * Requires action *org:Load*  * Requires action *org:LoadExpired* to load expired organizations  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org import Org
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load an organization
        api_response = api_instance.org_load(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_load: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |

### Return type

[**Org**](Org.md)

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

# **org_load_notification_policy**
> NotificationPolicy org_load_notification_policy(org_uid)

Load Notification Policy

 Loads the notification policy for an organization. The notification policy defines who and how the organization is notified.     * If the content-type is application/hjson the policy will be returned as hjson  * If the content-type is application/json the policy will be returned as json    * Requires the user have the action *org:LoadNotificationPolicy* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.notification_policy import NotificationPolicy
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Load Notification Policy
        api_response = api_instance.org_load_notification_policy(org_uid)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_load_notification_policy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |

### Return type

[**NotificationPolicy**](NotificationPolicy.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/hjson, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **org_test_notification_target**
> org_test_notification_target(org_uid)

Test Notification Target

 Sends a test notification to a target.   * Requires the user have the action *org:SendTestNotification* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org_test_notification_target_input import OrgTestNotificationTargetInput
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 
    org_test_notification_target_input = OrgTestNotificationTargetInput(
        target="target_example",
    ) # OrgTestNotificationTargetInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Test Notification Target
        api_instance.org_test_notification_target(org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_test_notification_target: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Test Notification Target
        api_instance.org_test_notification_target(org_uid, org_test_notification_target_input=org_test_notification_target_input)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_test_notification_target: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **org_test_notification_target_input** | [**OrgTestNotificationTargetInput**](OrgTestNotificationTargetInput.md)|  | [optional]

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

# **org_unassign_role**
> org_unassign_role(org_uid)

Unassign OrgRole

 Unassigns a role to a particular user on an organization   * Requires the user have the action *user.UnassignRole* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org_unassign_role_input import OrgUnassignRoleInput
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 
    org_unassign_role_input = OrgUnassignRoleInput(
        role_uid="role_uid_example",
        user_uid="user_uid_example",
    ) # OrgUnassignRoleInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Unassign OrgRole
        api_instance.org_unassign_role(org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_unassign_role: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Unassign OrgRole
        api_instance.org_unassign_role(org_uid, org_unassign_role_input=org_unassign_role_input)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_unassign_role: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **org_unassign_role_input** | [**OrgUnassignRoleInput**](OrgUnassignRoleInput.md)|  | [optional]

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

# **org_update**
> org_update(org_uid)

Update an organization

 Updates the organization    * Requires the user have the action *org:Update* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.org_update_input import OrgUpdateInput
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | Org UID
    org_update_input = OrgUpdateInput(
        name="name_example",
        org_type_uid="org_type_uid_example",
        owner_email="owner_email_example",
        owner_uid="owner_uid_example",
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
    ) # OrgUpdateInput |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an organization
        api_instance.org_update(org_uid)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_update: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an organization
        api_instance.org_update(org_uid, org_update_input=org_update_input)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| Org UID |
 **org_update_input** | [**OrgUpdateInput**](OrgUpdateInput.md)|  | [optional]

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

# **org_update_notification_policy**
> org_update_notification_policy(org_uid, notification_policy)

Update an organization's notification policy

 Updates the organization's notification policy  The policy contains targets, which are named destinations, for example a list of admins, and then rules for how notifications are routed. The most basic notification is the default policy assigned to all organizations, which is to send notifications to the organization owner.   Each notification has a schema which can be used to filter how notifications are routed.   Here is an example notification policy:  ``` {     \"targets\": {       \"admins\": {         \"emails\": [           \"admin1@foo.com\",           \"admin2@foo.com\"         ]       },       \"soc\": {         \"slack\": {           \"url\": \"http://app.slack.com/XXXX\"         }       },       \"sns_example\": {         \"data\": \"UI supplied data\",         \"description\": \"Cross-Account Notification\",         \"sns\": {           \"sns_topic_arn\": \"arn:aws:sns:CUSTOMER-AWS-REGION:CUSTOMER-AWS-ACCOUNT-ID:CUSTOMER-TOPIC-NAME\",           \"cross_account_iam_role\": \"arn:aws:iam::CUSTOMER-ACCOUNT-ID:role/CUSTOMER-ROLE-NAME\"         }       }     },     \"routes\": [       {         \"target\": \"admins\",         \"expr\": {           \"schema\": \"agent_offline\"         }       },       {         \"target\": \"soc\",         \"expr\": {           \"schema\": \"spydertrace_updated\"         }       },       {         \"destination\": {           \"users\": [             \"X23hs8234lks\"           ]         }       }     ]   } ```  This policy says that any notification with the schema \"agent_offline\" is send to the admin emails, and that any notifications with the schema \"spydertrace_update\" is send to the soc slack channel, and all other notifications are sent to a user specified by their UserUID. If a destination is an explicit UserUID then the users notification policy will applied, for example to notify them by their notification type of choice.     * The following destination types are currently supported, UserUID, Email, Slack, Webhook, SNS. See the associated destination definition in the notification policy for details.   * First match for a routing rule wins, and further processsing of the notification stops.   * If a UserUID is specified the users notification policy may be used to contact that user.  * Each notification has an associated schema, which is used to provide a consistent schema for the notification.  * Expressions are optional, if no expression is specified the route matches by default  * If the content-type is application/hjson the policy will be parsed and stored as hjson, all comments will be lost on an existing hjson policy if it is uploaded as json  This is an example notification generated using a notification policy and dashboardsearch:  ``` {  \"uid\": \"lQ0Q1lKm\",  \"org_uid\": \"your_org_uid\",  \"valid_from\": \"2021-10-14T19:22:00.869159169Z\",  \"title\": \"Spyderbat: Dashboard search notification Recent interactive (shell) Processes\",  \"message\": \"Spyderbat: Dashboard search notification Recent interactive (shell) Processes\",  \"data\": {    \"dashboardsearch\": {   \"data\": {     \"createTime\": 1634237497.142,     \"createdBy\": \"user@example.com\"   },   \"description\": \"Recent interactive (shell) Processes\",   \"notify\": true,   \"notify_frequency\": 300,   \"org_uid\": \"spyderbatuid\",   \"search\": \"schema:model_process AND interactive:true\",   \"uid\": \"RTBQoR3uucyjG8ZeHJmw\"    }  },  \"schema\": \"dashboard_saved_search\",  \"records\": [    {   \"total_hits\": 0    }  ]   } ```   * Requires the user have the action *org:UpdateNotificationPolicy* on the organization 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import org_api
from sbapi.model.notification_policy import NotificationPolicy
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
    api_instance = org_api.OrgApi(api_client)
    org_uid = "orgUID_example" # str | 
    notification_policy = NotificationPolicy(
        routes=[
            NotificationPolicyRoutesInner(
                data=None,
                description="description_example",
                destination=NotificationPolicyDestination(
                    data=None,
                    description="description_example",
                    email=[
                        "email_example",
                    ],
                    org_uid="org_uid_example",
                    slack=NotificationPolicyDestinationSlack(
                        url="url_example",
                    ),
                    users=[
                        "users_example",
                    ],
                    webhook=NotificationPolicyDestinationWebhook(
                        no_tls_validation=True,
                        url="url_example",
                    ),
                ),
                expr=Expr(
                    _and=[
                        Expr(),
                    ],
                    contains_str="contains_str_example",
                    equals=None,
                    exists=True,
                    greater_than=None,
                    has_prefix="has_prefix_example",
                    has_suffix="has_suffix_example",
                    _in=[
                        None,
                    ],
                    less_than=None,
                    _not=Expr(),
                    _or=[
                        Expr(),
                    ],
                    _property="_property_example",
                    re_match="re_match_example",
                    schema="schema_example",
                    time_range=RstreamTimeRange(
                        end_time=3.14,
                        start_time=3.14,
                    ),
                ),
                target="target_example",
            ),
        ],
        targets={
            "key": NotificationPolicyDestination(
                data=None,
                description="description_example",
                email=[
                    "email_example",
                ],
                org_uid="org_uid_example",
                slack=NotificationPolicyDestinationSlack(
                    url="url_example",
                ),
                users=[
                    "users_example",
                ],
                webhook=NotificationPolicyDestinationWebhook(
                    no_tls_validation=True,
                    url="url_example",
                ),
            ),
        },
    ) # NotificationPolicy | The notification policy

    # example passing only required values which don't have defaults set
    try:
        # Update an organization's notification policy
        api_instance.org_update_notification_policy(org_uid, notification_policy)
    except sbapi.ApiException as e:
        print("Exception when calling OrgApi->org_update_notification_policy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**|  |
 **notification_policy** | [**NotificationPolicy**](NotificationPolicy.md)| The notification policy |

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

