package main

import (
	"log"
	"nosvagor/cullyn.dev/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}
	defer s.Pool.Close()

	// PRODUCTION =============================================================
	if gin.Mode() == gin.ReleaseMode {
	} // ======================================================================

	// DEVELOPMENT ============================================================
	if gin.Mode() == gin.DebugMode {
		go s.HealthCheck()
	} // ======================================================================

	s.Start()
}
