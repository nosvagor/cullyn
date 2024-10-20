package server

import (
	"nosvagor/cullyn.dev/internal/api"
	"nosvagor/cullyn.dev/internal/api/files"
	"nosvagor/cullyn.dev/internal/api/web"
)

func (s *Server) routes() {

	// ⾕index {{{ ------------------------------------------------------------
	w := s.Router.Group("/")
	{
		w.GET("/", web.Home)
		w.GET("/robots.txt", files.Robots)
	} // ------------------------------------------------------------------ }}}

	// ⽕test {{{ -------------------------------------------------------------
	// t := s.Router.Group("/test")
	// {
	// t.GET("/colors", web.Colors)
	// } // ------------------------------------------------------------------ }}}

	// ⽔api {{{ --------------------------------------------------------------
	a := s.Router.Group("/api")
	{
		a.GET("/health", api.Health)
	} // ------------------------------------------------------------------ }}}

	s.Router.Static("/static", "./static")
}
