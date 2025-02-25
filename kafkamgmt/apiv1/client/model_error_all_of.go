/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// ErrorAllOf struct for ErrorAllOf
type ErrorAllOf struct {

	Code *string `json:"code,omitempty"`

	Reason *string `json:"reason,omitempty"`

	OperationId *string `json:"operation_id,omitempty"`

}

// NewErrorAllOf instantiates a new ErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorAllOf() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// NewErrorAllOfWithDefaults instantiates a new ErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorAllOfWithDefaults() *ErrorAllOf {
	this := ErrorAllOf{}




	return &this
}


// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorAllOf) SetCode(v string) {
	o.Code = &v
}


// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorAllOf) SetReason(v string) {
	o.Reason = &v
}


// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *ErrorAllOf) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *ErrorAllOf) HasOperationId() bool {
	if o != nil && o.OperationId != nil {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *ErrorAllOf) SetOperationId(v string) {
	o.OperationId = &v
}


func (o ErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
    
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
    
	if o.OperationId != nil {
		toSerialize["operation_id"] = o.OperationId
	}
    
	return json.Marshal(toSerialize)
}

type NullableErrorAllOf struct {
	value *ErrorAllOf
	isSet bool
}

func (v NullableErrorAllOf) Get() *ErrorAllOf {
	return v.value
}

func (v *NullableErrorAllOf) Set(val *ErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorAllOf(val *ErrorAllOf) *NullableErrorAllOf {
	return &NullableErrorAllOf{value: val, isSet: true}
}

func (v NullableErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

