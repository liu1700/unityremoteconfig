# Go API client for unityremoteconfig

# Overview
Unity Remote Config is a cloud service that allows you to tune your game design without deploying new versions of your application. With Remote Config, you can:

Adapt your game to different types of players.
Tune your game difficulty curve in near real time.
Alter graphic quality based on device to optimize performance.
Roll out new features gradually while monitoring impact.
Tailor game settings to different regions or other player segments.
Run campaign tests comparing colors, styles, prices, etc.
Turn seasonal, holiday, or other time-sensitive events on or off.
Enable or disable features for specific player segments or across the entire user base.

Define Game Overrides that control which players receive what settings updates, and when. Unity manages the delivery and assignment of those settings with minimal impact to performance. No update to your application is necessary. When a player launches your game, Remote Config detects contextual attributes used as game override conditions, based on Unity, the application, the user, or custom criteria that you define. The service then returns customized settings for each player according to the game overrides that apply to them. This allows different players using the same version of your game to have slightly different experiences. It also allows you to understand the impact each experience has on your business.

# Limitation
* A configuration can have a max payload size of 5MB
* A string limit is only of 65 535 characters


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import unityremoteconfig "github.com/liu1700/unityremoteconfig"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), unityremoteconfig.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), unityremoteconfig.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), unityremoteconfig.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), unityremoteconfig.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://services.api.unity.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdConfigsConfigIdDelete**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidconfigsconfigiddelete) | **Delete** /remote-config/v1/projects/{projectId}/configs/{configId} | Delete Config
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdConfigsConfigIdGet**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidconfigsconfigidget) | **Get** /remote-config/v1/projects/{projectId}/configs/{configId} | Get Config
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdConfigsConfigIdPut**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidconfigsconfigidput) | **Put** /remote-config/v1/projects/{projectId}/configs/{configId} | Update Config
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdConfigsGet**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidconfigsget) | **Get** /remote-config/v1/projects/{projectId}/configs | Get config with the default environment
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdConfigsPost**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidconfigspost) | **Post** /remote-config/v1/projects/{projectId}/configs | Create Config
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdConfigscopytoPost**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidconfigscopytopost) | **Post** /remote-config/v1/projects/{projectId}/configs:copyto | Copy Config Settings
*ConfigsAPI* | [**RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet**](docs/ConfigsAPI.md#remoteconfigv1projectsprojectidenvironmentsenvironmentidconfigsget) | **Get** /remote-config/v1/projects/{projectId}/environments/{environmentId}/configs | Get Configs by Environment ID
*EnvironmentsAPI* | [**RemoteConfigV1ProjectsProjectIdEnvironmentsDefaultGet**](docs/EnvironmentsAPI.md#remoteconfigv1projectsprojectidenvironmentsdefaultget) | **Get** /remote-config/v1/projects/{projectId}/environments/default | Get Default Environment
*EnvironmentsAPI* | [**RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdGet**](docs/EnvironmentsAPI.md#remoteconfigv1projectsprojectidenvironmentsenvironmentidget) | **Get** /remote-config/v1/projects/{projectId}/environments/{environmentId} | Get Environment
*EnvironmentsAPI* | [**RemoteConfigV1ProjectsProjectIdEnvironmentsGet**](docs/EnvironmentsAPI.md#remoteconfigv1projectsprojectidenvironmentsget) | **Get** /remote-config/v1/projects/{projectId}/environments | Get All Environments
*GameOverridesAPI* | [**CreateGameOverride**](docs/GameOverridesAPI.md#creategameoverride) | **Post** /remote-config/v1/projects/{projectId}/environments/{environmentId}/campaigns | Create a game override
*GameOverridesAPI* | [**DeleteGameOverride**](docs/GameOverridesAPI.md#deletegameoverride) | **Delete** /remote-config/v1/projects/{projectId}/campaigns/{campaignId} | Delete a single game override
*GameOverridesAPI* | [**GetAllGameOverrides**](docs/GameOverridesAPI.md#getallgameoverrides) | **Get** /remote-config/v1/projects/{projectId}/environments/{environmentId}/campaigns | Get all game overrides
*GameOverridesAPI* | [**GetGameOverride**](docs/GameOverridesAPI.md#getgameoverride) | **Get** /remote-config/v1/projects/{projectId}/campaigns/{campaignId} | Get a single game override
*GameOverridesAPI* | [**UpdateGameOverride**](docs/GameOverridesAPI.md#updategameoverride) | **Put** /remote-config/v1/projects/{projectId}/campaigns/{campaignId} | Update a single game override


## Documentation For Models

 - [BoolSetting](docs/BoolSetting.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignContentVariant](docs/CampaignContentVariant.md)
 - [CampaignVariants](docs/CampaignVariants.md)
 - [Config](docs/Config.md)
 - [ConfigCopyRequest](docs/ConfigCopyRequest.md)
 - [CreateCampaignResponse](docs/CreateCampaignResponse.md)
 - [CreateConfigResponse](docs/CreateConfigResponse.md)
 - [CreateEnvironmentResponse](docs/CreateEnvironmentResponse.md)
 - [DefaultEnvironmentResponse](docs/DefaultEnvironmentResponse.md)
 - [Environment](docs/Environment.md)
 - [FloatSetting](docs/FloatSetting.md)
 - [GetAllCampaignsResponse](docs/GetAllCampaignsResponse.md)
 - [GetCampaignResponse](docs/GetCampaignResponse.md)
 - [GetEnvironmentsResponse](docs/GetEnvironmentsResponse.md)
 - [IntSetting](docs/IntSetting.md)
 - [JsonSetting](docs/JsonSetting.md)
 - [LongSetting](docs/LongSetting.md)
 - [RemoteConfigV1ProjectsProjectIdConfigsCopytoPost404Response](docs/RemoteConfigV1ProjectsProjectIdConfigsCopytoPost404Response.md)
 - [RemoteConfigV1ProjectsProjectIdConfigsCopytoPost409Response](docs/RemoteConfigV1ProjectsProjectIdConfigsCopytoPost409Response.md)
 - [RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response](docs/RemoteConfigV1ProjectsProjectIdEnvironmentsEnvironmentIdConfigsGet200Response.md)
 - [Setting](docs/Setting.md)
 - [StringSetting](docs/StringSetting.md)
 - [UpdateConfig](docs/UpdateConfig.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ServiceAccount

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



