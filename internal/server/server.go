package server

import (
	"context"
	"fmt"
	"time"

	"nosvagor/llc/internal/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Pool   *pgxpool.Pool
	Config *config.Config
	Router *gin.Engine
}

func New() (*Server, error) {
	// load environment variables
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("unable to load config: %w", err)
	}

	// set the timezone to server timezone
	location, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, fmt.Errorf("unable to load timezone: %w", err)
	}
	time.Local = location

	// create new connection pool
	pool, err := pgxpool.New(context.Background(), cfg.DBUrl())
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &Server{
		Pool:   pool,
		Config: cfg,
		Router: gin.Default(),
	}, nil
}

func (s *Server) HealthCheck() {
	time.Sleep(3 * time.Second)

	fmt.Printf("========================================\n")
	stats := s.Pool.Stat()
	fmt.Printf("\tMax Connections: %v\n", stats.MaxConns())
	fmt.Printf("\tOpen Connections: %v\n", stats.AcquiredConns())
	fmt.Printf("\tIdle Connections: %v\n", stats.IdleConns())
	fmt.Printf("========================================\n")
}

func (s *Server) Start() {
	s.routes()
	s.Router.Run("0.0.0.0:" + s.Config.WebPort)
}
