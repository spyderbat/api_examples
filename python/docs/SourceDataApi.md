# spyderbat_api.SourceDataApi

All URIs are relative to *https://api.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**src_data_query**](SourceDataApi.md#src_data_query) | **POST** /api/v1/source/query/ | Query source data
[**src_data_query_v2**](SourceDataApi.md#src_data_query_v2) | **GET** /api/v1/org/{orgUID}/data/ | Query source data
[**src_send_data**](SourceDataApi.md#src_send_data) | **POST** /api/v1/org/{orgUID}/source/{sourceUID}/data/{dataType} | Send data to a source, this is expected to be gzip compressed nd-json. The &#39;Content-Encoding&#39; header should be specified with a value of &#39;gzip&#39;. Alternatively, a multi-part form upload may be used with gzipped data up to a maximum size of 1MB.


# **src_data_query**
> str src_data_query()

Query source data

 Allows querying of the source data, data is stored as 'records' which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 24 hour period of time, it is the callers responsibility to construct an appropriate 24 hour range. * Documentation for the returned spydergraph datatype can be found at https://app.spyderbat.com/schema/spydergraph/index.html  * The user must have both the *org.ListSourceData* action on the org and *source_data.Query* action on the source * To get a count of results (up to 10K) but no data, use querySize: 0  

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import source_data_api
from spyderbat_api.model.src_data_query_input import SrcDataQueryInput
from spyderbat_api.model.validation_error import ValidationError
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
    api_instance = source_data_api.SourceDataApi(api_client)
    src_data_query_input = SrcDataQueryInput(
        aggs={
            "key": None,
        },
        causal=RstoreCausalQuery(
            peer=[
                {},
            ],
            post=[
                {},
            ],
            pre=[
                {},
            ],
        ),
        context_uid="context_uid_example",
        data_type="data_type_example",
        end_time=3.14,
        expr=Expr(
            _and=[
                Expr(),
            ],
            contains_str="contains_str_example",
            equals=None,
            exists=True,
            greater_than=None,
            has_prefix="has_prefix_example",
            has_suffix="has_suffix_example",
            _in=[
                None,
            ],
            less_than=None,
            _not=Expr(),
            _or=[
                Expr(),
            ],
            _property="_property_example",
            re_match="re_match_example",
            schema="schema_example",
            time_range=RstreamTimeRange(
                end_time=3.14,
                start_time=3.14,
            ),
        ),
        ids=[
            "ids_example",
        ],
        org_uid="org_uid_example",
        projection=[
            "projection_example",
        ],
        query="query_example",
        query_from=1,
        query_size=1,
        return_rptrs=True,
        rptrs=[
            "rptrs_example",
        ],
        sort_col="sort_col_example",
        sort_order="sort_order_example",
        src_uid="src_uid_example",
        start_time=3.14,
    ) # SrcDataQueryInput |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query source data
        api_response = api_instance.src_data_query(src_data_query_input=src_data_query_input)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling SourceDataApi->src_data_query: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **src_data_query_input** | [**SrcDataQueryInput**](SrcDataQueryInput.md)|  | [optional]

### Return type

**str**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-ndjson, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | invalid input parameters |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **src_data_query_v2**
> str src_data_query_v2(org_uid, dt)

Query source data

Same as the post query above except results are cached

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import source_data_api
from spyderbat_api.model.validation_error import ValidationError
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
    api_instance = source_data_api.SourceDataApi(api_client)
    org_uid = "orgUID_example" # str | Organization UID to query
    dt = "dt_example" # str | DataType to query
    e = "e_example" # str | Data which matches this expression are returned, as a json object (optional)
    et = 3.14 # float | Time in unix epoch time, records < time are returned (optional)
    id = [
        "id_example",
    ] # [str] | List of IDs to resolve (optional)
    pj = [
        "pj_example",
    ] # [str] | Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj=id&pj=version (optional)
    q = "q_example" # str | Lucene based search query (optional)
    qf = 1 # int | Where to start the query in the result set from (optional)
    qs = 1 # int | Size of the query result set (optional)
    rr = True # bool | Return rptrs mixed with the data (optional)
    src = "src_example" # str | Source UID to query (optional)
    st = 3.14 # float | Time in unix epoch time, records >= time are returned (optional)

    # example passing only required values which don't have defaults set
    try:
        # Query source data
        api_response = api_instance.src_data_query_v2(org_uid, dt)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling SourceDataApi->src_data_query_v2: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query source data
        api_response = api_instance.src_data_query_v2(org_uid, dt, e=e, et=et, id=id, pj=pj, q=q, qf=qf, qs=qs, rr=rr, src=src, st=st)
        pprint(api_response)
    except spyderbat_api.ApiException as e:
        print("Exception when calling SourceDataApi->src_data_query_v2: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_uid** | **str**| Organization UID to query |
 **dt** | **str**| DataType to query |
 **e** | **str**| Data which matches this expression are returned, as a json object | [optional]
 **et** | **float**| Time in unix epoch time, records &lt; time are returned | [optional]
 **id** | **[str]**| List of IDs to resolve | [optional]
 **pj** | **[str]**| Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj&#x3D;id&amp;pj&#x3D;version | [optional]
 **q** | **str**| Lucene based search query | [optional]
 **qf** | **int**| Where to start the query in the result set from | [optional]
 **qs** | **int**| Size of the query result set | [optional]
 **rr** | **bool**| Return rptrs mixed with the data | [optional]
 **src** | **str**| Source UID to query | [optional]
 **st** | **float**| Time in unix epoch time, records &gt;&#x3D; time are returned | [optional]

### Return type

**str**

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-ndjson, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | invalid input parameters |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **src_send_data**
> src_send_data(data_type, org_uid, source_uid, encoding, file)

Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'. Alternatively, a multi-part form upload may be used with gzipped data up to a maximum size of 1MB.

Sends data to a source

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import spyderbat_api
from spyderbat_api.api import source_data_api
from spyderbat_api.model.validation_error import ValidationError
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
    api_instance = source_data_api.SourceDataApi(api_client)
    data_type = "dataType_example" # str | 
    org_uid = "orgUID_example" # str | 
    source_uid = "sourceUID_example" # str | 
    encoding = "encoding_example" # str | must be gzip
    file = open('/path/to/file', 'rb') # file_type | The file to upload. The file must be a valid gzip-ed JSON file.

    # example passing only required values which don't have defaults set
    try:
        # Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'. Alternatively, a multi-part form upload may be used with gzipped data up to a maximum size of 1MB.
        api_instance.src_send_data(data_type, org_uid, source_uid, encoding, file)
    except spyderbat_api.ApiException as e:
        print("Exception when calling SourceDataApi->src_send_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_type** | **str**|  |
 **org_uid** | **str**|  |
 **source_uid** | **str**|  |
 **encoding** | **str**| must be gzip |
 **file** | **file_type**| The file to upload. The file must be a valid gzip-ed JSON file. |

### Return type

void (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | invalid input parameters |  -  |
**403** | permission denied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

