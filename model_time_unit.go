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
	"fmt"
)

// TimeUnit the unit of time queried for (minute, hour, day, week, month)
type TimeUnit string

// List of TimeUnit
const (
	MINUTE TimeUnit = "minute"
	HOUR TimeUnit = "hour"
	DAY TimeUnit = "day"
	WEEK TimeUnit = "week"
	MONTH TimeUnit = "month"
)

func (v *TimeUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeUnit(value)
	for _, existing := range []TimeUnit{ "minute", "hour", "day", "week", "month",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeUnit", value)
}

// Ptr returns reference to TimeUnit value
func (v TimeUnit) Ptr() *TimeUnit {
	return &v
}

type NullableTimeUnit struct {
	value *TimeUnit
	isSet bool
}

func (v NullableTimeUnit) Get() *TimeUnit {
	return v.value
}

func (v *NullableTimeUnit) Set(val *TimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeUnit(val *TimeUnit) *NullableTimeUnit {
	return &NullableTimeUnit{value: val, isSet: true}
}

func (v NullableTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

