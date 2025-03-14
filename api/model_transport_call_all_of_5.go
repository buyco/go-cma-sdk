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

// TransportCallAllOf5 struct for TransportCallAllOf5
type TransportCallAllOf5 struct {
	// The vessel operator-specific identifier of the import Voyage.
	ImportVoyageNumber *string `json:"importVoyageNumber,omitempty"`
}

// NewTransportCallAllOf5 instantiates a new TransportCallAllOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportCallAllOf5() *TransportCallAllOf5 {
	this := TransportCallAllOf5{}
	return &this
}

// NewTransportCallAllOf5WithDefaults instantiates a new TransportCallAllOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportCallAllOf5WithDefaults() *TransportCallAllOf5 {
	this := TransportCallAllOf5{}
	return &this
}

// GetImportVoyageNumber returns the ImportVoyageNumber field value if set, zero value otherwise.
func (o *TransportCallAllOf5) GetImportVoyageNumber() string {
	if o == nil || o.ImportVoyageNumber == nil {
		var ret string
		return ret
	}
	return *o.ImportVoyageNumber
}

// GetImportVoyageNumberOk returns a tuple with the ImportVoyageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportCallAllOf5) GetImportVoyageNumberOk() (*string, bool) {
	if o == nil || o.ImportVoyageNumber == nil {
		return nil, false
	}
	return o.ImportVoyageNumber, true
}

// HasImportVoyageNumber returns a boolean if a field has been set.
func (o *TransportCallAllOf5) HasImportVoyageNumber() bool {
	if o != nil && o.ImportVoyageNumber != nil {
		return true
	}

	return false
}

// SetImportVoyageNumber gets a reference to the given string and assigns it to the ImportVoyageNumber field.
func (o *TransportCallAllOf5) SetImportVoyageNumber(v string) {
	o.ImportVoyageNumber = &v
}

func (o TransportCallAllOf5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportVoyageNumber != nil {
		toSerialize["importVoyageNumber"] = o.ImportVoyageNumber
	}
	return json.Marshal(toSerialize)
}

type NullableTransportCallAllOf5 struct {
	value *TransportCallAllOf5
	isSet bool
}

func (v NullableTransportCallAllOf5) Get() *TransportCallAllOf5 {
	return v.value
}

func (v *NullableTransportCallAllOf5) Set(val *TransportCallAllOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportCallAllOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportCallAllOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportCallAllOf5(val *TransportCallAllOf5) *NullableTransportCallAllOf5 {
	return &NullableTransportCallAllOf5{value: val, isSet: true}
}

func (v NullableTransportCallAllOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportCallAllOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
