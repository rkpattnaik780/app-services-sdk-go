/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmt

import (
	"encoding/json"
)

// KafkaRequestPayload Schema for the request body sent to /kafkas POST
type KafkaRequestPayload struct {
	// The cloud provider where the Kafka cluster will be created in
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// Set this to true to configure the Kafka cluster to be multiAZ
	MultiAz *bool `json:"multi_az,omitempty"`
	// The name of the Kafka cluster. It must consist of lower-case alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters.
	Name string `json:"name"`
	// The region where the Kafka cluster will be created in
	Region *string `json:"region,omitempty"`
}

// NewKafkaRequestPayload instantiates a new KafkaRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaRequestPayload(name string) *KafkaRequestPayload {
	this := KafkaRequestPayload{}
	this.Name = name
	return &this
}

// NewKafkaRequestPayloadWithDefaults instantiates a new KafkaRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaRequestPayloadWithDefaults() *KafkaRequestPayload {
	this := KafkaRequestPayload{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KafkaRequestPayload) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestPayload) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KafkaRequestPayload) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KafkaRequestPayload) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetMultiAz returns the MultiAz field value if set, zero value otherwise.
func (o *KafkaRequestPayload) GetMultiAz() bool {
	if o == nil || o.MultiAz == nil {
		var ret bool
		return ret
	}
	return *o.MultiAz
}

// GetMultiAzOk returns a tuple with the MultiAz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestPayload) GetMultiAzOk() (*bool, bool) {
	if o == nil || o.MultiAz == nil {
		return nil, false
	}
	return o.MultiAz, true
}

// HasMultiAz returns a boolean if a field has been set.
func (o *KafkaRequestPayload) HasMultiAz() bool {
	if o != nil && o.MultiAz != nil {
		return true
	}

	return false
}

// SetMultiAz gets a reference to the given bool and assigns it to the MultiAz field.
func (o *KafkaRequestPayload) SetMultiAz(v bool) {
	o.MultiAz = &v
}

// GetName returns the Name field value
func (o *KafkaRequestPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KafkaRequestPayload) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KafkaRequestPayload) SetName(v string) {
	o.Name = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *KafkaRequestPayload) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestPayload) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *KafkaRequestPayload) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *KafkaRequestPayload) SetRegion(v string) {
	o.Region = &v
}

func (o KafkaRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.MultiAz != nil {
		toSerialize["multi_az"] = o.MultiAz
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaRequestPayload struct {
	value *KafkaRequestPayload
	isSet bool
}

func (v NullableKafkaRequestPayload) Get() *KafkaRequestPayload {
	return v.value
}

func (v *NullableKafkaRequestPayload) Set(val *KafkaRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaRequestPayload(val *KafkaRequestPayload) *NullableKafkaRequestPayload {
	return &NullableKafkaRequestPayload{value: val, isSet: true}
}

func (v NullableKafkaRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


