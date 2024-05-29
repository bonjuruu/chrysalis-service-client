/*
Chrysalis Service API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// IUserRoleEnum User role.
type IUserRoleEnum string

// List of IUserRoleEnum
const (
	ADMIN IUserRoleEnum = "ADMIN"
	DEFAULT IUserRoleEnum = "DEFAULT"
)

// All allowed values of IUserRoleEnum enum
var AllowedIUserRoleEnumEnumValues = []IUserRoleEnum{
	"ADMIN",
	"DEFAULT",
}

func (v *IUserRoleEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IUserRoleEnum(value)
	for _, existing := range AllowedIUserRoleEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IUserRoleEnum", value)
}

// NewIUserRoleEnumFromValue returns a pointer to a valid IUserRoleEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIUserRoleEnumFromValue(v string) (*IUserRoleEnum, error) {
	ev := IUserRoleEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IUserRoleEnum: valid values are %v", v, AllowedIUserRoleEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IUserRoleEnum) IsValid() bool {
	for _, existing := range AllowedIUserRoleEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IUserRoleEnum value
func (v IUserRoleEnum) Ptr() *IUserRoleEnum {
	return &v
}

type NullableIUserRoleEnum struct {
	value *IUserRoleEnum
	isSet bool
}

func (v NullableIUserRoleEnum) Get() *IUserRoleEnum {
	return v.value
}

func (v *NullableIUserRoleEnum) Set(val *IUserRoleEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableIUserRoleEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableIUserRoleEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIUserRoleEnum(val *IUserRoleEnum) *NullableIUserRoleEnum {
	return &NullableIUserRoleEnum{value: val, isSet: true}
}

func (v NullableIUserRoleEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIUserRoleEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

