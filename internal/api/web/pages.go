package web

import (
	pages "nosvagor/llc/views/pages"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	renderPage(c, "", pages.Home("cullyn"))
}
