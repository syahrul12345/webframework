package api

import "github.com/gin-gonic/gin"

// Register the endpoint to the relevant routers
func Register(router *gin.RouterGroup) {
	router.POST("/craeteAcount", createAccountHandler)
}
