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

// AddressAllOf6 struct for AddressAllOf6
type AddressAllOf6 struct {
	// The state/region of the party’s address.
	StateRegion *string `json:"stateRegion,omitempty"`
}

// NewAddressAllOf6 instantiates a new AddressAllOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf6() *AddressAllOf6 {
	this := AddressAllOf6{}
	return &this
}

// NewAddressAllOf6WithDefaults instantiates a new AddressAllOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf6WithDefaults() *AddressAllOf6 {
	this := AddressAllOf6{}
	return &this
}

// GetStateRegion returns the StateRegion field value if set, zero value otherwise.
func (o *AddressAllOf6) GetStateRegion() string {
	if o == nil || o.StateRegion == nil {
		var ret string
		return ret
	}
	return *o.StateRegion
}

// GetStateRegionOk returns a tuple with the StateRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf6) GetStateRegionOk() (*string, bool) {
	if o == nil || o.StateRegion == nil {
		return nil, false
	}
	return o.StateRegion, true
}

// HasStateRegion returns a boolean if a field has been set.
func (o *AddressAllOf6) HasStateRegion() bool {
	if o != nil && o.StateRegion != nil {
		return true
	}

	return false
}

// SetStateRegion gets a reference to the given string and assigns it to the StateRegion field.
func (o *AddressAllOf6) SetStateRegion(v string) {
	o.StateRegion = &v
}

func (o AddressAllOf6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StateRegion != nil {
		toSerialize["stateRegion"] = o.StateRegion
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf6 struct {
	value *AddressAllOf6
	isSet bool
}

func (v NullableAddressAllOf6) Get() *AddressAllOf6 {
	return v.value
}

func (v *NullableAddressAllOf6) Set(val *AddressAllOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf6(val *AddressAllOf6) *NullableAddressAllOf6 {
	return &NullableAddressAllOf6{value: val, isSet: true}
}

func (v NullableAddressAllOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

