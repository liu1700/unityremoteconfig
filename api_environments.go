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


// EnvironmentsAPIService EnvironmentsAPI service
type EnvironmentsAPIService service

type ApiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest struct {
	ctx context.Context
	ApiService *EnvironmentsAPIService
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest) Execute() (*DefaultEnvironmentResponse, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet Get Default Environment

Obtain the id of the default environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest
*/
func (a *EnvironmentsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet(ctx context.Context, projectId string) ApiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest {
	return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return DefaultEnvironmentResponse
func (a *EnvironmentsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetExecute(r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest) (*DefaultEnvironmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DefaultEnvironmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/environments/default"
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

type ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest struct {
	ctx context.Context
	ApiService *EnvironmentsAPIService
	environmentId string
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest) Execute() (*Environment, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet Get Environment

Return environment for a given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment Id
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest
*/
func (a *EnvironmentsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet(ctx context.Context, environmentId string, projectId string) ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest {
	return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Environment
func (a *EnvironmentsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetExecute(r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest) (*Environment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Environment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/environments/{environmentId}"
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

type ApiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest struct {
	ctx context.Context
	ApiService *EnvironmentsAPIService
	projectId string
}

func (r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest) Execute() (*GetEnvironmentsResponse, *http.Response, error) {
	return r.ApiService.RemoteConfigV1ProjectsProjectIdEnvironmentsGetExecute(r)
}

/*
RemoteConfigV1ProjectsProjectIdEnvironmentsGet Get All Environments

Returns the list of environments for a Project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Genesis projectId
 @return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest
*/
func (a *EnvironmentsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsGet(ctx context.Context, projectId string) ApiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest {
	return ApiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return GetEnvironmentsResponse
func (a *EnvironmentsAPIService) RemoteConfigV1ProjectsProjectIdEnvironmentsGetExecute(r ApiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest) (*GetEnvironmentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetEnvironmentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.RemoteConfigV1ProjectsProjectIdEnvironmentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/remote-config/v1/projects/{projectId}/environments"
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