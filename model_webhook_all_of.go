/*
 * Bitly API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.0.0
 * Contact: api@bitly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WebhookAllOf struct for WebhookAllOf
type WebhookAllOf struct {
	Status *string `json:"status,omitempty"`
	ModifiedBy *string `json:"modified_by,omitempty"`
	Name *string `json:"name,omitempty"`
	Created *string `json:"created,omitempty"`
	Url *string `json:"url,omitempty"`
	Deactivated *string `json:"deactivated,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Modified *string `json:"modified,omitempty"`
	OrganizationGuid *string `json:"organization_guid,omitempty"`
	GroupGuid *string `json:"group_guid,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Event *string `json:"event,omitempty"`
}

// NewWebhookAllOf instantiates a new WebhookAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookAllOf() *WebhookAllOf {
	this := WebhookAllOf{}
	return &this
}

// NewWebhookAllOfWithDefaults instantiates a new WebhookAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookAllOfWithDefaults() *WebhookAllOf {
	this := WebhookAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhookAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhookAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebhookAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *WebhookAllOf) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *WebhookAllOf) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *WebhookAllOf) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookAllOf) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WebhookAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WebhookAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WebhookAllOf) SetCreated(v string) {
	o.Created = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhookAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebhookAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetDeactivated returns the Deactivated field value if set, zero value otherwise.
func (o *WebhookAllOf) GetDeactivated() string {
	if o == nil || o.Deactivated == nil {
		var ret string
		return ret
	}
	return *o.Deactivated
}

// GetDeactivatedOk returns a tuple with the Deactivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetDeactivatedOk() (*string, bool) {
	if o == nil || o.Deactivated == nil {
		return nil, false
	}
	return o.Deactivated, true
}

// HasDeactivated returns a boolean if a field has been set.
func (o *WebhookAllOf) HasDeactivated() bool {
	if o != nil && o.Deactivated != nil {
		return true
	}

	return false
}

// SetDeactivated gets a reference to the given string and assigns it to the Deactivated field.
func (o *WebhookAllOf) SetDeactivated(v string) {
	o.Deactivated = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *WebhookAllOf) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *WebhookAllOf) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *WebhookAllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *WebhookAllOf) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *WebhookAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *WebhookAllOf) SetModified(v string) {
	o.Modified = &v
}

// GetOrganizationGuid returns the OrganizationGuid field value if set, zero value otherwise.
func (o *WebhookAllOf) GetOrganizationGuid() string {
	if o == nil || o.OrganizationGuid == nil {
		var ret string
		return ret
	}
	return *o.OrganizationGuid
}

// GetOrganizationGuidOk returns a tuple with the OrganizationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetOrganizationGuidOk() (*string, bool) {
	if o == nil || o.OrganizationGuid == nil {
		return nil, false
	}
	return o.OrganizationGuid, true
}

// HasOrganizationGuid returns a boolean if a field has been set.
func (o *WebhookAllOf) HasOrganizationGuid() bool {
	if o != nil && o.OrganizationGuid != nil {
		return true
	}

	return false
}

// SetOrganizationGuid gets a reference to the given string and assigns it to the OrganizationGuid field.
func (o *WebhookAllOf) SetOrganizationGuid(v string) {
	o.OrganizationGuid = &v
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *WebhookAllOf) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *WebhookAllOf) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *WebhookAllOf) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *WebhookAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *WebhookAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *WebhookAllOf) SetGuid(v string) {
	o.Guid = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *WebhookAllOf) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAllOf) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *WebhookAllOf) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *WebhookAllOf) SetEvent(v string) {
	o.Event = &v
}

func (o WebhookAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Deactivated != nil {
		toSerialize["deactivated"] = o.Deactivated
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.OrganizationGuid != nil {
		toSerialize["organization_guid"] = o.OrganizationGuid
	}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookAllOf struct {
	value *WebhookAllOf
	isSet bool
}

func (v NullableWebhookAllOf) Get() *WebhookAllOf {
	return v.value
}

func (v *NullableWebhookAllOf) Set(val *WebhookAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookAllOf(val *WebhookAllOf) *NullableWebhookAllOf {
	return &NullableWebhookAllOf{value: val, isSet: true}
}

func (v NullableWebhookAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


