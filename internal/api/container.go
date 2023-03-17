package api

import (
	"github.com/gogovan/ggx-kr-payment-service/internal/api/http"
	"github.com/gogovan/ggx-kr-payment-service/internal/validation"
)

type ApiContainer struct {
	HttpServer       *http.Server
	ValidationEngine *validation.ValidationRulesEngine
}

func NewApiContainer(httpServer *http.Server, validationEngine *validation.ValidationRulesEngine) *ApiContainer {
	return &ApiContainer{HttpServer: httpServer, ValidationEngine: validationEngine}
}
