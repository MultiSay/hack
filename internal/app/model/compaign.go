package model

import (
	"github.com/go-playground/validator/v10"
)

type Compaign struct {
	ID          int    `json:"id"`
	UTMCampaign string `json:"utm_campaign"`
	Gender      string `json:"gender"`
	AgeFrom     int    `json:"age_from"`
	AgeTo       int    `json:"age_to"`
	City        string `json:"city"`
	Theme       string `json:"theme"`
}

func (u *Compaign) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
