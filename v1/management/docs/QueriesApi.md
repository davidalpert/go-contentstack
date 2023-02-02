# \QueriesApi

All URIs are relative to *https://cdn.contentstack.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperator**](QueriesApi.md#GetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperator) | **Get** /v3/content_types/{content_type_uid}/entries | Get all entries | Equals Operator | Equals Operator Within Group | Equals Operator Within Modular Blocks | Not-equals Operator.....



## GetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperator

> GetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperator(ctx, contentTypeUid).Locale(locale).IncludeFallback(includeFallback).IncludeBranch(includeBranch).Query(query).Include(include).Limit(limit).Skip(skip).Asc(asc).Desc(desc).OnlyBASE(onlyBASE).ExceptBASE(exceptBASE).IncludeCount(includeCount).Execute()

Get all entries | Equals Operator | Equals Operator Within Group | Equals Operator Within Modular Blocks | Not-equals Operator.....



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contentTypeUid := "product" // string | 
    locale := "en-us" // string | Enter the code of the language of which the entries needs to be included. Only the entries published in this locale will be displayed. (optional)
    includeFallback := "true" // string | Enter 'true' to include the published localized content from the fallback locale when the specified locale does not contain published content. (optional)
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)
    query := "{"title": "Redmi 3S"}" // string | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. (optional)
    include := "categories" // string | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. (optional)
    limit := "2" // string | Enter the maximum number of entries to be returned. (optional)
    skip := "2" // string | Enter the number of entries to be skipped. (optional)
    asc := "price_in_usd" // string | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. (optional)
    desc := "price_in_usd" // string | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. (optional)
    onlyBASE := "price_in_usd" // string | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. (optional)
    exceptBASE := "price_in_usd" // string | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. (optional)
    includeCount := "true" // string | Enter 'true'&nbsp;to get the total count of the entries. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.GetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperator(context.Background(), contentTypeUid).Locale(locale).IncludeFallback(includeFallback).IncludeBranch(includeBranch).Query(query).Include(include).Limit(limit).Skip(skip).Asc(asc).Desc(desc).OnlyBASE(onlyBASE).ExceptBASE(exceptBASE).IncludeCount(includeCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.GetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetallentriesEqualsoperatorEqualsoperatorwithingroupEqualsoperatorwithinmodularblocksNotEqualsoperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | Enter the code of the language of which the entries needs to be included. Only the entries published in this locale will be displayed. | 
 **includeFallback** | **string** | Enter &#39;true&#39; to include the published localized content from the fallback locale when the specified locale does not contain published content. | 
 **includeBranch** | **string** | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. | 
 **query** | **string** | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. | 
 **include** | **string** | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. | 
 **limit** | **string** | Enter the maximum number of entries to be returned. | 
 **skip** | **string** | Enter the number of entries to be skipped. | 
 **asc** | **string** | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. | 
 **desc** | **string** | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. | 
 **onlyBASE** | **string** | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. | 
 **exceptBASE** | **string** | Enter the actual query that will be executed to retrieve entries. This query should be in JSON format. | 
 **includeCount** | **string** | Enter &#39;true&#39;&amp;nbsp;to get the total count of the entries. | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token), [api_key](../README.md#api_key), [environment](../README.md#environment)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

