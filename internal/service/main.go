package service

import (
	"net"
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/internal/config"
	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"github.com/Distributed-Lab-Testing/example-svc/internal/data/postgres"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	db       data.DB
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "failed to register in cop")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		db:       postgres.NewDB(cfg.DB()),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(errors.Wrap(err, "failed to run service"))
	}
}
