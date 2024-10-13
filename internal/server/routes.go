package server

import (
	"nosvagor/llc/internal/api"
	"nosvagor/llc/internal/api/files"
	"nosvagor/llc/internal/api/web"
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
