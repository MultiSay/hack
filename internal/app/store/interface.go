package store

import (
	"context"
	"hack/internal/app/model"
)

// Инициализируем репозитории
//go:generate mockery --name=Store --structname=Store
type Store interface {
	File() FileRepository
	Region() RegionRepository
}

//go:generate mockery --name=FileRepository --structname=FileRepository
type FileRepository interface {
	Create(context.Context, *model.File) error
	Update(context.Context, *model.File) error
	GetByID(context.Context, int) (*model.File, error)
}

//go:generate mockery --name=RegionRepository --structname=RegionRepository
type RegionRepository interface {
	PredictList(context.Context) ([]model.RegionPredict, error)
}
