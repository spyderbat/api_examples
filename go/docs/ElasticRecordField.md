# ElasticRecordField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dynamic** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to [**map[string]ElasticRecordField**](ElasticRecordField.md) |  | [optional] 
**Index** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to [**map[string]ElasticRecordField**](ElasticRecordField.md) | The properties associated with this field | [optional] 
**Store** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | The type used for indexing, keyword&#x3D;matches entire seaerch term, text&#x3D;partial match, ip&#x3D;ip address, float&#x3D;number | [optional] 

## Methods

### NewElasticRecordField

`func NewElasticRecordField() *ElasticRecordField`

NewElasticRecordField instantiates a new ElasticRecordField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticRecordFieldWithDefaults

`func NewElasticRecordFieldWithDefaults() *ElasticRecordField`

NewElasticRecordFieldWithDefaults instantiates a new ElasticRecordField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamic

`func (o *ElasticRecordField) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *ElasticRecordField) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *ElasticRecordField) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *ElasticRecordField) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetFields

`func (o *ElasticRecordField) GetFields() map[string]ElasticRecordField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ElasticRecordField) GetFieldsOk() (*map[string]ElasticRecordField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ElasticRecordField) SetFields(v map[string]ElasticRecordField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ElasticRecordField) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIndex

`func (o *ElasticRecordField) GetIndex() bool`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ElasticRecordField) GetIndexOk() (*bool, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ElasticRecordField) SetIndex(v bool)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ElasticRecordField) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetProperties

`func (o *ElasticRecordField) GetProperties() map[string]ElasticRecordField`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ElasticRecordField) GetPropertiesOk() (*map[string]ElasticRecordField, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ElasticRecordField) SetProperties(v map[string]ElasticRecordField)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ElasticRecordField) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStore

`func (o *ElasticRecordField) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ElasticRecordField) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ElasticRecordField) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ElasticRecordField) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetType

`func (o *ElasticRecordField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ElasticRecordField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ElasticRecordField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ElasticRecordField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


