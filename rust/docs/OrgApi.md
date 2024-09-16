# \OrgApi

All URIs are relative to *https://api.spyderbat.com*

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
[**org_update_notification_policy**](OrgApi.md#org_update_notification_policy) | **PUT** /api/v1/org/{orgUID}/notification_policy | Update an organization's notification policy



## org_assign_role

> org_assign_role(org_uid, org_assign_role_input)
Assign OrgRole

 Assigns a role to a particular user on an organization   * Requires the user have the action *user.AssignRole* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**org_assign_role_input** | Option<[**OrgAssignRoleInput**](OrgAssignRoleInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_invite_users

> org_invite_users(org_uid, org_invite_users_input)
Invite users to an organization

 Invites users to an organization   * Requires action *org:InviteUsers* on the organization to invite users  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**org_invite_users_input** | Option<[**OrgInviteUsersInput**](OrgInviteUsersInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_list

> Vec<crate::models::Org> org_list(has_resource_policy, has_tags, include_expired, name_contains, owner_uid_equals, page, page_size, uid_equals)
List organizations

 Lists organizations   * Will list organizations which the user has the action *org:Load* or *org:LoadExpired* on 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**has_resource_policy** | Option<**bool**> |  |  |
**has_tags** | Option<[**Vec<String>**](String.md)> |  |  |
**include_expired** | Option<**bool**> |  |  |
**name_contains** | Option<**String**> |  |  |
**owner_uid_equals** | Option<**String**> |  |  |
**page** | Option<**i32**> |  |  |
**page_size** | Option<**i32**> |  |  |
**uid_equals** | Option<**String**> |  |  |

### Return type

[**Vec<crate::models::Org>**](Org.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_list_role

> Vec<crate::models::DaoOrgRoleResponse> org_list_role(org_uid, role_uid_equals, user_email_equals, user_uid_equals)
List OrgRoles

 Allows querying of roles on the organization   * Requires the user have the action *org:ListOrgRoles* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**role_uid_equals** | Option<**String**> |  |  |
**user_email_equals** | Option<**String**> |  |  |
**user_uid_equals** | Option<**String**> |  |  |

### Return type

[**Vec<crate::models::DaoOrgRoleResponse>**](DaoOrgRoleResponse.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_load

> crate::models::Org org_load(org_uid)
Load an organization

 Loads an organization by UID.    * Requires action *org:Load*  * Requires action *org:LoadExpired* to load expired organizations  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_load_notification_policy

> crate::models::NotificationPolicy org_load_notification_policy(org_uid)
Load Notification Policy

 Loads the notification policy for an organization. The notification policy defines who and how the organization is notified.     * If the content-type is application/hjson the policy will be returned as hjson  * If the content-type is application/json the policy will be returned as json    * Requires the user have the action *org:LoadNotificationPolicy* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::NotificationPolicy**](NotificationPolicy.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_test_notification_target

> org_test_notification_target(org_uid, org_test_notification_target_input)
Test Notification Target

 Sends a test notification to a target.   * Requires the user have the action *org:SendTestNotification* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**org_test_notification_target_input** | Option<[**OrgTestNotificationTargetInput**](OrgTestNotificationTargetInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_unassign_role

> org_unassign_role(org_uid, org_unassign_role_input)
Unassign OrgRole

 Unassigns a role to a particular user on an organization   * Requires the user have the action *user.UnassignRole* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**org_unassign_role_input** | Option<[**OrgUnassignRoleInput**](OrgUnassignRoleInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_update

> org_update(org_uid, org_update_input)
Update an organization

 Updates the organization    * Requires the user have the action *org:Update* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | Org UID | [required] |
**org_update_input** | Option<[**OrgUpdateInput**](OrgUpdateInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_update_notification_policy

> org_update_notification_policy(org_uid, notification_policy)
Update an organization's notification policy

 Updates the organization's notification policy  The policy contains targets, which are named destinations, for example a list of admins, and then rules for how notifications are routed. The most basic notification is the default policy assigned to all organizations, which is to send notifications to the organization owner.   Each notification has a schema which can be used to filter how notifications are routed.   Here is an example notification policy:  ``` {     \"targets\": {       \"admins\": {         \"emails\": [           \"admin1@foo.com\",           \"admin2@foo.com\"         ]       },       \"soc\": {         \"slack\": {           \"url\": \"http://app.slack.com/XXXX\"         }       },       \"sns_example\": {         \"data\": \"UI supplied data\",         \"description\": \"Cross-Account Notification\",         \"sns\": {           \"sns_topic_arn\": \"arn:aws:sns:CUSTOMER-AWS-REGION:CUSTOMER-AWS-ACCOUNT-ID:CUSTOMER-TOPIC-NAME\",           \"cross_account_iam_role\": \"arn:aws:iam::CUSTOMER-ACCOUNT-ID:role/CUSTOMER-ROLE-NAME\"         }       }     },     \"routes\": [       {         \"target\": \"admins\",         \"expr\": {           \"schema\": \"agent_offline\"         }       },       {         \"target\": \"soc\",         \"expr\": {           \"schema\": \"spydertrace_updated\"         }       },       {         \"destination\": {           \"users\": [             \"X23hs8234lks\"           ]         }       }     ]   } ```  This policy says that any notification with the schema \"agent_offline\" is send to the admin emails, and that any notifications with the schema \"spydertrace_update\" is send to the soc slack channel, and all other notifications are sent to a user specified by their UserUID. If a destination is an explicit UserUID then the users notification policy will applied, for example to notify them by their notification type of choice.     * The following destination types are currently supported, UserUID, Email, Slack, Webhook, SNS. See the associated destination definition in the notification policy for details.   * First match for a routing rule wins, and further processsing of the notification stops.   * If a UserUID is specified the users notification policy may be used to contact that user.  * Each notification has an associated schema, which is used to provide a consistent schema for the notification.  * Expressions are optional, if no expression is specified the route matches by default  * If the content-type is application/hjson the policy will be parsed and stored as hjson, all comments will be lost on an existing hjson policy if it is uploaded as json  This is an example notification generated using a notification policy and dashboardsearch:  ``` {  \"uid\": \"lQ0Q1lKm\",  \"org_uid\": \"your_org_uid\",  \"valid_from\": \"2021-10-14T19:22:00.869159169Z\",  \"title\": \"Spyderbat: Dashboard search notification Recent interactive (shell) Processes\",  \"message\": \"Spyderbat: Dashboard search notification Recent interactive (shell) Processes\",  \"data\": {    \"dashboardsearch\": {   \"data\": {     \"createTime\": 1634237497.142,     \"createdBy\": \"user@example.com\"   },   \"description\": \"Recent interactive (shell) Processes\",   \"notify\": true,   \"notify_frequency\": 300,   \"org_uid\": \"spyderbatuid\",   \"search\": \"schema:model_process AND interactive:true\",   \"uid\": \"RTBQoR3uucyjG8ZeHJmw\"    }  },  \"schema\": \"dashboard_saved_search\",  \"records\": [    {   \"total_hits\": 0    }  ]   } ```   * Requires the user have the action *org:UpdateNotificationPolicy* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**notification_policy** | [**NotificationPolicy**](NotificationPolicy.md) | The notification policy | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

