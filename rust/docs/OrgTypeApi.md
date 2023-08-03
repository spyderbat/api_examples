# \OrgTypeApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**org_type_limit_active_sources**](OrgTypeApi.md#org_type_limit_active_sources) | **GET** /api/v1/org/{orgUID}/type/limit/active_sources | Loads limits regarding active sources
[**org_type_limit_org_roles**](OrgTypeApi.md#org_type_limit_org_roles) | **GET** /api/v1/org/{orgUID}/type/limit/org_roles | Loads limits regarding org roles
[**org_type_load**](OrgTypeApi.md#org_type_load) | **GET** /api/v1/org/{orgUID}/type | Load the org type for the organization



## org_type_limit_active_sources

> crate::models::SessionOrgTypeMaxLimit org_type_limit_active_sources(org_uid)
Loads limits regarding active sources

 Loads the limits regarding active sources allowed on the organization, the active sources in an org are calculated on a 5m0s basis.    * Requires action *org:Load* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_type_limit_org_roles

> crate::models::SessionOrgTypeMaxLimit org_type_limit_org_roles(org_uid)
Loads limits regarding org roles

 Loads the limit information regarding the number of associated roles allowed per an organization   * Requires action *org:Load* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::SessionOrgTypeMaxLimit**](SessionOrgTypeMaxLimit.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## org_type_load

> crate::models::DaoOrgType org_type_load(org_uid)
Load the org type for the organization

 Loads the org type for the organiation   * Requires action *org:Load* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::DaoOrgType**](DaoOrgType.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

