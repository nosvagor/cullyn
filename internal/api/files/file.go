package files

import "github.com/gin-gonic/gin"

func Robots(c *gin.Context) {
	c.File("./static/file/robots.txt")
}
