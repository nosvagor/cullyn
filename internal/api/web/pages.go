package web

import (
	"nosvagor/cullyn.dev/views/pages"
	"nosvagor/cullyn.dev/views/pages/home"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	heroText := "cullyn"

	inputs := []home.Input{}
	outputs := []home.Output{}

	renderPage(c, "", pages.Home(home.HomeContent{
		HeroText: heroText,
		Inputs:   inputs,
		Outputs:  outputs,
	}))
}
