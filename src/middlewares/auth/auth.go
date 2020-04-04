package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware implements the authentication
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// secretKey := os.Getenv("INTERNAL_SECRET_KEY")
		// apiKey := c.GetHeader("api-key")
		fullPath := c.FullPath()

		noAuthRoutes := []string{
			"/api/v1/createAccount",
			"api/v1/loginAccount",
			"api/v1/forgetPassword",
		}
		for _, noAuthRoute := range noAuthRoutes {
			if fullPath == noAuthRoute {
				log.Println("no auth needed")
				return
			}
		}
		// Auth needed
		log.Println("Auth needed")

	}
}
