# \BlockedServicesAPI

All URIs are relative to */control*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockedServicesAll**](BlockedServicesAPI.md#BlockedServicesAll) | **Get** /blocked_services/all | Get available services to use for blocking
[**BlockedServicesAvailableServices**](BlockedServicesAPI.md#BlockedServicesAvailableServices) | **Get** /blocked_services/services | Get available services to use for blocking
[**BlockedServicesList**](BlockedServicesAPI.md#BlockedServicesList) | **Get** /blocked_services/list | Get blocked services list
[**BlockedServicesSchedule**](BlockedServicesAPI.md#BlockedServicesSchedule) | **Get** /blocked_services/get | Get blocked services
[**BlockedServicesScheduleUpdate**](BlockedServicesAPI.md#BlockedServicesScheduleUpdate) | **Put** /blocked_services/update | Update blocked services
[**BlockedServicesSet**](BlockedServicesAPI.md#BlockedServicesSet) | **Post** /blocked_services/set | Set blocked services list



## BlockedServicesAll

> BlockedServicesAll BlockedServicesAll(ctx).Execute()

Get available services to use for blocking

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
	resp, r, err := apiClient.BlockedServicesAPI.BlockedServicesAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockedServicesAPI.BlockedServicesAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockedServicesAll`: BlockedServicesAll
	fmt.Fprintf(os.Stdout, "Response from `BlockedServicesAPI.BlockedServicesAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBlockedServicesAllRequest struct via the builder pattern


### Return type

[**BlockedServicesAll**](BlockedServicesAll.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockedServicesAvailableServices

> []string BlockedServicesAvailableServices(ctx).Execute()

Get available services to use for blocking



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
	resp, r, err := apiClient.BlockedServicesAPI.BlockedServicesAvailableServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockedServicesAPI.BlockedServicesAvailableServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockedServicesAvailableServices`: []string
	fmt.Fprintf(os.Stdout, "Response from `BlockedServicesAPI.BlockedServicesAvailableServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBlockedServicesAvailableServicesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockedServicesList

> []string BlockedServicesList(ctx).Execute()

Get blocked services list



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
	resp, r, err := apiClient.BlockedServicesAPI.BlockedServicesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockedServicesAPI.BlockedServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockedServicesList`: []string
	fmt.Fprintf(os.Stdout, "Response from `BlockedServicesAPI.BlockedServicesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBlockedServicesListRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockedServicesSchedule

> BlockedServicesSchedule BlockedServicesSchedule(ctx).Execute()

Get blocked services

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
	resp, r, err := apiClient.BlockedServicesAPI.BlockedServicesSchedule(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockedServicesAPI.BlockedServicesSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockedServicesSchedule`: BlockedServicesSchedule
	fmt.Fprintf(os.Stdout, "Response from `BlockedServicesAPI.BlockedServicesSchedule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBlockedServicesScheduleRequest struct via the builder pattern


### Return type

[**BlockedServicesSchedule**](BlockedServicesSchedule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockedServicesScheduleUpdate

> BlockedServicesScheduleUpdate(ctx).BlockedServicesSchedule(blockedServicesSchedule).Execute()

Update blocked services

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
	blockedServicesSchedule := *openapiclient.NewBlockedServicesSchedule() // BlockedServicesSchedule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlockedServicesAPI.BlockedServicesScheduleUpdate(context.Background()).BlockedServicesSchedule(blockedServicesSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockedServicesAPI.BlockedServicesScheduleUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockedServicesScheduleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockedServicesSchedule** | [**BlockedServicesSchedule**](BlockedServicesSchedule.md) |  | 

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


## BlockedServicesSet

> BlockedServicesSet(ctx).RequestBody(requestBody).Execute()

Set blocked services list



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
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlockedServicesAPI.BlockedServicesSet(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockedServicesAPI.BlockedServicesSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockedServicesSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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

