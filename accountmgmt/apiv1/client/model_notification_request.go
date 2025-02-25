/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// NotificationRequest struct for NotificationRequest
type NotificationRequest struct {
	BccAddress *string `json:"bcc_address,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	ClusterUuid *string `json:"cluster_uuid,omitempty"`
	IncludeRedHatAssociates *bool `json:"include_red_hat_associates,omitempty"`
	Subject *string `json:"subject,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	TemplateName string `json:"template_name"`
	TemplateParameters *[]TemplateParameter `json:"template_parameters,omitempty"`
}

// NewNotificationRequest instantiates a new NotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRequest(templateName string) *NotificationRequest {
	this := NotificationRequest{}
	this.TemplateName = templateName
	return &this
}

// NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRequestWithDefaults() *NotificationRequest {
	this := NotificationRequest{}
	return &this
}

// GetBccAddress returns the BccAddress field value if set, zero value otherwise.
func (o *NotificationRequest) GetBccAddress() string {
	if o == nil || o.BccAddress == nil {
		var ret string
		return ret
	}
	return *o.BccAddress
}

// GetBccAddressOk returns a tuple with the BccAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetBccAddressOk() (*string, bool) {
	if o == nil || o.BccAddress == nil {
		return nil, false
	}
	return o.BccAddress, true
}

// HasBccAddress returns a boolean if a field has been set.
func (o *NotificationRequest) HasBccAddress() bool {
	if o != nil && o.BccAddress != nil {
		return true
	}

	return false
}

// SetBccAddress gets a reference to the given string and assigns it to the BccAddress field.
func (o *NotificationRequest) SetBccAddress(v string) {
	o.BccAddress = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *NotificationRequest) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *NotificationRequest) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *NotificationRequest) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *NotificationRequest) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *NotificationRequest) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *NotificationRequest) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetIncludeRedHatAssociates returns the IncludeRedHatAssociates field value if set, zero value otherwise.
func (o *NotificationRequest) GetIncludeRedHatAssociates() bool {
	if o == nil || o.IncludeRedHatAssociates == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRedHatAssociates
}

// GetIncludeRedHatAssociatesOk returns a tuple with the IncludeRedHatAssociates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetIncludeRedHatAssociatesOk() (*bool, bool) {
	if o == nil || o.IncludeRedHatAssociates == nil {
		return nil, false
	}
	return o.IncludeRedHatAssociates, true
}

// HasIncludeRedHatAssociates returns a boolean if a field has been set.
func (o *NotificationRequest) HasIncludeRedHatAssociates() bool {
	if o != nil && o.IncludeRedHatAssociates != nil {
		return true
	}

	return false
}

// SetIncludeRedHatAssociates gets a reference to the given bool and assigns it to the IncludeRedHatAssociates field.
func (o *NotificationRequest) SetIncludeRedHatAssociates(v bool) {
	o.IncludeRedHatAssociates = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *NotificationRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *NotificationRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *NotificationRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *NotificationRequest) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *NotificationRequest) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *NotificationRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTemplateName returns the TemplateName field value
func (o *NotificationRequest) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetTemplateNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *NotificationRequest) SetTemplateName(v string) {
	o.TemplateName = v
}

// GetTemplateParameters returns the TemplateParameters field value if set, zero value otherwise.
func (o *NotificationRequest) GetTemplateParameters() []TemplateParameter {
	if o == nil || o.TemplateParameters == nil {
		var ret []TemplateParameter
		return ret
	}
	return *o.TemplateParameters
}

// GetTemplateParametersOk returns a tuple with the TemplateParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetTemplateParametersOk() (*[]TemplateParameter, bool) {
	if o == nil || o.TemplateParameters == nil {
		return nil, false
	}
	return o.TemplateParameters, true
}

// HasTemplateParameters returns a boolean if a field has been set.
func (o *NotificationRequest) HasTemplateParameters() bool {
	if o != nil && o.TemplateParameters != nil {
		return true
	}

	return false
}

// SetTemplateParameters gets a reference to the given []TemplateParameter and assigns it to the TemplateParameters field.
func (o *NotificationRequest) SetTemplateParameters(v []TemplateParameter) {
	o.TemplateParameters = &v
}

func (o NotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BccAddress != nil {
		toSerialize["bcc_address"] = o.BccAddress
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.ClusterUuid != nil {
		toSerialize["cluster_uuid"] = o.ClusterUuid
	}
	if o.IncludeRedHatAssociates != nil {
		toSerialize["include_red_hat_associates"] = o.IncludeRedHatAssociates
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.SubscriptionId != nil {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if true {
		toSerialize["template_name"] = o.TemplateName
	}
	if o.TemplateParameters != nil {
		toSerialize["template_parameters"] = o.TemplateParameters
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationRequest struct {
	value *NotificationRequest
	isSet bool
}

func (v NullableNotificationRequest) Get() *NotificationRequest {
	return v.value
}

func (v *NullableNotificationRequest) Set(val *NotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRequest(val *NotificationRequest) *NullableNotificationRequest {
	return &NullableNotificationRequest{value: val, isSet: true}
}

func (v NullableNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


