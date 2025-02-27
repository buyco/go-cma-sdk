/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org.  <i>Please note that shipment events and subscription management are not covered yet by CMA CGM API.</i>  <br> This API is accessible through <ul> <li> <b> Public </b> connection (authentication with API Key): this allows to retrieve standard equipment moves (e.g. ready to be loaded, discharged), Transhipment moves (e.g. discharged and re-loaded onboard) and planned vessel departure and arrival dates. </li> <li> <b> Private </b> connection (authentication with Oauth2 client credentials): this allows to retrieve additional private events such as actual rail and ramp moves, planned departure and arrival dates for inland moves. These private events are available only if you (or your end customer) are identified on the booking as: Booking Party, Deciding Party, Consignee, Notifier or Shipper. </li> </ul> </br> For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// FacilityCodeListProvider The provider used for identifying the facility Code
type FacilityCodeListProvider string

// List of facilityCodeListProvider
const (
	BIC  FacilityCodeListProvider = "BIC"
	SMDG FacilityCodeListProvider = "SMDG"
)

// All allowed values of FacilityCodeListProvider enum
var AllowedFacilityCodeListProviderEnumValues = []FacilityCodeListProvider{
	"BIC",
	"SMDG",
}

func (v *FacilityCodeListProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FacilityCodeListProvider(value)
	for _, existing := range AllowedFacilityCodeListProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FacilityCodeListProvider", value)
}

// NewFacilityCodeListProviderFromValue returns a pointer to a valid FacilityCodeListProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFacilityCodeListProviderFromValue(v string) (*FacilityCodeListProvider, error) {
	ev := FacilityCodeListProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FacilityCodeListProvider: valid values are %v", v, AllowedFacilityCodeListProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FacilityCodeListProvider) IsValid() bool {
	for _, existing := range AllowedFacilityCodeListProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to facilityCodeListProvider value
func (v FacilityCodeListProvider) Ptr() *FacilityCodeListProvider {
	return &v
}

type NullableFacilityCodeListProvider struct {
	value *FacilityCodeListProvider
	isSet bool
}

func (v NullableFacilityCodeListProvider) Get() *FacilityCodeListProvider {
	return v.value
}

func (v *NullableFacilityCodeListProvider) Set(val *FacilityCodeListProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableFacilityCodeListProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableFacilityCodeListProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacilityCodeListProvider(val *FacilityCodeListProvider) *NullableFacilityCodeListProvider {
	return &NullableFacilityCodeListProvider{value: val, isSet: true}
}

func (v NullableFacilityCodeListProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacilityCodeListProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
