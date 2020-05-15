package rest

import (
	"github.com/gin-gonic/gin"
)

// StartServer start REST server
func StartServer() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/isgood", isGoodHandler)

	return r
}
