package model

import (
	"github.com/go-playground/validator/v10"
)

type RegionPredict struct {
	ID                 int     `json:"id"`
	Position           string  `json:"position" validate:"required"`
	City               string  `json:"city" validate:"required"`
	CurrentClientIndex int     `json:"currentClientIndex,omitempty"`
	PredictClientIndex int     `json:"predictClientIndex,omitempty"`
	PredictArpu        int     `json:"predictArpu,omitempty"`
	PredictScore       float32 `json:"predictScore,omitempty"`
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
