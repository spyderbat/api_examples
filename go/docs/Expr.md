# Expr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]Expr**](Expr.md) | returns true if all of the contained expressions are true | [optional] 
**ContainsStr** | Pointer to **string** | returns true if the property is a string and contains the specified value | [optional] 
**Equals** | Pointer to **interface{}** | returns true if the property matches the supplied value | [optional] 
**Exists** | Pointer to **bool** | returns true if the property exists | [optional] 
**GreaterThan** | Pointer to **interface{}** | returns true if the property is a number and is greater than this value | [optional] 
**HasPrefix** | Pointer to **string** | returns true if the property is a string and has the specified prefix | [optional] 
**HasSuffix** | Pointer to **string** | returns true if the property is a string and has the specified suffix | [optional] 
**In** | Pointer to **[]interface{}** | returns true if the property matches any of the values specified | [optional] 
**LessThan** | Pointer to **interface{}** | returns true if te property is a number and is less than this value | [optional] 
**Not** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Or** | Pointer to [**[]Expr**](Expr.md) | returns true if any of the contained expressions are true | [optional] 
**Property** | Pointer to **string** | property to match against, in dotted property notation | [optional] 
**ReMatch** | Pointer to **string** | returns true if the property is a string and matches the specified regex | [optional] 
**Schema** | Pointer to **string** | matches only records with the specified schema | [optional] 
**TimeRange** | Pointer to [**RstreamTimeRange**](RstreamTimeRange.md) |  | [optional] 

## Methods

### NewExpr

`func NewExpr() *Expr`

NewExpr instantiates a new Expr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExprWithDefaults

`func NewExprWithDefaults() *Expr`

NewExprWithDefaults instantiates a new Expr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *Expr) GetAnd() []Expr`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Expr) GetAndOk() (*[]Expr, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Expr) SetAnd(v []Expr)`

SetAnd sets And field to given value.

### HasAnd

`func (o *Expr) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetContainsStr

`func (o *Expr) GetContainsStr() string`

GetContainsStr returns the ContainsStr field if non-nil, zero value otherwise.

### GetContainsStrOk

`func (o *Expr) GetContainsStrOk() (*string, bool)`

GetContainsStrOk returns a tuple with the ContainsStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsStr

`func (o *Expr) SetContainsStr(v string)`

SetContainsStr sets ContainsStr field to given value.

### HasContainsStr

`func (o *Expr) HasContainsStr() bool`

HasContainsStr returns a boolean if a field has been set.

### GetEquals

`func (o *Expr) GetEquals() interface{}`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *Expr) GetEqualsOk() (*interface{}, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *Expr) SetEquals(v interface{})`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *Expr) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### SetEqualsNil

`func (o *Expr) SetEqualsNil(b bool)`

 SetEqualsNil sets the value for Equals to be an explicit nil

### UnsetEquals
`func (o *Expr) UnsetEquals()`

UnsetEquals ensures that no value is present for Equals, not even an explicit nil
### GetExists

`func (o *Expr) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *Expr) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *Expr) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *Expr) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetGreaterThan

`func (o *Expr) GetGreaterThan() interface{}`

GetGreaterThan returns the GreaterThan field if non-nil, zero value otherwise.

### GetGreaterThanOk

`func (o *Expr) GetGreaterThanOk() (*interface{}, bool)`

GetGreaterThanOk returns a tuple with the GreaterThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreaterThan

`func (o *Expr) SetGreaterThan(v interface{})`

SetGreaterThan sets GreaterThan field to given value.

### HasGreaterThan

`func (o *Expr) HasGreaterThan() bool`

HasGreaterThan returns a boolean if a field has been set.

### SetGreaterThanNil

`func (o *Expr) SetGreaterThanNil(b bool)`

 SetGreaterThanNil sets the value for GreaterThan to be an explicit nil

### UnsetGreaterThan
`func (o *Expr) UnsetGreaterThan()`

UnsetGreaterThan ensures that no value is present for GreaterThan, not even an explicit nil
### GetHasPrefix

`func (o *Expr) GetHasPrefix() string`

GetHasPrefix returns the HasPrefix field if non-nil, zero value otherwise.

### GetHasPrefixOk

`func (o *Expr) GetHasPrefixOk() (*string, bool)`

GetHasPrefixOk returns a tuple with the HasPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrefix

`func (o *Expr) SetHasPrefix(v string)`

SetHasPrefix sets HasPrefix field to given value.

### HasHasPrefix

`func (o *Expr) HasHasPrefix() bool`

HasHasPrefix returns a boolean if a field has been set.

### GetHasSuffix

`func (o *Expr) GetHasSuffix() string`

GetHasSuffix returns the HasSuffix field if non-nil, zero value otherwise.

### GetHasSuffixOk

`func (o *Expr) GetHasSuffixOk() (*string, bool)`

GetHasSuffixOk returns a tuple with the HasSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSuffix

`func (o *Expr) SetHasSuffix(v string)`

SetHasSuffix sets HasSuffix field to given value.

### HasHasSuffix

`func (o *Expr) HasHasSuffix() bool`

HasHasSuffix returns a boolean if a field has been set.

### GetIn

`func (o *Expr) GetIn() []interface{}`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *Expr) GetInOk() (*[]interface{}, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *Expr) SetIn(v []interface{})`

SetIn sets In field to given value.

### HasIn

`func (o *Expr) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetLessThan

`func (o *Expr) GetLessThan() interface{}`

GetLessThan returns the LessThan field if non-nil, zero value otherwise.

### GetLessThanOk

`func (o *Expr) GetLessThanOk() (*interface{}, bool)`

GetLessThanOk returns a tuple with the LessThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessThan

`func (o *Expr) SetLessThan(v interface{})`

SetLessThan sets LessThan field to given value.

### HasLessThan

`func (o *Expr) HasLessThan() bool`

HasLessThan returns a boolean if a field has been set.

### SetLessThanNil

`func (o *Expr) SetLessThanNil(b bool)`

 SetLessThanNil sets the value for LessThan to be an explicit nil

### UnsetLessThan
`func (o *Expr) UnsetLessThan()`

UnsetLessThan ensures that no value is present for LessThan, not even an explicit nil
### GetNot

`func (o *Expr) GetNot() Expr`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *Expr) GetNotOk() (*Expr, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *Expr) SetNot(v Expr)`

SetNot sets Not field to given value.

### HasNot

`func (o *Expr) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetOr

`func (o *Expr) GetOr() []Expr`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *Expr) GetOrOk() (*[]Expr, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *Expr) SetOr(v []Expr)`

SetOr sets Or field to given value.

### HasOr

`func (o *Expr) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *Expr) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Expr) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Expr) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Expr) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetReMatch

`func (o *Expr) GetReMatch() string`

GetReMatch returns the ReMatch field if non-nil, zero value otherwise.

### GetReMatchOk

`func (o *Expr) GetReMatchOk() (*string, bool)`

GetReMatchOk returns a tuple with the ReMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReMatch

`func (o *Expr) SetReMatch(v string)`

SetReMatch sets ReMatch field to given value.

### HasReMatch

`func (o *Expr) HasReMatch() bool`

HasReMatch returns a boolean if a field has been set.

### GetSchema

`func (o *Expr) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Expr) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Expr) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Expr) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTimeRange

`func (o *Expr) GetTimeRange() RstreamTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *Expr) GetTimeRangeOk() (*RstreamTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *Expr) SetTimeRange(v RstreamTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *Expr) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


