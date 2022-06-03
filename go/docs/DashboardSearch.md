# DashboardSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | UI supplied JSON object | [optional] 
**Description** | Pointer to **string** | Description of query (max 64 chars) | [optional] 
**Notify** | Pointer to **bool** | Are notifications generated from this search | [optional] 
**NotifyFrequency** | Pointer to **int32** | Frequency of notification in seconds. One of 300, 3600, or 86400. | [optional] 
**OrgUid** | Pointer to **string** | Org UID | [optional] 
**Search** | Pointer to **string** | Search query to run | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 
**Uid** | Pointer to **string** | UID for the DashboardSearch | [optional] 

## Methods

### NewDashboardSearch

`func NewDashboardSearch() *DashboardSearch`

NewDashboardSearch instantiates a new DashboardSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSearchWithDefaults

`func NewDashboardSearchWithDefaults() *DashboardSearch`

NewDashboardSearchWithDefaults instantiates a new DashboardSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DashboardSearch) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardSearch) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardSearch) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DashboardSearch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *DashboardSearch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardSearch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardSearch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardSearch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotify

`func (o *DashboardSearch) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *DashboardSearch) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *DashboardSearch) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *DashboardSearch) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyFrequency

`func (o *DashboardSearch) GetNotifyFrequency() int32`

GetNotifyFrequency returns the NotifyFrequency field if non-nil, zero value otherwise.

### GetNotifyFrequencyOk

`func (o *DashboardSearch) GetNotifyFrequencyOk() (*int32, bool)`

GetNotifyFrequencyOk returns a tuple with the NotifyFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrequency

`func (o *DashboardSearch) SetNotifyFrequency(v int32)`

SetNotifyFrequency sets NotifyFrequency field to given value.

### HasNotifyFrequency

`func (o *DashboardSearch) HasNotifyFrequency() bool`

HasNotifyFrequency returns a boolean if a field has been set.

### GetOrgUid

`func (o *DashboardSearch) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *DashboardSearch) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *DashboardSearch) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *DashboardSearch) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetSearch

`func (o *DashboardSearch) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *DashboardSearch) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *DashboardSearch) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *DashboardSearch) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTags

`func (o *DashboardSearch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardSearch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardSearch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardSearch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUid

`func (o *DashboardSearch) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DashboardSearch) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DashboardSearch) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DashboardSearch) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


