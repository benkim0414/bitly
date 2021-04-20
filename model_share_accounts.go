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

// ShareAccounts struct for ShareAccounts
type ShareAccounts struct {
	ShareAccounts *[]ShareAccount `json:"share_accounts,omitempty"`
}

// NewShareAccounts instantiates a new ShareAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareAccounts() *ShareAccounts {
	this := ShareAccounts{}
	return &this
}

// NewShareAccountsWithDefaults instantiates a new ShareAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareAccountsWithDefaults() *ShareAccounts {
	this := ShareAccounts{}
	return &this
}

// GetShareAccounts returns the ShareAccounts field value if set, zero value otherwise.
func (o *ShareAccounts) GetShareAccounts() []ShareAccount {
	if o == nil || o.ShareAccounts == nil {
		var ret []ShareAccount
		return ret
	}
	return *o.ShareAccounts
}

// GetShareAccountsOk returns a tuple with the ShareAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareAccounts) GetShareAccountsOk() (*[]ShareAccount, bool) {
	if o == nil || o.ShareAccounts == nil {
		return nil, false
	}
	return o.ShareAccounts, true
}

// HasShareAccounts returns a boolean if a field has been set.
func (o *ShareAccounts) HasShareAccounts() bool {
	if o != nil && o.ShareAccounts != nil {
		return true
	}

	return false
}

// SetShareAccounts gets a reference to the given []ShareAccount and assigns it to the ShareAccounts field.
func (o *ShareAccounts) SetShareAccounts(v []ShareAccount) {
	o.ShareAccounts = &v
}

func (o ShareAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShareAccounts != nil {
		toSerialize["share_accounts"] = o.ShareAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableShareAccounts struct {
	value *ShareAccounts
	isSet bool
}

func (v NullableShareAccounts) Get() *ShareAccounts {
	return v.value
}

func (v *NullableShareAccounts) Set(val *ShareAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableShareAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableShareAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareAccounts(val *ShareAccounts) *NullableShareAccounts {
	return &NullableShareAccounts{value: val, isSet: true}
}

func (v NullableShareAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


