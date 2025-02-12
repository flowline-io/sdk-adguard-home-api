/*
AdGuard Home

Testing GlobalAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/flowline-io/sdk-adguard-home-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_GlobalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GlobalAPIService BeginUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.BeginUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService CacheClear", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.CacheClear(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService DnsConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.DnsConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService DnsInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalAPI.DnsInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService GetProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalAPI.GetProfile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService GetVersionJson", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalAPI.GetVersionJson(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService Login", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.Login(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService Logout", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.Logout(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService MobileConfigDoH", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.MobileConfigDoH(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService MobileConfigDoT", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.MobileConfigDoT(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService SetProtection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.SetProtection(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService Status", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalAPI.Status(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService TestUpstreamDNS", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalAPI.TestUpstreamDNS(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAPIService UpdateProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GlobalAPI.UpdateProfile(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
