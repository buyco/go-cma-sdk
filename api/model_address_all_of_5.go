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

// AddressAllOf5 struct for AddressAllOf5
type AddressAllOf5 struct {
	// The city name of the party’s address.
	City *string `json:"city,omitempty"`
}

// NewAddressAllOf5 instantiates a new AddressAllOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf5() *AddressAllOf5 {
	this := AddressAllOf5{}
	return &this
}

// NewAddressAllOf5WithDefaults instantiates a new AddressAllOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf5WithDefaults() *AddressAllOf5 {
	this := AddressAllOf5{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressAllOf5) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf5) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressAllOf5) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressAllOf5) SetCity(v string) {
	o.City = &v
}

func (o AddressAllOf5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf5 struct {
	value *AddressAllOf5
	isSet bool
}

func (v NullableAddressAllOf5) Get() *AddressAllOf5 {
	return v.value
}

func (v *NullableAddressAllOf5) Set(val *AddressAllOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf5(val *AddressAllOf5) *NullableAddressAllOf5 {
	return &NullableAddressAllOf5{value: val, isSet: true}
}

func (v NullableAddressAllOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

