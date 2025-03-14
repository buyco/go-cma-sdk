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

// Location generally used to capture location-related data, also for locations without UN Location Codes.
type Location struct {
	// The name of the location.
	LocationName *string `json:"locationName,omitempty"`
	// Geographic coordinate that specifies the north–south position of a point on the Earth&apos;s surface.
	Latitude *string `json:"latitude,omitempty"`
	// Geographic coordinate that specifies the east–west position of a point on the Earth&apos;s surface.
	Longitude *string `json:"longitude,omitempty"`
	// The UN Location code specifying where the place is located.
	UNLocationCode *string `json:"UNLocationCode,omitempty"`
	// Address related information
	Address NullableAddress `json:"address,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *Location) GetLocationName() string {
	if o == nil || o.LocationName == nil {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLocationNameOk() (*string, bool) {
	if o == nil || o.LocationName == nil {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *Location) HasLocationName() bool {
	if o != nil && o.LocationName != nil {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *Location) SetLocationName(v string) {
	o.LocationName = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Location) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLatitudeOk() (*string, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *Location) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *Location) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Location) GetLongitude() string {
	if o == nil || o.Longitude == nil {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLongitudeOk() (*string, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *Location) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *Location) SetLongitude(v string) {
	o.Longitude = &v
}

// GetUNLocationCode returns the UNLocationCode field value if set, zero value otherwise.
func (o *Location) GetUNLocationCode() string {
	if o == nil || o.UNLocationCode == nil {
		var ret string
		return ret
	}
	return *o.UNLocationCode
}

// GetUNLocationCodeOk returns a tuple with the UNLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetUNLocationCodeOk() (*string, bool) {
	if o == nil || o.UNLocationCode == nil {
		return nil, false
	}
	return o.UNLocationCode, true
}

// HasUNLocationCode returns a boolean if a field has been set.
func (o *Location) HasUNLocationCode() bool {
	if o != nil && o.UNLocationCode != nil {
		return true
	}

	return false
}

// SetUNLocationCode gets a reference to the given string and assigns it to the UNLocationCode field.
func (o *Location) SetUNLocationCode(v string) {
	o.UNLocationCode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetAddress() Address {
	if o == nil || o.Address.Get() == nil {
		var ret Address
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *Location) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAddress and assigns it to the Address field.
func (o *Location) SetAddress(v Address) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *Location) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *Location) UnsetAddress() {
	o.Address.Unset()
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocationName != nil {
		toSerialize["locationName"] = o.LocationName
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.UNLocationCode != nil {
		toSerialize["UNLocationCode"] = o.UNLocationCode
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
