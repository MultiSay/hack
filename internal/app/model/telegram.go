package model

import (
	"github.com/go-playground/validator/v10"
)

type Telegram struct {
	ID           int    `json:"id"`
	NameID       string `json:"name_id" `
	Name         string `json:"name"`
	NSubscribers int    `json:"n_subscribers"`
	Category     string `json:"category"`
}

func (u *Telegram) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
