# GetCampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EnvironmentId** | Pointer to **string** |  | [optional] [readonly] 
**ProjectId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Variants** | Pointer to [**[]CampaignVariants**](CampaignVariants.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**RolloutPercentage** | Pointer to **int32** | The percentage of your user base that will adhere to this game override. For values less than 100, Unity randomly assigns the game override to that percent of your players on a user ID basis. While experiences may differ from player to player, players will have a consistent experience each session.  | [optional] [default to 100]
**Priority** | Pointer to **int32** | Integer between 0 (highest priority) and 1000 (lowest priority). In the event of a conflict, priority is awarded to the game overrides that was created first. | [optional] [default to 1000]
**Condition** | Pointer to **string** | The Condition is a JEXL expression of contextual data attributes that defines the target audience for a game override.  ### JEXL support  - &#x60;JexlExpression&#x60; from the [Java Expression Language (JEXL) spec](https://commons.apache.org/proper/commons-jexl/reference/syntax.html). - &#x60;JexlScripts&#x60; are not supported.  ### Categories  Remote Config supports the following attribute categories.  Category | Description -------- | ----------- &#x60;user&#x60;   | Custom developer-defined attributes that describe the user. &#x60;app&#x60;    | Custom developer-defined attributes that describe the application. &#x60;unity&#x60;  | Predefined attributes (detailed below).  #### Custom User ID  When you call &#x60;SetCustomUserId&#x60; in our C# APIs, this automatically exposes the ID you pass in to the &#x60;user&#x60; attributes.  Example from C#: &#x60;&#x60;&#x60; ConfigManager.SetCustomUserID(\&quot;some-user-id\&quot;); &#x60;&#x60;&#x60;  Allows you to create a game override condition like:  &#x60;&#x60;&#x60; user.customUserId &#x3D;&#x3D; \&quot;some-user-id\&quot; &#x60;&#x60;&#x60;  #### Predefined &#x60;unity&#x60; attributes  Attribute              | Type |   Description | Example ---------              | ----   | ----------- | ------- &#x60;appBuildVersion&#x60;      | String | The build number your application is running. | &#x60;\&quot;1\&quot;&#x60; &#x60;appVersion&#x60;           | String | The version your application is running. | &#x60;\&quot;1.0\&quot;&#x60; &#x60;cpu&#x60;                  | String | The name of the CPU processor. | &#x60;\&quot;Intel(R) Core(TM) i7-7920HQ CPU @ 3.10GHz\&quot;&#x60; &#x60;cpuFrequency&#x60;         | Int    | The processor frequency in MHz of the device running your app. | &#x60;3100&#x60; &#x60;country&#x60;              | String | The applicable country. This attribute uses ISO 3166-1 alpha2 country codes. | &#x60;\&quot;US\&quot;&#x60; &#x60;language&#x60;             | String | The applicable language. This attribute uses ISO 639-1 language codes. | &#x60;\&quot;en\&quot;&#x60; &#x60;osVersion&#x60;            | String | The operating system version of the device running your app. | &#x60;\&quot;Mac OS X 10.14.3\&quot;&#x60; &#x60;platform&#x60;             | String | The applicable device or platform (see below). | &#x60;\&quot;iOS\&quot;&#x60; &#x60;timeSinceStart&#x60;       | Int    | The time in milliseconds since a session of your app has begun. | &#x60;2473741&#x60; &#x60;graphicsDeviceVendor&#x60; | String | Vendor of the user&#39;s graphics card. | &#x60;\&quot;ATI Technologies Inc.\&quot;&#x60; &#x60;ram&#x60;                  | Int    | Amount of RAM in MB on the device. | &#x60;16384&#x60;  #### Supported platforms    - &#x60;Android&#x60;   - &#x60;iOS&#x60;   - &#x60;Linux&#x60;   - &#x60;macOS&#x60;   - &#x60;Metro&#x60;   - &#x60;SamsungTV&#x60;   - &#x60;Switch&#x60;   - &#x60;Tizen&#x60;   - &#x60;tvOS&#x60;   - &#x60;WebGL&#x60;   - &#x60;Wii&#x60;   - &#x60;Windows&#x60;  | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Timestamp in ISO 8601 UTC format (YYYY-MM-DDThh:mm:ssZ). Lower bound for when enabled game overrides will go into effect. | [optional] 
**EndDate** | Pointer to **time.Time** | Timestamp in ISO 8601 UTC format (YYYY-MM-DDThh:mm:ssZ). Upper bound for when enabled game overrides will cease to be active. | [optional] 
**Kpi** | Pointer to **[]string** | List of key performing indicators | [optional] 
**Audience** | Pointer to **[]string** | A list of audience whom you want to target your campaign | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCampaignResponse

`func NewGetCampaignResponse() *GetCampaignResponse`

NewGetCampaignResponse instantiates a new GetCampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignResponseWithDefaults

`func NewGetCampaignResponseWithDefaults() *GetCampaignResponse`

NewGetCampaignResponseWithDefaults instantiates a new GetCampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCampaignResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCampaignResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCampaignResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetCampaignResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *GetCampaignResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *GetCampaignResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *GetCampaignResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *GetCampaignResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetProjectId

`func (o *GetCampaignResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetCampaignResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetCampaignResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetCampaignResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *GetCampaignResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCampaignResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCampaignResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCampaignResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariants

`func (o *GetCampaignResponse) GetVariants() []CampaignVariants`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *GetCampaignResponse) GetVariantsOk() (*[]CampaignVariants, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *GetCampaignResponse) SetVariants(v []CampaignVariants)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *GetCampaignResponse) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetEnabled

`func (o *GetCampaignResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCampaignResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCampaignResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCampaignResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRolloutPercentage

`func (o *GetCampaignResponse) GetRolloutPercentage() int32`

GetRolloutPercentage returns the RolloutPercentage field if non-nil, zero value otherwise.

### GetRolloutPercentageOk

`func (o *GetCampaignResponse) GetRolloutPercentageOk() (*int32, bool)`

GetRolloutPercentageOk returns a tuple with the RolloutPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentage

`func (o *GetCampaignResponse) SetRolloutPercentage(v int32)`

SetRolloutPercentage sets RolloutPercentage field to given value.

### HasRolloutPercentage

`func (o *GetCampaignResponse) HasRolloutPercentage() bool`

HasRolloutPercentage returns a boolean if a field has been set.

### GetPriority

`func (o *GetCampaignResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetCampaignResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetCampaignResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetCampaignResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCondition

`func (o *GetCampaignResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetCampaignResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetCampaignResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetCampaignResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetStartDate

`func (o *GetCampaignResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetCampaignResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetCampaignResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetCampaignResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetCampaignResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetCampaignResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetCampaignResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetCampaignResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetKpi

`func (o *GetCampaignResponse) GetKpi() []string`

GetKpi returns the Kpi field if non-nil, zero value otherwise.

### GetKpiOk

`func (o *GetCampaignResponse) GetKpiOk() (*[]string, bool)`

GetKpiOk returns a tuple with the Kpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpi

`func (o *GetCampaignResponse) SetKpi(v []string)`

SetKpi sets Kpi field to given value.

### HasKpi

`func (o *GetCampaignResponse) HasKpi() bool`

HasKpi returns a boolean if a field has been set.

### GetAudience

`func (o *GetCampaignResponse) GetAudience() []string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *GetCampaignResponse) GetAudienceOk() (*[]string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *GetCampaignResponse) SetAudience(v []string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *GetCampaignResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetCampaignResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCampaignResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCampaignResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetCampaignResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetCampaignResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetCampaignResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetCampaignResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetCampaignResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GetCampaignResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCampaignResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCampaignResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCampaignResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


