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

// OrganizationUpdate struct for OrganizationUpdate
type OrganizationUpdate struct {
	Name string `json:"name"`
}

// NewOrganizationUpdate instantiates a new OrganizationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUpdate(name string) *OrganizationUpdate {
	this := OrganizationUpdate{}
	this.Name = name
	return &this
}

// NewOrganizationUpdateWithDefaults instantiates a new OrganizationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUpdateWithDefaults() *OrganizationUpdate {
	this := OrganizationUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationUpdate) SetName(v string) {
	o.Name = v
}

func (o OrganizationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUpdate struct {
	value *OrganizationUpdate
	isSet bool
}

func (v NullableOrganizationUpdate) Get() *OrganizationUpdate {
	return v.value
}

func (v *NullableOrganizationUpdate) Set(val *OrganizationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUpdate(val *OrganizationUpdate) *NullableOrganizationUpdate {
	return &NullableOrganizationUpdate{value: val, isSet: true}
}

func (v NullableOrganizationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


