# Sbapi.MetricsDataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**metricsDataQuery**](MetricsDataApi.md#metricsDataQuery) | **POST** /api/v1/metrics/query/ | Query metrics data



## metricsDataQuery

> metricsDataQuery(opts)

Query metrics data

 Allows querying of the metrics for metricss, data is stored as &#39;records&#39; which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 1 week period of time, it is the callers responsibility to construct an appropriate 24 hour range.  * The user must have both the *metrics:QueryData* action on the org 

### Example

```javascript
import Sbapi from 'sbapi';
let defaultClient = Sbapi.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: apiToken
let apiToken = defaultClient.authentications['apiToken'];
apiToken.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new Sbapi.MetricsDataApi();
let opts = {
  'metricsDataQueryInput': new Sbapi.MetricsDataQueryInput() // MetricsDataQueryInput | 
};
apiInstance.metricsDataQuery(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsDataQueryInput** | [**MetricsDataQueryInput**](MetricsDataQueryInput.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

