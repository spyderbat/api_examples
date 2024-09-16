# Expr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**and** | Option<[**Vec<crate::models::Expr>**](Expr.md)> | returns true if all of the contained expressions are true | [optional]
**contains_str** | Option<**String**> | returns true if the property is a string and contains the specified value | [optional]
**equals** | Option<[**serde_json::Value**](.md)> | returns true if the property matches the supplied value | [optional]
**exists** | Option<**bool**> | returns true if the property exists | [optional]
**greater_than** | Option<[**serde_json::Value**](.md)> | returns true if the property is a number and is greater than this value | [optional]
**has_prefix** | Option<**String**> | returns true if the property is a string and has the specified prefix | [optional]
**has_suffix** | Option<**String**> | returns true if the property is a string and has the specified suffix | [optional]
**_in** | Option<[**Vec<serde_json::Value>**](serde_json::Value.md)> | returns true if the property matches any of the values specified | [optional]
**less_than** | Option<[**serde_json::Value**](.md)> | returns true if te property is a number and is less than this value | [optional]
**not** | Option<[**crate::models::Expr**](Expr.md)> |  | [optional]
**or** | Option<[**Vec<crate::models::Expr>**](Expr.md)> | returns true if any of the contained expressions are true | [optional]
**property** | Option<**String**> | property to match against, in dotted property notation | [optional]
**re_match** | Option<**String**> | returns true if the property is a string and matches the specified regex | [optional]
**schema** | Option<**String**> | matches only records with the specified schema | [optional]
**time_range** | Option<[**crate::models::RstreamTimeRange**](RstreamTimeRange.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


