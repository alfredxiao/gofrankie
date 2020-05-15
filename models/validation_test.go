package models

import (
	"testing"

	"github.com/alfredxiao/gofrankie/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyValuePairObjectNotEmptyAndTypeIsLegal(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Key, Value string
		Type       EnumKVPType
		IsValid    bool
	}{
		{"key", "value", "general.string", true},
		{"", "value", "general.string", false},
		{"key", "", "general.string", false},
		{"key", "value", "", false},
		{"key", "value", "general.other", false},
	}

	for _, c := range cases {
		err := validateKeyValuePairObject(KeyValuePairObject{c.Key, c.Value, c.Type})
		assert.Equal(c.IsValid, err == nil)
	}
}

func TestKeyValuePairObjectValueMatchingItsType(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Key, Value string
		Type       EnumKVPType
		IsValid    bool
	}{
		{"key", "value", "general.string", true},
		{"key", "12345", "general.integer", true},
		{"key", "x1234", "general.integer", false},
		{"key", "12345", "general.float", true},
		{"key", "12.34", "general.float", true},
		{"key", "x12.3", "general.float", false},
		{"key", "true", "general.bool", true},
		{"key", "false", "general.bool", true},
		{"key", "yes", "general.bool", false},
	}

	for _, c := range cases {
		err := validateKeyValuePairObject(KeyValuePairObject{c.Key, c.Value, c.Type})
		assert.Equal(c.IsValid, err == nil)
	}
}

func TestValidateDeviceCheckDetailsObjectValidatesNestedActivityData(t *testing.T) {
	assert := assert.New(t)

	var detail = DeviceCheckDetailsObject{
		ActivityType:    "SIGNUP",
		CheckSessionKey: "key",
		CheckType:       "DEVICE",
		ActivityData: []*KeyValuePairObject{
			{},
		},
	}

	err := validateDeviceCheckDetailsObject(detail)
	assert.True(err != nil)
}

func TestValidateDeviceCheckDetailsObjectValidatesActivityDataKeysAreUnique(t *testing.T) {
	assert := assert.New(t)

	var kvp1 = &KeyValuePairObject{
		KvpKey:   "k1",
		KvpValue: "v1",
		KvpType:  "general.string",
	}

	var kvp2 = &KeyValuePairObject{
		KvpKey:   "k2",
		KvpValue: "v2",
		KvpType:  "general.string",
	}

	var detail = DeviceCheckDetailsObject{
		ActivityType:    "SIGNUP",
		CheckSessionKey: "key",
		CheckType:       "DEVICE",
	}

	var validDetail = detail
	validDetail.ActivityData = []*KeyValuePairObject{kvp1, kvp2}

	err := validateDeviceCheckDetailsObject(validDetail)
	assert.True(err == nil)

	var invalidDetail = detail
	invalidDetail.ActivityData = []*KeyValuePairObject{kvp1, kvp1}

	err = validateDeviceCheckDetailsObject(invalidDetail)
	require.True(t, err != nil, "kvp keys should be unique")
	assert.Contains(err.Error(), "kvpKey not unique within all activityData")
}

func TestValidateDeviceCheckDetailsObjectCollectionValidatesEmptiness(t *testing.T) {
	assert := assert.New(t)

	_, err := ValidateDeviceCheckDetailsObjectCollection(DeviceCheckDetailsObjectCollection{})
	assert.True(err != nil)
}

func TestValidateDeviceCheckDetailsObjectCollectionValidatesNestedDetailObject(t *testing.T) {
	assert := assert.New(t)

	var validDetail = &DeviceCheckDetailsObject{
		ActivityType:    "SIGNUP",
		CheckSessionKey: "key",
		CheckType:       "DEVICE",
	}

	var invalidDetail = &DeviceCheckDetailsObject{
		ActivityType:    "SIGNUP",
		CheckSessionKey: "key",
		CheckType:       "UNKNOWN",
	}

	_, err := ValidateDeviceCheckDetailsObjectCollection(DeviceCheckDetailsObjectCollection{validDetail})
	assert.True(err == nil)

	_, err = ValidateDeviceCheckDetailsObjectCollection(DeviceCheckDetailsObjectCollection{invalidDetail})
	assert.True(err != nil)
}

func TestValidateDeviceCheckDetailsObjectCollectionValidatesSessionKeysAreUnique(t *testing.T) {
	assert := assert.New(t)

	var detail1 = &DeviceCheckDetailsObject{
		ActivityType:    "SIGNUP",
		CheckSessionKey: "k1",
		CheckType:       "DEVICE",
	}

	var detail2 = &DeviceCheckDetailsObject{
		ActivityType:    "LOGIN",
		CheckSessionKey: "k2",
		CheckType:       "COMBO",
	}

	_, err := ValidateDeviceCheckDetailsObjectCollection(DeviceCheckDetailsObjectCollection{detail1, detail2})
	assert.True(err == nil)

	_, err = ValidateDeviceCheckDetailsObjectCollection(DeviceCheckDetailsObjectCollection{detail1, detail1})
	require.True(t, err != nil, "session keys must be unique within the same request payload")
	assert.Contains(err.Error(), "checkSessionKey not unique within this request payload")
}

func TestValidateDeviceCheckDetailsObjectCollectionReturnsSessionKeys(t *testing.T) {
	assert := assert.New(t)

	var detail1 = &DeviceCheckDetailsObject{
		ActivityType:    "SIGNUP",
		CheckSessionKey: "k1",
		CheckType:       "DEVICE",
	}

	var detail2 = &DeviceCheckDetailsObject{
		ActivityType:    "LOGIN",
		CheckSessionKey: "k2",
		CheckType:       "COMBO",
	}

	var expectedSet = make(set.Set)
	expectedSet.Add("k1", "k2")

	sessionKeys, err := ValidateDeviceCheckDetailsObjectCollection(DeviceCheckDetailsObjectCollection{detail1, detail2})
	assert.True(err == nil)
	assert.Equal(expectedSet, sessionKeys)
}
