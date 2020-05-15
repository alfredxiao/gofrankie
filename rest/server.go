package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var server *http.Server

// StartThenWait starts http server then waits on signal to stop it
func StartThenWait(address string) {
	router := setupRouter()
	server = &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen error: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}
	log.Println("Server exit complete")
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/isgood", isGoodHandler)

	return router
}
