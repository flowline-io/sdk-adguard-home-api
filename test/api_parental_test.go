/*
AdGuard Home

Testing ParentalAPIService

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

func Test_openapi_ParentalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ParentalAPIService ParentalDisable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ParentalAPI.ParentalDisable(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ParentalAPIService ParentalEnable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ParentalAPI.ParentalEnable(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ParentalAPIService ParentalStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ParentalAPI.ParentalStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
