package main

import (
	"scratchuniversity/apps/api"
	_ "scratchuniversity/apps/db"
	"scratchuniversity/middlewares/auth"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter will set up all the required router
func SetupRouter() *gin.Engine {
	app := gin.Default()

	// Middlewares
	app.Use(auth.AuthenticationMiddleware())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
