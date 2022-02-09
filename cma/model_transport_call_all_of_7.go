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

// TransportCallAllOf7 struct for TransportCallAllOf7
type TransportCallAllOf7 struct {
	// The UN Location code specifying where the place is located.
	UNLocationCode *string `json:"UNLocationCode,omitempty"`
}

// NewTransportCallAllOf7 instantiates a new TransportCallAllOf7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf7() *TransportCallAllOf7 {
	this := TransportCallAllOf7{}
	return &this
}

// NewTransportCallAllOf7WithDefaults instantiates a new TransportCallAllOf7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf7WithDefaults() *TransportCallAllOf7 {
	this := TransportCallAllOf7{}
	return &this
}

// GetUNLocationCode returns the UNLocationCode field value if set, zero value otherwise.
func (o *TransportCallAllOf7) GetUNLocationCode() string {
	if o == nil || o.UNLocationCode == nil {
		var ret string
		return ret
	}
	return *o.UNLocationCode
}

// GetUNLocationCodeOk returns a tuple with the UNLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf7) GetUNLocationCodeOk() (*string, bool) {
	if o == nil || o.UNLocationCode == nil {
		return nil, false
	}
	return o.UNLocationCode, true
}

// HasUNLocationCode returns a boolean if a field has been set.
func (o *TransportCallAllOf7) HasUNLocationCode() bool {
	if o != nil && o.UNLocationCode != nil {
		return true
	}

	return false
}

// SetUNLocationCode gets a reference to the given string and assigns it to the UNLocationCode field.
func (o *TransportCallAllOf7) SetUNLocationCode(v string) {
	o.UNLocationCode = &v
}

func (o TransportCallAllOf7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UNLocationCode != nil {
		toSerialize["UNLocationCode"] = o.UNLocationCode
	}
	return json.Marshal(toSerialize)
}

type NullableTransportCallAllOf7 struct {
	value *TransportCallAllOf7
	isSet bool
}

func (v NullableTransportCallAllOf7) Get() *TransportCallAllOf7 {
	return v.value
}

func (v *NullableTransportCallAllOf7) Set(val *TransportCallAllOf7) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf7) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf7(val *TransportCallAllOf7) *NullableTransportCallAllOf7 {
	return &NullableTransportCallAllOf7{value: val, isSet: true}
}

func (v NullableTransportCallAllOf7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


