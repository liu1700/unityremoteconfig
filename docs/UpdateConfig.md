# UpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | We currently only support one config type. | 
**Value** | [**[]Setting**](Setting.md) | Array of &#x60;Setting&#x60; objects that represent the default values in your game. | 

## Methods

### NewUpdateConfig

`func NewUpdateConfig(type_ string, value []Setting, ) *UpdateConfig`

NewUpdateConfig instantiates a new UpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigWithDefaults

`func NewUpdateConfigWithDefaults() *UpdateConfig`

NewUpdateConfigWithDefaults instantiates a new UpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateConfig) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UpdateConfig) GetValue() []Setting`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateConfig) GetValueOk() (*[]Setting, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateConfig) SetValue(v []Setting)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


