package main

import (
	"fmt"
	"log"
	"os"
	"scratchuniversity/apps/api"
	_ "scratchuniversity/apps/db"
	"scratchuniversity/apps/website"
	"scratchuniversity/middlewares/auth"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	port string = fmt.Sprintf(":%s", os.Getenv("PORT"))
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
	isProduction := os.Getenv("is_production")
	if isProduction == "false" {
		// Build folder will be in the /apps/website/build.
		// We will build using a multistage docker build which will send the html files to this folder
		log.Println("Non-production build")
		app.Use(static.Serve("/", static.LocalFile("./website/build", true)))
	} else {
		// Non docker build, use the build outside of the folder. This will be in alpine linux
		log.Println("Production build")
		app.Use(static.Serve("/", static.LocalFile("./build", true)))
	}

	apiRouter := app.Group("/api/v1")
	websiteRouter := app.Group("/")

	api.Register(apiRouter)
	website.Register(websiteRouter)

	return app
}

func main() {
	app := SetupRouter()
	app.Run(port)
}
