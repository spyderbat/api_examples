# spyderbat_api.MetadataAPIApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metadata_search_schema**](MetadataAPIApi.md#metadata_search_schema) | **GET** /api/v1/meta/search/schema | Returns the schema used for search


# **metadata_search_schema**
> {str: (ElasticRecordSchema,)} metadata_search_schema()

Returns the schema used for search

 Returns meta data around which model and event properties are indexed for search, i.e. the search schema 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import metadata_api_api
from spyderbat_api.model.elastic_record_schema import ElasticRecordSchema
from pprint import pprint
# Defining the host is optional and defaults to https://api.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = spyderbat_api.Configuration(
    host = "https://api.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = spyderbat_api.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with spyderbat_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = metadata_api_api.MetadataAPIApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Returns the schema used for search
        api_response = api_instance.metadata_search_schema()
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling MetadataAPIApi->metadata_search_schema: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**{str: (ElasticRecordSchema,)}**](ElasticRecordSchema.md)

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

