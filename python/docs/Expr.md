# Expr

Data which matches this expression are returned

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_and** | [**[Expr]**](Expr.md) | returns true if all of the contained expressions are true | [optional] 
**contains_str** | **str** | returns true if the property is a string and contains the specified value | [optional] 
**equals** | **bool, date, datetime, dict, float, int, list, str, none_type** | returns true if the property matches the supplied value | [optional] 
**exists** | **bool** | returns true if the property exists | [optional] 
**greater_than** | **bool, date, datetime, dict, float, int, list, str, none_type** | returns true if the property is a number and is greater than this value | [optional] 
**has_prefix** | **str** | returns true if the property is a string and has the specified prefix | [optional] 
**has_suffix** | **str** | returns true if the property is a string and has the specified suffix | [optional] 
**_in** | **[bool, date, datetime, dict, float, int, list, str, none_type]** | returns true if the property matches any of the values specified | [optional] 
**less_than** | **bool, date, datetime, dict, float, int, list, str, none_type** | returns true if te property is a number and is less than this value | [optional] 
**_not** | [**Expr**](Expr.md) |  | [optional] 
**_or** | [**[Expr]**](Expr.md) | returns true if any of the contained expressions are true | [optional] 
**_property** | **str** | property to match against, in dotted property notation | [optional] 
**re_match** | **str** | returns true if the property is a string and matches the specified regex | [optional] 
**schema** | **str** | matches only records with the specified schema | [optional] 
**time_range** | [**RstreamTimeRange**](RstreamTimeRange.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


