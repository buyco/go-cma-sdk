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
)

// ModelError struct for ModelError
type ModelError struct {
	// The HTTP request method type
	HttpMethod string `json:"httpMethod"`
	// The request URI.
	RequestUri string           `json:"requestUri"`
	Errors     []SubErrorsInner `json:"errors"`
	// The HTTP status code
	StatusCode int32 `json:"statusCode"`
	// The textual representation of the response status.
	StatusCodeText string `json:"statusCodeText"`
	// The date and time (in ISO 8601 format) the error occurred.
	ErrorDateTime string `json:"errorDateTime"`
}

// NewModelError instantiates a new ModelError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelError(httpMethod string, requestUri string, errors []SubErrorsInner, statusCode int32, statusCodeText string, errorDateTime string) *ModelError {
	this := ModelError{}
	this.HttpMethod = httpMethod
	this.RequestUri = requestUri
	this.Errors = errors
	this.StatusCode = statusCode
	this.StatusCodeText = statusCodeText
	this.ErrorDateTime = errorDateTime
	return &this
}

// NewModelErrorWithDefaults instantiates a new ModelError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelErrorWithDefaults() *ModelError {
	this := ModelError{}
	return &this
}

// GetHttpMethod returns the HttpMethod field value
func (o *ModelError) GetHttpMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetHttpMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpMethod, true
}

// SetHttpMethod sets field value
func (o *ModelError) SetHttpMethod(v string) {
	o.HttpMethod = v
}

// GetRequestUri returns the RequestUri field value
func (o *ModelError) GetRequestUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUri
}

// GetRequestUriOk returns a tuple with the RequestUri field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetRequestUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUri, true
}

// SetRequestUri sets field value
func (o *ModelError) SetRequestUri(v string) {
	o.RequestUri = v
}

// GetErrors returns the Errors field value
func (o *ModelError) GetErrors() []SubErrorsInner {
	if o == nil {
		var ret []SubErrorsInner
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorsOk() ([]SubErrorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ModelError) SetErrors(v []SubErrorsInner) {
	o.Errors = v
}

// GetStatusCode returns the StatusCode field value
func (o *ModelError) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *ModelError) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetStatusCodeText returns the StatusCodeText field value
func (o *ModelError) GetStatusCodeText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusCodeText
}

// GetStatusCodeTextOk returns a tuple with the StatusCodeText field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetStatusCodeTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCodeText, true
}

// SetStatusCodeText sets field value
func (o *ModelError) SetStatusCodeText(v string) {
	o.StatusCodeText = v
}

// GetErrorDateTime returns the ErrorDateTime field value
func (o *ModelError) GetErrorDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorDateTime
}

// GetErrorDateTimeOk returns a tuple with the ErrorDateTime field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorDateTime, true
}

// SetErrorDateTime sets field value
func (o *ModelError) SetErrorDateTime(v string) {
	o.ErrorDateTime = v
}

func (o ModelError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if true {
		toSerialize["requestUri"] = o.RequestUri
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["statusCode"] = o.StatusCode
	}
	if true {
		toSerialize["statusCodeText"] = o.StatusCodeText
	}
	if true {
		toSerialize["errorDateTime"] = o.ErrorDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableModelError struct {
	value *ModelError
	isSet bool
}

func (v NullableModelError) Get() *ModelError {
	return v.value
}

func (v *NullableModelError) Set(val *ModelError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelError(val *ModelError) *NullableModelError {
	return &NullableModelError{value: val, isSet: true}
}

func (v NullableModelError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
