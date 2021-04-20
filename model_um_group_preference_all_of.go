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

// UMGroupPreferenceAllOf struct for UMGroupPreferenceAllOf
type UMGroupPreferenceAllOf struct {
	Created *string `json:"created,omitempty"`
	GroupGuid *string `json:"group_guid,omitempty"`
	Modified *string `json:"modified,omitempty"`
	Value *string `json:"value,omitempty"`
	Preference *string `json:"preference,omitempty"`
	Expired *bool `json:"expired,omitempty"`
	LastSeen *string `json:"last_seen,omitempty"`
}

// NewUMGroupPreferenceAllOf instantiates a new UMGroupPreferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUMGroupPreferenceAllOf() *UMGroupPreferenceAllOf {
	this := UMGroupPreferenceAllOf{}
	return &this
}

// NewUMGroupPreferenceAllOfWithDefaults instantiates a new UMGroupPreferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUMGroupPreferenceAllOfWithDefaults() *UMGroupPreferenceAllOf {
	this := UMGroupPreferenceAllOf{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *UMGroupPreferenceAllOf) SetCreated(v string) {
	o.Created = &v
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *UMGroupPreferenceAllOf) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *UMGroupPreferenceAllOf) SetModified(v string) {
	o.Modified = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UMGroupPreferenceAllOf) SetValue(v string) {
	o.Value = &v
}

// GetPreference returns the Preference field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetPreference() string {
	if o == nil || o.Preference == nil {
		var ret string
		return ret
	}
	return *o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetPreferenceOk() (*string, bool) {
	if o == nil || o.Preference == nil {
		return nil, false
	}
	return o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasPreference() bool {
	if o != nil && o.Preference != nil {
		return true
	}

	return false
}

// SetPreference gets a reference to the given string and assigns it to the Preference field.
func (o *UMGroupPreferenceAllOf) SetPreference(v string) {
	o.Preference = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetExpired() bool {
	if o == nil || o.Expired == nil {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetExpiredOk() (*bool, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *UMGroupPreferenceAllOf) SetExpired(v bool) {
	o.Expired = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *UMGroupPreferenceAllOf) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UMGroupPreferenceAllOf) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *UMGroupPreferenceAllOf) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *UMGroupPreferenceAllOf) SetLastSeen(v string) {
	o.LastSeen = &v
}

func (o UMGroupPreferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Preference != nil {
		toSerialize["preference"] = o.Preference
	}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	return json.Marshal(toSerialize)
}

type NullableUMGroupPreferenceAllOf struct {
	value *UMGroupPreferenceAllOf
	isSet bool
}

func (v NullableUMGroupPreferenceAllOf) Get() *UMGroupPreferenceAllOf {
	return v.value
}

func (v *NullableUMGroupPreferenceAllOf) Set(val *UMGroupPreferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUMGroupPreferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUMGroupPreferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUMGroupPreferenceAllOf(val *UMGroupPreferenceAllOf) *NullableUMGroupPreferenceAllOf {
	return &NullableUMGroupPreferenceAllOf{value: val, isSet: true}
}

func (v NullableUMGroupPreferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUMGroupPreferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


