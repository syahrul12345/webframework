package website

import "github.com/gin-gonic/gin"

func websiteHandler(c *gin.Context) {
	c.Next()
}
