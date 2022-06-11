package v1

import (
	"hack/internal/app/config"
	"hack/internal/app/store"
	"hack/internal/app/worker"
)

type Api struct {
	store  store.Store
	cfg    config.Config
	worker *worker.Worker
}

func New(s store.Store, cfg config.Config, worker *worker.Worker) *Api {
	return &Api{
		store:  s,
		cfg:    cfg,
		worker: worker,
	}
}
