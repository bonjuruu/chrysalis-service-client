/*
Chrysalis Service API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the IUserDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IUserDTO{}

// IUserDTO struct for IUserDTO
type IUserDTO struct {
	// User id.
	Id string `json:"id"`
	// User name.
	Username string `json:"username"`
	// User email.
	Email string `json:"email"`
	Role IUserRoleEnum `json:"role"`
	// User creation date.
	CreatedAt string `json:"createdAt"`
	// User update date.
	UpdatedAt string `json:"updatedAt"`
}

type _IUserDTO IUserDTO

// NewIUserDTO instantiates a new IUserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIUserDTO(id string, username string, email string, role IUserRoleEnum, createdAt string, updatedAt string) *IUserDTO {
	this := IUserDTO{}
	this.Id = id
	this.Username = username
	this.Email = email
	this.Role = role
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewIUserDTOWithDefaults instantiates a new IUserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIUserDTOWithDefaults() *IUserDTO {
	this := IUserDTO{}
	return &this
}

// GetId returns the Id field value
func (o *IUserDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IUserDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IUserDTO) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *IUserDTO) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *IUserDTO) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *IUserDTO) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *IUserDTO) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IUserDTO) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IUserDTO) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *IUserDTO) GetRole() IUserRoleEnum {
	if o == nil {
		var ret IUserRoleEnum
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *IUserDTO) GetRoleOk() (*IUserRoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *IUserDTO) SetRole(v IUserRoleEnum) {
	o.Role = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IUserDTO) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IUserDTO) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IUserDTO) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *IUserDTO) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *IUserDTO) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *IUserDTO) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o IUserDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IUserDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	toSerialize["email"] = o.Email
	toSerialize["role"] = o.Role
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *IUserDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"email",
		"role",
		"createdAt",
		"updatedAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIUserDTO := _IUserDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIUserDTO)

	if err != nil {
		return err
	}

	*o = IUserDTO(varIUserDTO)

	return err
}

type NullableIUserDTO struct {
	value *IUserDTO
	isSet bool
}

func (v NullableIUserDTO) Get() *IUserDTO {
	return v.value
}

func (v *NullableIUserDTO) Set(val *IUserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIUserDTO(val *IUserDTO) *NullableIUserDTO {
	return &NullableIUserDTO{value: val, isSet: true}
}

func (v NullableIUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


