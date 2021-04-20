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

// ConsumableFeatures struct for ConsumableFeatures
type ConsumableFeatures struct {
	ConsumableFeatures *[]ConsumableFeature `json:"consumable_features,omitempty"`
}

// NewConsumableFeatures instantiates a new ConsumableFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumableFeatures() *ConsumableFeatures {
	this := ConsumableFeatures{}
	return &this
}

// NewConsumableFeaturesWithDefaults instantiates a new ConsumableFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumableFeaturesWithDefaults() *ConsumableFeatures {
	this := ConsumableFeatures{}
	return &this
}

// GetConsumableFeatures returns the ConsumableFeatures field value if set, zero value otherwise.
func (o *ConsumableFeatures) GetConsumableFeatures() []ConsumableFeature {
	if o == nil || o.ConsumableFeatures == nil {
		var ret []ConsumableFeature
		return ret
	}
	return *o.ConsumableFeatures
}

// GetConsumableFeaturesOk returns a tuple with the ConsumableFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumableFeatures) GetConsumableFeaturesOk() (*[]ConsumableFeature, bool) {
	if o == nil || o.ConsumableFeatures == nil {
		return nil, false
	}
	return o.ConsumableFeatures, true
}

// HasConsumableFeatures returns a boolean if a field has been set.
func (o *ConsumableFeatures) HasConsumableFeatures() bool {
	if o != nil && o.ConsumableFeatures != nil {
		return true
	}

	return false
}

// SetConsumableFeatures gets a reference to the given []ConsumableFeature and assigns it to the ConsumableFeatures field.
func (o *ConsumableFeatures) SetConsumableFeatures(v []ConsumableFeature) {
	o.ConsumableFeatures = &v
}

func (o ConsumableFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumableFeatures != nil {
		toSerialize["consumable_features"] = o.ConsumableFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableConsumableFeatures struct {
	value *ConsumableFeatures
	isSet bool
}

func (v NullableConsumableFeatures) Get() *ConsumableFeatures {
	return v.value
}

func (v *NullableConsumableFeatures) Set(val *ConsumableFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumableFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumableFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumableFeatures(val *ConsumableFeatures) *NullableConsumableFeatures {
	return &NullableConsumableFeatures{value: val, isSet: true}
}

func (v NullableConsumableFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumableFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


