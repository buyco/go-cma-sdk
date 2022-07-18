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

// TransportCallAllOf13 struct for TransportCallAllOf13
type TransportCallAllOf13 struct {
	Location NullableLocation `json:"location,omitempty"`
}

// NewTransportCallAllOf13 instantiates a new TransportCallAllOf13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf13() *TransportCallAllOf13 {
	this := TransportCallAllOf13{}
	return &this
}

// NewTransportCallAllOf13WithDefaults instantiates a new TransportCallAllOf13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf13WithDefaults() *TransportCallAllOf13 {
	this := TransportCallAllOf13{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransportCallAllOf13) GetLocation() Location {
	if o == nil || o.Location.Get() == nil {
		var ret Location
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransportCallAllOf13) GetLocationOk() (*Location, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *TransportCallAllOf13) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableLocation and assigns it to the Location field.
func (o *TransportCallAllOf13) SetLocation(v Location) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *TransportCallAllOf13) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *TransportCallAllOf13) UnsetLocation() {
	o.Location.Unset()
}

func (o TransportCallAllOf13) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransportCallAllOf13 struct {
	value *TransportCallAllOf13
	isSet bool
}

func (v NullableTransportCallAllOf13) Get() *TransportCallAllOf13 {
	return v.value
}

func (v *NullableTransportCallAllOf13) Set(val *TransportCallAllOf13) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf13) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf13(val *TransportCallAllOf13) *NullableTransportCallAllOf13 {
	return &NullableTransportCallAllOf13{value: val, isSet: true}
}

func (v NullableTransportCallAllOf13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

