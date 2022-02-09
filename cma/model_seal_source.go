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
	"fmt"
)

// SealSource The source of the seal, namely who has affixed the seal. This attribute links to the Seal Source ID defined in the Seal Source reference data entity. - CAR (Carrier) - SHI (Shipper) - PHY (Phytosanitary) - VET (Veterinary) - CUS (Customs)
type SealSource string

// List of sealSource
const (
	SS_CAR SealSource = "CAR"
	SS_SHI SealSource = "SHI"
	SS_PHY SealSource = "PHY"
	SS_VET SealSource = "VET"
	SS_CUS SealSource = "CUS"
)

// All allowed values of SealSource enum
var AllowedSealSourceEnumValues = []SealSource{
	"CAR",
	"SHI",
	"PHY",
	"VET",
	"CUS",
}

func (v *SealSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SealSource(value)
	for _, existing := range AllowedSealSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SealSource", value)
}

// NewSealSourceFromValue returns a pointer to a valid SealSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSealSourceFromValue(v string) (*SealSource, error) {
	ev := SealSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SealSource: valid values are %v", v, AllowedSealSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SealSource) IsValid() bool {
	for _, existing := range AllowedSealSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sealSource value
func (v SealSource) Ptr() *SealSource {
	return &v
}

type NullableSealSource struct {
	value *SealSource
	isSet bool
}

func (v NullableSealSource) Get() *SealSource {
	return v.value
}

func (v *NullableSealSource) Set(val *SealSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSealSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSealSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSealSource(val *SealSource) *NullableSealSource {
	return &NullableSealSource{value: val, isSet: true}
}

func (v NullableSealSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSealSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

