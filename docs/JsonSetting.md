# JsonSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** | Limited to 10000 characters. | [optional] 

## Methods

### NewJsonSetting

`func NewJsonSetting() *JsonSetting`

NewJsonSetting instantiates a new JsonSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSettingWithDefaults

`func NewJsonSettingWithDefaults() *JsonSetting`

NewJsonSettingWithDefaults instantiates a new JsonSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *JsonSetting) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JsonSetting) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JsonSetting) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *JsonSetting) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *JsonSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *JsonSetting) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JsonSetting) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JsonSetting) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *JsonSetting) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


