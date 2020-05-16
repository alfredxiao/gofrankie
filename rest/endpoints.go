package rest

import (
	"github.com/alfredxiao/gofrankie/models"
	"github.com/alfredxiao/gofrankie/set"
	"github.com/gin-gonic/gin"
)

const errorJSONBinding = 10
const errorDataValidation = 20
const errorSessionKeyNonUnique = 30

var sessions = make(set.Set)

func isGoodHandler(c *gin.Context) {
	var details models.DeviceCheckDetailsObjectCollection

	// due to limitation of Gin/its validator, only struct is validated, array is not
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(400, errorObject(errorJSONBinding, err.Error()))
		return
	}

	// TODO: Consider return individual validation error code rather than general code as errorDataValidation
	sessionKeysInThisRequest, err := models.ValidateDeviceCheckDetailsObjectCollection(details)
	if err != nil {
		c.JSON(400, errorObject(errorDataValidation, err.Error()))
		return
	}

	// check each session key against saved sessions
	for _, detail := range details {
		if sessions.Contains(detail.CheckSessionKey) {
			c.JSON(400, errorObject(errorSessionKeyNonUnique, "checkSessionKey not unique across calls:"+detail.CheckSessionKey))
			return
		}
	}

	// save session keys once accepted
	for k := range sessionKeysInThisRequest {
		sessions.Add(k)
	}

	c.JSON(200, models.PuppyObject{
		Puppy: true,
	})
}

func errorObject(code int, message string) models.ErrorObject {
	return models.ErrorObject{
		Code:    code,
		Message: message,
	}
}
