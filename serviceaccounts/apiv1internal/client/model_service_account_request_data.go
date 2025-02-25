/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 5.0.19
 * Contact: it-user-team-list@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// ServiceAccountRequestData struct for ServiceAccountRequestData
type ServiceAccountRequestData struct {

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

}

// NewServiceAccountRequestData instantiates a new ServiceAccountRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountRequestData() *ServiceAccountRequestData {
	this := ServiceAccountRequestData{}
	return &this
}

// NewServiceAccountRequestDataWithDefaults instantiates a new ServiceAccountRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountRequestDataWithDefaults() *ServiceAccountRequestData {
	this := ServiceAccountRequestData{}



	return &this
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountRequestData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequestData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountRequestData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountRequestData) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountRequestData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequestData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountRequestData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountRequestData) SetDescription(v string) {
	o.Description = &v
}


func (o ServiceAccountRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
    
	return json.Marshal(toSerialize)
}

type NullableServiceAccountRequestData struct {
	value *ServiceAccountRequestData
	isSet bool
}

func (v NullableServiceAccountRequestData) Get() *ServiceAccountRequestData {
	return v.value
}

func (v *NullableServiceAccountRequestData) Set(val *ServiceAccountRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountRequestData(val *ServiceAccountRequestData) *NullableServiceAccountRequestData {
	return &NullableServiceAccountRequestData{value: val, isSet: true}
}

func (v NullableServiceAccountRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

