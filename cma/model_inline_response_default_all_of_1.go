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

// InlineResponseDefaultAllOf1 struct for InlineResponseDefaultAllOf1
type InlineResponseDefaultAllOf1 struct {
	RequestUri interface{} `json:"requestUri,omitempty"`
}

// NewInlineResponseDefaultAllOf1 instantiates a new InlineResponseDefaultAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponseDefaultAllOf1() *InlineResponseDefaultAllOf1 {
	this := InlineResponseDefaultAllOf1{}
	return &this
}

// NewInlineResponseDefaultAllOf1WithDefaults instantiates a new InlineResponseDefaultAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponseDefaultAllOf1WithDefaults() *InlineResponseDefaultAllOf1 {
	this := InlineResponseDefaultAllOf1{}
	return &this
}

// GetRequestUri returns the RequestUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponseDefaultAllOf1) GetRequestUri() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequestUri
}

// GetRequestUriOk returns a tuple with the RequestUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponseDefaultAllOf1) GetRequestUriOk() (*interface{}, bool) {
	if o == nil || o.RequestUri == nil {
		return nil, false
	}
	return &o.RequestUri, true
}

// HasRequestUri returns a boolean if a field has been set.
func (o *InlineResponseDefaultAllOf1) HasRequestUri() bool {
	if o != nil && o.RequestUri != nil {
		return true
	}

	return false
}

// SetRequestUri gets a reference to the given interface{} and assigns it to the RequestUri field.
func (o *InlineResponseDefaultAllOf1) SetRequestUri(v interface{}) {
	o.RequestUri = v
}

func (o InlineResponseDefaultAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestUri != nil {
		toSerialize["requestUri"] = o.RequestUri
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponseDefaultAllOf1 struct {
	value *InlineResponseDefaultAllOf1
	isSet bool
}

func (v NullableInlineResponseDefaultAllOf1) Get() *InlineResponseDefaultAllOf1 {
	return v.value
}

func (v *NullableInlineResponseDefaultAllOf1) Set(val *InlineResponseDefaultAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponseDefaultAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponseDefaultAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponseDefaultAllOf1(val *InlineResponseDefaultAllOf1) *NullableInlineResponseDefaultAllOf1 {
	return &NullableInlineResponseDefaultAllOf1{value: val, isSet: true}
}

func (v NullableInlineResponseDefaultAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponseDefaultAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


