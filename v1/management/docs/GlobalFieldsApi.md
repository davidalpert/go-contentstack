# \GlobalFieldsApi

All URIs are relative to *https://cdn.contentstack.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getallglobalfields**](GlobalFieldsApi.md#Getallglobalfields) | **Get** /v3/global_fields | Get all global fields
[**Getasingleglobalfield**](GlobalFieldsApi.md#Getasingleglobalfield) | **Get** /v3/global_fields/{global_field_uid} | Get a single global field



## Getallglobalfields

> Getallglobalfields(ctx).IncludeBranch(includeBranch).Execute()

Get all global fields



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
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalFieldsApi.Getallglobalfields(context.Background()).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalFieldsApi.Getallglobalfields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetallglobalfieldsRequest struct via the builder pattern


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


## Getasingleglobalfield

> Getasingleglobalfield(ctx, globalFieldUid).IncludeBranch(includeBranch).Execute()

Get a single global field



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
    globalFieldUid := "seo" // string | 
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalFieldsApi.Getasingleglobalfield(context.Background(), globalFieldUid).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalFieldsApi.Getasingleglobalfield``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalFieldUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetasingleglobalfieldRequest struct via the builder pattern


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

