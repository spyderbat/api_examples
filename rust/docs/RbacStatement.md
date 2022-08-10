# RbacStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**actions** | Option<**Vec<String>**> | The actions that may be performed | [optional]
**condition** | Option<[**serde_json::Value**](.md)> | Additional conditions which may be applied | [optional]
**effect** | **String** | The effect of the statement, i.e. allow or deny | 
**resources** | Option<**Vec<String>**> | The resource the statement applies to | [optional]
**sid** | **String** | Statement ID, used to identify the statement in audit logs | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


