/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.8.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ConsumerGroupResetOffsetParameters struct for ConsumerGroupResetOffsetParameters
type ConsumerGroupResetOffsetParameters struct {

	Offset OffsetType `json:"offset"`

	// Value associated with the given `offset`. Not used for `offset` values `earliest` and `latest`. When `offset` is `timestamp` then `value` must be a valid timestamp representing the point in time to reset the consumer group. When `offset` is `absolute` then `value` must be the integer offset to which the consumer group will be reset.
	Value *string `json:"value,omitempty"`

	Topics *[]TopicsToResetOffset `json:"topics,omitempty"`

}

// NewConsumerGroupResetOffsetParameters instantiates a new ConsumerGroupResetOffsetParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupResetOffsetParameters(offset OffsetType) *ConsumerGroupResetOffsetParameters {
	this := ConsumerGroupResetOffsetParameters{}
	this.Offset = offset
	return &this
}

// NewConsumerGroupResetOffsetParametersWithDefaults instantiates a new ConsumerGroupResetOffsetParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupResetOffsetParametersWithDefaults() *ConsumerGroupResetOffsetParameters {
	this := ConsumerGroupResetOffsetParameters{}




	return &this
}


// GetOffset returns the Offset field value
func (o *ConsumerGroupResetOffsetParameters) GetOffset() OffsetType {
	if o == nil {
		var ret OffsetType
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetParameters) GetOffsetOk() (*OffsetType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ConsumerGroupResetOffsetParameters) SetOffset(v OffsetType) {
	o.Offset = v
}


// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConsumerGroupResetOffsetParameters) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetParameters) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConsumerGroupResetOffsetParameters) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConsumerGroupResetOffsetParameters) SetValue(v string) {
	o.Value = &v
}


// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *ConsumerGroupResetOffsetParameters) GetTopics() []TopicsToResetOffset {
	if o == nil || o.Topics == nil {
		var ret []TopicsToResetOffset
		return ret
	}
	return *o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetParameters) GetTopicsOk() (*[]TopicsToResetOffset, bool) {
	if o == nil || o.Topics == nil {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *ConsumerGroupResetOffsetParameters) HasTopics() bool {
	if o != nil && o.Topics != nil {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []TopicsToResetOffset and assigns it to the Topics field.
func (o *ConsumerGroupResetOffsetParameters) SetTopics(v []TopicsToResetOffset) {
	o.Topics = &v
}


func (o ConsumerGroupResetOffsetParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["offset"] = o.Offset
	}
    
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
    
	if o.Topics != nil {
		toSerialize["topics"] = o.Topics
	}
    
	return json.Marshal(toSerialize)
}

type NullableConsumerGroupResetOffsetParameters struct {
	value *ConsumerGroupResetOffsetParameters
	isSet bool
}

func (v NullableConsumerGroupResetOffsetParameters) Get() *ConsumerGroupResetOffsetParameters {
	return v.value
}

func (v *NullableConsumerGroupResetOffsetParameters) Set(val *ConsumerGroupResetOffsetParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupResetOffsetParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupResetOffsetParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupResetOffsetParameters(val *ConsumerGroupResetOffsetParameters) *NullableConsumerGroupResetOffsetParameters {
	return &NullableConsumerGroupResetOffsetParameters{value: val, isSet: true}
}

func (v NullableConsumerGroupResetOffsetParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupResetOffsetParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

