# InvestigationCreateInput


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_by** | **str** | UID of user who created the investigation | [optional] 
**data** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | JSON Object associated with the investigation | [optional] 
**modified_by** | **str** | UID of the user who last modified the investigation | [optional] 
**modified_on** | **datetime** | Date the investigation was last modified | [optional] 
**name** | **str** | Name of the investigation | [optional] 
**resource_name** | **str** | Resource name used for RBAC | [optional] 
**resource_policy** | [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**tags** | **[str]** | User supplied tags | [optional] 
**valid_from** | **datetime** | Valid from date, the first date this object was valid | [optional] 
**valid_to** | **datetime** | Valid to date, the date this object is valid to | [optional] 
**version** | **int** | Version of the investigation | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


