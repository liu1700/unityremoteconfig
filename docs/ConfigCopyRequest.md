# ConfigCopyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceConfigId** | **string** |  | 
**DestinationConfigId** | **string** |  | 
**Settings** | **[]string** | Array of &#x60;Setting&#x60; keys that represent the Settings to be copied. | 

## Methods

### NewConfigCopyRequest

`func NewConfigCopyRequest(sourceConfigId string, destinationConfigId string, settings []string, ) *ConfigCopyRequest`

NewConfigCopyRequest instantiates a new ConfigCopyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCopyRequestWithDefaults

`func NewConfigCopyRequestWithDefaults() *ConfigCopyRequest`

NewConfigCopyRequestWithDefaults instantiates a new ConfigCopyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceConfigId

`func (o *ConfigCopyRequest) GetSourceConfigId() string`

GetSourceConfigId returns the SourceConfigId field if non-nil, zero value otherwise.

### GetSourceConfigIdOk

`func (o *ConfigCopyRequest) GetSourceConfigIdOk() (*string, bool)`

GetSourceConfigIdOk returns a tuple with the SourceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigId

`func (o *ConfigCopyRequest) SetSourceConfigId(v string)`

SetSourceConfigId sets SourceConfigId field to given value.


### GetDestinationConfigId

`func (o *ConfigCopyRequest) GetDestinationConfigId() string`

GetDestinationConfigId returns the DestinationConfigId field if non-nil, zero value otherwise.

### GetDestinationConfigIdOk

`func (o *ConfigCopyRequest) GetDestinationConfigIdOk() (*string, bool)`

GetDestinationConfigIdOk returns a tuple with the DestinationConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationConfigId

`func (o *ConfigCopyRequest) SetDestinationConfigId(v string)`

SetDestinationConfigId sets DestinationConfigId field to given value.


### GetSettings

`func (o *ConfigCopyRequest) GetSettings() []string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ConfigCopyRequest) GetSettingsOk() (*[]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ConfigCopyRequest) SetSettings(v []string)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


