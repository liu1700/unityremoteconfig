/*
Remote Config Admin API

Testing EnvironmentsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package unityremoteconfig

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/liu1700/unityremoteconfig"
)

func Test_unityremoteconfig_EnvironmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentsAPIService RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsAPIService RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentId string
		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet(context.Background(), environmentId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsAPIService RemoteConfigV1ProjectsProjectIdEnvironmentsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsGet(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}