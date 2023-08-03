# SpyderbatApi.OrgUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**activeSources** | **Number** | Total number of active sources within the last 5 minutes | [optional] 
**activeUsers** | **Number** | Total number of active users within the last 7 days (which might be active on a different org) | [optional] 
**name** | **String** | Name of the organization | 
**orgTypeUid** | **String** | Organization Type | [optional] 
**ownerEmail** | **String** | The email address of the user who owns this org | 
**ownerUid** | **String** | The user UID who owns this organization | [optional] 
**resourceName** | **String** | Resource name utilized by RBAC | [optional] 
**resourcePolicy** | [**ResourcePolicy**](ResourcePolicy.md) |  | [optional] 
**tags** | **[String]** | User supplied tags | [optional] 
**totalSources** | **Number** | Total number of sources | [optional] 
**totalUsers** | **Number** | Total number of users | [optional] 
**validFrom** | **Date** | Valid from date, the first date this object was valid | [optional] 
**validTo** | **Date** | Valid to date, the date this object is valid to | [optional] 


