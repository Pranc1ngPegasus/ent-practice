package persistence

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Pranc1ngPegasus/ent-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/ent-practice/ent"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func NewRDBConnector(
	logger *zap.Logger,
	config configuration.Database,
) (*ent.Client, error) {
	logger.Info("Start RDB connector")

	drv, err := sql.Open(dialect.SQLite, config.DSN)
	if err != nil {
		logger.Error("Failed to open RDB connection", zap.Error(err))

		return nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(config.MaxIdle)
	db.SetMaxOpenConns(config.MaxOpen)
	db.SetConnMaxLifetime(config.MaxLifetime)

	client := ent.NewClient(ent.Driver(drv))

	if config.Debug {
		client.Debug()
	}

	return client, nil
}
