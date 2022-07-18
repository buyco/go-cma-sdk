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

// AddressAllOf1 struct for AddressAllOf1
type AddressAllOf1 struct {
	// The name of the street of the party’s address.
	Street *string `json:"street,omitempty"`
}

// NewAddressAllOf1 instantiates a new AddressAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf1() *AddressAllOf1 {
	this := AddressAllOf1{}
	return &this
}

// NewAddressAllOf1WithDefaults instantiates a new AddressAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf1WithDefaults() *AddressAllOf1 {
	this := AddressAllOf1{}
	return &this
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *AddressAllOf1) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf1) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *AddressAllOf1) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *AddressAllOf1) SetStreet(v string) {
	o.Street = &v
}

func (o AddressAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf1 struct {
	value *AddressAllOf1
	isSet bool
}

func (v NullableAddressAllOf1) Get() *AddressAllOf1 {
	return v.value
}

func (v *NullableAddressAllOf1) Set(val *AddressAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf1(val *AddressAllOf1) *NullableAddressAllOf1 {
	return &NullableAddressAllOf1{value: val, isSet: true}
}

func (v NullableAddressAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

