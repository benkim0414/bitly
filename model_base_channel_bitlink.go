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

// BaseChannelBitlink struct for BaseChannelBitlink
type BaseChannelBitlink struct {
	CampaignGuid *string `json:"campaign_guid,omitempty"`
	BitlinkId *string `json:"bitlink_id,omitempty"`
}

// NewBaseChannelBitlink instantiates a new BaseChannelBitlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseChannelBitlink() *BaseChannelBitlink {
	this := BaseChannelBitlink{}
	return &this
}

// NewBaseChannelBitlinkWithDefaults instantiates a new BaseChannelBitlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseChannelBitlinkWithDefaults() *BaseChannelBitlink {
	this := BaseChannelBitlink{}
	return &this
}

// GetCampaignGuid returns the CampaignGuid field value if set, zero value otherwise.
func (o *BaseChannelBitlink) GetCampaignGuid() string {
	if o == nil || o.CampaignGuid == nil {
		var ret string
		return ret
	}
	return *o.CampaignGuid
}

// GetCampaignGuidOk returns a tuple with the CampaignGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseChannelBitlink) GetCampaignGuidOk() (*string, bool) {
	if o == nil || o.CampaignGuid == nil {
		return nil, false
	}
	return o.CampaignGuid, true
}

// HasCampaignGuid returns a boolean if a field has been set.
func (o *BaseChannelBitlink) HasCampaignGuid() bool {
	if o != nil && o.CampaignGuid != nil {
		return true
	}

	return false
}

// SetCampaignGuid gets a reference to the given string and assigns it to the CampaignGuid field.
func (o *BaseChannelBitlink) SetCampaignGuid(v string) {
	o.CampaignGuid = &v
}

// GetBitlinkId returns the BitlinkId field value if set, zero value otherwise.
func (o *BaseChannelBitlink) GetBitlinkId() string {
	if o == nil || o.BitlinkId == nil {
		var ret string
		return ret
	}
	return *o.BitlinkId
}

// GetBitlinkIdOk returns a tuple with the BitlinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseChannelBitlink) GetBitlinkIdOk() (*string, bool) {
	if o == nil || o.BitlinkId == nil {
		return nil, false
	}
	return o.BitlinkId, true
}

// HasBitlinkId returns a boolean if a field has been set.
func (o *BaseChannelBitlink) HasBitlinkId() bool {
	if o != nil && o.BitlinkId != nil {
		return true
	}

	return false
}

// SetBitlinkId gets a reference to the given string and assigns it to the BitlinkId field.
func (o *BaseChannelBitlink) SetBitlinkId(v string) {
	o.BitlinkId = &v
}

func (o BaseChannelBitlink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CampaignGuid != nil {
		toSerialize["campaign_guid"] = o.CampaignGuid
	}
	if o.BitlinkId != nil {
		toSerialize["bitlink_id"] = o.BitlinkId
	}
	return json.Marshal(toSerialize)
}

type NullableBaseChannelBitlink struct {
	value *BaseChannelBitlink
	isSet bool
}

func (v NullableBaseChannelBitlink) Get() *BaseChannelBitlink {
	return v.value
}

func (v *NullableBaseChannelBitlink) Set(val *BaseChannelBitlink) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseChannelBitlink) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseChannelBitlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseChannelBitlink(val *BaseChannelBitlink) *NullableBaseChannelBitlink {
	return &NullableBaseChannelBitlink{value: val, isSet: true}
}

func (v NullableBaseChannelBitlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseChannelBitlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


