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

// CustomBitlinkHistory struct for CustomBitlinkHistory
type CustomBitlinkHistory struct {
	Hash *string `json:"hash,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Keyword *string `json:"keyword,omitempty"`
	Created *string `json:"created,omitempty"`
	GroupGuid *string `json:"group_guid,omitempty"`
	FirstCreated *string `json:"first_created,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	LongUrl *string `json:"long_url,omitempty"`
	Deactivated *string `json:"deactivated,omitempty"`
	Bsd *string `json:"bsd,omitempty"`
	Login *string `json:"login,omitempty"`
}

// NewCustomBitlinkHistory instantiates a new CustomBitlinkHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBitlinkHistory() *CustomBitlinkHistory {
	this := CustomBitlinkHistory{}
	return &this
}

// NewCustomBitlinkHistoryWithDefaults instantiates a new CustomBitlinkHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBitlinkHistoryWithDefaults() *CustomBitlinkHistory {
	this := CustomBitlinkHistory{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *CustomBitlinkHistory) SetHash(v string) {
	o.Hash = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CustomBitlinkHistory) SetUuid(v string) {
	o.Uuid = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetKeyword() string {
	if o == nil || o.Keyword == nil {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetKeywordOk() (*string, bool) {
	if o == nil || o.Keyword == nil {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasKeyword() bool {
	if o != nil && o.Keyword != nil {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *CustomBitlinkHistory) SetKeyword(v string) {
	o.Keyword = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CustomBitlinkHistory) SetCreated(v string) {
	o.Created = &v
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *CustomBitlinkHistory) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetFirstCreated returns the FirstCreated field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetFirstCreated() string {
	if o == nil || o.FirstCreated == nil {
		var ret string
		return ret
	}
	return *o.FirstCreated
}

// GetFirstCreatedOk returns a tuple with the FirstCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetFirstCreatedOk() (*string, bool) {
	if o == nil || o.FirstCreated == nil {
		return nil, false
	}
	return o.FirstCreated, true
}

// HasFirstCreated returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasFirstCreated() bool {
	if o != nil && o.FirstCreated != nil {
		return true
	}

	return false
}

// SetFirstCreated gets a reference to the given string and assigns it to the FirstCreated field.
func (o *CustomBitlinkHistory) SetFirstCreated(v string) {
	o.FirstCreated = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CustomBitlinkHistory) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLongUrl returns the LongUrl field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetLongUrl() string {
	if o == nil || o.LongUrl == nil {
		var ret string
		return ret
	}
	return *o.LongUrl
}

// GetLongUrlOk returns a tuple with the LongUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetLongUrlOk() (*string, bool) {
	if o == nil || o.LongUrl == nil {
		return nil, false
	}
	return o.LongUrl, true
}

// HasLongUrl returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasLongUrl() bool {
	if o != nil && o.LongUrl != nil {
		return true
	}

	return false
}

// SetLongUrl gets a reference to the given string and assigns it to the LongUrl field.
func (o *CustomBitlinkHistory) SetLongUrl(v string) {
	o.LongUrl = &v
}

// GetDeactivated returns the Deactivated field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetDeactivated() string {
	if o == nil || o.Deactivated == nil {
		var ret string
		return ret
	}
	return *o.Deactivated
}

// GetDeactivatedOk returns a tuple with the Deactivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetDeactivatedOk() (*string, bool) {
	if o == nil || o.Deactivated == nil {
		return nil, false
	}
	return o.Deactivated, true
}

// HasDeactivated returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasDeactivated() bool {
	if o != nil && o.Deactivated != nil {
		return true
	}

	return false
}

// SetDeactivated gets a reference to the given string and assigns it to the Deactivated field.
func (o *CustomBitlinkHistory) SetDeactivated(v string) {
	o.Deactivated = &v
}

// GetBsd returns the Bsd field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetBsd() string {
	if o == nil || o.Bsd == nil {
		var ret string
		return ret
	}
	return *o.Bsd
}

// GetBsdOk returns a tuple with the Bsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetBsdOk() (*string, bool) {
	if o == nil || o.Bsd == nil {
		return nil, false
	}
	return o.Bsd, true
}

// HasBsd returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasBsd() bool {
	if o != nil && o.Bsd != nil {
		return true
	}

	return false
}

// SetBsd gets a reference to the given string and assigns it to the Bsd field.
func (o *CustomBitlinkHistory) SetBsd(v string) {
	o.Bsd = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *CustomBitlinkHistory) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlinkHistory) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *CustomBitlinkHistory) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *CustomBitlinkHistory) SetLogin(v string) {
	o.Login = &v
}

func (o CustomBitlinkHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Keyword != nil {
		toSerialize["keyword"] = o.Keyword
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.FirstCreated != nil {
		toSerialize["first_created"] = o.FirstCreated
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LongUrl != nil {
		toSerialize["long_url"] = o.LongUrl
	}
	if o.Deactivated != nil {
		toSerialize["deactivated"] = o.Deactivated
	}
	if o.Bsd != nil {
		toSerialize["bsd"] = o.Bsd
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	return json.Marshal(toSerialize)
}

type NullableCustomBitlinkHistory struct {
	value *CustomBitlinkHistory
	isSet bool
}

func (v NullableCustomBitlinkHistory) Get() *CustomBitlinkHistory {
	return v.value
}

func (v *NullableCustomBitlinkHistory) Set(val *CustomBitlinkHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBitlinkHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBitlinkHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBitlinkHistory(val *CustomBitlinkHistory) *NullableCustomBitlinkHistory {
	return &NullableCustomBitlinkHistory{value: val, isSet: true}
}

func (v NullableCustomBitlinkHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBitlinkHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


