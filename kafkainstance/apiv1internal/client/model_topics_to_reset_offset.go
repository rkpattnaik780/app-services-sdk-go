/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicsToResetOffset struct for TopicsToResetOffset
type TopicsToResetOffset struct {

	Topic string `json:"topic"`

	Partitions *[]int32 `json:"partitions,omitempty"`

}

// NewTopicsToResetOffset instantiates a new TopicsToResetOffset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicsToResetOffset(topic string) *TopicsToResetOffset {
	this := TopicsToResetOffset{}
	this.Topic = topic
	return &this
}

// NewTopicsToResetOffsetWithDefaults instantiates a new TopicsToResetOffset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicsToResetOffsetWithDefaults() *TopicsToResetOffset {
	this := TopicsToResetOffset{}



	return &this
}


// GetTopic returns the Topic field value
func (o *TopicsToResetOffset) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *TopicsToResetOffset) GetTopicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *TopicsToResetOffset) SetTopic(v string) {
	o.Topic = v
}


// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *TopicsToResetOffset) GetPartitions() []int32 {
	if o == nil || o.Partitions == nil {
		var ret []int32
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsToResetOffset) GetPartitionsOk() (*[]int32, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *TopicsToResetOffset) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []int32 and assigns it to the Partitions field.
func (o *TopicsToResetOffset) SetPartitions(v []int32) {
	o.Partitions = &v
}


func (o TopicsToResetOffset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["topic"] = o.Topic
	}
    
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
    
	return json.Marshal(toSerialize)
}

type NullableTopicsToResetOffset struct {
	value *TopicsToResetOffset
	isSet bool
}

func (v NullableTopicsToResetOffset) Get() *TopicsToResetOffset {
	return v.value
}

func (v *NullableTopicsToResetOffset) Set(val *TopicsToResetOffset) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicsToResetOffset) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicsToResetOffset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicsToResetOffset(val *TopicsToResetOffset) *NullableTopicsToResetOffset {
	return &NullableTopicsToResetOffset{value: val, isSet: true}
}

func (v NullableTopicsToResetOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicsToResetOffset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

