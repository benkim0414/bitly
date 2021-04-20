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

// BillingAddress struct for BillingAddress
type BillingAddress struct {
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	State *string `json:"state,omitempty"`
	Address1 *string `json:"address_1,omitempty"`
	Address2 *string `json:"address_2,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
}

// NewBillingAddress instantiates a new BillingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddress() *BillingAddress {
	this := BillingAddress{}
	return &this
}

// NewBillingAddressWithDefaults instantiates a new BillingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressWithDefaults() *BillingAddress {
	this := BillingAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *BillingAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *BillingAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *BillingAddress) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *BillingAddress) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *BillingAddress) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *BillingAddress) SetCountry(v string) {
	o.Country = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BillingAddress) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BillingAddress) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BillingAddress) SetState(v string) {
	o.State = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *BillingAddress) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *BillingAddress) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *BillingAddress) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *BillingAddress) GetAddress2() string {
	if o == nil || o.Address2 == nil {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || o.Address2 == nil {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *BillingAddress) HasAddress2() bool {
	if o != nil && o.Address2 != nil {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *BillingAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *BillingAddress) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *BillingAddress) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *BillingAddress) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o BillingAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Address1 != nil {
		toSerialize["address_1"] = o.Address1
	}
	if o.Address2 != nil {
		toSerialize["address_2"] = o.Address2
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableBillingAddress struct {
	value *BillingAddress
	isSet bool
}

func (v NullableBillingAddress) Get() *BillingAddress {
	return v.value
}

func (v *NullableBillingAddress) Set(val *BillingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingAddress(val *BillingAddress) *NullableBillingAddress {
	return &NullableBillingAddress{value: val, isSet: true}
}

func (v NullableBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

