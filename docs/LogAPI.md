# \LogAPI

All URIs are relative to */control*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueryLogConfig**](LogAPI.md#GetQueryLogConfig) | **Get** /querylog/config | Get query log parameters
[**PutQueryLogConfig**](LogAPI.md#PutQueryLogConfig) | **Put** /querylog/config/update | Set query log parameters
[**QueryLog**](LogAPI.md#QueryLog) | **Get** /querylog | Get DNS server query log.
[**QueryLogConfig**](LogAPI.md#QueryLogConfig) | **Post** /querylog_config | Set query log parameters
[**QueryLogInfo**](LogAPI.md#QueryLogInfo) | **Get** /querylog_info | Get query log parameters
[**QuerylogClear**](LogAPI.md#QuerylogClear) | **Post** /querylog_clear | Clear query log



## GetQueryLogConfig

> GetQueryLogConfigResponse GetQueryLogConfig(ctx).Execute()

Get query log parameters

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
	resp, r, err := apiClient.LogAPI.GetQueryLogConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetQueryLogConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryLogConfig`: GetQueryLogConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetQueryLogConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryLogConfigRequest struct via the builder pattern


### Return type

[**GetQueryLogConfigResponse**](GetQueryLogConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutQueryLogConfig

> PutQueryLogConfig(ctx).Body(body).Execute()

Set query log parameters

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
	body := GetQueryLogConfigResponse(987) // GetQueryLogConfigResponse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogAPI.PutQueryLogConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.PutQueryLogConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutQueryLogConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **GetQueryLogConfigResponse** |  | 

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


## QueryLog

> QueryLog QueryLog(ctx).OlderThan(olderThan).Offset(offset).Limit(limit).Search(search).ResponseStatus(responseStatus).Execute()

Get DNS server query log.

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
	olderThan := "olderThan_example" // string | Filter by older than (optional)
	offset := int32(56) // int32 | Specify the ranking number of the first item on the page.  Even though it is possible to use \"offset\" and \"older_than\", we recommend choosing one of them and sticking to it.  (optional)
	limit := int32(56) // int32 | Limit the number of records to be returned (optional)
	search := "search_example" // string | Filter by domain name or client IP (optional)
	responseStatus := "responseStatus_example" // string | Filter by response status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.QueryLog(context.Background()).OlderThan(olderThan).Offset(offset).Limit(limit).Search(search).ResponseStatus(responseStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.QueryLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryLog`: QueryLog
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.QueryLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **olderThan** | **string** | Filter by older than | 
 **offset** | **int32** | Specify the ranking number of the first item on the page.  Even though it is possible to use \&quot;offset\&quot; and \&quot;older_than\&quot;, we recommend choosing one of them and sticking to it.  | 
 **limit** | **int32** | Limit the number of records to be returned | 
 **search** | **string** | Filter by domain name or client IP | 
 **responseStatus** | **string** | Filter by response status | 

### Return type

[**QueryLog**](QueryLog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryLogConfig

> QueryLogConfig(ctx).QueryLogConfig(queryLogConfig).Execute()

Set query log parameters



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
	queryLogConfig := *openapiclient.NewQueryLogConfig() // QueryLogConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogAPI.QueryLogConfig(context.Background()).QueryLogConfig(queryLogConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.QueryLogConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryLogConfig** | [**QueryLogConfig**](QueryLogConfig.md) |  | 

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


## QueryLogInfo

> QueryLogConfig QueryLogInfo(ctx).Execute()

Get query log parameters



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
	resp, r, err := apiClient.LogAPI.QueryLogInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.QueryLogInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryLogInfo`: QueryLogConfig
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.QueryLogInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogInfoRequest struct via the builder pattern


### Return type

[**QueryLogConfig**](QueryLogConfig.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerylogClear

> QuerylogClear(ctx).Execute()

Clear query log

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
	r, err := apiClient.LogAPI.QuerylogClear(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.QuerylogClear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQuerylogClearRequest struct via the builder pattern


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

