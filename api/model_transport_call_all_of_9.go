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

// TransportCallAllOf9 struct for TransportCallAllOf9
type TransportCallAllOf9 struct {
	FacilityCodeListProvider *FacilityCodeListProvider `json:"facilityCodeListProvider,omitempty"`
}

// NewTransportCallAllOf9 instantiates a new TransportCallAllOf9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf9() *TransportCallAllOf9 {
	this := TransportCallAllOf9{}
	return &this
}

// NewTransportCallAllOf9WithDefaults instantiates a new TransportCallAllOf9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf9WithDefaults() *TransportCallAllOf9 {
	this := TransportCallAllOf9{}
	return &this
}

// GetFacilityCodeListProvider returns the FacilityCodeListProvider field value if set, zero value otherwise.
func (o *TransportCallAllOf9) GetFacilityCodeListProvider() FacilityCodeListProvider {
	if o == nil || o.FacilityCodeListProvider == nil {
		var ret FacilityCodeListProvider
		return ret
	}
	return *o.FacilityCodeListProvider
}

// GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf9) GetFacilityCodeListProviderOk() (*FacilityCodeListProvider, bool) {
	if o == nil || o.FacilityCodeListProvider == nil {
		return nil, false
	}
	return o.FacilityCodeListProvider, true
}

// HasFacilityCodeListProvider returns a boolean if a field has been set.
func (o *TransportCallAllOf9) HasFacilityCodeListProvider() bool {
	if o != nil && o.FacilityCodeListProvider != nil {
		return true
	}

	return false
}

// SetFacilityCodeListProvider gets a reference to the given FacilityCodeListProvider and assigns it to the FacilityCodeListProvider field.
func (o *TransportCallAllOf9) SetFacilityCodeListProvider(v FacilityCodeListProvider) {
	o.FacilityCodeListProvider = &v
}

func (o TransportCallAllOf9) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FacilityCodeListProvider != nil {
		toSerialize["facilityCodeListProvider"] = o.FacilityCodeListProvider
	}
	return json.Marshal(toSerialize)
}

type NullableTransportCallAllOf9 struct {
	value *TransportCallAllOf9
	isSet bool
}

func (v NullableTransportCallAllOf9) Get() *TransportCallAllOf9 {
	return v.value
}

func (v *NullableTransportCallAllOf9) Set(val *TransportCallAllOf9) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf9) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf9(val *TransportCallAllOf9) *NullableTransportCallAllOf9 {
	return &NullableTransportCallAllOf9{value: val, isSet: true}
}

func (v NullableTransportCallAllOf9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
