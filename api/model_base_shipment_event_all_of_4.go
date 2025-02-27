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

// BaseShipmentEventAllOf4 struct for BaseShipmentEventAllOf4
type BaseShipmentEventAllOf4 struct {
	DocumentTypeCode *DocumentTypeCode `json:"documentTypeCode,omitempty"`
}

// NewBaseShipmentEventAllOf4 instantiates a new BaseShipmentEventAllOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShipmentEventAllOf4() *BaseShipmentEventAllOf4 {
	this := BaseShipmentEventAllOf4{}
	return &this
}

// NewBaseShipmentEventAllOf4WithDefaults instantiates a new BaseShipmentEventAllOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShipmentEventAllOf4WithDefaults() *BaseShipmentEventAllOf4 {
	this := BaseShipmentEventAllOf4{}
	return &this
}

// GetDocumentTypeCode returns the DocumentTypeCode field value if set, zero value otherwise.
func (o *BaseShipmentEventAllOf4) GetDocumentTypeCode() DocumentTypeCode {
	if o == nil || o.DocumentTypeCode == nil {
		var ret DocumentTypeCode
		return ret
	}
	return *o.DocumentTypeCode
}

// GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEventAllOf4) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool) {
	if o == nil || o.DocumentTypeCode == nil {
		return nil, false
	}
	return o.DocumentTypeCode, true
}

// HasDocumentTypeCode returns a boolean if a field has been set.
func (o *BaseShipmentEventAllOf4) HasDocumentTypeCode() bool {
	if o != nil && o.DocumentTypeCode != nil {
		return true
	}

	return false
}

// SetDocumentTypeCode gets a reference to the given DocumentTypeCode and assigns it to the DocumentTypeCode field.
func (o *BaseShipmentEventAllOf4) SetDocumentTypeCode(v DocumentTypeCode) {
	o.DocumentTypeCode = &v
}

func (o BaseShipmentEventAllOf4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentTypeCode != nil {
		toSerialize["documentTypeCode"] = o.DocumentTypeCode
	}
	return json.Marshal(toSerialize)
}

type NullableBaseShipmentEventAllOf4 struct {
	value *BaseShipmentEventAllOf4
	isSet bool
}

func (v NullableBaseShipmentEventAllOf4) Get() *BaseShipmentEventAllOf4 {
	return v.value
}

func (v *NullableBaseShipmentEventAllOf4) Set(val *BaseShipmentEventAllOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShipmentEventAllOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShipmentEventAllOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShipmentEventAllOf4(val *BaseShipmentEventAllOf4) *NullableBaseShipmentEventAllOf4 {
	return &NullableBaseShipmentEventAllOf4{value: val, isSet: true}
}

func (v NullableBaseShipmentEventAllOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShipmentEventAllOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
