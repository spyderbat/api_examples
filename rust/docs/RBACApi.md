# \RBACApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**can_user_perform**](RBACApi.md#can_user_perform) | **POST** /api/v1/rbac/capabilities/ | Query allows actions on objects



## can_user_perform

> crate::models::ApiRbacActions can_user_perform(can_user_perform_input)
Query allows actions on objects

Allows for querying of what actions a user can perform; results may be cached for a short period of time.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**can_user_perform_input** | Option<[**CanUserPerformInput**](CanUserPerformInput.md)> |  |  |

### Return type

[**crate::models::ApiRbacActions**](ApiRBACActions.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

