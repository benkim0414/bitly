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

// GroupBitlinksCountRollup struct for GroupBitlinksCountRollup
type GroupBitlinksCountRollup struct {
	Units *int32 `json:"units,omitempty"`
	TotalClicks *int32 `json:"total_clicks,omitempty"`
	UnitReference *string `json:"unit_reference,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewGroupBitlinksCountRollup instantiates a new GroupBitlinksCountRollup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupBitlinksCountRollup() *GroupBitlinksCountRollup {
	this := GroupBitlinksCountRollup{}
	return &this
}

// NewGroupBitlinksCountRollupWithDefaults instantiates a new GroupBitlinksCountRollup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupBitlinksCountRollupWithDefaults() *GroupBitlinksCountRollup {
	this := GroupBitlinksCountRollup{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GroupBitlinksCountRollup) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBitlinksCountRollup) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GroupBitlinksCountRollup) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *GroupBitlinksCountRollup) SetUnits(v int32) {
	o.Units = &v
}

// GetTotalClicks returns the TotalClicks field value if set, zero value otherwise.
func (o *GroupBitlinksCountRollup) GetTotalClicks() int32 {
	if o == nil || o.TotalClicks == nil {
		var ret int32
		return ret
	}
	return *o.TotalClicks
}

// GetTotalClicksOk returns a tuple with the TotalClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBitlinksCountRollup) GetTotalClicksOk() (*int32, bool) {
	if o == nil || o.TotalClicks == nil {
		return nil, false
	}
	return o.TotalClicks, true
}

// HasTotalClicks returns a boolean if a field has been set.
func (o *GroupBitlinksCountRollup) HasTotalClicks() bool {
	if o != nil && o.TotalClicks != nil {
		return true
	}

	return false
}

// SetTotalClicks gets a reference to the given int32 and assigns it to the TotalClicks field.
func (o *GroupBitlinksCountRollup) SetTotalClicks(v int32) {
	o.TotalClicks = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *GroupBitlinksCountRollup) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBitlinksCountRollup) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *GroupBitlinksCountRollup) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *GroupBitlinksCountRollup) SetUnitReference(v string) {
	o.UnitReference = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GroupBitlinksCountRollup) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBitlinksCountRollup) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GroupBitlinksCountRollup) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GroupBitlinksCountRollup) SetUnit(v string) {
	o.Unit = &v
}

func (o GroupBitlinksCountRollup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.TotalClicks != nil {
		toSerialize["total_clicks"] = o.TotalClicks
	}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGroupBitlinksCountRollup struct {
	value *GroupBitlinksCountRollup
	isSet bool
}

func (v NullableGroupBitlinksCountRollup) Get() *GroupBitlinksCountRollup {
	return v.value
}

func (v *NullableGroupBitlinksCountRollup) Set(val *GroupBitlinksCountRollup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupBitlinksCountRollup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupBitlinksCountRollup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupBitlinksCountRollup(val *GroupBitlinksCountRollup) *NullableGroupBitlinksCountRollup {
	return &NullableGroupBitlinksCountRollup{value: val, isSet: true}
}

func (v NullableGroupBitlinksCountRollup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupBitlinksCountRollup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


