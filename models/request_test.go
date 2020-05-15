package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateVendorSpecific(t *testing.T) {
	assert := assert.New(t)

	assert.True(isVendorSpecific("_LOGIN3"))
	assert.False(isVendorSpecific("LOGIN"))
}

func TestKVPTypeEnum(t *testing.T) {
	assert := assert.New(t)

	assert.True(isValidKvpType("general.string"))
	assert.True(isValidKvpType("general.integer"))
	assert.True(isValidKvpType("general.float"))
	assert.True(isValidKvpType("general.bool"))

	assert.False(isValidKvpType("general.other"))
	assert.False(isValidKvpType("other.string"))
}

func TestActivityTypeFromEnumAsWellAsVendorSpecific(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		ActivityType string
		IsValid      bool
	}{
		{"SIGNUP", true},
		{"LOGIN", true},
		{"PAYMENT", true},
		{"CONFIRMATION", true},
		{"_VENDOR", true},
		{"OTHER", false},
		{"", false},
	}

	for _, c := range cases {
		err := Validator.Struct(DeviceCheckDetailsObject{
			ActivityType:    c.ActivityType,
			CheckSessionKey: "key",
			CheckType:       "DEVICE",
		})
		assert.Equal(c.IsValid, err == nil)
	}
}

func TestCheckTypeFromEnum(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		CheckType string
		IsValid   bool
	}{
		{"DEVICE", true},
		{"BIOMETRIC", true},
		{"COMBO", true},
		{"OTHER", false},
		{"", false},
	}

	for _, c := range cases {
		err := Validator.Struct(DeviceCheckDetailsObject{
			ActivityType:    "LOGIN",
			CheckSessionKey: "key",
			CheckType:       c.CheckType,
		})
		assert.Equal(c.IsValid, err == nil)
	}
}

func TestCheckSessionKeyNotEmpty(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		CheckSessionKey string
		IsValid         bool
	}{
		{"akey", true},
		{"", false},
	}

	for _, c := range cases {
		err := Validator.Struct(DeviceCheckDetailsObject{
			ActivityType:    "LOGIN",
			CheckSessionKey: c.CheckSessionKey,
			CheckType:       "DEVICE",
		})
		assert.Equal(c.IsValid, err == nil)
	}
}
