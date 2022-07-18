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

// TransportEventAllOf struct for TransportEventAllOf
type TransportEventAllOf struct {
	// An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The <b>documentReferenceType</b>-field is used to describe where the <b>documentReferenceValue</b>-field is pointing to.
	DocumentReferences []DocumentReferencesInner `json:"documentReferences,omitempty"`
}

// NewTransportEventAllOf instantiates a new TransportEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportEventAllOf() *TransportEventAllOf {
	this := TransportEventAllOf{}
	return &this
}

// NewTransportEventAllOfWithDefaults instantiates a new TransportEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportEventAllOfWithDefaults() *TransportEventAllOf {
	this := TransportEventAllOf{}
	return &this
}

// GetDocumentReferences returns the DocumentReferences field value if set, zero value otherwise.
func (o *TransportEventAllOf) GetDocumentReferences() []DocumentReferencesInner {
	if o == nil || o.DocumentReferences == nil {
		var ret []DocumentReferencesInner
		return ret
	}
	return o.DocumentReferences
}

// GetDocumentReferencesOk returns a tuple with the DocumentReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEventAllOf) GetDocumentReferencesOk() ([]DocumentReferencesInner, bool) {
	if o == nil || o.DocumentReferences == nil {
		return nil, false
	}
	return o.DocumentReferences, true
}

// HasDocumentReferences returns a boolean if a field has been set.
func (o *TransportEventAllOf) HasDocumentReferences() bool {
	if o != nil && o.DocumentReferences != nil {
		return true
	}

	return false
}

// SetDocumentReferences gets a reference to the given []DocumentReferencesInner and assigns it to the DocumentReferences field.
func (o *TransportEventAllOf) SetDocumentReferences(v []DocumentReferencesInner) {
	o.DocumentReferences = v
}

func (o TransportEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentReferences != nil {
		toSerialize["documentReferences"] = o.DocumentReferences
	}
	return json.Marshal(toSerialize)
}

type NullableTransportEventAllOf struct {
	value *TransportEventAllOf
	isSet bool
}

func (v NullableTransportEventAllOf) Get() *TransportEventAllOf {
	return v.value
}

func (v *NullableTransportEventAllOf) Set(val *TransportEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportEventAllOf(val *TransportEventAllOf) *NullableTransportEventAllOf {
	return &NullableTransportEventAllOf{value: val, isSet: true}
}

func (v NullableTransportEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
