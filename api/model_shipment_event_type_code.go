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

// ShipmentEventTypeCode The status of the document in the process. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released) More details can be found on <a href=\"https://github.com/dcsaorg/DCSA-Information-Model/blob/master/datamodel/referencedata.d/shipmenteventtypecodes.csv\">GitHub</a>
type ShipmentEventTypeCode string

// List of shipmentEventTypeCode
const (
	RECE ShipmentEventTypeCode = "RECE"
	DRFT ShipmentEventTypeCode = "DRFT"
	PENA ShipmentEventTypeCode = "PENA"
	PENU ShipmentEventTypeCode = "PENU"
	REJE ShipmentEventTypeCode = "REJE"
	APPR ShipmentEventTypeCode = "APPR"
	ISSU ShipmentEventTypeCode = "ISSU"
	SURR ShipmentEventTypeCode = "SURR"
	SUBM ShipmentEventTypeCode = "SUBM"
	VOID ShipmentEventTypeCode = "VOID"
	CONF ShipmentEventTypeCode = "CONF"
	REQS ShipmentEventTypeCode = "REQS"
	CMPL ShipmentEventTypeCode = "CMPL"
	HOLD ShipmentEventTypeCode = "HOLD"
	RELS ShipmentEventTypeCode = "RELS"
)

// All allowed values of ShipmentEventTypeCode enum
var AllowedShipmentEventTypeCodeEnumValues = []ShipmentEventTypeCode{
	"RECE",
	"DRFT",
	"PENA",
	"PENU",
	"REJE",
	"APPR",
	"ISSU",
	"SURR",
	"SUBM",
	"VOID",
	"CONF",
	"REQS",
	"CMPL",
	"HOLD",
	"RELS",
}

func (v *ShipmentEventTypeCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipmentEventTypeCode(value)
	for _, existing := range AllowedShipmentEventTypeCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipmentEventTypeCode", value)
}

// NewShipmentEventTypeCodeFromValue returns a pointer to a valid ShipmentEventTypeCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentEventTypeCodeFromValue(v string) (*ShipmentEventTypeCode, error) {
	ev := ShipmentEventTypeCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentEventTypeCode: valid values are %v", v, AllowedShipmentEventTypeCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentEventTypeCode) IsValid() bool {
	for _, existing := range AllowedShipmentEventTypeCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shipmentEventTypeCode value
func (v ShipmentEventTypeCode) Ptr() *ShipmentEventTypeCode {
	return &v
}

type NullableShipmentEventTypeCode struct {
	value *ShipmentEventTypeCode
	isSet bool
}

func (v NullableShipmentEventTypeCode) Get() *ShipmentEventTypeCode {
	return v.value
}

func (v *NullableShipmentEventTypeCode) Set(val *ShipmentEventTypeCode) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentEventTypeCode) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentEventTypeCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentEventTypeCode(val *ShipmentEventTypeCode) *NullableShipmentEventTypeCode {
	return &NullableShipmentEventTypeCode{value: val, isSet: true}
}

func (v NullableShipmentEventTypeCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentEventTypeCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
