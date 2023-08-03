# InvestigationCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_by** | Option<**String**> | UID of user who created the investigation | [optional]
**data** | Option<[**::std::collections::HashMap<String, serde_json::Value>**](serde_json::Value.md)> | JSON Object associated with the investigation | [optional]
**modified_by** | Option<**String**> | UID of the user who last modified the investigation | [optional]
**modified_on** | Option<**String**> | Date the investigation was last modified | [optional]
**name** | Option<**String**> | Name of the investigation | [optional]
**resource_name** | Option<**String**> | Resource name used for RBAC | [optional]
**resource_policy** | Option<[**crate::models::ResourcePolicy**](ResourcePolicy.md)> |  | [optional]
**tags** | Option<**Vec<String>**> | User supplied tags | [optional]
**valid_from** | Option<**String**> | Valid from date, the first date this object was valid | [optional]
**valid_to** | Option<**String**> | Valid to date, the date this object is valid to | [optional]
**version** | Option<**i32**> | Version of the investigation | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


