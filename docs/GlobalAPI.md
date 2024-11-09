# \GlobalAPI

All URIs are relative to */control*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BeginUpdate**](GlobalAPI.md#BeginUpdate) | **Post** /update | Begin auto-upgrade procedure
[**CacheClear**](GlobalAPI.md#CacheClear) | **Post** /cache_clear | Clear DNS cache
[**DnsConfig**](GlobalAPI.md#DnsConfig) | **Post** /dns_config | Set general DNS parameters
[**DnsInfo**](GlobalAPI.md#DnsInfo) | **Get** /dns_info | Get general DNS parameters
[**GetProfile**](GlobalAPI.md#GetProfile) | **Get** /profile | 
[**GetVersionJson**](GlobalAPI.md#GetVersionJson) | **Post** /version.json | Gets information about the latest available version of AdGuard 
[**Login**](GlobalAPI.md#Login) | **Post** /login | Perform administrator log-in
[**Logout**](GlobalAPI.md#Logout) | **Get** /logout | Perform administrator log-out
[**MobileConfigDoH**](GlobalAPI.md#MobileConfigDoH) | **Get** /apple/doh.mobileconfig | Get DNS over HTTPS .mobileconfig.
[**MobileConfigDoT**](GlobalAPI.md#MobileConfigDoT) | **Get** /apple/dot.mobileconfig | Get DNS over TLS .mobileconfig.
[**SetProtection**](GlobalAPI.md#SetProtection) | **Post** /protection | Set protection state and duration
[**Status**](GlobalAPI.md#Status) | **Get** /status | Get DNS server current status and general settings
[**TestUpstreamDNS**](GlobalAPI.md#TestUpstreamDNS) | **Post** /test_upstream_dns | Test upstream configuration
[**UpdateProfile**](GlobalAPI.md#UpdateProfile) | **Put** /profile/update | Updates current user info



## BeginUpdate

> BeginUpdate(ctx).Execute()

Begin auto-upgrade procedure

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
	r, err := apiClient.GlobalAPI.BeginUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.BeginUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBeginUpdateRequest struct via the builder pattern


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


## CacheClear

> CacheClear(ctx).Execute()

Clear DNS cache

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
	r, err := apiClient.GlobalAPI.CacheClear(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.CacheClear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCacheClearRequest struct via the builder pattern


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


## DnsConfig

> DnsConfig(ctx).DNSConfig(dNSConfig).Execute()

Set general DNS parameters

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
	dNSConfig := *openapiclient.NewDNSConfig() // DNSConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalAPI.DnsConfig(context.Background()).DNSConfig(dNSConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.DnsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dNSConfig** | [**DNSConfig**](DNSConfig.md) |  | 

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


## DnsInfo

> DnsInfo200Response DnsInfo(ctx).Execute()

Get general DNS parameters

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
	resp, r, err := apiClient.GlobalAPI.DnsInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.DnsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsInfo`: DnsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.DnsInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDnsInfoRequest struct via the builder pattern


### Return type

[**DnsInfo200Response**](DnsInfo200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> ProfileInfo GetProfile(ctx).Execute()



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
	resp, r, err := apiClient.GlobalAPI.GetProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: ProfileInfo
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


### Return type

[**ProfileInfo**](ProfileInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionJson

> VersionInfo GetVersionJson(ctx).GetVersionRequest(getVersionRequest).Execute()

Gets information about the latest available version of AdGuard 

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
	getVersionRequest := *openapiclient.NewGetVersionRequest() // GetVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAPI.GetVersionJson(context.Background()).GetVersionRequest(getVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GetVersionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionJson`: VersionInfo
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GetVersionJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getVersionRequest** | [**GetVersionRequest**](GetVersionRequest.md) |  | 

### Return type

[**VersionInfo**](VersionInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> Login(ctx).Login(login).Execute()

Perform administrator log-in

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
	login := *openapiclient.NewLogin() // Login | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalAPI.Login(context.Background()).Login(login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | [**Login**](Login.md) |  | 

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


## Logout

> Logout(ctx).Execute()

Perform administrator log-out

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
	r, err := apiClient.GlobalAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## MobileConfigDoH

> MobileConfigDoH(ctx).Host(host).ClientId(clientId).Execute()

Get DNS over HTTPS .mobileconfig.

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
	host := "example.org" // string | Host for which the config is generated.  If no host is provided, `tls.server_name` from the configuration file is used.  If `tls.server_name` is not set, the API returns an error with a 500 status. 
	clientId := "client-1" // string | ClientID.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalAPI.MobileConfigDoH(context.Background()).Host(host).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.MobileConfigDoH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMobileConfigDoHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | Host for which the config is generated.  If no host is provided, &#x60;tls.server_name&#x60; from the configuration file is used.  If &#x60;tls.server_name&#x60; is not set, the API returns an error with a 500 status.  | 
 **clientId** | **string** | ClientID.  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobileConfigDoT

> MobileConfigDoT(ctx).Host(host).ClientId(clientId).Execute()

Get DNS over TLS .mobileconfig.

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
	host := "example.org" // string | Host for which the config is generated.  If no host is provided, `tls.server_name` from the configuration file is used.  If `tls.server_name` is not set, the API returns an error with a 500 status. 
	clientId := "client-1" // string | ClientID.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalAPI.MobileConfigDoT(context.Background()).Host(host).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.MobileConfigDoT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMobileConfigDoTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | Host for which the config is generated.  If no host is provided, &#x60;tls.server_name&#x60; from the configuration file is used.  If &#x60;tls.server_name&#x60; is not set, the API returns an error with a 500 status.  | 
 **clientId** | **string** | ClientID.  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProtection

> SetProtection(ctx).SetProtectionRequest(setProtectionRequest).Execute()

Set protection state and duration

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
	setProtectionRequest := *openapiclient.NewSetProtectionRequest(false) // SetProtectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalAPI.SetProtection(context.Background()).SetProtectionRequest(setProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.SetProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setProtectionRequest** | [**SetProtectionRequest**](SetProtectionRequest.md) |  | 

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


## Status

> ServerStatus Status(ctx).Execute()

Get DNS server current status and general settings

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
	resp, r, err := apiClient.GlobalAPI.Status(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.Status``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Status`: ServerStatus
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.Status`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


### Return type

[**ServerStatus**](ServerStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUpstreamDNS

> map[string]string TestUpstreamDNS(ctx).UpstreamsConfig(upstreamsConfig).Execute()

Test upstream configuration

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
	upstreamsConfig := *openapiclient.NewUpstreamsConfig([]string{"BootstrapDns_example"}, []string{"UpstreamDns_example"}) // UpstreamsConfig | Upstream configuration to be tested (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAPI.TestUpstreamDNS(context.Background()).UpstreamsConfig(upstreamsConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.TestUpstreamDNS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestUpstreamDNS`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.TestUpstreamDNS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestUpstreamDNSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upstreamsConfig** | [**UpstreamsConfig**](UpstreamsConfig.md) | Upstream configuration to be tested | 

### Return type

**map[string]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile(ctx).ProfileInfo(profileInfo).Execute()

Updates current user info

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
	profileInfo := *openapiclient.NewProfileInfo("Name_example", "Language_example", "Theme_example") // ProfileInfo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalAPI.UpdateProfile(context.Background()).ProfileInfo(profileInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileInfo** | [**ProfileInfo**](ProfileInfo.md) |  | 

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

