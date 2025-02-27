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

// TransportCallAllOf2 struct for TransportCallAllOf2
type TransportCallAllOf2 struct {
	// The vessel operator-specific identifier of the Voyage.
	CarrierVoyageNumber *string `json:"carrierVoyageNumber,omitempty"`
}

// NewTransportCallAllOf2 instantiates a new TransportCallAllOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf2() *TransportCallAllOf2 {
	this := TransportCallAllOf2{}
	return &this
}

// NewTransportCallAllOf2WithDefaults instantiates a new TransportCallAllOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf2WithDefaults() *TransportCallAllOf2 {
	this := TransportCallAllOf2{}
	return &this
}

// GetCarrierVoyageNumber returns the CarrierVoyageNumber field value if set, zero value otherwise.
func (o *TransportCallAllOf2) GetCarrierVoyageNumber() string {
	if o == nil || o.CarrierVoyageNumber == nil {
		var ret string
		return ret
	}
	return *o.CarrierVoyageNumber
}

// GetCarrierVoyageNumberOk returns a tuple with the CarrierVoyageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf2) GetCarrierVoyageNumberOk() (*string, bool) {
	if o == nil || o.CarrierVoyageNumber == nil {
		return nil, false
	}
	return o.CarrierVoyageNumber, true
}

// HasCarrierVoyageNumber returns a boolean if a field has been set.
func (o *TransportCallAllOf2) HasCarrierVoyageNumber() bool {
	if o != nil && o.CarrierVoyageNumber != nil {
		return true
	}

	return false
}

// SetCarrierVoyageNumber gets a reference to the given string and assigns it to the CarrierVoyageNumber field.
func (o *TransportCallAllOf2) SetCarrierVoyageNumber(v string) {
	o.CarrierVoyageNumber = &v
}

func (o TransportCallAllOf2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CarrierVoyageNumber != nil {
		toSerialize["carrierVoyageNumber"] = o.CarrierVoyageNumber
	}
	return json.Marshal(toSerialize)
}

type NullableTransportCallAllOf2 struct {
	value *TransportCallAllOf2
	isSet bool
}

func (v NullableTransportCallAllOf2) Get() *TransportCallAllOf2 {
	return v.value
}

func (v *NullableTransportCallAllOf2) Set(val *TransportCallAllOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf2(val *TransportCallAllOf2) *NullableTransportCallAllOf2 {
	return &NullableTransportCallAllOf2{value: val, isSet: true}
}

func (v NullableTransportCallAllOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
