package web

import (
	"fmt"
	"nosvagor/cullyn.dev/views/pages"
	"nosvagor/cullyn.dev/views/pages/home"
	"strings"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	listToRoutes := func(parent string, pages []string) []home.Route {
		routes := make([]home.Route, len(pages))
		for i, page := range pages {
			routes[i] = home.Route{
				Title: page,
				Route: fmt.Sprintf("/%s/%s", parent, strings.ToLower(page)),
				ImagePath: fmt.Sprintf(
					"/static/imgs/home/%s/%s.jpg",
					parent,
					strings.ToLower(page),
				),
			}
		}
		return routes
	}

	inputs := listToRoutes("inputs",
		[]string{
			"books",
			"podcasts",
			"videos",
			"essays",
			"film",
			"music",
			"games",
			"links",
			"story",
		},
	)
	outputs := listToRoutes("outputs",
		[]string{
			"essays",
			"videos",
			"software",
			"clothing",
			"art",
		},
	)

	renderPage(c, "", pages.Home(home.Content{
		Inputs:   inputs,
		HeroText: "cullyn",
		Outputs:  outputs,
	}))
}
