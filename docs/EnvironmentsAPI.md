# \EnvironmentsAPI

All URIs are relative to *https://services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet**](EnvironmentsAPI.md#RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet) | **Get** /remote-config/v1/projects/{projectId}/environments/default | Get Default Environment
[**RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet**](EnvironmentsAPI.md#RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet) | **Get** /remote-config/v1/projects/{projectId}/environments/{environmentId} | Get Environment
[**RemoteConfigV1ProjectsProjectIdEnvironmentsGet**](EnvironmentsAPI.md#RemoteConfigV1ProjectsProjectIdEnvironmentsGet) | **Get** /remote-config/v1/projects/{projectId}/environments | Get All Environments



## RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet

> DefaultEnvironmentResponse RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet(ctx, projectId).Execute()

Get Default Environment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityremoteconfig"
)

func main() {
    projectId := "projectId_example" // string | Genesis projectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet`: DefaultEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefaultEnvironmentResponse**](DefaultEnvironmentResponse.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet

> Environment RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet(ctx, environmentId, projectId).Execute()

Get Environment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityremoteconfig"
)

func main() {
    environmentId := "environmentId_example" // string | Environment Id
    projectId := "projectId_example" // string | Genesis projectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet(context.Background(), environmentId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment Id | 
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Environment**](Environment.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdEnvironmentsGet

> GetEnvironmentsResponse RemoteConfigV1ProjectsProjectIdEnvironmentsGet(ctx, projectId).Execute()

Get All Environments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liu1700/unityremoteconfig"
)

func main() {
    projectId := "projectId_example" // string | Genesis projectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdEnvironmentsGet`: GetEnvironmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdEnvironmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnvironmentsResponse**](GetEnvironmentsResponse.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

