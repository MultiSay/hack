package store

import (
	"context"
	"hack/internal/app/model"
)

//go:generate mockery --name=FileRepository --structname=FileRepository
type FileRepository interface {
	Create(context.Context, *model.File) error
	Update(context.Context, *model.File) error
	GetByID(context.Context, int) (*model.File, error)
}

type RegionRepository interface {
	PredictList(context.Context) ([]model.RegionPredict, error)
}
