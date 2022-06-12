package model

import (
	"github.com/go-playground/validator/v10"
)

type RegionPredict struct {
	ID                 int     `json:"id"`
	Position           int     `json:"position" validate:"required"`
	City               string  `json:"city" validate:"required"`
	CurrentClientIndex float32 `json:"currentClientIndex,omitempty"`
	PredictClientIndex float32 `json:"predictClientIndex,omitempty"`
	PredictArpu        float32 `json:"predictArpu,omitempty"`
	PredictScore       float32 `json:"predictScore,omitempty"`
	Product            string  `json:"product"`
}

type PredictResult struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []RegionPredict `json:"data"`
}

func (u *RegionPredict) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
