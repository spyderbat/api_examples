# DashboardSearchUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | UI supplied JSON object | [optional] 
**Description** | Pointer to **string** | Description of query (max 64 chars) | [optional] 
**Notify** | Pointer to **bool** | Are notifications generated from this search | [optional] 
**NotifyFrequency** | Pointer to **int32** | Frequency of notification in seconds. One of 300, 3600, or 86400. | [optional] 
**Search** | Pointer to **string** | Search query to run | [optional] 
**Tags** | Pointer to **[]string** | User supplied tags | [optional] 

## Methods

### NewDashboardSearchUpdateInput

`func NewDashboardSearchUpdateInput() *DashboardSearchUpdateInput`

NewDashboardSearchUpdateInput instantiates a new DashboardSearchUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSearchUpdateInputWithDefaults

`func NewDashboardSearchUpdateInputWithDefaults() *DashboardSearchUpdateInput`

NewDashboardSearchUpdateInputWithDefaults instantiates a new DashboardSearchUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DashboardSearchUpdateInput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardSearchUpdateInput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardSearchUpdateInput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DashboardSearchUpdateInput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *DashboardSearchUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardSearchUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardSearchUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardSearchUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotify

`func (o *DashboardSearchUpdateInput) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *DashboardSearchUpdateInput) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *DashboardSearchUpdateInput) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *DashboardSearchUpdateInput) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyFrequency

`func (o *DashboardSearchUpdateInput) GetNotifyFrequency() int32`

GetNotifyFrequency returns the NotifyFrequency field if non-nil, zero value otherwise.

### GetNotifyFrequencyOk

`func (o *DashboardSearchUpdateInput) GetNotifyFrequencyOk() (*int32, bool)`

GetNotifyFrequencyOk returns a tuple with the NotifyFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrequency

`func (o *DashboardSearchUpdateInput) SetNotifyFrequency(v int32)`

SetNotifyFrequency sets NotifyFrequency field to given value.

### HasNotifyFrequency

`func (o *DashboardSearchUpdateInput) HasNotifyFrequency() bool`

HasNotifyFrequency returns a boolean if a field has been set.

### GetSearch

`func (o *DashboardSearchUpdateInput) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *DashboardSearchUpdateInput) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *DashboardSearchUpdateInput) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *DashboardSearchUpdateInput) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTags

`func (o *DashboardSearchUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardSearchUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardSearchUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardSearchUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


