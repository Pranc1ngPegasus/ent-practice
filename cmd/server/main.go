package main

import (
	"context"

	"go.uber.org/zap"
)

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := app.dbConn.Schema.Create(ctx); err != nil {
		app.logger.Error("Failed to create schema", zap.Error(err))
	}

	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Error("Failed to start server", zap.Error(err))
	}
}
