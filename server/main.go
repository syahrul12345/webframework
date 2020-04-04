package main

import (
	"scratchuniversity/apps/api"
	"github.com/gin-gonic/gin"
)

// SetupRouter will set up all the required router
func SetupRouter() *gin.Engine {
	app := gin.Default()

	// Routers
	apiRouter := app.Group("/api/v1")
	// Register the routes
	api.Register(apiRouter)
	return app
}

func main() {
	app := SetupRouter()
	app.Run(config.Web.Port)
}