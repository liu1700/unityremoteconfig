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

// checks if the GetAllCampaignsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllCampaignsResponse{}

// GetAllCampaignsResponse A list of all campaigns for a given project
type GetAllCampaignsResponse struct {
	Campaigns []GetCampaignResponse `json:"campaigns,omitempty"`
}

// NewGetAllCampaignsResponse instantiates a new GetAllCampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllCampaignsResponse() *GetAllCampaignsResponse {
	this := GetAllCampaignsResponse{}
	return &this
}

// NewGetAllCampaignsResponseWithDefaults instantiates a new GetAllCampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllCampaignsResponseWithDefaults() *GetAllCampaignsResponse {
	this := GetAllCampaignsResponse{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *GetAllCampaignsResponse) GetCampaigns() []GetCampaignResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret []GetCampaignResponse
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCampaignsResponse) GetCampaignsOk() ([]GetCampaignResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *GetAllCampaignsResponse) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []GetCampaignResponse and assigns it to the Campaigns field.
func (o *GetAllCampaignsResponse) SetCampaigns(v []GetCampaignResponse) {
	o.Campaigns = v
}

func (o GetAllCampaignsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllCampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableGetAllCampaignsResponse struct {
	value *GetAllCampaignsResponse
	isSet bool
}

func (v NullableGetAllCampaignsResponse) Get() *GetAllCampaignsResponse {
	return v.value
}

func (v *NullableGetAllCampaignsResponse) Set(val *GetAllCampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllCampaignsResponse(val *GetAllCampaignsResponse) *NullableGetAllCampaignsResponse {
	return &NullableGetAllCampaignsResponse{value: val, isSet: true}
}

func (v NullableGetAllCampaignsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

