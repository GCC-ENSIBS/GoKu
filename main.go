package main

import (
	"goku/deploy"
	"goku/refresh"
	"goku/status"
	"goku/stop"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	errProxies := router.SetTrustedProxies(nil)
	if errProxies != nil {
		return
	}

	// Routes declaration
	router.GET("/status/:id", status.Status)
	router.GET("/refresh/:id", refresh.Refresh)
	router.GET("/deploy/:id", deploy.Deploy)
	router.GET("/stop/:id", stop.Stop)

	errRun := router.Run()
	if errRun != nil {
		return
	}
}
