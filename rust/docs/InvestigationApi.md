# \InvestigationApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**investigation_create**](InvestigationApi.md#investigation_create) | **POST** /api/v1/org/{orgUID}/investigation/ | Create an investigation
[**investigation_delete**](InvestigationApi.md#investigation_delete) | **DELETE** /api/v1/org/{orgUID}/investigation/{investigationUID} | Delete an investigation
[**investigation_list**](InvestigationApi.md#investigation_list) | **GET** /api/v1/org/{orgUID}/investigation/ | List investigations
[**investigation_list_versions**](InvestigationApi.md#investigation_list_versions) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/ | List Investigation Versions
[**investigation_load**](InvestigationApi.md#investigation_load) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID} | Load an investigation
[**investigation_load_version**](InvestigationApi.md#investigation_load_version) | **GET** /api/v1/org/{orgUID}/investigation/{investigationUID}/version/{version} | Load Investigation Version
[**investigation_update**](InvestigationApi.md#investigation_update) | **PUT** /api/v1/org/{orgUID}/investigation/{investigationUID} | Update an investigation



## investigation_create

> crate::models::ApiInvestigationCreateOutput investigation_create(org_uid, investigation_create_input)
Create an investigation

 Create an investigationan   * Requires the user have the action *investigation:Create* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** | Investigation OrgUID | [required] |
**investigation_create_input** | Option<[**InvestigationCreateInput**](InvestigationCreateInput.md)> |  |  |

### Return type

[**crate::models::ApiInvestigationCreateOutput**](ApiInvestigationCreateOutput.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## investigation_delete

> investigation_delete(investigation_uid, org_uid)
Delete an investigation

 Deletes an investigation, by setting valid_to=now so that the investigation is virtually deleted.   * Requires the user have the action *investigation:Delete* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investigation_uid** | **String** | Investigation UID | [required] |
**org_uid** | **String** | Investigation OrgUID | [required] |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## investigation_list

> Vec<crate::models::DaoInvestigation> investigation_list(org_uid)
List investigations

 Lists investigations   * Will list investigations which the user has the action *investigation:Load* or *investigation:LoadExpired* on 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**org_uid** | **String** |  | [required] |

### Return type

[**Vec<crate::models::DaoInvestigation>**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## investigation_list_versions

> Vec<crate::models::DaoInvestigation> investigation_list_versions(investigation_uid, org_uid)
List Investigation Versions

 Lists prior version of this investigation   * Requires the user have the action *investigation:ListVersions* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investigation_uid** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |

### Return type

[**Vec<crate::models::DaoInvestigation>**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## investigation_load

> crate::models::DaoInvestigation investigation_load(investigation_uid, org_uid)
Load an investigation

 Loads an investigation by UID.    * Requires action  *investigation:Load* to load an active investigation  * Requires action *investigation:LoadExpired* to load expired investigations  

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investigation_uid** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |

### Return type

[**crate::models::DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## investigation_load_version

> crate::models::DaoInvestigation investigation_load_version(investigation_uid, org_uid, version)
Load Investigation Version

 Loads a specific version of an investigation   * Requires the user have the action *investigation:LoadVersion* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investigation_uid** | **String** |  | [required] |
**org_uid** | **String** |  | [required] |
**version** | **i32** |  | [required] |

### Return type

[**crate::models::DaoInvestigation**](DaoInvestigation.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## investigation_update

> investigation_update(investigation_uid, org_uid, investigation_update_input)
Update an investigation

 Updates the investigationan   * Requires the user have the action *investigation:Update* 

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investigation_uid** | **String** | Investigation UID | [required] |
**org_uid** | **String** | Investigation OrgUID | [required] |
**investigation_update_input** | Option<[**InvestigationUpdateInput**](InvestigationUpdateInput.md)> |  |  |

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

