package startup

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogovan/ggx-kr-payment-service/config"
	"github.com/gogovan/ggx-kr-payment-service/database"
	"github.com/gogovan/ggx-kr-payment-service/internal"
	"github.com/gogovan/ggx-kr-payment-service/internal/api"
	"github.com/gogovan/ggx-kr-payment-service/pkg/healthcheck"
	"github.com/gogovan/ggx-kr-payment-service/pkg/metrics"
	"go.uber.org/zap"

	"github.com/gammazero/workerpool"
	"github.com/gogovan/ggx-kr-service-utils/command"
	"github.com/gogovan/ggx-kr-service-utils/localization"
	"github.com/gogovan/ggx-kr-service-utils/logger"
	"github.com/gogovan/ggx-kr-service-utils/tracing"
	"github.com/gogovan/ggx-kr-service-utils/validation"
)

func runServer(
	ctx context.Context,
	logger logger.Logger,
	container *api.ApiContainer,
	readDb database.ReadDb,
	// writeDb database.WriteDb,
) {
	wp := workerpool.New(4)
	// Run healthcheck
	wp.Submit(healthcheck.RunHealthCheck(ctx, logger, cfg, readDb, nil))
	// Run metrics
	wp.Submit(metrics.RunMetrics(logger, cfg))
	// Run Grpc server
	//wp.Submit(container.GrpcServer.Run)
	// Run Http server
	wp.Submit(container.HttpServer.Run)

	wp.StopWait()
}

func registerDependencies(ctx context.Context, logger logger.Logger) (*api.ApiContainer, database.ReadDb) {
	// Open database connection
	readDb, _ := database.Open(cfg.Database, logger)

	return internal.InitializeContainer(cfg, logger, readDb), readDb
}

var cfg *config.AppConfig

func Execute() {
	// Init AppConfig
	cfg = &config.AppConfig{}

	// Init commands
	command.UseCommands(
		//TODO temporary disable start auto migrate
		command.WithStartCommand(start, cfg),
		//command.WithStartCommand(start, cfg, "database.writeDb"),
		command.WithMigrationCommand("database.writeDb"),
	)
}

func start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	// Init logger
	logger := logger.NewZapLogger(cfg.Logger.LogLevel)
	// Register dependencies
	container, readDb := registerDependencies(ctx, logger)
	// Init resources for localization
	if err := localization.InitResources(cfg.Http.Resources); err != nil {
		logger.Fatal("Fail to init resources", zap.Error(err))
	}

	// Init tracing
	tracing.UseOpenTelemetry(tracing.Config(*cfg.Tracing))
	// Init validation
	validation.UseValidation(container.ValidationEngine.GetValidations()...)
	// Run server
	runServer(ctx, logger, container, readDb)
}
