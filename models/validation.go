package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alfredxiao/gofrankie/set"
	"github.com/go-playground/validator/v10"
)

// Validator validates entites in models package
var Validator = validator.New()

// validateKeyValuePairObject validates a KeyValuePairObject, fields that are required, and the value matches its claimed type
func validateKeyValuePairObject(kvp KeyValuePairObject) error {
	if err := Validator.Struct(kvp); err != nil {
		return err
	}

	var err error
	switch kvp.KvpType {
	case generalInteger:
		_, err = strconv.Atoi(kvp.KvpValue)
	case generalFloat:
		_, err = strconv.ParseFloat(kvp.KvpValue, 64)
	case generalBool:
		_, err = strconv.ParseBool(kvp.KvpValue)
	}

	if err != nil {
		return fmt.Errorf("Invalid data type in KeyValuePairObject{%s, %s, %s}", kvp.KvpKey, kvp.KvpValue, kvp.KvpType)
	}

	return nil
}

// validateDeviceCheckDetailsObject validates a details object, including its nested array objects.
// Note that it does not collect all errors, but rather returns first encountered validation error.
// And it validates uniqueness of keys in its ActivityData array
func validateDeviceCheckDetailsObject(detail DeviceCheckDetailsObject) error {
	if err := Validator.Struct(detail); err != nil {
		return err
	}

	var activityKeys = make(set.Set)
	for _, activityData := range detail.ActivityData {
		if err := validateKeyValuePairObject(*activityData); err != nil {
			return err
		}
		activityKeys.Add(activityData.KvpKey)
	}

	if len(detail.ActivityData) != len(activityKeys) {
		return errors.New("kvpKey not unique within all activityData")
	}

	return nil
}

// ValidateDeviceCheckDetailsObjectCollection validates top level request obect DeviceCheckDetailsObjectCollection
// It recursively validates nested objects, and returns the first validation error encountered.
// Note it also validates uniqueness of session keys (uniqueness within this collection, not across calls)
func ValidateDeviceCheckDetailsObjectCollection(detailsCollection DeviceCheckDetailsObjectCollection) (set.Set, error) {
	if len(detailsCollection) == 0 {
		return nil, errors.New("DeviceCheckDetailsObjectCollection cannot be empty")
	}

	var sessionKeys = make(set.Set)
	for _, detail := range detailsCollection {
		if err := validateDeviceCheckDetailsObject(*detail); err != nil {
			return nil, err
		}
		sessionKeys.Add(detail.CheckSessionKey)
	}

	if len(detailsCollection) != len(sessionKeys) {
		return nil, errors.New("checkSessionKey not unique within this request payload")
	}

	return sessionKeys, nil
}
