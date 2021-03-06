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

// Payments struct for Payments
type Payments struct {
	PaymentNumber *string `json:"payment_number,omitempty"`
	PaymentDate *string `json:"payment_date,omitempty"`
	PaymentAmount *float32 `json:"payment_amount,omitempty"`
}

// NewPayments instantiates a new Payments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayments() *Payments {
	this := Payments{}
	return &this
}

// NewPaymentsWithDefaults instantiates a new Payments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsWithDefaults() *Payments {
	this := Payments{}
	return &this
}

// GetPaymentNumber returns the PaymentNumber field value if set, zero value otherwise.
func (o *Payments) GetPaymentNumber() string {
	if o == nil || o.PaymentNumber == nil {
		var ret string
		return ret
	}
	return *o.PaymentNumber
}

// GetPaymentNumberOk returns a tuple with the PaymentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payments) GetPaymentNumberOk() (*string, bool) {
	if o == nil || o.PaymentNumber == nil {
		return nil, false
	}
	return o.PaymentNumber, true
}

// HasPaymentNumber returns a boolean if a field has been set.
func (o *Payments) HasPaymentNumber() bool {
	if o != nil && o.PaymentNumber != nil {
		return true
	}

	return false
}

// SetPaymentNumber gets a reference to the given string and assigns it to the PaymentNumber field.
func (o *Payments) SetPaymentNumber(v string) {
	o.PaymentNumber = &v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *Payments) GetPaymentDate() string {
	if o == nil || o.PaymentDate == nil {
		var ret string
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payments) GetPaymentDateOk() (*string, bool) {
	if o == nil || o.PaymentDate == nil {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *Payments) HasPaymentDate() bool {
	if o != nil && o.PaymentDate != nil {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given string and assigns it to the PaymentDate field.
func (o *Payments) SetPaymentDate(v string) {
	o.PaymentDate = &v
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise.
func (o *Payments) GetPaymentAmount() float32 {
	if o == nil || o.PaymentAmount == nil {
		var ret float32
		return ret
	}
	return *o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payments) GetPaymentAmountOk() (*float32, bool) {
	if o == nil || o.PaymentAmount == nil {
		return nil, false
	}
	return o.PaymentAmount, true
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *Payments) HasPaymentAmount() bool {
	if o != nil && o.PaymentAmount != nil {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given float32 and assigns it to the PaymentAmount field.
func (o *Payments) SetPaymentAmount(v float32) {
	o.PaymentAmount = &v
}

func (o Payments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentNumber != nil {
		toSerialize["payment_number"] = o.PaymentNumber
	}
	if o.PaymentDate != nil {
		toSerialize["payment_date"] = o.PaymentDate
	}
	if o.PaymentAmount != nil {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	return json.Marshal(toSerialize)
}

type NullablePayments struct {
	value *Payments
	isSet bool
}

func (v NullablePayments) Get() *Payments {
	return v.value
}

func (v *NullablePayments) Set(val *Payments) {
	v.value = val
	v.isSet = true
}

func (v NullablePayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayments(val *Payments) *NullablePayments {
	return &NullablePayments{value: val, isSet: true}
}

func (v NullablePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


