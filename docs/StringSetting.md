# StringSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Strings are limited to 10000 characters. | [optional] 

## Methods

### NewStringSetting

`func NewStringSetting() *StringSetting`

NewStringSetting instantiates a new StringSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSettingWithDefaults

`func NewStringSettingWithDefaults() *StringSetting`

NewStringSettingWithDefaults instantiates a new StringSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *StringSetting) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StringSetting) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StringSetting) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StringSetting) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *StringSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StringSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *StringSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringSetting) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringSetting) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


