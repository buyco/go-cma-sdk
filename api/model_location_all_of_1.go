/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org. For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.0.3
Contact: ho.support-api@cma-cgm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LocationAllOf1 struct for LocationAllOf1
type LocationAllOf1 struct {
	// Geographic coordinate that specifies the north–south position of a point on the Earth&apos;s surface.
	Latitude *string `json:"latitude,omitempty"`
}

// NewLocationAllOf1 instantiates a new LocationAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAllOf1() *LocationAllOf1 {
	this := LocationAllOf1{}
	return &this
}

// NewLocationAllOf1WithDefaults instantiates a new LocationAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAllOf1WithDefaults() *LocationAllOf1 {
	this := LocationAllOf1{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *LocationAllOf1) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAllOf1) GetLatitudeOk() (*string, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *LocationAllOf1) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *LocationAllOf1) SetLatitude(v string) {
	o.Latitude = &v
}

func (o LocationAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	return json.Marshal(toSerialize)
}

type NullableLocationAllOf1 struct {
	value *LocationAllOf1
	isSet bool
}

func (v NullableLocationAllOf1) Get() *LocationAllOf1 {
	return v.value
}

func (v *NullableLocationAllOf1) Set(val *LocationAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAllOf1(val *LocationAllOf1) *NullableLocationAllOf1 {
	return &NullableLocationAllOf1{value: val, isSet: true}
}

func (v NullableLocationAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
