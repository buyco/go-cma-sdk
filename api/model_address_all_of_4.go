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

// AddressAllOf4 struct for AddressAllOf4
type AddressAllOf4 struct {
	// The post code of the party’s address.
	PostCode *string `json:"postCode,omitempty"`
}

// NewAddressAllOf4 instantiates a new AddressAllOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf4() *AddressAllOf4 {
	this := AddressAllOf4{}
	return &this
}

// NewAddressAllOf4WithDefaults instantiates a new AddressAllOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOf4WithDefaults() *AddressAllOf4 {
	this := AddressAllOf4{}
	return &this
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *AddressAllOf4) GetPostCode() string {
	if o == nil || o.PostCode == nil {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf4) GetPostCodeOk() (*string, bool) {
	if o == nil || o.PostCode == nil {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *AddressAllOf4) HasPostCode() bool {
	if o != nil && o.PostCode != nil {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *AddressAllOf4) SetPostCode(v string) {
	o.PostCode = &v
}

func (o AddressAllOf4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PostCode != nil {
		toSerialize["postCode"] = o.PostCode
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf4 struct {
	value *AddressAllOf4
	isSet bool
}

func (v NullableAddressAllOf4) Get() *AddressAllOf4 {
	return v.value
}

func (v *NullableAddressAllOf4) Set(val *AddressAllOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf4(val *AddressAllOf4) *NullableAddressAllOf4 {
	return &NullableAddressAllOf4{value: val, isSet: true}
}

func (v NullableAddressAllOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

