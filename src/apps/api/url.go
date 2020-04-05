package api

import "github.com/gin-gonic/gin"

// Register the endpoint to the relevant routers
func Register(router *gin.RouterGroup) {
	router.POST("/createAccount", createAccountHandler)
	router.POST("/loginAccount", loginAccountHandler)
	router.POST("/forgetPassword", forgetPasswordHandler)
	router.POST("/changePassword", changePasswordHandler)
	router.POST("/getAccountDetails", getAccountDetailHandler)
}
