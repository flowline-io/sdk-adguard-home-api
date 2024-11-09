# \SafesearchAPI

All URIs are relative to */control*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SafesearchDisable**](SafesearchAPI.md#SafesearchDisable) | **Post** /safesearch/disable | Disable safesearch
[**SafesearchEnable**](SafesearchAPI.md#SafesearchEnable) | **Post** /safesearch/enable | Enable safesearch
[**SafesearchSettings**](SafesearchAPI.md#SafesearchSettings) | **Put** /safesearch/settings | Update safesearch settings
[**SafesearchStatus**](SafesearchAPI.md#SafesearchStatus) | **Get** /safesearch/status | Get safesearch status



## SafesearchDisable

> SafesearchDisable(ctx).Execute()

Disable safesearch

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
	r, err := apiClient.SafesearchAPI.SafesearchDisable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafesearchAPI.SafesearchDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSafesearchDisableRequest struct via the builder pattern


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


## SafesearchEnable

> SafesearchEnable(ctx).Execute()

Enable safesearch

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
	r, err := apiClient.SafesearchAPI.SafesearchEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafesearchAPI.SafesearchEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSafesearchEnableRequest struct via the builder pattern


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


## SafesearchSettings

> SafesearchSettings(ctx).SafeSearchConfig(safeSearchConfig).Execute()

Update safesearch settings

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
	safeSearchConfig := *openapiclient.NewSafeSearchConfig() // SafeSearchConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SafesearchAPI.SafesearchSettings(context.Background()).SafeSearchConfig(safeSearchConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafesearchAPI.SafesearchSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSafesearchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safeSearchConfig** | [**SafeSearchConfig**](SafeSearchConfig.md) |  | 

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


## SafesearchStatus

> SafeSearchConfig SafesearchStatus(ctx).Execute()

Get safesearch status

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
	resp, r, err := apiClient.SafesearchAPI.SafesearchStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SafesearchAPI.SafesearchStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SafesearchStatus`: SafeSearchConfig
	fmt.Fprintf(os.Stdout, "Response from `SafesearchAPI.SafesearchStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSafesearchStatusRequest struct via the builder pattern


### Return type

[**SafeSearchConfig**](SafeSearchConfig.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

