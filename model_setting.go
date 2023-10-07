/*
Remote Config Admin API

# Overview Unity Remote Config is a cloud service that allows you to tune your game design without deploying new versions of your application. With Remote Config, you can:  Adapt your game to different types of players. Tune your game difficulty curve in near real time. Alter graphic quality based on device to optimize performance. Roll out new features gradually while monitoring impact. Tailor game settings to different regions or other player segments. Run campaign tests comparing colors, styles, prices, etc. Turn seasonal, holiday, or other time-sensitive events on or off. Enable or disable features for specific player segments or across the entire user base.  Define Game Overrides that control which players receive what settings updates, and when. Unity manages the delivery and assignment of those settings with minimal impact to performance. No update to your application is necessary. When a player launches your game, Remote Config detects contextual attributes used as game override conditions, based on Unity, the application, the user, or custom criteria that you define. The service then returns customized settings for each player according to the game overrides that apply to them. This allows different players using the same version of your game to have slightly different experiences. It also allows you to understand the impact each experience has on your business.  # Limitation * A configuration can have a max payload size of 5MB * A string limit is only of 65 535 characters 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unityremoteconfig

import (
	"encoding/json"
	"fmt"
)

// Setting - Each Setting consists of a Key, Type, and Value:   - `key` the name of the setting   - `type` C# variable data type of the setting value   - `value` the value for this setting 
type Setting struct {
	BoolSetting *BoolSetting
	FloatSetting *FloatSetting
	IntSetting *IntSetting
	JsonSetting *JsonSetting
	LongSetting *LongSetting
	StringSetting *StringSetting
}

// BoolSettingAsSetting is a convenience function that returns BoolSetting wrapped in Setting
func BoolSettingAsSetting(v *BoolSetting) Setting {
	return Setting{
		BoolSetting: v,
	}
}

// FloatSettingAsSetting is a convenience function that returns FloatSetting wrapped in Setting
func FloatSettingAsSetting(v *FloatSetting) Setting {
	return Setting{
		FloatSetting: v,
	}
}

// IntSettingAsSetting is a convenience function that returns IntSetting wrapped in Setting
func IntSettingAsSetting(v *IntSetting) Setting {
	return Setting{
		IntSetting: v,
	}
}

// JsonSettingAsSetting is a convenience function that returns JsonSetting wrapped in Setting
func JsonSettingAsSetting(v *JsonSetting) Setting {
	return Setting{
		JsonSetting: v,
	}
}

// LongSettingAsSetting is a convenience function that returns LongSetting wrapped in Setting
func LongSettingAsSetting(v *LongSetting) Setting {
	return Setting{
		LongSetting: v,
	}
}

// StringSettingAsSetting is a convenience function that returns StringSetting wrapped in Setting
func StringSettingAsSetting(v *StringSetting) Setting {
	return Setting{
		StringSetting: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Setting) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BoolSetting
	err = newStrictDecoder(data).Decode(&dst.BoolSetting)
	if err == nil {
		jsonBoolSetting, _ := json.Marshal(dst.BoolSetting)
		if string(jsonBoolSetting) == "{}" { // empty struct
			dst.BoolSetting = nil
		} else {
			match++
		}
	} else {
		dst.BoolSetting = nil
	}

	// try to unmarshal data into FloatSetting
	err = newStrictDecoder(data).Decode(&dst.FloatSetting)
	if err == nil {
		jsonFloatSetting, _ := json.Marshal(dst.FloatSetting)
		if string(jsonFloatSetting) == "{}" { // empty struct
			dst.FloatSetting = nil
		} else {
			match++
		}
	} else {
		dst.FloatSetting = nil
	}

	// try to unmarshal data into IntSetting
	err = newStrictDecoder(data).Decode(&dst.IntSetting)
	if err == nil {
		jsonIntSetting, _ := json.Marshal(dst.IntSetting)
		if string(jsonIntSetting) == "{}" { // empty struct
			dst.IntSetting = nil
		} else {
			match++
		}
	} else {
		dst.IntSetting = nil
	}

	// try to unmarshal data into JsonSetting
	err = newStrictDecoder(data).Decode(&dst.JsonSetting)
	if err == nil {
		jsonJsonSetting, _ := json.Marshal(dst.JsonSetting)
		if string(jsonJsonSetting) == "{}" { // empty struct
			dst.JsonSetting = nil
		} else {
			match++
		}
	} else {
		dst.JsonSetting = nil
	}

	// try to unmarshal data into LongSetting
	err = newStrictDecoder(data).Decode(&dst.LongSetting)
	if err == nil {
		jsonLongSetting, _ := json.Marshal(dst.LongSetting)
		if string(jsonLongSetting) == "{}" { // empty struct
			dst.LongSetting = nil
		} else {
			match++
		}
	} else {
		dst.LongSetting = nil
	}

	// try to unmarshal data into StringSetting
	err = newStrictDecoder(data).Decode(&dst.StringSetting)
	if err == nil {
		jsonStringSetting, _ := json.Marshal(dst.StringSetting)
		if string(jsonStringSetting) == "{}" { // empty struct
			dst.StringSetting = nil
		} else {
			match++
		}
	} else {
		dst.StringSetting = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BoolSetting = nil
		dst.FloatSetting = nil
		dst.IntSetting = nil
		dst.JsonSetting = nil
		dst.LongSetting = nil
		dst.StringSetting = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Setting)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Setting)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Setting) MarshalJSON() ([]byte, error) {
	if src.BoolSetting != nil {
		return json.Marshal(&src.BoolSetting)
	}

	if src.FloatSetting != nil {
		return json.Marshal(&src.FloatSetting)
	}

	if src.IntSetting != nil {
		return json.Marshal(&src.IntSetting)
	}

	if src.JsonSetting != nil {
		return json.Marshal(&src.JsonSetting)
	}

	if src.LongSetting != nil {
		return json.Marshal(&src.LongSetting)
	}

	if src.StringSetting != nil {
		return json.Marshal(&src.StringSetting)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Setting) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BoolSetting != nil {
		return obj.BoolSetting
	}

	if obj.FloatSetting != nil {
		return obj.FloatSetting
	}

	if obj.IntSetting != nil {
		return obj.IntSetting
	}

	if obj.JsonSetting != nil {
		return obj.JsonSetting
	}

	if obj.LongSetting != nil {
		return obj.LongSetting
	}

	if obj.StringSetting != nil {
		return obj.StringSetting
	}

	// all schemas are nil
	return nil
}

type NullableSetting struct {
	value *Setting
	isSet bool
}

func (v NullableSetting) Get() *Setting {
	return v.value
}

func (v *NullableSetting) Set(val *Setting) {
	v.value = val
	v.isSet = true
}

func (v NullableSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetting(val *Setting) *NullableSetting {
	return &NullableSetting{value: val, isSet: true}
}

func (v NullableSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


