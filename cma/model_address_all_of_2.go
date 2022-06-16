/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org. For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.0.3
Contact: ho.support-api@cma-cgm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cma

import (
	"encoding/json"
)

// AddressAllOf2 struct for AddressAllOf2
type AddressAllOf2 struct {
	// The number of the street of the party’s address.
	StreetNumber *string `json:"streetNumber,omitempty"`
}

// NewAddressAllOf2 instantiates a new AddressAllOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf2() *AddressAllOf2 {
	this := AddressAllOf2{}
	return &this
}

// NewAddressAllOf2WithDefaults instantiates a new AddressAllOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf2WithDefaults() *AddressAllOf2 {
	this := AddressAllOf2{}
	return &this
}

// GetStreetNumber returns the StreetNumber field value if set, zero value otherwise.
func (o *AddressAllOf2) GetStreetNumber() string {
	if o == nil || o.StreetNumber == nil {
		var ret string
		return ret
	}
	return *o.StreetNumber
}

// GetStreetNumberOk returns a tuple with the StreetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf2) GetStreetNumberOk() (*string, bool) {
	if o == nil || o.StreetNumber == nil {
		return nil, false
	}
	return o.StreetNumber, true
}

// HasStreetNumber returns a boolean if a field has been set.
func (o *AddressAllOf2) HasStreetNumber() bool {
	if o != nil && o.StreetNumber != nil {
		return true
	}

	return false
}

// SetStreetNumber gets a reference to the given string and assigns it to the StreetNumber field.
func (o *AddressAllOf2) SetStreetNumber(v string) {
	o.StreetNumber = &v
}

func (o AddressAllOf2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StreetNumber != nil {
		toSerialize["streetNumber"] = o.StreetNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf2 struct {
	value *AddressAllOf2
	isSet bool
}

func (v NullableAddressAllOf2) Get() *AddressAllOf2 {
	return v.value
}

func (v *NullableAddressAllOf2) Set(val *AddressAllOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf2(val *AddressAllOf2) *NullableAddressAllOf2 {
	return &NullableAddressAllOf2{value: val, isSet: true}
}

func (v NullableAddressAllOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


