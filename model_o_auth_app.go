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

// OAuthApp struct for OAuthApp
type OAuthApp struct {
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewOAuthApp instantiates a new OAuthApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthApp() *OAuthApp {
	this := OAuthApp{}
	return &this
}

// NewOAuthAppWithDefaults instantiates a new OAuthApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthAppWithDefaults() *OAuthApp {
	this := OAuthApp{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *OAuthApp) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApp) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *OAuthApp) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *OAuthApp) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OAuthApp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OAuthApp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OAuthApp) SetName(v string) {
	o.Name = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuthApp) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApp) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuthApp) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuthApp) SetClientId(v string) {
	o.ClientId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OAuthApp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApp) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OAuthApp) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OAuthApp) SetDescription(v string) {
	o.Description = &v
}

func (o OAuthApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthApp struct {
	value *OAuthApp
	isSet bool
}

func (v NullableOAuthApp) Get() *OAuthApp {
	return v.value
}

func (v *NullableOAuthApp) Set(val *OAuthApp) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthApp) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthApp(val *OAuthApp) *NullableOAuthApp {
	return &NullableOAuthApp{value: val, isSet: true}
}

func (v NullableOAuthApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


