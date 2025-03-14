/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org.  <i>Please note that shipment events and subscription management are not covered yet by CMA CGM API.</i>  <br> This API is accessible through <ul> <li> <b> Public </b> connection (authentication with API Key): this allows to retrieve standard equipment moves (e.g. ready to be loaded, discharged), Transhipment moves (e.g. discharged and re-loaded onboard) and planned vessel departure and arrival dates. </li> <li> <b> Private </b> connection (authentication with Oauth2 client credentials): this allows to retrieve additional private events such as actual rail and ramp moves, planned departure and arrival dates for inland moves. These private events are available only if you (or your end customer) are identified on the booking as: Booking Party, Deciding Party, Consignee, Notifier or Shipper. </li> </ul> </br> For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DocumentReferencesInner struct for DocumentReferencesInner
type DocumentReferencesInner struct {
	// Describes where the documentReferenceValue is pointing to
	DocumentReferenceType *string `json:"documentReferenceType,omitempty"`
	// The value of the identifier the documentReferenceType is describing
	DocumentReferenceValue *string `json:"documentReferenceValue,omitempty"`
}

// NewDocumentReferencesInner instantiates a new DocumentReferencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentReferencesInner() *DocumentReferencesInner {
	this := DocumentReferencesInner{}
	return &this
}

// NewDocumentReferencesInnerWithDefaults instantiates a new DocumentReferencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentReferencesInnerWithDefaults() *DocumentReferencesInner {
	this := DocumentReferencesInner{}
	return &this
}

// GetDocumentReferenceType returns the DocumentReferenceType field value if set, zero value otherwise.
func (o *DocumentReferencesInner) GetDocumentReferenceType() string {
	if o == nil || o.DocumentReferenceType == nil {
		var ret string
		return ret
	}
	return *o.DocumentReferenceType
}

// GetDocumentReferenceTypeOk returns a tuple with the DocumentReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferencesInner) GetDocumentReferenceTypeOk() (*string, bool) {
	if o == nil || o.DocumentReferenceType == nil {
		return nil, false
	}
	return o.DocumentReferenceType, true
}

// HasDocumentReferenceType returns a boolean if a field has been set.
func (o *DocumentReferencesInner) HasDocumentReferenceType() bool {
	if o != nil && o.DocumentReferenceType != nil {
		return true
	}

	return false
}

// SetDocumentReferenceType gets a reference to the given string and assigns it to the DocumentReferenceType field.
func (o *DocumentReferencesInner) SetDocumentReferenceType(v string) {
	o.DocumentReferenceType = &v
}

// GetDocumentReferenceValue returns the DocumentReferenceValue field value if set, zero value otherwise.
func (o *DocumentReferencesInner) GetDocumentReferenceValue() string {
	if o == nil || o.DocumentReferenceValue == nil {
		var ret string
		return ret
	}
	return *o.DocumentReferenceValue
}

// GetDocumentReferenceValueOk returns a tuple with the DocumentReferenceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReferencesInner) GetDocumentReferenceValueOk() (*string, bool) {
	if o == nil || o.DocumentReferenceValue == nil {
		return nil, false
	}
	return o.DocumentReferenceValue, true
}

// HasDocumentReferenceValue returns a boolean if a field has been set.
func (o *DocumentReferencesInner) HasDocumentReferenceValue() bool {
	if o != nil && o.DocumentReferenceValue != nil {
		return true
	}

	return false
}

// SetDocumentReferenceValue gets a reference to the given string and assigns it to the DocumentReferenceValue field.
func (o *DocumentReferencesInner) SetDocumentReferenceValue(v string) {
	o.DocumentReferenceValue = &v
}

func (o DocumentReferencesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentReferenceType != nil {
		toSerialize["documentReferenceType"] = o.DocumentReferenceType
	}
	if o.DocumentReferenceValue != nil {
		toSerialize["documentReferenceValue"] = o.DocumentReferenceValue
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentReferencesInner struct {
	value *DocumentReferencesInner
	isSet bool
}

func (v NullableDocumentReferencesInner) Get() *DocumentReferencesInner {
	return v.value
}

func (v *NullableDocumentReferencesInner) Set(val *DocumentReferencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentReferencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentReferencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentReferencesInner(val *DocumentReferencesInner) *NullableDocumentReferencesInner {
	return &NullableDocumentReferencesInner{value: val, isSet: true}
}

func (v NullableDocumentReferencesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentReferencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
