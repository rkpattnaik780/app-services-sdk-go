/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// ServiceAccountRequest Schema for the request to create a service account
type ServiceAccountRequest struct {

	// The name of the service account
	Name string `json:"name"`

	// A description for the service account
	Description *string `json:"description,omitempty"`

}

// NewServiceAccountRequest instantiates a new ServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountRequest(name string) *ServiceAccountRequest {
	this := ServiceAccountRequest{}
	this.Name = name
	return &this
}

// NewServiceAccountRequestWithDefaults instantiates a new ServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountRequestWithDefaults() *ServiceAccountRequest {
	this := ServiceAccountRequest{}



	return &this
}


// GetName returns the Name field value
func (o *ServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceAccountRequest) SetName(v string) {
	o.Name = v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountRequest) SetDescription(v string) {
	o.Description = &v
}


func (o ServiceAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["name"] = o.Name
	}
    
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
    
	return json.Marshal(toSerialize)
}

type NullableServiceAccountRequest struct {
	value *ServiceAccountRequest
	isSet bool
}

func (v NullableServiceAccountRequest) Get() *ServiceAccountRequest {
	return v.value
}

func (v *NullableServiceAccountRequest) Set(val *ServiceAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountRequest(val *ServiceAccountRequest) *NullableServiceAccountRequest {
	return &NullableServiceAccountRequest{value: val, isSet: true}
}

func (v NullableServiceAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

