# DashboardSearchCreateInput

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

### NewDashboardSearchCreateInput

`func NewDashboardSearchCreateInput() *DashboardSearchCreateInput`

NewDashboardSearchCreateInput instantiates a new DashboardSearchCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSearchCreateInputWithDefaults

`func NewDashboardSearchCreateInputWithDefaults() *DashboardSearchCreateInput`

NewDashboardSearchCreateInputWithDefaults instantiates a new DashboardSearchCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DashboardSearchCreateInput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardSearchCreateInput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardSearchCreateInput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DashboardSearchCreateInput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *DashboardSearchCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardSearchCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardSearchCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardSearchCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotify

`func (o *DashboardSearchCreateInput) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *DashboardSearchCreateInput) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *DashboardSearchCreateInput) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *DashboardSearchCreateInput) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyFrequency

`func (o *DashboardSearchCreateInput) GetNotifyFrequency() int32`

GetNotifyFrequency returns the NotifyFrequency field if non-nil, zero value otherwise.

### GetNotifyFrequencyOk

`func (o *DashboardSearchCreateInput) GetNotifyFrequencyOk() (*int32, bool)`

GetNotifyFrequencyOk returns a tuple with the NotifyFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrequency

`func (o *DashboardSearchCreateInput) SetNotifyFrequency(v int32)`

SetNotifyFrequency sets NotifyFrequency field to given value.

### HasNotifyFrequency

`func (o *DashboardSearchCreateInput) HasNotifyFrequency() bool`

HasNotifyFrequency returns a boolean if a field has been set.

### GetSearch

`func (o *DashboardSearchCreateInput) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *DashboardSearchCreateInput) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *DashboardSearchCreateInput) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *DashboardSearchCreateInput) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTags

`func (o *DashboardSearchCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardSearchCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardSearchCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardSearchCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


