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

// SubErrorsInner struct for SubErrorsInner
type SubErrorsInner struct {
	// High level error message.
	Reason string `json:"reason"`
	// Detailed error message.
	Message string `json:"message"`
}

// NewSubErrorsInner instantiates a new SubErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubErrorsInner(reason string, message string) *SubErrorsInner {
	this := SubErrorsInner{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewSubErrorsInnerWithDefaults instantiates a new SubErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubErrorsInnerWithDefaults() *SubErrorsInner {
	this := SubErrorsInner{}
	return &this
}

// GetReason returns the Reason field value
func (o *SubErrorsInner) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SubErrorsInner) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SubErrorsInner) SetReason(v string) {
	o.Reason = v
}

// GetMessage returns the Message field value
func (o *SubErrorsInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SubErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SubErrorsInner) SetMessage(v string) {
	o.Message = v
}

func (o SubErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableSubErrorsInner struct {
	value *SubErrorsInner
	isSet bool
}

func (v NullableSubErrorsInner) Get() *SubErrorsInner {
	return v.value
}

func (v *NullableSubErrorsInner) Set(val *SubErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubErrorsInner(val *SubErrorsInner) *NullableSubErrorsInner {
	return &NullableSubErrorsInner{value: val, isSet: true}
}

func (v NullableSubErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
