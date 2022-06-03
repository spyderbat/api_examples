# RstoreCausalQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peer** | Pointer to **[]map[string]interface{}** | Causal tree matching rules for peer causal records | [optional] 
**Post** | Pointer to **[]map[string]interface{}** | Causal tree matching rules for post causal records | [optional] 
**Pre** | Pointer to **[]map[string]interface{}** | Causal tree matching rules for pre causal records | [optional] 

## Methods

### NewRstoreCausalQuery

`func NewRstoreCausalQuery() *RstoreCausalQuery`

NewRstoreCausalQuery instantiates a new RstoreCausalQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRstoreCausalQueryWithDefaults

`func NewRstoreCausalQueryWithDefaults() *RstoreCausalQuery`

NewRstoreCausalQueryWithDefaults instantiates a new RstoreCausalQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeer

`func (o *RstoreCausalQuery) GetPeer() []map[string]interface{}`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *RstoreCausalQuery) GetPeerOk() (*[]map[string]interface{}, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *RstoreCausalQuery) SetPeer(v []map[string]interface{})`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *RstoreCausalQuery) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPost

`func (o *RstoreCausalQuery) GetPost() []map[string]interface{}`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *RstoreCausalQuery) GetPostOk() (*[]map[string]interface{}, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *RstoreCausalQuery) SetPost(v []map[string]interface{})`

SetPost sets Post field to given value.

### HasPost

`func (o *RstoreCausalQuery) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPre

`func (o *RstoreCausalQuery) GetPre() []map[string]interface{}`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *RstoreCausalQuery) GetPreOk() (*[]map[string]interface{}, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *RstoreCausalQuery) SetPre(v []map[string]interface{})`

SetPre sets Pre field to given value.

### HasPre

`func (o *RstoreCausalQuery) HasPre() bool`

HasPre returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


