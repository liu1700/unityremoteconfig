# \ConfigsAPI

All URIs are relative to *https://services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete) | **Delete** /remote-config/v1/projects/{projectId}/configs/{configId} | Delete Config
[**RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet) | **Get** /remote-config/v1/projects/{projectId}/configs/{configId} | Get Config
[**RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut) | **Put** /remote-config/v1/projects/{projectId}/configs/{configId} | Update Config
[**RemoteConfigV1ProjectsProjectIdConfigsGet**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdConfigsGet) | **Get** /remote-config/v1/projects/{projectId}/configs | Get config with the default environment
[**RemoteConfigV1ProjectsProjectIdConfigsPost**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdConfigsPost) | **Post** /remote-config/v1/projects/{projectId}/configs | Create Config
[**RemoteConfigV1ProjectsProjectIdConfigscopytoPost**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdConfigscopytoPost) | **Post** /remote-config/v1/projects/{projectId}/configs:copyto | Copy Config Settings
[**RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet**](ConfigsAPI.md#RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet) | **Get** /remote-config/v1/projects/{projectId}/environments/{environmentId}/configs | Get Configs by Environment ID



## RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete

> RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete(ctx, configId, projectId).Execute()

Delete Config



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
    configId := "configId_example" // string | Configuration Id
    projectId := "projectId_example" // string | Genesis projectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete(context.Background(), configId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration Id | 
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdConfigsConfigIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet

> Config RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet(ctx, configId, projectId).Execute()

Get Config



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
    configId := "configId_example" // string | Configuration Id
    projectId := "projectId_example" // string | Genesis projectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet(context.Background(), configId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration Id | 
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdConfigsConfigIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Config**](Config.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut

> RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut(ctx, configId, projectId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Update Config



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
    configId := "configId_example" // string | Configuration Id
    projectId := "projectId_example" // string | Genesis projectId
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Updated Config values

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut(context.Background(), configId, projectId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration Id | 
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdConfigsConfigIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Updated Config values | 

### Return type

 (empty response body)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdConfigsGet

> Config RemoteConfigV1ProjectsProjectIdConfigsGet(ctx, projectId).Execute()

Get config with the default environment



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
    resp, r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdConfigsGet`: Config
    fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdConfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Config**](Config.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdConfigsPost

> CreateConfigResponse RemoteConfigV1ProjectsProjectIdConfigsPost(ctx, projectId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Create Config



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Config object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsPost(context.Background(), projectId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdConfigsPost`: CreateConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdConfigsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Config object | 

### Return type

[**CreateConfigResponse**](CreateConfigResponse.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdConfigscopytoPost

> RemoteConfigV1ProjectsProjectIdConfigscopytoPost(ctx, projectId).ConfigCopyRequest(configCopyRequest).Execute()

Copy Config Settings



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
    configCopyRequest := *openapiclient.NewConfigCopyRequest("SourceConfigId_example", "DestinationConfigId_example", []string{"Settings_example"}) // ConfigCopyRequest | IDs for source and destination Config, and Settings keys to be copied

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigscopytoPost(context.Background(), projectId).ConfigCopyRequest(configCopyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdConfigscopytoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdConfigscopytoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configCopyRequest** | [**ConfigCopyRequest**](ConfigCopyRequest.md) | IDs for source and destination Config, and Settings keys to be copied | 

### Return type

 (empty response body)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet

> RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet(ctx, environmentId, projectId).Execute()

Get Configs by Environment ID



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
    resp, r, err := apiClient.ConfigsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet(context.Background(), environmentId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet`: RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment Id | 
**projectId** | **string** | Genesis projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response**](RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

