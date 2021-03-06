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

// ChannelBitlinks struct for ChannelBitlinks
type ChannelBitlinks struct {
	Total *int32 `json:"total,omitempty"`
	Bitlinks *[]ChannelBitlink `json:"bitlinks,omitempty"`
}

// NewChannelBitlinks instantiates a new ChannelBitlinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelBitlinks() *ChannelBitlinks {
	this := ChannelBitlinks{}
	return &this
}

// NewChannelBitlinksWithDefaults instantiates a new ChannelBitlinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelBitlinksWithDefaults() *ChannelBitlinks {
	this := ChannelBitlinks{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ChannelBitlinks) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBitlinks) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ChannelBitlinks) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ChannelBitlinks) SetTotal(v int32) {
	o.Total = &v
}

// GetBitlinks returns the Bitlinks field value if set, zero value otherwise.
func (o *ChannelBitlinks) GetBitlinks() []ChannelBitlink {
	if o == nil || o.Bitlinks == nil {
		var ret []ChannelBitlink
		return ret
	}
	return *o.Bitlinks
}

// GetBitlinksOk returns a tuple with the Bitlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBitlinks) GetBitlinksOk() (*[]ChannelBitlink, bool) {
	if o == nil || o.Bitlinks == nil {
		return nil, false
	}
	return o.Bitlinks, true
}

// HasBitlinks returns a boolean if a field has been set.
func (o *ChannelBitlinks) HasBitlinks() bool {
	if o != nil && o.Bitlinks != nil {
		return true
	}

	return false
}

// SetBitlinks gets a reference to the given []ChannelBitlink and assigns it to the Bitlinks field.
func (o *ChannelBitlinks) SetBitlinks(v []ChannelBitlink) {
	o.Bitlinks = &v
}

func (o ChannelBitlinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Bitlinks != nil {
		toSerialize["bitlinks"] = o.Bitlinks
	}
	return json.Marshal(toSerialize)
}

type NullableChannelBitlinks struct {
	value *ChannelBitlinks
	isSet bool
}

func (v NullableChannelBitlinks) Get() *ChannelBitlinks {
	return v.value
}

func (v *NullableChannelBitlinks) Set(val *ChannelBitlinks) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelBitlinks) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelBitlinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelBitlinks(val *ChannelBitlinks) *NullableChannelBitlinks {
	return &NullableChannelBitlinks{value: val, isSet: true}
}

func (v NullableChannelBitlinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelBitlinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


