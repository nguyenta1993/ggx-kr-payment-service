package metrics

import (
	"bytes"
	"github.com/gogovan/ggx-kr-payment-service/config"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogovan/ggx-kr-service-utils/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func RunMetrics(logger logger.Logger, cfg *config.AppConfig) func() {
	return func() {
		gin.SetMode(gin.ReleaseMode)
		metricsServer := gin.New()
		metricsServer.Use(LoggerMiddleware(logger))
		metricsServer.GET(cfg.Metrics.PrometheusPath, prometheusHandler())
		logger.Info("Metrics server is running on port", zap.String("Metrics port", cfg.Metrics.PrometheusPort))
		if err := metricsServer.Run(cfg.Metrics.PrometheusPort); err != nil {
			logger.Error("metricsServer.Run", zap.Error(err))
		}
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func LoggerMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		if !strings.Contains(c.FullPath(), "swagger") {
			statusCode := c.Writer.Status()
			logger.Info(
				"Response information",
				zap.String("status_code", strconv.Itoa(statusCode)),
				zap.String("Method", c.Request.Method),
				zap.String("URL", c.Request.RequestURI),
			)
		}

	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
