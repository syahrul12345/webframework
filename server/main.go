package main

import (
	"scratchuniversity/apps/api"
	_ "scratchuniversity/apps/db"
	"scratchuniversity/middlewares/auth"

	"github.com/gin-gonic/gin"
)

// SetupRouter will set up all the required router
func SetupRouter() *gin.Engine {
	app := gin.Default()

	// Middlewares
	app.Use(auth.AuthenticationMiddleware())

	// Routers
	apiRouter := app.Group("/api/v1")
	// Register the routes
	api.Register(apiRouter)
	return app
}

func main() {
	app := SetupRouter()
	app.Run("localhost:8000")
}
