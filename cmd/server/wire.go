//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Pranc1ngPegasus/ent-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/ent-practice/adapter/logger"
	"github.com/Pranc1ngPegasus/ent-practice/adapter/persistence"
	"github.com/Pranc1ngPegasus/ent-practice/adapter/server"
	"github.com/Pranc1ngPegasus/ent-practice/ent"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type app struct {
	logger *zap.Logger
	dbConn *ent.Client
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		logger.NewLogger,

		configuration.Get,
		wire.FieldsOf(new(configuration.Config), "Database"),
		wire.FieldsOf(new(configuration.Config), "Server"),

		persistence.NewRDBConnector,

		server.NewServer,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
