package v1

import (
	"hack/internal/app/config"
	"hack/internal/app/store"
)

type Api struct {
	store store.Store
	cfg   config.Config
}

func New(s store.Store, cfg config.Config) *Api {
	return &Api{
		store: s,
		cfg:   cfg,
	}
}
