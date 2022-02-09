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

// AddressAllOf3 struct for AddressAllOf3
type AddressAllOf3 struct {
	// The floor of the party’s street number.
	Floor *string `json:"floor,omitempty"`
}

// NewAddressAllOf3 instantiates a new AddressAllOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf3() *AddressAllOf3 {
	this := AddressAllOf3{}
	return &this
}

// NewAddressAllOf3WithDefaults instantiates a new AddressAllOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf3WithDefaults() *AddressAllOf3 {
	this := AddressAllOf3{}
	return &this
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *AddressAllOf3) GetFloor() string {
	if o == nil || o.Floor == nil {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf3) GetFloorOk() (*string, bool) {
	if o == nil || o.Floor == nil {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *AddressAllOf3) HasFloor() bool {
	if o != nil && o.Floor != nil {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *AddressAllOf3) SetFloor(v string) {
	o.Floor = &v
}

func (o AddressAllOf3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Floor != nil {
		toSerialize["floor"] = o.Floor
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf3 struct {
	value *AddressAllOf3
	isSet bool
}

func (v NullableAddressAllOf3) Get() *AddressAllOf3 {
	return v.value
}

func (v *NullableAddressAllOf3) Set(val *AddressAllOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf3(val *AddressAllOf3) *NullableAddressAllOf3 {
	return &NullableAddressAllOf3{value: val, isSet: true}
}

func (v NullableAddressAllOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


