//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gogovan/ggx-kr-payment-service/config"
	"github.com/gogovan/ggx-kr-payment-service/database"
	"github.com/gogovan/ggx-kr-payment-service/internal/api"
	"github.com/gogovan/ggx-kr-payment-service/internal/api/http"
	"github.com/gogovan/ggx-kr-payment-service/internal/application/payment"
	"github.com/gogovan/ggx-kr-payment-service/internal/validation"

	"github.com/gogovan/ggx-kr-service-utils/logger"
	"github.com/google/wire"
)

var container = wire.NewSet(
	api.NewApiContainer,
)

var apiSet = wire.NewSet(
	http.NewServer,
	validation.NewValidationRulesEngine,
)

var httpHandlerSet = wire.NewSet()

var specificServiceSet = wire.NewSet(
	payment.NewService,
)

var handlerSet = wire.NewSet()

var infraSets = wire.NewSet()

func InitializeContainer(
	appCfg *config.AppConfig,
	logger logger.Logger,
	db database.ReadDb,
) *api.ApiContainer {
	wire.Build(apiSet, container)
	return &api.ApiContainer{}
}
