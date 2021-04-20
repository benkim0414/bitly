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

// Organization struct for Organization
type Organization struct {
	References *map[string]string `json:"references,omitempty"`
	Name string `json:"name"`
	Bsds []string `json:"bsds"`
	Created string `json:"created"`
	IsActive bool `json:"is_active"`
	Modified string `json:"modified"`
	TierDisplayName string `json:"tier_display_name"`
	TierFamily string `json:"tier_family"`
	Tier string `json:"tier"`
	Role string `json:"role"`
	Guid string `json:"guid"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(name string, bsds []string, created string, isActive bool, modified string, tierDisplayName string, tierFamily string, tier string, role string, guid string) *Organization {
	this := Organization{}
	this.Name = name
	this.Bsds = bsds
	this.Created = created
	this.IsActive = isActive
	this.Modified = modified
	this.TierDisplayName = tierDisplayName
	this.TierFamily = tierFamily
	this.Tier = tier
	this.Role = role
	this.Guid = guid
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *Organization) GetReferences() map[string]string {
	if o == nil || o.References == nil {
		var ret map[string]string
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetReferencesOk() (*map[string]string, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *Organization) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given map[string]string and assigns it to the References field.
func (o *Organization) SetReferences(v map[string]string) {
	o.References = &v
}

// GetName returns the Name field value
func (o *Organization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Organization) SetName(v string) {
	o.Name = v
}

// GetBsds returns the Bsds field value
func (o *Organization) GetBsds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Bsds
}

// GetBsdsOk returns a tuple with the Bsds field value
// and a boolean to check if the value has been set.
func (o *Organization) GetBsdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bsds, true
}

// SetBsds sets field value
func (o *Organization) SetBsds(v []string) {
	o.Bsds = v
}

// GetCreated returns the Created field value
func (o *Organization) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Organization) SetCreated(v string) {
	o.Created = v
}

// GetIsActive returns the IsActive field value
func (o *Organization) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *Organization) SetIsActive(v bool) {
	o.IsActive = v
}

// GetModified returns the Modified field value
func (o *Organization) GetModified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *Organization) GetModifiedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *Organization) SetModified(v string) {
	o.Modified = v
}

// GetTierDisplayName returns the TierDisplayName field value
func (o *Organization) GetTierDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierDisplayName
}

// GetTierDisplayNameOk returns a tuple with the TierDisplayName field value
// and a boolean to check if the value has been set.
func (o *Organization) GetTierDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TierDisplayName, true
}

// SetTierDisplayName sets field value
func (o *Organization) SetTierDisplayName(v string) {
	o.TierDisplayName = v
}

// GetTierFamily returns the TierFamily field value
func (o *Organization) GetTierFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierFamily
}

// GetTierFamilyOk returns a tuple with the TierFamily field value
// and a boolean to check if the value has been set.
func (o *Organization) GetTierFamilyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TierFamily, true
}

// SetTierFamily sets field value
func (o *Organization) SetTierFamily(v string) {
	o.TierFamily = v
}

// GetTier returns the Tier field value
func (o *Organization) GetTier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *Organization) GetTierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *Organization) SetTier(v string) {
	o.Tier = v
}

// GetRole returns the Role field value
func (o *Organization) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Organization) GetRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Organization) SetRole(v string) {
	o.Role = v
}

// GetGuid returns the Guid field value
func (o *Organization) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *Organization) GetGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *Organization) SetGuid(v string) {
	o.Guid = v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["bsds"] = o.Bsds
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["is_active"] = o.IsActive
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if true {
		toSerialize["tier_display_name"] = o.TierDisplayName
	}
	if true {
		toSerialize["tier_family"] = o.TierFamily
	}
	if true {
		toSerialize["tier"] = o.Tier
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["guid"] = o.Guid
	}
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


