/*
Remote Config Admin API

# Overview Unity Remote Config is a cloud service that allows you to tune your game design without deploying new versions of your application. With Remote Config, you can:  Adapt your game to different types of players. Tune your game difficulty curve in near real time. Alter graphic quality based on device to optimize performance. Roll out new features gradually while monitoring impact. Tailor game settings to different regions or other player segments. Run campaign tests comparing colors, styles, prices, etc. Turn seasonal, holiday, or other time-sensitive events on or off. Enable or disable features for specific player segments or across the entire user base.  Define Game Overrides that control which players receive what settings updates, and when. Unity manages the delivery and assignment of those settings with minimal impact to performance. No update to your application is necessary. When a player launches your game, Remote Config detects contextual attributes used as game override conditions, based on Unity, the application, the user, or custom criteria that you define. The service then returns customized settings for each player according to the game overrides that apply to them. This allows different players using the same version of your game to have slightly different experiences. It also allows you to understand the impact each experience has on your business.  # Limitation * A configuration can have a max payload size of 5MB * A string limit is only of 65 535 characters 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityremoteconfig

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ConfigsAPIService ConfigsAPI service
type ConfigsAPIService service

type ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	configId string
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete Delete Config

Deletes an existing Config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId Configuration Id
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete(ctx context.Context, configId string, projectId string) ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest {
	return ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		projectId: projectId,
	}
}

// Execute executes the request
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteExecute(r ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/configs/{configId}"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	configId string
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest) Execute() (*Config, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdConfigsConfigIdGetExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet Get Config

Returns Config for a given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId Configuration Id
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet(ctx context.Context, configId string, projectId string) ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest {
	return ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Config
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsConfigIdGetExecute(r ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest) (*Config, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Config
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/configs/{configId}"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	configId string
	projectId string
}

// Updated Config values
func (r ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdConfigsConfigIdPutExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut Update Config

Updates the value of an existing Config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId Configuration Id
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut(ctx context.Context, configId string, projectId string) ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest {
	return ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		projectId: projectId,
	}
}

// Execute executes the request
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsConfigIdPutExecute(r ApiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/configs/{configId}"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoteConfigV1ProjectsProjectIdConfigsGetRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdConfigsGetRequest) Execute() (*Config, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdConfigsGetExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdConfigsGet Get config with the default environment

Obtain the configs with the default environment 'production' or specified environment id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdConfigsGetRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsGet(ctx context.Context, projectId string) ApiRemoteConfigV1ProjectsProjectIdConfigsGetRequest {
	return ApiRemoteConfigV1ProjectsProjectIdConfigsGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Config
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsGetExecute(r ApiRemoteConfigV1ProjectsProjectIdConfigsGetRequest) (*Config, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Config
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdConfigsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoteConfigV1ProjectsProjectIdConfigsPostRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	projectId string
}

// Config object
func (r ApiRemoteConfigV1ProjectsProjectIdConfigsPostRequest) Execute() (*CreateConfigResponse, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdConfigsPostExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdConfigsPost Create Config

Create Config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdConfigsPostRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsPost(ctx context.Context, projectId string) ApiRemoteConfigV1ProjectsProjectIdConfigsPostRequest {
	return ApiRemoteConfigV1ProjectsProjectIdConfigsPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return CreateConfigResponse
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigsPostExecute(r ApiRemoteConfigV1ProjectsProjectIdConfigsPostRequest) (*CreateConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdConfigsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	projectId string
	configCopyRequest *ConfigCopyRequest
}

// IDs for source and destination Config, and Settings keys to be copied
func (r ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest) ConfigCopyRequest(configCopyRequest ConfigCopyRequest) ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest {
	r.configCopyRequest = &configCopyRequest
	return r
}

func (r ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdConfigscopytoPostExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdConfigscopytoPost Copy Config Settings

Copies specified settings from a source Config to a destination Config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigscopytoPost(ctx context.Context, projectId string) ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest {
	return ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdConfigscopytoPostExecute(r ApiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdConfigscopytoPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/configs:copyto"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.configCopyRequest == nil {
		return nil, reportError("configCopyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.configCopyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RemoteConfigV1ProjectsProjectIdConfigsCopytoPost404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v RemoteConfigV1ProjectsProjectIdConfigsCopytoPost409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	environmentId string
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest) Execute() (*RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet Get Configs by Environment ID

Returns the list of all configs for a Environment. Currently, the only supported config type is "settings"

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment Id
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest
*/
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet(ctx context.Context, environmentId string, projectId string) ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest {
	return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response
func (a *ConfigsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetExecute(r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest) (*RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/environments/{environmentId}/configs"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
