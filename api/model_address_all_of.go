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

// AddressAllOf struct for AddressAllOf
type AddressAllOf struct {
	// Name of the address
	Name *string `json:"name,omitempty"`
}

// NewAddressAllOf instantiates a new AddressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAllOf() *AddressAllOf {
	this := AddressAllOf{}
	return &this
}

// NewAddressAllOfWithDefaults instantiates a new AddressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressAllOfWithDefaults() *AddressAllOf {
	this := AddressAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddressAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddressAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddressAllOf) SetName(v string) {
	o.Name = &v
}

func (o AddressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAddressAllOf struct {
	value *AddressAllOf
	isSet bool
}

func (v NullableAddressAllOf) Get() *AddressAllOf {
	return v.value
}

func (v *NullableAddressAllOf) Set(val *AddressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressAllOf(val *AddressAllOf) *NullableAddressAllOf {
	return &NullableAddressAllOf{value: val, isSet: true}
}

func (v NullableAddressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
