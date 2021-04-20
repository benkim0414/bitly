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

// Deeplink struct for Deeplink
type Deeplink struct {
	AppUriPath *string `json:"app_uri_path,omitempty"`
	InstallType *string `json:"install_type,omitempty"`
	InstallUrl *string `json:"install_url,omitempty"`
	AppId *string `json:"app_id,omitempty"`
}

// NewDeeplink instantiates a new Deeplink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeeplink() *Deeplink {
	this := Deeplink{}
	return &this
}

// NewDeeplinkWithDefaults instantiates a new Deeplink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeeplinkWithDefaults() *Deeplink {
	this := Deeplink{}
	return &this
}

// GetAppUriPath returns the AppUriPath field value if set, zero value otherwise.
func (o *Deeplink) GetAppUriPath() string {
	if o == nil || o.AppUriPath == nil {
		var ret string
		return ret
	}
	return *o.AppUriPath
}

// GetAppUriPathOk returns a tuple with the AppUriPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deeplink) GetAppUriPathOk() (*string, bool) {
	if o == nil || o.AppUriPath == nil {
		return nil, false
	}
	return o.AppUriPath, true
}

// HasAppUriPath returns a boolean if a field has been set.
func (o *Deeplink) HasAppUriPath() bool {
	if o != nil && o.AppUriPath != nil {
		return true
	}

	return false
}

// SetAppUriPath gets a reference to the given string and assigns it to the AppUriPath field.
func (o *Deeplink) SetAppUriPath(v string) {
	o.AppUriPath = &v
}

// GetInstallType returns the InstallType field value if set, zero value otherwise.
func (o *Deeplink) GetInstallType() string {
	if o == nil || o.InstallType == nil {
		var ret string
		return ret
	}
	return *o.InstallType
}

// GetInstallTypeOk returns a tuple with the InstallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deeplink) GetInstallTypeOk() (*string, bool) {
	if o == nil || o.InstallType == nil {
		return nil, false
	}
	return o.InstallType, true
}

// HasInstallType returns a boolean if a field has been set.
func (o *Deeplink) HasInstallType() bool {
	if o != nil && o.InstallType != nil {
		return true
	}

	return false
}

// SetInstallType gets a reference to the given string and assigns it to the InstallType field.
func (o *Deeplink) SetInstallType(v string) {
	o.InstallType = &v
}

// GetInstallUrl returns the InstallUrl field value if set, zero value otherwise.
func (o *Deeplink) GetInstallUrl() string {
	if o == nil || o.InstallUrl == nil {
		var ret string
		return ret
	}
	return *o.InstallUrl
}

// GetInstallUrlOk returns a tuple with the InstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deeplink) GetInstallUrlOk() (*string, bool) {
	if o == nil || o.InstallUrl == nil {
		return nil, false
	}
	return o.InstallUrl, true
}

// HasInstallUrl returns a boolean if a field has been set.
func (o *Deeplink) HasInstallUrl() bool {
	if o != nil && o.InstallUrl != nil {
		return true
	}

	return false
}

// SetInstallUrl gets a reference to the given string and assigns it to the InstallUrl field.
func (o *Deeplink) SetInstallUrl(v string) {
	o.InstallUrl = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Deeplink) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deeplink) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Deeplink) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Deeplink) SetAppId(v string) {
	o.AppId = &v
}

func (o Deeplink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppUriPath != nil {
		toSerialize["app_uri_path"] = o.AppUriPath
	}
	if o.InstallType != nil {
		toSerialize["install_type"] = o.InstallType
	}
	if o.InstallUrl != nil {
		toSerialize["install_url"] = o.InstallUrl
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	return json.Marshal(toSerialize)
}

type NullableDeeplink struct {
	value *Deeplink
	isSet bool
}

func (v NullableDeeplink) Get() *Deeplink {
	return v.value
}

func (v *NullableDeeplink) Set(val *Deeplink) {
	v.value = val
	v.isSet = true
}

func (v NullableDeeplink) IsSet() bool {
	return v.isSet
}

func (v *NullableDeeplink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeeplink(val *Deeplink) *NullableDeeplink {
	return &NullableDeeplink{value: val, isSet: true}
}

func (v NullableDeeplink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeeplink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

