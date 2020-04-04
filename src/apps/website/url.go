package website

import "github.com/gin-gonic/gin"

// Register all website routes
func Register(router *gin.RouterGroup) {
	router.GET("/forgetPassword", websiteHandler)
	router.GET("/create", websiteHandler)
	router.GET("/dashboard", websiteHandler)
}
