# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active_sources** | Option<**i32**> | Total number of active sources within the last 5 minutes | [optional]
**active_users** | Option<**i32**> | Total number of active users within the last 7 days (which might be active on a different org) | [optional]
**name** | **String** | Name of the organization | 
**org_type_uid** | Option<**String**> | Organization Type | [optional]
**owner_email** | **String** | The email address of the user who owns this org | 
**owner_uid** | Option<**String**> | The user UID who owns this organization | [optional]
**resource_name** | Option<**String**> | Resource name utilized by RBAC | [optional]
**resource_policy** | Option<[**crate::models::ResourcePolicy**](ResourcePolicy.md)> |  | [optional]
**tags** | Option<**Vec<String>**> | User supplied tags | [optional]
**total_sources** | Option<**i32**> | Total number of sources | [optional]
**total_users** | Option<**i32**> | Total number of users | [optional]
**uid** | Option<**String**> | Org UID | [optional]
**valid_from** | Option<**String**> | Valid from date, the first date this object was valid | [optional]
**valid_to** | Option<**String**> | Valid to date, the date this object is valid to | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


