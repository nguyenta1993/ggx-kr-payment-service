package http

import (
	"github.com/gogovan/ggx-kr-payment-service/config"
	v1 "github.com/gogovan/ggx-kr-payment-service/internal/api/http/v1"
	"os"

	"github.com/swaggo/swag"

	httpserver "github.com/gogovan/ggx-kr-service-utils/http/server"
	"github.com/gogovan/ggx-kr-service-utils/logger"
)

type Server struct {
	logger logger.Logger
	cfg    *config.AppConfig
}

func NewServer(logger logger.Logger, cfg *config.AppConfig) *Server {
	return &Server{logger: logger, cfg: cfg}
}

func (s *Server) Run() {
	httpServer, router := httpserver.NewServer(s.logger, httpserver.HttpServerConfig{
		Port:            s.cfg.Http.Port,
		Development:     s.cfg.Http.Development,
		ShutdownTimeout: s.cfg.Http.ShutdownTimeout,
		Resources:       s.cfg.Http.Resources,
		AllowOrigins:    s.cfg.Http.AllowOrigins,
		RateLimiting: &httpserver.RateLimitingConfig{
			RateFormat: s.cfg.Http.RateLimiting.RateFormat,
		},
	})
	// In the future, if we have v2, v3..., we will add at here
	v1.MapRoutes(router)
	httpServer.Run()
}

func init() {
	dat, err := os.ReadFile("./docs/swagger.json")
	if err != nil {
		println("error when reading specs, please regenerate swagger")
	}
	spec := &swag.Spec{
		Version:          "1.0",
		BasePath:         "/api/v1/",
		Schemes:          []string{},
		Title:            "Common Service API",
		Description:      "Service for common api",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(dat),
	}
	swag.Register(spec.InstanceName(), spec)
}
