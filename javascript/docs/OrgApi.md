# SpyderbatApi.OrgApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**orgAssignRole**](OrgApi.md#orgAssignRole) | **POST** /api/v1/org/{orgUID}/assignedrole/add | Assign OrgRole
[**orgInviteUsers**](OrgApi.md#orgInviteUsers) | **POST** /api/v1/org/{orgUID}/invite | Invite users to an organization
[**orgList**](OrgApi.md#orgList) | **GET** /api/v1/org/ | List organizations
[**orgListRole**](OrgApi.md#orgListRole) | **GET** /api/v1/org/{orgUID}/assignedrole/ | List OrgRoles
[**orgLoad**](OrgApi.md#orgLoad) | **GET** /api/v1/org/{orgUID} | Load an organization
[**orgLoadNotificationPolicy**](OrgApi.md#orgLoadNotificationPolicy) | **GET** /api/v1/org/{orgUID}/notification_policy/ | Load Notification Policy
[**orgTestNotificationTarget**](OrgApi.md#orgTestNotificationTarget) | **POST** /api/v1/org/{orgUID}/notification_policy/test_target | Test Notification Target
[**orgUnassignRole**](OrgApi.md#orgUnassignRole) | **POST** /api/v1/org/{orgUID}/assignedrole/del | Unassign OrgRole
[**orgUpdate**](OrgApi.md#orgUpdate) | **PUT** /api/v1/org/{orgUID} | Update an organization
[**orgUpdateNotificationPolicy**](OrgApi.md#orgUpdateNotificationPolicy) | **PUT** /api/v1/org/{orgUID}/notification_policy | Update an organization&#39;s notification policy



## orgAssignRole

> orgAssignRole(orgUID, opts)

Assign OrgRole

 Assigns a role to a particular user on an organization   * Requires the user have the action *user.AssignRole* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'orgAssignRoleInput': new SpyderbatApi.OrgAssignRoleInput() // OrgAssignRoleInput | 
};
apiInstance.orgAssignRole(orgUID, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 
 **orgAssignRoleInput** | [**OrgAssignRoleInput**](OrgAssignRoleInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## orgInviteUsers

> orgInviteUsers(orgUID, opts)

Invite users to an organization

 Invites users to an organization   * Requires action *org:InviteUsers* on the organization to invite users  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'orgInviteUsersInput': new SpyderbatApi.OrgInviteUsersInput() // OrgInviteUsersInput | 
};
apiInstance.orgInviteUsers(orgUID, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 
 **orgInviteUsersInput** | [**OrgInviteUsersInput**](OrgInviteUsersInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## orgList

> [Org] orgList(opts)

List organizations

 Lists organizations   * Will list organizations which the user has the action *org:Load* or *org:LoadExpired* on 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let opts = {
  'hasResourcePolicy': true, // Boolean | 
  'hasTags': ["null"], // [String] | 
  'includeExpired': true, // Boolean | 
  'nameContains': "nameContains_example", // String | 
  'ownerUidEquals': "ownerUidEquals_example", // String | 
  'page': 56, // Number | 
  'pageSize': 56, // Number | 
  'uidEquals': "uidEquals_example" // String | 
};
apiInstance.orgList(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hasResourcePolicy** | **Boolean**|  | [optional] 
 **hasTags** | [**[String]**](String.md)|  | [optional] 
 **includeExpired** | **Boolean**|  | [optional] 
 **nameContains** | **String**|  | [optional] 
 **ownerUidEquals** | **String**|  | [optional] 
 **page** | **Number**|  | [optional] 
 **pageSize** | **Number**|  | [optional] 
 **uidEquals** | **String**|  | [optional] 

### Return type

[**[Org]**](Org.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgListRole

> [DaoOrgRoleResponse] orgListRole(orgUID, opts)

List OrgRoles

 Allows querying of roles on the organization   * Requires the user have the action *org:ListOrgRoles* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'roleUidEquals': "roleUidEquals_example", // String | 
  'userEmailEquals': "userEmailEquals_example", // String | 
  'userUidEquals': "userUidEquals_example" // String | 
};
apiInstance.orgListRole(orgUID, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 
 **roleUidEquals** | **String**|  | [optional] 
 **userEmailEquals** | **String**|  | [optional] 
 **userUidEquals** | **String**|  | [optional] 

### Return type

[**[DaoOrgRoleResponse]**](DaoOrgRoleResponse.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgLoad

> Org orgLoad(orgUID)

Load an organization

 Loads an organization by UID.    * Requires action *org:Load*  * Requires action *org:LoadExpired* to load expired organizations  

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
apiInstance.orgLoad(orgUID, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 

### Return type

[**Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## orgLoadNotificationPolicy

> NotificationPolicy orgLoadNotificationPolicy(orgUID)

Load Notification Policy

 Loads the notification policy for an organization. The notification policy defines who and how the organization is notified.     * If the content-type is application/hjson the policy will be returned as hjson  * If the content-type is application/json the policy will be returned as json    * Requires the user have the action *org:LoadNotificationPolicy* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
apiInstance.orgLoadNotificationPolicy(orgUID, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 

### Return type

[**NotificationPolicy**](NotificationPolicy.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hjson, application/json


## orgTestNotificationTarget

> orgTestNotificationTarget(orgUID, opts)

Test Notification Target

 Sends a test notification to a target.   * Requires the user have the action *org:SendTestNotification* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'orgTestNotificationTargetInput': new SpyderbatApi.OrgTestNotificationTargetInput() // OrgTestNotificationTargetInput | 
};
apiInstance.orgTestNotificationTarget(orgUID, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 
 **orgTestNotificationTargetInput** | [**OrgTestNotificationTargetInput**](OrgTestNotificationTargetInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## orgUnassignRole

> orgUnassignRole(orgUID, opts)

Unassign OrgRole

 Unassigns a role to a particular user on an organization   * Requires the user have the action *user.UnassignRole* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
let opts = {
  'orgUnassignRoleInput': new SpyderbatApi.OrgUnassignRoleInput() // OrgUnassignRoleInput | 
};
apiInstance.orgUnassignRole(orgUID, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 
 **orgUnassignRoleInput** | [**OrgUnassignRoleInput**](OrgUnassignRoleInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## orgUpdate

> orgUpdate(orgUID, opts)

Update an organization

 Updates the organization    * Requires the user have the action *org:Update* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | Org UID
let opts = {
  'orgUpdateInput': new SpyderbatApi.OrgUpdateInput() // OrgUpdateInput | 
};
apiInstance.orgUpdate(orgUID, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**| Org UID | 
 **orgUpdateInput** | [**OrgUpdateInput**](OrgUpdateInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## orgUpdateNotificationPolicy

> orgUpdateNotificationPolicy(orgUID, notificationPolicy)

Update an organization&#39;s notification policy

 Updates the organization&#39;s notification policy  The policy contains targets, which are named destinations, for example a list of admins, and then rules for how notifications are routed. The most basic notification is the default policy assigned to all organizations, which is to send notifications to the organization owner.   Each notification has a schema which can be used to filter how notifications are routed.   Here is an example notification policy:  &#x60;&#x60;&#x60; {     \&quot;targets\&quot;: {       \&quot;admins\&quot;: {         \&quot;emails\&quot;: [           \&quot;admin1@foo.com\&quot;,           \&quot;admin2@foo.com\&quot;         ]       },       \&quot;soc\&quot;: {         \&quot;slack\&quot;: {           \&quot;url\&quot;: \&quot;http://app.slack.com/XXXX\&quot;         }       },       \&quot;sns_example\&quot;: {         \&quot;data\&quot;: \&quot;UI supplied data\&quot;,         \&quot;description\&quot;: \&quot;Cross-Account Notification\&quot;,         \&quot;sns\&quot;: {           \&quot;sns_topic_arn\&quot;: \&quot;arn:aws:sns:CUSTOMER-AWS-REGION:CUSTOMER-AWS-ACCOUNT-ID:CUSTOMER-TOPIC-NAME\&quot;,           \&quot;cross_account_iam_role\&quot;: \&quot;arn:aws:iam::CUSTOMER-ACCOUNT-ID:role/CUSTOMER-ROLE-NAME\&quot;         }       }     },     \&quot;routes\&quot;: [       {         \&quot;target\&quot;: \&quot;admins\&quot;,         \&quot;expr\&quot;: {           \&quot;schema\&quot;: \&quot;agent_offline\&quot;         }       },       {         \&quot;target\&quot;: \&quot;soc\&quot;,         \&quot;expr\&quot;: {           \&quot;schema\&quot;: \&quot;spydertrace_updated\&quot;         }       },       {         \&quot;destination\&quot;: {           \&quot;users\&quot;: [             \&quot;X23hs8234lks\&quot;           ]         }       }     ]   } &#x60;&#x60;&#x60;  This policy says that any notification with the schema \&quot;agent_offline\&quot; is send to the admin emails, and that any notifications with the schema \&quot;spydertrace_update\&quot; is send to the soc slack channel, and all other notifications are sent to a user specified by their UserUID. If a destination is an explicit UserUID then the users notification policy will applied, for example to notify them by their notification type of choice.     * The following destination types are currently supported, UserUID, Email, Slack, Webhook, SNS. See the associated destination definition in the notification policy for details.   * First match for a routing rule wins, and further processsing of the notification stops.   * If a UserUID is specified the users notification policy may be used to contact that user.  * Each notification has an associated schema, which is used to provide a consistent schema for the notification.  * Expressions are optional, if no expression is specified the route matches by default  * If the content-type is application/hjson the policy will be parsed and stored as hjson, all comments will be lost on an existing hjson policy if it is uploaded as json  This is an example notification generated using a notification policy and dashboardsearch:  &#x60;&#x60;&#x60; {  \&quot;uid\&quot;: \&quot;lQ0Q1lKm\&quot;,  \&quot;org_uid\&quot;: \&quot;your_org_uid\&quot;,  \&quot;valid_from\&quot;: \&quot;2021-10-14T19:22:00.869159169Z\&quot;,  \&quot;title\&quot;: \&quot;Spyderbat: Dashboard search notification Recent interactive (shell) Processes\&quot;,  \&quot;message\&quot;: \&quot;Spyderbat: Dashboard search notification Recent interactive (shell) Processes\&quot;,  \&quot;data\&quot;: {    \&quot;dashboardsearch\&quot;: {   \&quot;data\&quot;: {     \&quot;createTime\&quot;: 1634237497.142,     \&quot;createdBy\&quot;: \&quot;user@example.com\&quot;   },   \&quot;description\&quot;: \&quot;Recent interactive (shell) Processes\&quot;,   \&quot;notify\&quot;: true,   \&quot;notify_frequency\&quot;: 300,   \&quot;org_uid\&quot;: \&quot;spyderbatuid\&quot;,   \&quot;search\&quot;: \&quot;schema:model_process AND interactive:true\&quot;,   \&quot;uid\&quot;: \&quot;RTBQoR3uucyjG8ZeHJmw\&quot;    }  },  \&quot;schema\&quot;: \&quot;dashboard_saved_search\&quot;,  \&quot;records\&quot;: [    {   \&quot;total_hits\&quot;: 0    }  ]   } &#x60;&#x60;&#x60;   * Requires the user have the action *org:UpdateNotificationPolicy* on the organization 

### Example

```javascript
import SpyderbatApi from 'spyderbat-api';
let defaultClient = SpyderbatApi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new SpyderbatApi.OrgApi();
let orgUID = "orgUID_example"; // String | 
let notificationPolicy = new SpyderbatApi.NotificationPolicy(); // NotificationPolicy | The notification policy
apiInstance.orgUpdateNotificationPolicy(orgUID, notificationPolicy, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUID** | **String**|  | 
 **notificationPolicy** | [**NotificationPolicy**](NotificationPolicy.md)| The notification policy | 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

