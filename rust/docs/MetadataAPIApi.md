# \MetadataAPIApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metadata_search_schema**](MetadataAPIApi.md#metadata_search_schema) | **GET** /api/v1/meta/search/schema | Returns the schema used for search



## metadata_search_schema

> ::std::collections::HashMap<String, crate::models::ElasticRecordSchema> metadata_search_schema()
Returns the schema used for search

 Returns meta data around which model and event properties are indexed for search, i.e. the search schema 

### Parameters

This endpoint does not need any parameter.

### Return type

[**::std::collections::HashMap<String, crate::models::ElasticRecordSchema>**](ElasticRecordSchema.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

