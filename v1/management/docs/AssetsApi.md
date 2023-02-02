# \AssetsApi

All URIs are relative to *https://cdn.contentstack.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getallassets**](AssetsApi.md#Getallassets) | **Get** /v3/assets | Get all assets
[**Getasingleasset**](AssetsApi.md#Getasingleasset) | **Get** /v3/assets/{asset_uid} | Get a single asset



## Getallassets

> Getallassets(ctx).IncludeFallback(includeFallback).RelativeUrls(relativeUrls).IncludeDimension(includeDimension).IncludeBranch(includeBranch).Execute()

Get all assets



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
    includeFallback := "true" // string | Enter 'true' to include the published asset details from the fallback locale when the specified locale does not contain published content. (optional)
    relativeUrls := "false" // string | Enter 'true' to include the relative URLs of the assets in the response. (optional)
    includeDimension := "true" // string | Enter 'true' to include the dimensions (height and width) of the image in the response. Supported image types: JPG, GIF, PNG, WebP, BMP, TIFF, SVG, and PSD. (optional)
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.Getallassets(context.Background()).IncludeFallback(includeFallback).RelativeUrls(relativeUrls).IncludeDimension(includeDimension).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.Getallassets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetallassetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeFallback** | **string** | Enter &#39;true&#39; to include the published asset details from the fallback locale when the specified locale does not contain published content. | 
 **relativeUrls** | **string** | Enter &#39;true&#39; to include the relative URLs of the assets in the response. | 
 **includeDimension** | **string** | Enter &#39;true&#39; to include the dimensions (height and width) of the image in the response. Supported image types: JPG, GIF, PNG, WebP, BMP, TIFF, SVG, and PSD. | 
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


## Getasingleasset

> Getasingleasset(ctx, assetUid).Version(version).IncludeFallback(includeFallback).IncludeDimension(includeDimension).RelativeUrls(relativeUrls).IncludeBranch(includeBranch).Execute()

Get a single asset



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
    assetUid := "blt19c34e5374418484" // string | 
    version := "1" // string | Specify the version number of the asset that you want to retrieve. To retrieve a specific version, keep the environment parameter blank.&nbsp;If the version is not specified, the details of the latest version will be retrieved. (optional)
    includeFallback := "true" // string | Enter 'true' to include published asset details from the fallback locale when the specified locale does not contain published information. (optional)
    includeDimension := "true" // string | Enter 'true' to include the dimensions (height and width)&nbsp;of the image&nbsp;in the response.&nbsp;Supported image&nbsp;types: JPG, GIF, PNG, WebP, BMP, TIFF, SVG, and PSD. (optional)
    relativeUrls := "false" // string | Enter 'true' to include the relative URL of the asset in the response. (optional)
    includeBranch := "false" // string | Set this to true to include the _branch top-level key in the response. This key states the unique ID of the branch where the concerned Contentstack module resides. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.Getasingleasset(context.Background(), assetUid).Version(version).IncludeFallback(includeFallback).IncludeDimension(includeDimension).RelativeUrls(relativeUrls).IncludeBranch(includeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.Getasingleasset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetasingleassetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Specify the version number of the asset that you want to retrieve. To retrieve a specific version, keep the environment parameter blank.&amp;nbsp;If the version is not specified, the details of the latest version will be retrieved. | 
 **includeFallback** | **string** | Enter &#39;true&#39; to include published asset details from the fallback locale when the specified locale does not contain published information. | 
 **includeDimension** | **string** | Enter &#39;true&#39; to include the dimensions (height and width)&amp;nbsp;of the image&amp;nbsp;in the response.&amp;nbsp;Supported image&amp;nbsp;types: JPG, GIF, PNG, WebP, BMP, TIFF, SVG, and PSD. | 
 **relativeUrls** | **string** | Enter &#39;true&#39; to include the relative URL of the asset in the response. | 
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

