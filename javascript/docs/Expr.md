# SpyderbatApi.Expr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**and** | [**[Expr]**](Expr.md) | returns true if all of the contained expressions are true | [optional] 
**containsStr** | **String** | returns true if the property is a string and contains the specified value | [optional] 
**equals** | **Object** | returns true if the property matches the supplied value | [optional] 
**exists** | **Boolean** | returns true if the property exists | [optional] 
**greaterThan** | **Object** | returns true if the property is a number and is greater than this value | [optional] 
**hasPrefix** | **String** | returns true if the property is a string and has the specified prefix | [optional] 
**hasSuffix** | **String** | returns true if the property is a string and has the specified suffix | [optional] 
**_in** | **[Object]** | returns true if the property matches any of the values specified | [optional] 
**lessThan** | **Object** | returns true if te property is a number and is less than this value | [optional] 
**not** | [**Expr**](Expr.md) |  | [optional] 
**or** | [**[Expr]**](Expr.md) | returns true if any of the contained expressions are true | [optional] 
**property** | **String** | property to match against, in dotted property notation | [optional] 
**reMatch** | **String** | returns true if the property is a string and matches the specified regex | [optional] 
**schema** | **String** | matches only records with the specified schema | [optional] 
**timeRange** | [**RstreamTimeRange**](RstreamTimeRange.md) |  | [optional] 


