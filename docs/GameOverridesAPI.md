# \GameOverridesAPI

All URIs are relative to *https://services.api.unity.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGameOverride**](GameOverridesAPI.md#DeleteGameOverride) | **Delete** /remote-config/v1/projects/{projectId}/campaigns/{campaignId} | Delete a single game override
[**GetGameOverride**](GameOverridesAPI.md#GetGameOverride) | **Get** /remote-config/v1/projects/{projectId}/campaigns/{campaignId} | Get a single game override
[**UpdateGameOverride**](GameOverridesAPI.md#UpdateGameOverride) | **Put** /remote-config/v1/projects/{projectId}/campaigns/{campaignId} | Update a single game override



## DeleteGameOverride

> DeleteGameOverride(ctx, campaignId, projectId).Execute()

Delete a single game override



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
    campaignId := "campaignId_example" // string | Campaign ID
    projectId := "projectId_example" // string | Project Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameOverridesAPI.DeleteGameOverride(context.Background(), campaignId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameOverridesAPI.DeleteGameOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign ID | 
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGameOverrideRequest struct via the builder pattern


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


## GetGameOverride

> GetCampaignResponse GetGameOverride(ctx, campaignId, projectId).Execute()

Get a single game override



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
    campaignId := "campaignId_example" // string | Campaign ID
    projectId := "projectId_example" // string | Project Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameOverridesAPI.GetGameOverride(context.Background(), campaignId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameOverridesAPI.GetGameOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGameOverride`: GetCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `GameOverridesAPI.GetGameOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign ID | 
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCampaignResponse**](GetCampaignResponse.md)

### Authorization

[ServiceAccount](../README.md#ServiceAccount)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGameOverride

> UpdateGameOverride(ctx, campaignId, projectId).Campaign(campaign).Execute()

Update a single game override



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
    campaignId := "campaignId_example" // string | Campaign ID
    projectId := "projectId_example" // string | Project Id
    campaign := *openapiclient.NewCampaign() // Campaign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameOverridesAPI.UpdateGameOverride(context.Background(), campaignId, projectId).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameOverridesAPI.UpdateGameOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign ID | 
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGameOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **campaign** | [**Campaign**](Campaign.md) |  | 

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

