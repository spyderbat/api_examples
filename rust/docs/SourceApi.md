# \SourceApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**integration_soar_src_list**](SourceApi.md#integration_soar_src_list) | **GET** /api/v1/integration/soar/org/{orgUID}/source/ | List sources for integration with SOARs
[**src_create**](SourceApi.md#src_create) | **POST** /api/v1/org/{orgUID}/source/ | Create a source
[**src_delete**](SourceApi.md#src_delete) | **DELETE** /api/v1/org/{orgUID}/source/{sourceUID} | Delete a source
[**src_list**](SourceApi.md#src_list) | **GET** /api/v1/org/{orgUID}/source/ | List sources
[**src_load**](SourceApi.md#src_load) | **GET** /api/v1/org/{orgUID}/source/{sourceUID} | Load a source
[**src_update**](SourceApi.md#src_update) | **PUT** /api/v1/org/{orgUID}/source/{sourceUID} | Update a source



## integration_soar_src_list

> Vec<crate::models::ApiSoarListHandlerOutput> integration_soar_src_list(org_uid, et, hostname, ip_address, mac_address, page, page_size, st)
List sources for integration with SOARs

 Lists the sources of data that match the specified query parameters, and return  URL entry points into the UI for matching sources.   * Requires the action  *org:ListSources* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**et** | Option<**i64**> | optional end time of the query |  |
**hostname** | Option<**String**> | A single hostname to match |  |
**ip_address** | Option<**String**> | A single IP address to match |  |
**mac_address** | Option<**String**> | A single mac address to match |  |
**page** | Option<**i32**> |  |  |
**page_size** | Option<**i32**> |  |  |
**st** | Option<**i64**> | optional start time of the query, if only a start time is provided, end time will be start+10m |  |

### Return type

[**Vec<crate::models::ApiSoarListHandlerOutput>**](ApiSOARListHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_create

> crate::models::ApiSourceCreateHandlerOutput src_create(org_uid, src_create_input)
Create a source

 Creates a new source of data  * Requires the action  *org:CreateSource* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | The org this source is associated with | [required] |
**src_create_input** | Option<[**SrcCreateInput**](SrcCreateInput.md)> |  |  |

### Return type

[**crate::models::ApiSourceCreateHandlerOutput**](ApiSourceCreateHandlerOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_delete

> src_delete(org_uid, source_uid)
Delete a source

 Delete a source  * Requires the action  *org:DeleteSource* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**source_uid** | **String** |  | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_list

> Vec<crate::models::Source> src_list(org_uid, agent_uid_equals, description_contains, has_tags, original_association, page, page_size)
List sources

 Lists the sources of data for an organization  * Requires the action  *org:ListSources* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**agent_uid_equals** | Option<**String**> |  |  |
**description_contains** | Option<**String**> |  |  |
**has_tags** | Option<[**Vec<String>**](String.md)> |  |  |
**original_association** | Option<**bool**> |  |  |
**page** | Option<**i32**> |  |  |
**page_size** | Option<**i32**> |  |  |

### Return type

[**Vec<crate::models::Source>**](Source.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_load

> crate::models::Source src_load(org_uid, source_uid)
Load a source

 Loads the meta data for a source of data  * Requires the action  *org:LoadSource* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |
**source_uid** | **String** |  | [required] |

### Return type

[**crate::models::Source**](Source.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## src_update

> src_update(org_uid, source_uid, src_update_input)
Update a source

 Updates the meta data for a source of data  * Requires the action  *org:UpdateSource* on the organization 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | The org this source is associated with | [required] |
**source_uid** | **String** | The UID of the source | [required] |
**src_update_input** | Option<[**SrcUpdateInput**](SrcUpdateInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

