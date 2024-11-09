# \SafebrowsingAPI

All URIs are relative to */control*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SafebrowsingDisable**](SafebrowsingAPI.md#SafebrowsingDisable) | **Post** /safebrowsing/disable | Disable safebrowsing
[**SafebrowsingEnable**](SafebrowsingAPI.md#SafebrowsingEnable) | **Post** /safebrowsing/enable | Enable safebrowsing
[**SafebrowsingStatus**](SafebrowsingAPI.md#SafebrowsingStatus) | **Get** /safebrowsing/status | Get safebrowsing status



## SafebrowsingDisable

> SafebrowsingDisable(ctx).Execute()

Disable safebrowsing

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
	r, err := apiClient.SafebrowsingAPI.SafebrowsingDisable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafebrowsingAPI.SafebrowsingDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSafebrowsingDisableRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafebrowsingEnable

> SafebrowsingEnable(ctx).Execute()

Enable safebrowsing

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
	r, err := apiClient.SafebrowsingAPI.SafebrowsingEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafebrowsingAPI.SafebrowsingEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSafebrowsingEnableRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafebrowsingStatus

> SafebrowsingStatus200Response SafebrowsingStatus(ctx).Execute()

Get safebrowsing status

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
	resp, r, err := apiClient.SafebrowsingAPI.SafebrowsingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafebrowsingAPI.SafebrowsingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SafebrowsingStatus`: SafebrowsingStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `SafebrowsingAPI.SafebrowsingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSafebrowsingStatusRequest struct via the builder pattern


### Return type

[**SafebrowsingStatus200Response**](SafebrowsingStatus200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

