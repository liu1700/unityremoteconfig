# IntSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] 

## Methods

### NewIntSetting

`func NewIntSetting() *IntSetting`

NewIntSetting instantiates a new IntSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntSettingWithDefaults

`func NewIntSettingWithDefaults() *IntSetting`

NewIntSettingWithDefaults instantiates a new IntSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IntSetting) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IntSetting) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IntSetting) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IntSetting) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *IntSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IntSetting) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IntSetting) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IntSetting) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *IntSetting) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


