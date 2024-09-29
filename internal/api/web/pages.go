package web

import (
	pages "nosvagor/llc/views/pages"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	renderPage(c, "home", pages.Home("@nosvagor"))
}

func Colors(c *gin.Context) {
	renderPage(c, "colors", pages.Colors())
}
