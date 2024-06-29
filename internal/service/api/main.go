package api

import (
	"context"
	"fmt"
	"net"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/black-pepper-team/community-indexer/internal/config"
	"github.com/black-pepper-team/community-indexer/internal/service/core"
)

type service struct {
	log      *logan.Entry
	listener net.Listener
	core     *core.Core
	mockAPI  bool
}

func newService(ctx context.Context, cfg config.Config) (*service, error) {
	coreService, err := core.New(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create new issuer: %w", err)
	}

	return &service{
		log:      cfg.Log().WithField("service", "api"),
		listener: cfg.Listener(),
		core:     coreService,
		mockAPI:  cfg.API().MockAPI,
	}, nil
}

func Run(ctx context.Context, cfg config.Config) {
	svc, err := newService(ctx, cfg)
	if err != nil {
		panic(err)
	}

	svc.log.Info("Service started")
	ape.Serve(ctx, svc.router(), cfg, ape.ServeOpts{})
}
