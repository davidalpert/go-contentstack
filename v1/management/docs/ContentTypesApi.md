# \ContentTypesApi

All URIs are relative to *https://cdn.contentstack.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getallcontenttypes**](ContentTypesApi.md#Getallcontenttypes) | **Get** /v3/content_types | Get all content types
[**Getasinglecontenttype**](ContentTypesApi.md#Getasinglecontenttype) | **Get** /v3/content_types/{content_type_uid} | Get a single content type



## Getallcontenttypes

> Getallcontenttypes(ctx).IncludeCount(includeCount).IncludeBranch(includeBranch).Execute()

Get all content types



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
    includeCount := "false" // string | Set this to 'true' to include in response the total count of content types available in your stack.
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentTypesApi.Getallcontenttypes(context.Background()).IncludeCount(includeCount).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentTypesApi.Getallcontenttypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetallcontenttypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCount** | **string** | Set this to &#39;true&#39; to include in response the total count of content types available in your stack. | 
 **includeBranch** | **string** | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token), [api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getasinglecontenttype

> Getasinglecontenttype(ctx, contentTypeUid).IncludeBranch(includeBranch).Execute()

Get a single content type



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
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentTypesApi.Getasinglecontenttype(context.Background(), contentTypeUid).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentTypesApi.Getasinglecontenttype``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetasinglecontenttypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeBranch** | **string** | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token), [api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

