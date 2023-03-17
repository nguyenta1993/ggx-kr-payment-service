package healthcheck

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gogovan/ggx-kr-payment-service/config"
	"github.com/gogovan/ggx-kr-payment-service/database"
	"github.com/gogovan/ggx-kr-payment-service/pkg/constants"

	"github.com/gogovan/ggx-kr-service-utils/logger"
	"github.com/heptiolabs/healthcheck"
	"go.uber.org/zap"
)

func RunHealthCheck(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.AppConfig,
	readDb database.ReadDb,
	writeDb database.WriteDb,
) func() {
	return func() {
		health := healthcheck.NewHandler()

		interval := time.Duration(cfg.Heathcheck.Interval) * time.Second
		livenessCheck(ctx, cfg, health)
		readinessCheck(ctx, logger, health, interval, readDb, writeDb)

		logMd := func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Our middleware logic goes here...
				rw := NewResponseWriter(w)
				next.ServeHTTP(rw, r)
				statusCode := rw.Code()
				logger.Info(
					"Response information",
					zap.String("status_code", strconv.Itoa(statusCode)),
					zap.String("Method", r.Method),
					zap.String("URL", r.RequestURI),
				)
			})
		}

		logger.Info("Heathcheck server listening on port", zap.String("Port", cfg.Heathcheck.Port))
		if err := http.ListenAndServe(cfg.Heathcheck.Port, logMd(health)); err != nil {
			logger.Warn("Heathcheck server", zap.Error(err))
		}
	}
}

func livenessCheck(ctx context.Context, cfg *config.AppConfig, health healthcheck.Handler) {
	health.AddLivenessCheck(constants.GoroutineThreshold, healthcheck.GoroutineCountCheck(cfg.Heathcheck.GoroutineThreshold))
}

func readinessCheck(
	ctx context.Context,
	logger logger.Logger,
	health healthcheck.Handler,
	interval time.Duration,
	readDb database.ReadDb,
	writeDb database.WriteDb,
) {
	health.AddReadinessCheck(constants.ReadDatabase, healthcheck.AsyncWithContext(ctx, func() error {
		return readDb.DB.PingContext(ctx)
	}, interval))

	//health.AddReadinessCheck(constants.WriteDatabase, healthcheck.AsyncWithContext(ctx, func() error {
	//	return writeDb.DB.PingContext(ctx)
	//}, interval))

}

type ResponseWriter struct {
	http.ResponseWriter

	code int
	size int
}

// Returns a new `ResponseWriter` type by decorating `http.ResponseWriter` type.
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
	}
}

// Overrides `http.ResponseWriter` type.
func (r *ResponseWriter) WriteHeader(code int) {
	if r.Code() == 0 {
		r.code = code
		r.ResponseWriter.WriteHeader(code)
	}
}

// Overrides `http.ResponseWriter` type.
func (r *ResponseWriter) Write(body []byte) (int, error) {
	if r.Code() == 0 {
		r.WriteHeader(http.StatusOK)
	}

	var err error
	r.size, err = r.ResponseWriter.Write(body)

	return r.size, err
}

// Overrides `http.Flusher` type.
func (r *ResponseWriter) Flush() {
	if fl, ok := r.ResponseWriter.(http.Flusher); ok {
		if r.Code() == 0 {
			r.WriteHeader(http.StatusOK)
		}

		fl.Flush()
	}
}

// Overrides `http.Hijacker` type.
func (r *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := r.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the hijacker interface is not supported")
	}

	return hj.Hijack()
}

// Returns response status code.
func (r *ResponseWriter) Code() int {
	return r.code
}

// Returns response size.
func (r *ResponseWriter) Size() int {
	return r.size
}
