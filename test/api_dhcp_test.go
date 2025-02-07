/*
AdGuard Home

Testing DhcpAPIService

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

func Test_openapi_DhcpAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DhcpAPIService CheckActiveDhcp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DhcpAPI.CheckActiveDhcp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpAddStaticLease", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DhcpAPI.DhcpAddStaticLease(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpInterfaces", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DhcpAPI.DhcpInterfaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpRemoveStaticLease", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DhcpAPI.DhcpRemoveStaticLease(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpReset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DhcpAPI.DhcpReset(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpResetLeases", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DhcpAPI.DhcpResetLeases(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpSetConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DhcpAPI.DhcpSetConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DhcpAPI.DhcpStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpAPIService DhcpUpdateStaticLease", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DhcpAPI.DhcpUpdateStaticLease(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
