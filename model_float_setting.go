/*
Remote Config Admin API

# Overview Unity Remote Config is a cloud service that allows you to tune your game design without deploying new versions of your application. With Remote Config, you can:  Adapt your game to different types of players. Tune your game difficulty curve in near real time. Alter graphic quality based on device to optimize performance. Roll out new features gradually while monitoring impact. Tailor game settings to different regions or other player segments. Run campaign tests comparing colors, styles, prices, etc. Turn seasonal, holiday, or other time-sensitive events on or off. Enable or disable features for specific player segments or across the entire user base.  Define Game Overrides that control which players receive what settings updates, and when. Unity manages the delivery and assignment of those settings with minimal impact to performance. No update to your application is necessary. When a player launches your game, Remote Config detects contextual attributes used as game override conditions, based on Unity, the application, the user, or custom criteria that you define. The service then returns customized settings for each player according to the game overrides that apply to them. This allows different players using the same version of your game to have slightly different experiences. It also allows you to understand the impact each experience has on your business.  # Limitation * A configuration can have a max payload size of 5MB * A string limit is only of 65 535 characters 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityremoteconfig

import (
	"encoding/json"
)

// checks if the FloatSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FloatSetting{}

// FloatSetting struct for FloatSetting
type FloatSetting struct {
	Key *string `json:"key,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewFloatSetting instantiates a new FloatSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFloatSetting() *FloatSetting {
	this := FloatSetting{}
	return &this
}

// NewFloatSettingWithDefaults instantiates a new FloatSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFloatSettingWithDefaults() *FloatSetting {
	this := FloatSetting{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FloatSetting) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatSetting) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FloatSetting) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FloatSetting) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FloatSetting) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatSetting) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FloatSetting) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FloatSetting) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FloatSetting) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatSetting) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FloatSetting) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *FloatSetting) SetValue(v float32) {
	o.Value = &v
}

func (o FloatSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FloatSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFloatSetting struct {
	value *FloatSetting
	isSet bool
}

func (v NullableFloatSetting) Get() *FloatSetting {
	return v.value
}

func (v *NullableFloatSetting) Set(val *FloatSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableFloatSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableFloatSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloatSetting(val *FloatSetting) *NullableFloatSetting {
	return &NullableFloatSetting{value: val, isSet: true}
}

func (v NullableFloatSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloatSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


