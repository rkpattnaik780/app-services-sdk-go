/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.1.0-SNAPSHOT
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
	"fmt"
)

// SortBy the model 'SortBy'
type SortBy string

// List of SortBy
const (
	SORTBY_NAME SortBy = "name"
	SORTBY_CREATED_ON SortBy = "createdOn"
)

var allowedSortByEnumValues = []SortBy{
	"name",
	"createdOn",
}

func (v *SortBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortBy(value)
	for _, existing := range allowedSortByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortBy", value)
}

// NewSortByFromValue returns a pointer to a valid SortBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortByFromValue(v string) (*SortBy, error) {
	ev := SortBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortBy: valid values are %v", v, allowedSortByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortBy) IsValid() bool {
	for _, existing := range allowedSortByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortBy value
func (v SortBy) Ptr() *SortBy {
	return &v
}

type NullableSortBy struct {
	value *SortBy
	isSet bool
}

func (v NullableSortBy) Get() *SortBy {
	return v.value
}

func (v *NullableSortBy) Set(val *SortBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSortBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSortBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortBy(val *SortBy) *NullableSortBy {
	return &NullableSortBy{value: val, isSet: true}
}

func (v NullableSortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

