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

// TransportEventTypeCode Identifier for type of Transport event - ARRI (Arrived) - DEPA (Departed) More details can be found on <a href=\"https://github.com/dcsaorg/DCSA-Information-Model/blob/master/datamodel/referencedata.d/transporteventtypecodes.csv\">GitHub</a>
type TransportEventTypeCode string

// List of transportEventTypeCode
const (
	ARRI TransportEventTypeCode = "ARRI"
	DEPA TransportEventTypeCode = "DEPA"
)

// All allowed values of TransportEventTypeCode enum
var AllowedTransportEventTypeCodeEnumValues = []TransportEventTypeCode{
	"ARRI",
	"DEPA",
}

func (v *TransportEventTypeCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransportEventTypeCode(value)
	for _, existing := range AllowedTransportEventTypeCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransportEventTypeCode", value)
}

// NewTransportEventTypeCodeFromValue returns a pointer to a valid TransportEventTypeCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransportEventTypeCodeFromValue(v string) (*TransportEventTypeCode, error) {
	ev := TransportEventTypeCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransportEventTypeCode: valid values are %v", v, AllowedTransportEventTypeCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransportEventTypeCode) IsValid() bool {
	for _, existing := range AllowedTransportEventTypeCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to transportEventTypeCode value
func (v TransportEventTypeCode) Ptr() *TransportEventTypeCode {
	return &v
}

type NullableTransportEventTypeCode struct {
	value *TransportEventTypeCode
	isSet bool
}

func (v NullableTransportEventTypeCode) Get() *TransportEventTypeCode {
	return v.value
}

func (v *NullableTransportEventTypeCode) Set(val *TransportEventTypeCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportEventTypeCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportEventTypeCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportEventTypeCode(val *TransportEventTypeCode) *NullableTransportEventTypeCode {
	return &NullableTransportEventTypeCode{value: val, isSet: true}
}

func (v NullableTransportEventTypeCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportEventTypeCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
