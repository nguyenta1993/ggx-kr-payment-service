package database

import (
	"github.com/gogovan/ggx-kr-payment-service/config"

	"github.com/gogovan/ggx-kr-service-utils/database"
	"github.com/gogovan/ggx-kr-service-utils/logger"

	"github.com/jmoiron/sqlx"
)

type ReadDb *sqlx.DB
type WriteDb *sqlx.DB

func Open(cfg *config.DatabaseConfig, logger logger.Logger) (ReadDb, WriteDb) {
	readDb := database.MustConnect(cfg.ReadDbCfg.DbType, cfg.ReadDbCfg.ConnectionString)
	readDb.MapperFunc(func(s string) string { return s })
	writeDb := database.MustConnect(cfg.WriteDbCfg.DbType, cfg.WriteDbCfg.ConnectionString)
	writeDb.MapperFunc(func(s string) string { return s })
	logger.Info("Connected to read & write database!")
	return readDb, writeDb
}
