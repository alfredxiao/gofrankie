package models

import (
	"strings"

	"github.com/alfredxiao/gofrankie/set"
	"github.com/go-playground/validator/v10"
)

// DeviceCheckDetailsObject carries individual details to be checked
type DeviceCheckDetailsObject struct {
	ActivityData []*KeyValuePairObject `json:"activityData"`

	// Enum: [SIGNUP LOGIN PAYMENT CONFIRMATION _<Vendor Specific List>]
	// Here we assume _<Vendor Specific List> is anything starting with underscore
	ActivityType string `json:"activityType" binding:"required" validate:"required,eq=SIGNUP|eq=LOGIN|eq=PAYMENT|eq=CONFIRMATION|_VendorSpecific"`

	// Made REQUIRED (different from how go-swagger tool interprets though)
	CheckSessionKey string `json:"checkSessionKey" binding:"required" validate:"required"`

	CheckType string `json:"checkType" binding:"required" validate:"required,eq=DEVICE|eq=BIOMETRIC|eq=COMBO"`
}

// DeviceCheckDetailsObjectCollection is an array of DeviceCheckDetailsObject
type DeviceCheckDetailsObjectCollection []*DeviceCheckDetailsObject

// KeyValuePairObject represents a key/value pair plus its kind/type
type KeyValuePairObject struct {
	KvpKey   string      `json:"kvpKey" binding:"required" validate:"required"`
	KvpValue string      `json:"kvpValue" binding:"required" validate:"required"`
	KvpType  EnumKVPType `json:"kvpType" binding:"required" validate:"required,KVPType"`
}

// EnumKVPType represents enum of valid key/value pair type
type EnumKVPType string

// Valid KVP types
const generalString = "general.string"
const generalInteger = "general.integer"
const generalFloat = "general.float"
const generalBool = "general.bool"

var validKvpTypes = make(set.Set)

func init() {
	validKvpTypes.Add(generalString, generalInteger, generalFloat, generalBool)

	Validator.RegisterValidation("_VendorSpecific", func(fl validator.FieldLevel) bool {
		return isVendorSpecific(fl.Field().String())
	})

	Validator.RegisterValidation("KVPType", func(fl validator.FieldLevel) bool {
		return isValidKvpType(fl.Field().String())
	})
}

func isVendorSpecific(activityType string) bool {
	return strings.HasPrefix(activityType, "_")
}

func isValidKvpType(kvpType string) bool {
	return validKvpTypes.Contains(kvpType)
}
