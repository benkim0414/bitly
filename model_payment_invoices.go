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

// PaymentInvoices struct for PaymentInvoices
type PaymentInvoices struct {
	PaymentInvoices *[]PaymentInvoice `json:"payment_invoices,omitempty"`
}

// NewPaymentInvoices instantiates a new PaymentInvoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInvoices() *PaymentInvoices {
	this := PaymentInvoices{}
	return &this
}

// NewPaymentInvoicesWithDefaults instantiates a new PaymentInvoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInvoicesWithDefaults() *PaymentInvoices {
	this := PaymentInvoices{}
	return &this
}

// GetPaymentInvoices returns the PaymentInvoices field value if set, zero value otherwise.
func (o *PaymentInvoices) GetPaymentInvoices() []PaymentInvoice {
	if o == nil || o.PaymentInvoices == nil {
		var ret []PaymentInvoice
		return ret
	}
	return *o.PaymentInvoices
}

// GetPaymentInvoicesOk returns a tuple with the PaymentInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInvoices) GetPaymentInvoicesOk() (*[]PaymentInvoice, bool) {
	if o == nil || o.PaymentInvoices == nil {
		return nil, false
	}
	return o.PaymentInvoices, true
}

// HasPaymentInvoices returns a boolean if a field has been set.
func (o *PaymentInvoices) HasPaymentInvoices() bool {
	if o != nil && o.PaymentInvoices != nil {
		return true
	}

	return false
}

// SetPaymentInvoices gets a reference to the given []PaymentInvoice and assigns it to the PaymentInvoices field.
func (o *PaymentInvoices) SetPaymentInvoices(v []PaymentInvoice) {
	o.PaymentInvoices = &v
}

func (o PaymentInvoices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentInvoices != nil {
		toSerialize["payment_invoices"] = o.PaymentInvoices
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInvoices struct {
	value *PaymentInvoices
	isSet bool
}

func (v NullablePaymentInvoices) Get() *PaymentInvoices {
	return v.value
}

func (v *NullablePaymentInvoices) Set(val *PaymentInvoices) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInvoices) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInvoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInvoices(val *PaymentInvoices) *NullablePaymentInvoices {
	return &NullablePaymentInvoices{value: val, isSet: true}
}

func (v NullablePaymentInvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


