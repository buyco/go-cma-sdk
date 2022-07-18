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
	"fmt"
)

// EmptyIndicatorCode Code to denote whether the equipment is empty or laden.
type EmptyIndicatorCode string

// List of emptyIndicatorCode
const (
	EMPTY EmptyIndicatorCode = "EMPTY"
	LADEN EmptyIndicatorCode = "LADEN"
)

// All allowed values of EmptyIndicatorCode enum
var AllowedEmptyIndicatorCodeEnumValues = []EmptyIndicatorCode{
	"EMPTY",
	"LADEN",
}

func (v *EmptyIndicatorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmptyIndicatorCode(value)
	for _, existing := range AllowedEmptyIndicatorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmptyIndicatorCode", value)
}

// NewEmptyIndicatorCodeFromValue returns a pointer to a valid EmptyIndicatorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmptyIndicatorCodeFromValue(v string) (*EmptyIndicatorCode, error) {
	ev := EmptyIndicatorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmptyIndicatorCode: valid values are %v", v, AllowedEmptyIndicatorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmptyIndicatorCode) IsValid() bool {
	for _, existing := range AllowedEmptyIndicatorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to emptyIndicatorCode value
func (v EmptyIndicatorCode) Ptr() *EmptyIndicatorCode {
	return &v
}

type NullableEmptyIndicatorCode struct {
	value *EmptyIndicatorCode
	isSet bool
}

func (v NullableEmptyIndicatorCode) Get() *EmptyIndicatorCode {
	return v.value
}

func (v *NullableEmptyIndicatorCode) Set(val *EmptyIndicatorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEmptyIndicatorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEmptyIndicatorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmptyIndicatorCode(val *EmptyIndicatorCode) *NullableEmptyIndicatorCode {
	return &NullableEmptyIndicatorCode{value: val, isSet: true}
}

func (v NullableEmptyIndicatorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmptyIndicatorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
