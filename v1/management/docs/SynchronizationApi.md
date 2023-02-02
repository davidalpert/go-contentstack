# \SynchronizationApi

All URIs are relative to *https://cdn.contentstack.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitialsyncSyncusingpaginationtokenSubsequentsync**](SynchronizationApi.md#InitialsyncSyncusingpaginationtokenSubsequentsync) | **Get** /v3/stacks/sync | Initial Sync | Sync using pagination token | Subsequent Sync



## InitialsyncSyncusingpaginationtokenSubsequentsync

> InitialsyncSyncusingpaginationtokenSubsequentsync(ctx).Init(init).ContentTypeUid(contentTypeUid).Locale(locale).StartFrom(startFrom).Type_(type_).PaginationToken(paginationToken).SyncToken(syncToken).Execute()

Initial Sync | Sync using pagination token | Subsequent Sync



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
    init := "true" // string | Enter ‘true’ to perform a complete sync of all your app data. (optional)
    contentTypeUid := "product" // string | Enter the content type UID, if you want to sync entries of specific content types. (optional)
    locale := "en-us" // string | Enter the locale to retrieve and sync the content published on a specific locale. (optional)
    startFrom := "2018-01-14T00:00:00.000Z" // string | Specify the start date, if you want to retrieve and sync data starting from a specific date. (optional)
    type_ := "entry_published, entry_unpublished, asset_published" // string | Enter the type(s) of content you want to retrieve and sync. You can pass multiple types as comma-separated values. (optional)
    paginationToken := "bltd8d8b3d427917796048622" // string | Enter the pagination token that you received in the response body of the previous sync process. (optional)
    syncToken := "blt42d8feee86ff1eb6d0de5c" // string | Enter the sync token that you received in the response body of the previous completed Synchronization process to get the delta updates (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SynchronizationApi.InitialsyncSyncusingpaginationtokenSubsequentsync(context.Background()).Init(init).ContentTypeUid(contentTypeUid).Locale(locale).StartFrom(startFrom).Type_(type_).PaginationToken(paginationToken).SyncToken(syncToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationApi.InitialsyncSyncusingpaginationtokenSubsequentsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitialsyncSyncusingpaginationtokenSubsequentsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **init** | **string** | Enter ‘true’ to perform a complete sync of all your app data. | 
 **contentTypeUid** | **string** | Enter the content type UID, if you want to sync entries of specific content types. | 
 **locale** | **string** | Enter the locale to retrieve and sync the content published on a specific locale. | 
 **startFrom** | **string** | Specify the start date, if you want to retrieve and sync data starting from a specific date. | 
 **type_** | **string** | Enter the type(s) of content you want to retrieve and sync. You can pass multiple types as comma-separated values. | 
 **paginationToken** | **string** | Enter the pagination token that you received in the response body of the previous sync process. | 
 **syncToken** | **string** | Enter the sync token that you received in the response body of the previous completed Synchronization process to get the delta updates | 

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

