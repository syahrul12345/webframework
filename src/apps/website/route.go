package website

import "github.com/gin-gonic/gin"

func websiteHandler(c *gin.Context) {
	// let react handle the website
	c.File("./website/build/index.html")
}
