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

// TransportCallAllOf struct for TransportCallAllOf
type TransportCallAllOf struct {
	// The unique identifier for a transport call
	TransportCallID *string `json:"transportCallID,omitempty"`
}

// NewTransportCallAllOf instantiates a new TransportCallAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf() *TransportCallAllOf {
	this := TransportCallAllOf{}
	return &this
}

// NewTransportCallAllOfWithDefaults instantiates a new TransportCallAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOfWithDefaults() *TransportCallAllOf {
	this := TransportCallAllOf{}
	return &this
}

// GetTransportCallID returns the TransportCallID field value if set, zero value otherwise.
func (o *TransportCallAllOf) GetTransportCallID() string {
	if o == nil || o.TransportCallID == nil {
		var ret string
		return ret
	}
	return *o.TransportCallID
}

// GetTransportCallIDOk returns a tuple with the TransportCallID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf) GetTransportCallIDOk() (*string, bool) {
	if o == nil || o.TransportCallID == nil {
		return nil, false
	}
	return o.TransportCallID, true
}

// HasTransportCallID returns a boolean if a field has been set.
func (o *TransportCallAllOf) HasTransportCallID() bool {
	if o != nil && o.TransportCallID != nil {
		return true
	}

	return false
}

// SetTransportCallID gets a reference to the given string and assigns it to the TransportCallID field.
func (o *TransportCallAllOf) SetTransportCallID(v string) {
	o.TransportCallID = &v
}

func (o TransportCallAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransportCallID != nil {
		toSerialize["transportCallID"] = o.TransportCallID
	}
	return json.Marshal(toSerialize)
}

type NullableTransportCallAllOf struct {
	value *TransportCallAllOf
	isSet bool
}

func (v NullableTransportCallAllOf) Get() *TransportCallAllOf {
	return v.value
}

func (v *NullableTransportCallAllOf) Set(val *TransportCallAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf(val *TransportCallAllOf) *NullableTransportCallAllOf {
	return &NullableTransportCallAllOf{value: val, isSet: true}
}

func (v NullableTransportCallAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
