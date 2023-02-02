# \EntriesApi

All URIs are relative to *https://cdn.contentstack.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getasingleentry**](EntriesApi.md#Getasingleentry) | **Get** /v3/content_types/{content_type_uid}/entries/{entry_uid} | Get a single entry



## Getasingleentry

> Getasingleentry(ctx, contentTypeUid, entryUid).Version(version).Locale(locale).IncludeFallback(includeFallback).IncludeBranch(includeBranch).Execute()

Get a single entry



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
    entryUid := "blta250054cfa4f5aab" // string | 
    version := "4" // string | Enter the latest published version number. Passing a draft version number will return a blank response. (optional)
    locale := "en-us" // string | Enter the code of the language of which you want to include&nbsp;the entries. Only the published localized entries will be displayed. (optional)
    includeFallback := "true" // string | Enter 'true' to include the published localized content from the fallback locale when the specified locale does not contain published content. (optional)
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.Getasingleentry(context.Background(), contentTypeUid, entryUid).Version(version).Locale(locale).IncludeFallback(includeFallback).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.Getasingleentry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeUid** | **string** |  | 
**entryUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetasingleentryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | Enter the latest published version number. Passing a draft version number will return a blank response. | 
 **locale** | **string** | Enter the code of the language of which you want to include&amp;nbsp;the entries. Only the published localized entries will be displayed. | 
 **includeFallback** | **string** | Enter &#39;true&#39; to include the published localized content from the fallback locale when the specified locale does not contain published content. | 
 **includeBranch** | **string** | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. | 

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

