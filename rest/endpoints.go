package rest

import (
	"github.com/alfredxiao/gofrankie/models"
	"github.com/alfredxiao/gofrankie/set"
	"github.com/gin-gonic/gin"
)

var sessions = make(set.Set)

func isGoodHandler(c *gin.Context) {
	var details models.DeviceCheckDetailsObjectCollection

	// due to limitation of Gin/its validator, only struct is validated, array is not
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(400, invalidRequest("JSONBinding", err.Error()))
		return
	}

	sessionKeysInThisRequest, err := models.ValidateDeviceCheckDetailsObjectCollection(details)
	if err != nil {
		c.JSON(400, invalidRequest("DataValidation", err.Error()))
		return
	}

	// check each session key against saved sessions
	for _, detail := range details {
		if sessions.Contains(detail.CheckSessionKey) {
			c.JSON(400, invalidRequest("SessionKeyNonUnique", "checkSessionKey not unique across calls:"+detail.CheckSessionKey))
			return
		}
	}

	// save session keys once accepted
	for k := range sessionKeysInThisRequest {
		sessions.Add(k)
	}

	puppy := models.PuppyObject{
		Puppy: true,
	}
	c.JSON(200, puppy)
}

func invalidRequest(errorType, errorMsg string) gin.H {
	return gin.H{
		"status":    "error",
		"errorType": errorType,
		"errorMsg":  errorMsg,
	}
}
