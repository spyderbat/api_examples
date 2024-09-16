# \DashboardSearchApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard_search_create**](DashboardSearchApi.md#dashboard_search_create) | **POST** /api/v1/org/{orgUID}/dashboard_search/ | Create a dashboard search
[**dashboard_search_delete**](DashboardSearchApi.md#dashboard_search_delete) | **DELETE** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**dashboard_search_get**](DashboardSearchApi.md#dashboard_search_get) | **GET** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Get a dashboard search
[**dashboard_search_list**](DashboardSearchApi.md#dashboard_search_list) | **GET** /api/v1/org/{orgUID}/dashboard_search/ | List dashboard searches
[**dashboard_search_update**](DashboardSearchApi.md#dashboard_search_update) | **PUT** /api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID} | Update a dashboard search



## dashboard_search_create

> String dashboard_search_create(org_uid, dashboard_search_create_input)
Create a dashboard search

 Create a dashboard search in an org.   * Requires action dashboard_search:Create

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | Org UID | [required] |
**dashboard_search_create_input** | Option<[**DashboardSearchCreateInput**](DashboardSearchCreateInput.md)> |  |  |

### Return type

**String**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## dashboard_search_delete

> dashboard_search_delete(dashboard_search_uid, org_uid)
Get a dashboard search

 Get a dashboard search in an org.   * Requires action dashboard_search:Delete

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**dashboard_search_uid** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## dashboard_search_get

> crate::models::DashboardSearch dashboard_search_get(dashboard_search_uid, org_uid)
Get a dashboard search

 Get a dashboard search in an org.   * Requires action dashboard_search:Get

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**dashboard_search_uid** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::DashboardSearch**](DashboardSearch.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## dashboard_search_list

> Vec<crate::models::DashboardSearch> dashboard_search_list(org_uid)
List dashboard searches

 Lists dashboard searches by org.   * Requires action dashboard_search:Query

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | Org UID | [required] |

### Return type

[**Vec<crate::models::DashboardSearch>**](DashboardSearch.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## dashboard_search_update

> dashboard_search_update(dashboard_search_uid, org_uid, dashboard_search_update_input)
Update a dashboard search

 Update a dashboard search in an org.   * Requires action dashboard_search:Update

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**dashboard_search_uid** | **String** | UID for the DashboardSearch | [required] |
**org_uid** | **String** | Org UID | [required] |
**dashboard_search_update_input** | Option<[**DashboardSearchUpdateInput**](DashboardSearchUpdateInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

