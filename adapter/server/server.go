package server

import (
	"net/http"

	"github.com/Pranc1ngPegasus/ent-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/ent-practice/ent"
	"github.com/Pranc1ngPegasus/ent-practice/ent/ogent"
	"go.uber.org/zap"
)

func NewServer(
	logger *zap.Logger,
	config configuration.Server,
	client *ent.Client,
) (*http.Server, error) {
	srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		return nil, err
	}

	logger.Info("Start server", zap.String("port", config.Port))

	return &http.Server{
		Addr:    ":" + config.Port,
		Handler: srv,
	}, nil
}
