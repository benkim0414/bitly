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

// Share struct for Share
type Share struct {
	NumericId *int32 `json:"numeric_id,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	Text *string `json:"text,omitempty"`
	Link *string `json:"link,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
}

// NewShare instantiates a new Share object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShare() *Share {
	this := Share{}
	return &this
}

// NewShareWithDefaults instantiates a new Share object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareWithDefaults() *Share {
	this := Share{}
	return &this
}

// GetNumericId returns the NumericId field value if set, zero value otherwise.
func (o *Share) GetNumericId() int32 {
	if o == nil || o.NumericId == nil {
		var ret int32
		return ret
	}
	return *o.NumericId
}

// GetNumericIdOk returns a tuple with the NumericId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetNumericIdOk() (*int32, bool) {
	if o == nil || o.NumericId == nil {
		return nil, false
	}
	return o.NumericId, true
}

// HasNumericId returns a boolean if a field has been set.
func (o *Share) HasNumericId() bool {
	if o != nil && o.NumericId != nil {
		return true
	}

	return false
}

// SetNumericId gets a reference to the given int32 and assigns it to the NumericId field.
func (o *Share) SetNumericId(v int32) {
	o.NumericId = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *Share) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *Share) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *Share) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Share) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Share) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Share) SetAccountId(v string) {
	o.AccountId = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Share) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Share) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Share) SetText(v string) {
	o.Text = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Share) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Share) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *Share) SetLink(v string) {
	o.Link = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *Share) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *Share) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *Share) SetAccountName(v string) {
	o.AccountName = &v
}

func (o Share) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumericId != nil {
		toSerialize["numeric_id"] = o.NumericId
	}
	if o.AccountType != nil {
		toSerialize["account_type"] = o.AccountType
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	return json.Marshal(toSerialize)
}

type NullableShare struct {
	value *Share
	isSet bool
}

func (v NullableShare) Get() *Share {
	return v.value
}

func (v *NullableShare) Set(val *Share) {
	v.value = val
	v.isSet = true
}

func (v NullableShare) IsSet() bool {
	return v.isSet
}

func (v *NullableShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShare(val *Share) *NullableShare {
	return &NullableShare{value: val, isSet: true}
}

func (v NullableShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

