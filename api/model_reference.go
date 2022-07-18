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

// Reference references provided by the shipper or freight forwarder at the time of booking or at the time of providing shipping instruction. Carriers share it back when providing track and trace event updates, some are also printed on the B/L. Customers can use these references to track shipments in their internal systems.
type Reference struct {
	ReferenceType ReferenceType `json:"referenceType"`
	// The actual value of the reference.
	ReferenceValue string `json:"referenceValue"`
}

// NewReference instantiates a new Reference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReference(referenceType ReferenceType, referenceValue string) *Reference {
	this := Reference{}
	this.ReferenceType = referenceType
	this.ReferenceValue = referenceValue
	return &this
}

// NewReferenceWithDefaults instantiates a new Reference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceWithDefaults() *Reference {
	this := Reference{}
	return &this
}

// GetReferenceType returns the ReferenceType field value
func (o *Reference) GetReferenceType() ReferenceType {
	if o == nil {
		var ret ReferenceType
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *Reference) GetReferenceTypeOk() (*ReferenceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *Reference) SetReferenceType(v ReferenceType) {
	o.ReferenceType = v
}

// GetReferenceValue returns the ReferenceValue field value
func (o *Reference) GetReferenceValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceValue
}

// GetReferenceValueOk returns a tuple with the ReferenceValue field value
// and a boolean to check if the value has been set.
func (o *Reference) GetReferenceValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceValue, true
}

// SetReferenceValue sets field value
func (o *Reference) SetReferenceValue(v string) {
	o.ReferenceValue = v
}

func (o Reference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["referenceType"] = o.ReferenceType
	}
	if true {
		toSerialize["referenceValue"] = o.ReferenceValue
	}
	return json.Marshal(toSerialize)
}

type NullableReference struct {
	value *Reference
	isSet bool
}

func (v NullableReference) Get() *Reference {
	return v.value
}

func (v *NullableReference) Set(val *Reference) {
	v.value = val
	v.isSet = true
}

func (v NullableReference) IsSet() bool {
	return v.isSet
}

func (v *NullableReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReference(val *Reference) *NullableReference {
	return &NullableReference{value: val, isSet: true}
}

func (v NullableReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
