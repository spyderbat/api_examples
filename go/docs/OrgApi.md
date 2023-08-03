# \OrgApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgAssignRole**](OrgApi.md#OrgAssignRole) | **Post** /api/v1/org/{orgUID}/assignedrole/add | Assign OrgRole
[**OrgInviteUsers**](OrgApi.md#OrgInviteUsers) | **Post** /api/v1/org/{orgUID}/invite | Invite users to an organization
[**OrgList**](OrgApi.md#OrgList) | **Get** /api/v1/org/ | List organizations
[**OrgListRole**](OrgApi.md#OrgListRole) | **Get** /api/v1/org/{orgUID}/assignedrole/ | List OrgRoles
[**OrgLoad**](OrgApi.md#OrgLoad) | **Get** /api/v1/org/{orgUID} | Load an organization
[**OrgLoadNotificationPolicy**](OrgApi.md#OrgLoadNotificationPolicy) | **Get** /api/v1/org/{orgUID}/notification_policy/ | Load Notification Policy
[**OrgTestNotificationTarget**](OrgApi.md#OrgTestNotificationTarget) | **Post** /api/v1/org/{orgUID}/notification_policy/test_target | Test Notification Target
[**OrgUnassignRole**](OrgApi.md#OrgUnassignRole) | **Post** /api/v1/org/{orgUID}/assignedrole/del | Unassign OrgRole
[**OrgUpdate**](OrgApi.md#OrgUpdate) | **Put** /api/v1/org/{orgUID} | Update an organization
[**OrgUpdateNotificationPolicy**](OrgApi.md#OrgUpdateNotificationPolicy) | **Put** /api/v1/org/{orgUID}/notification_policy | Update an organization&#39;s notification policy



## OrgAssignRole

> OrgAssignRole(ctx, orgUID).OrgAssignRoleInput(orgAssignRoleInput).Execute()

Assign OrgRole



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 
    orgAssignRoleInput := *openapiclient.NewOrgAssignRoleInput() // OrgAssignRoleInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgAssignRole(context.Background(), orgUID).OrgAssignRoleInput(orgAssignRoleInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgAssignRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgAssignRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgAssignRoleInput** | [**OrgAssignRoleInput**](OrgAssignRoleInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgInviteUsers

> OrgInviteUsers(ctx, orgUID).OrgInviteUsersInput(orgInviteUsersInput).Execute()

Invite users to an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 
    orgInviteUsersInput := *openapiclient.NewOrgInviteUsersInput() // OrgInviteUsersInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgInviteUsers(context.Background(), orgUID).OrgInviteUsersInput(orgInviteUsersInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgInviteUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgInviteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgInviteUsersInput** | [**OrgInviteUsersInput**](OrgInviteUsersInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgList

> []Org OrgList(ctx).HasResourcePolicy(hasResourcePolicy).HasTags(hasTags).IncludeExpired(includeExpired).NameContains(nameContains).OwnerUidEquals(ownerUidEquals).Page(page).PageSize(pageSize).UidEquals(uidEquals).Execute()

List organizations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hasResourcePolicy := true // bool |  (optional)
    hasTags := []string{"Inner_example"} // []string |  (optional)
    includeExpired := true // bool |  (optional)
    nameContains := "nameContains_example" // string |  (optional)
    ownerUidEquals := "ownerUidEquals_example" // string |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    uidEquals := "uidEquals_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgList(context.Background()).HasResourcePolicy(hasResourcePolicy).HasTags(hasTags).IncludeExpired(includeExpired).NameContains(nameContains).OwnerUidEquals(ownerUidEquals).Page(page).PageSize(pageSize).UidEquals(uidEquals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgList`: []Org
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.OrgList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hasResourcePolicy** | **bool** |  | 
 **hasTags** | **[]string** |  | 
 **includeExpired** | **bool** |  | 
 **nameContains** | **string** |  | 
 **ownerUidEquals** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **uidEquals** | **string** |  | 

### Return type

[**[]Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgListRole

> []DaoOrgRoleResponse OrgListRole(ctx, orgUID).RoleUidEquals(roleUidEquals).UserEmailEquals(userEmailEquals).UserUidEquals(userUidEquals).Execute()

List OrgRoles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 
    roleUidEquals := "roleUidEquals_example" // string |  (optional)
    userEmailEquals := "userEmailEquals_example" // string |  (optional)
    userUidEquals := "userUidEquals_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgListRole(context.Background(), orgUID).RoleUidEquals(roleUidEquals).UserEmailEquals(userEmailEquals).UserUidEquals(userUidEquals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgListRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgListRole`: []DaoOrgRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.OrgListRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgListRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleUidEquals** | **string** |  | 
 **userEmailEquals** | **string** |  | 
 **userUidEquals** | **string** |  | 

### Return type

[**[]DaoOrgRoleResponse**](DaoOrgRoleResponse.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgLoad

> Org OrgLoad(ctx, orgUID).Execute()

Load an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgLoad(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgLoad`: Org
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.OrgLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgLoadNotificationPolicy

> NotificationPolicy OrgLoadNotificationPolicy(ctx, orgUID).Execute()

Load Notification Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgLoadNotificationPolicy(context.Background(), orgUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgLoadNotificationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgLoadNotificationPolicy`: NotificationPolicy
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.OrgLoadNotificationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgLoadNotificationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationPolicy**](NotificationPolicy.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgTestNotificationTarget

> OrgTestNotificationTarget(ctx, orgUID).OrgTestNotificationTargetInput(orgTestNotificationTargetInput).Execute()

Test Notification Target



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 
    orgTestNotificationTargetInput := *openapiclient.NewOrgTestNotificationTargetInput() // OrgTestNotificationTargetInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgTestNotificationTarget(context.Background(), orgUID).OrgTestNotificationTargetInput(orgTestNotificationTargetInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgTestNotificationTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgTestNotificationTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgTestNotificationTargetInput** | [**OrgTestNotificationTargetInput**](OrgTestNotificationTargetInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgUnassignRole

> OrgUnassignRole(ctx, orgUID).OrgUnassignRoleInput(orgUnassignRoleInput).Execute()

Unassign OrgRole



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 
    orgUnassignRoleInput := *openapiclient.NewOrgUnassignRoleInput() // OrgUnassignRoleInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgUnassignRole(context.Background(), orgUID).OrgUnassignRoleInput(orgUnassignRoleInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgUnassignRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgUnassignRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgUnassignRoleInput** | [**OrgUnassignRoleInput**](OrgUnassignRoleInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgUpdate

> OrgUpdate(ctx, orgUID).OrgUpdateInput(orgUpdateInput).Execute()

Update an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | Org UID
    orgUpdateInput := *openapiclient.NewOrgUpdateInput("Name_example", "OwnerEmail_example") // OrgUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgUpdate(context.Background(), orgUID).OrgUpdateInput(orgUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** | Org UID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgUpdateInput** | [**OrgUpdateInput**](OrgUpdateInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgUpdateNotificationPolicy

> OrgUpdateNotificationPolicy(ctx, orgUID).NotificationPolicy(notificationPolicy).Execute()

Update an organization's notification policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgUID := "orgUID_example" // string | 
    notificationPolicy := *openapiclient.NewNotificationPolicy([]openapiclient.NotificationPolicyRoutesInner{*openapiclient.NewNotificationPolicyRoutesInner()}, map[string]NotificationPolicyDestination{"key": *openapiclient.NewNotificationPolicyDestination()}) // NotificationPolicy | The notification policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.OrgUpdateNotificationPolicy(context.Background(), orgUID).NotificationPolicy(notificationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.OrgUpdateNotificationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgUpdateNotificationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationPolicy** | [**NotificationPolicy**](NotificationPolicy.md) | The notification policy | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

