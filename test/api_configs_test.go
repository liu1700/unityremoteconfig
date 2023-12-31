/*
Remote Config Admin API

Testing ConfigsAPIService

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

func Test_unityremoteconfig_ConfigsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configId string
		var projectId string

		httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete(context.Background(), configId, projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configId string
		var projectId string

		resp, httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet(context.Background(), configId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configId string
		var projectId string

		httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut(context.Background(), configId, projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdConfigsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsGet(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdConfigsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsPost(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdConfigscopytoPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigscopytoPost(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigsAPIService RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentId string
		var projectId string

		resp, httpRes, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet(context.Background(), environmentId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
