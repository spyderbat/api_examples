# ElasticRecordSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dynamic** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to [**map[string]ElasticRecordField**](ElasticRecordField.md) | List of properties and how they are stored | [optional] 

## Methods

### NewElasticRecordSchema

`func NewElasticRecordSchema() *ElasticRecordSchema`

NewElasticRecordSchema instantiates a new ElasticRecordSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticRecordSchemaWithDefaults

`func NewElasticRecordSchemaWithDefaults() *ElasticRecordSchema`

NewElasticRecordSchemaWithDefaults instantiates a new ElasticRecordSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamic

`func (o *ElasticRecordSchema) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *ElasticRecordSchema) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *ElasticRecordSchema) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *ElasticRecordSchema) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetProperties

`func (o *ElasticRecordSchema) GetProperties() map[string]ElasticRecordField`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ElasticRecordSchema) GetPropertiesOk() (*map[string]ElasticRecordField, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ElasticRecordSchema) SetProperties(v map[string]ElasticRecordField)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ElasticRecordSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


