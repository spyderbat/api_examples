# RbacStatement


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**effect** | **str** | The effect of the statement, i.e. allow or deny | 
**sid** | **str** | Statement ID, used to identify the statement in audit logs | 
**actions** | **[str]** | The actions that may be performed | [optional] 
**condition** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Additional conditions which may be applied | [optional] 
**resources** | **[str]** | The resource the statement applies to | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


