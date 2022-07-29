package configuration

import (
	"io"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type (
	Config struct {
		Server   Server
		Database Database
	}

	Server struct {
		Port string `required:"false" envconfig:"SERVER_PORT" default:"3000"`
	}

	Database struct {
		DSN         string        `required:"true" envconfig:"DATABASE_DSN" default:""`
		Debug       bool          `required:"false" envconfig:"DATABASE_DEBUG" default:"true"`
		MaxOpen     int           `required:"false" envconfig:"DATABASE_CONNECTION_MAX_OPEN" default:"1"`
		MaxIdle     int           `required:"false" envconfig:"DATABASE_CONNECTION_MAX_IDLE" default:"2"`
		MaxLifetime time.Duration `required:"false" envconfig:"DATABASE_CONNECTION_MAX_LIFETIME"`
	}
)

var globalConfig Config

func Usage(output io.Writer) {
	if err := envconfig.Usagef("", &globalConfig, output, envconfig.DefaultTableFormat); err != nil {
		panic(err.Error())
	}
}

func Get(
	logger *zap.Logger,
) (Config, error) {
	var emptyconfig Config
	if globalConfig == emptyconfig {
		if err := envconfig.Process("", &globalConfig); err != nil {
			logger.Error("Failed to load environment variables", zap.Error(err))

			return Config{}, err
		}
	}

	return globalConfig, nil
}
