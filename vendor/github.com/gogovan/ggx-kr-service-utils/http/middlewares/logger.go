package middlewares

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/gogovan/ggx-kr-service-utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
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
				zap.String("response_body", blw.body.String()),
			)
		}

	}
}
