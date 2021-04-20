# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to **map[string]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ModifiedBy** | **string** |  | 
**Name** | **string** |  | 
**Created** | **string** |  | 
**Url** | **string** |  | 
**Deactivated** | **string** |  | 
**IsActive** | **bool** |  | 
**Modified** | **string** |  | 
**OrganizationGuid** | **string** |  | 
**GroupGuid** | Pointer to **string** |  | [optional] 
**Guid** | **string** |  | 
**Event** | **string** |  | 

## Methods

### NewWebhook

`func NewWebhook(modifiedBy string, name string, created string, url string, deactivated string, isActive bool, modified string, organizationGuid string, guid string, event string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *Webhook) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Webhook) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Webhook) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Webhook) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetStatus

`func (o *Webhook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Webhook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Webhook) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Webhook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Webhook) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Webhook) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Webhook) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *Webhook) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Webhook) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Webhook) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDeactivated

`func (o *Webhook) GetDeactivated() string`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *Webhook) GetDeactivatedOk() (*string, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *Webhook) SetDeactivated(v string)`

SetDeactivated sets Deactivated field to given value.


### GetIsActive

`func (o *Webhook) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Webhook) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Webhook) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetModified

`func (o *Webhook) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Webhook) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Webhook) SetModified(v string)`

SetModified sets Modified field to given value.


### GetOrganizationGuid

`func (o *Webhook) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *Webhook) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *Webhook) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.


### GetGroupGuid

`func (o *Webhook) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *Webhook) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *Webhook) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *Webhook) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetGuid

`func (o *Webhook) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Webhook) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Webhook) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetEvent

`func (o *Webhook) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Webhook) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Webhook) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


