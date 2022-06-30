# sbapi.MetricsDataApi

All URIs are relative to *https://api.prod.spyderbat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metrics_data_query**](MetricsDataApi.md#metrics_data_query) | **POST** /api/v1/metrics/query/ | Query metrics data


# **metrics_data_query**
> str metrics_data_query()

Query metrics data

 Allows querying of the metrics for metricss, data is stored as 'records' which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 1 week period of time, it is the callers responsibility to construct an appropriate 24 hour range.  * The user must have both the *metrics:QueryData* action on the org 

### Example

* Bearer (JWT) Authentication (apiToken):

```python
import time
import sbapi
from sbapi.api import metrics_data_api
from sbapi.model.metrics_data_query_input import MetricsDataQueryInput
from sbapi.model.validation_error import ValidationError
from pprint import pprint
# Defining the host is optional and defaults to https://api.prod.spyderbat.com
# See configuration.py for a list of all supported configuration parameters.
configuration = sbapi.Configuration(
    host = "https://api.prod.spyderbat.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): apiToken
configuration = sbapi.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with sbapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = metrics_data_api.MetricsDataApi(api_client)
    metrics_data_query_input = MetricsDataQueryInput(
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
    ) # MetricsDataQueryInput |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Query metrics data
        api_response = api_instance.metrics_data_query(metrics_data_query_input=metrics_data_query_input)
        pprint(api_response)
    except sbapi.ApiException as e:
        print("Exception when calling MetricsDataApi->metrics_data_query: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metrics_data_query_input** | [**MetricsDataQueryInput**](MetricsDataQueryInput.md)|  | [optional]

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

