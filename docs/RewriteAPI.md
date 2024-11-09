# \RewriteAPI

All URIs are relative to */control*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RewriteAdd**](RewriteAPI.md#RewriteAdd) | **Post** /rewrite/add | Add a new Rewrite rule
[**RewriteDelete**](RewriteAPI.md#RewriteDelete) | **Post** /rewrite/delete | Remove a Rewrite rule
[**RewriteList**](RewriteAPI.md#RewriteList) | **Get** /rewrite/list | Get list of Rewrite rules
[**RewriteUpdate**](RewriteAPI.md#RewriteUpdate) | **Put** /rewrite/update | Update a Rewrite rule



## RewriteAdd

> RewriteAdd(ctx).RewriteEntry(rewriteEntry).Execute()

Add a new Rewrite rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	rewriteEntry := *openapiclient.NewRewriteEntry() // RewriteEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RewriteAPI.RewriteAdd(context.Background()).RewriteEntry(rewriteEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewriteAPI.RewriteAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRewriteAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewriteEntry** | [**RewriteEntry**](RewriteEntry.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriteDelete

> RewriteDelete(ctx).RewriteEntry(rewriteEntry).Execute()

Remove a Rewrite rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	rewriteEntry := *openapiclient.NewRewriteEntry() // RewriteEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RewriteAPI.RewriteDelete(context.Background()).RewriteEntry(rewriteEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewriteAPI.RewriteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRewriteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewriteEntry** | [**RewriteEntry**](RewriteEntry.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriteList

> []RewriteEntry RewriteList(ctx).Execute()

Get list of Rewrite rules

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewriteAPI.RewriteList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewriteAPI.RewriteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RewriteList`: []RewriteEntry
	fmt.Fprintf(os.Stdout, "Response from `RewriteAPI.RewriteList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRewriteListRequest struct via the builder pattern


### Return type

[**[]RewriteEntry**](RewriteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriteUpdate

> RewriteUpdate(ctx).RewriteUpdate(rewriteUpdate).Execute()

Update a Rewrite rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	rewriteUpdate := *openapiclient.NewRewriteUpdate() // RewriteUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RewriteAPI.RewriteUpdate(context.Background()).RewriteUpdate(rewriteUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewriteAPI.RewriteUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRewriteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewriteUpdate** | [**RewriteUpdate**](RewriteUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

