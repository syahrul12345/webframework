package auth

import (
	"log"
	"net/http"
	"os"
	"scratchuniversity/apps/db"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware implements the authentication
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// secretKey := os.Getenv("INTERNAL_SECRET_KEY")
		// apiKey := c.GetHeader("api-key")
		url := c.Request.URL
		fullPath := url.EscapedPath()
		noAuthRoutes := []string{
			"/",
			"/api/v1/createAccount",
			"/api/v1/loginAccount",
			"/api/v1/forgetPassword",
		}
		for _, noAuthRoute := range noAuthRoutes {
			if fullPath == noAuthRoute {
				c.Next()
				return
			}
			// Check if it's achunk.css or chunk.js or chunk.css.map
			if strings.Contains(fullPath, ".chunk.js") || strings.Contains(fullPath, ".chunk.css") || strings.Contains(fullPath, "manifest.json") || strings.Contains(fullPath, "favicon.ico") || strings.Contains(fullPath, ".png") || strings.Contains(fullPath, ".svg") {
				c.Next()
				return
			}
		}

		// Auth needed. Get from cookies or headers
		cookieXToken, _ := c.Cookie("x-token")
		headerXToken := c.GetHeader("x-token")

		// No cookie and header
		if cookieXToken == "" && headerXToken == "" {
			// redirect to login page
			c.Redirect(http.StatusPermanentRedirect, "/")
		}
		// Got cookie no header
		if cookieXToken != "" && headerXToken == "" {
			verifyToken(cookieXToken, c)
		}
		// No cookie got header
		if cookieXToken == "" && headerXToken != "" {
			verifyToken(headerXToken, c)
		}
		// if Both have
		if cookieXToken != "" && headerXToken != "" {
			verifyToken(cookieXToken, c)
		}

		return
	}
}

func verifyToken(tokenFromClient string, c *gin.Context) {
	tk := &db.Token{}
	if tokenFromClient != "" {
		splitArray := strings.Split(tokenFromClient, " ")
		tokenString := splitArray[1]
		token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			log.Printf("Failed to go to dashboard as %s\n", err.Error())
			c.Redirect(http.StatusPermanentRedirect, "/")
			return
		}
		if !token.Valid {
			log.Println("Failed to go to dashbaord as the jwt token is invalid")
			c.Redirect(http.StatusPermanentRedirect, "/")
		}
		c.Next()
		return
	}
	c.Next()
}
