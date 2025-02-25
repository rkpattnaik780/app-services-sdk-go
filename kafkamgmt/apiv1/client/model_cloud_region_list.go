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

// CloudRegionList struct for CloudRegionList
type CloudRegionList struct {

	Kind string `json:"kind"`

	Page int32 `json:"page"`

	Size int32 `json:"size"`

	Total int32 `json:"total"`

	Items []CloudRegion `json:"items"`

}

// NewCloudRegionList instantiates a new CloudRegionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegionList(kind string, page int32, size int32, total int32, items []CloudRegion) *CloudRegionList {
	this := CloudRegionList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewCloudRegionListWithDefaults instantiates a new CloudRegionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionListWithDefaults() *CloudRegionList {
	this := CloudRegionList{}






	return &this
}


// GetKind returns the Kind field value
func (o *CloudRegionList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CloudRegionList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CloudRegionList) SetKind(v string) {
	o.Kind = v
}


// GetPage returns the Page field value
func (o *CloudRegionList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *CloudRegionList) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *CloudRegionList) SetPage(v int32) {
	o.Page = v
}


// GetSize returns the Size field value
func (o *CloudRegionList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CloudRegionList) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CloudRegionList) SetSize(v int32) {
	o.Size = v
}


// GetTotal returns the Total field value
func (o *CloudRegionList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CloudRegionList) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CloudRegionList) SetTotal(v int32) {
	o.Total = v
}


// GetItems returns the Items field value
func (o *CloudRegionList) GetItems() []CloudRegion {
	if o == nil {
		var ret []CloudRegion
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CloudRegionList) GetItemsOk() (*[]CloudRegion, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CloudRegionList) SetItems(v []CloudRegion) {
	o.Items = v
}


func (o CloudRegionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["kind"] = o.Kind
	}
    
	if true {
		toSerialize["page"] = o.Page
	}
    
	if true {
		toSerialize["size"] = o.Size
	}
    
	if true {
		toSerialize["total"] = o.Total
	}
    
	if true {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableCloudRegionList struct {
	value *CloudRegionList
	isSet bool
}

func (v NullableCloudRegionList) Get() *CloudRegionList {
	return v.value
}

func (v *NullableCloudRegionList) Set(val *CloudRegionList) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRegionList) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRegionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRegionList(val *CloudRegionList) *NullableCloudRegionList {
	return &NullableCloudRegionList{value: val, isSet: true}
}

func (v NullableCloudRegionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRegionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

