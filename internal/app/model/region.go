package model

import (
	"github.com/go-playground/validator/v10"
)

type RegionPredict struct {
	ID                 int    `json:"id"`
	Position           string `json:"position" validate:"required"`
	City               string `json:"city" validate:"required"`
	CurrentClientIndex int    `json:"currentClientIndex"`
	PredictClientIndex int    `json:"predictClientIndex"`
	PredictArpu        int    `json:"predictArpu"`
}

func (u *RegionPredict) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
