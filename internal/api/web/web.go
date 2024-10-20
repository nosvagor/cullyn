package web

import (
	"bytes"
	"log"
	"net/http"
	"nosvagor/cullyn.dev/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, cmp templ.Component) {
	var b bytes.Buffer
	err := cmp.Render(c, &b)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(200, "text/html", b.Bytes())
}

func renderPage(c *gin.Context, title string, cmp templ.Component) {
	fromHTMX := c.GetHeader("HX-Request") == "true"

	if !fromHTMX {
		render(c, views.FullPage(title, cmp))
	} else {
		render(c, cmp)
	}
}
