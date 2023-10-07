# CampaignVariants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Conditionally optional. If no variants have a weight then Remote Config will treat each variant as equally distributed. However, if one variant has a weight then all variants must define a weight.  If defined, the sum of all variant&#39;s weights must equal 100  | [optional] 
**Content** | Pointer to [**[]CampaignContentVariant**](CampaignContentVariant.md) |  | [optional] 

## Methods

### NewCampaignVariants

`func NewCampaignVariants() *CampaignVariants`

NewCampaignVariants instantiates a new CampaignVariants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignVariantsWithDefaults

`func NewCampaignVariantsWithDefaults() *CampaignVariants`

NewCampaignVariantsWithDefaults instantiates a new CampaignVariants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignVariants) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignVariants) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignVariants) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignVariants) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeight

`func (o *CampaignVariants) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CampaignVariants) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CampaignVariants) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CampaignVariants) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetContent

`func (o *CampaignVariants) GetContent() []CampaignContentVariant`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignVariants) GetContentOk() (*[]CampaignContentVariant, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignVariants) SetContent(v []CampaignContentVariant)`

SetContent sets Content field to given value.

### HasContent

`func (o *CampaignVariants) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


