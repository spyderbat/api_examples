# DashboardSearchCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | Option<[**::std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> | UI supplied JSON object | [optional]
**description** | Option<**String**> | Description of query (max 64 chars) | [optional]
**notify** | Option<**bool**> | Are notifications generated from this search | [optional]
**notify_frequency** | Option<**i32**> | Frequency of notification in seconds. One of 300, 3600, or 86400. | [optional]
**search** | Option<**String**> | Search query to run | [optional]
**tags** | Option<**Vec<String>**> | User supplied tags | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


