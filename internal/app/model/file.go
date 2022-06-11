package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type File struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Size      float32   `json:"size"`
	CreateAt  time.Time `json:"createAt"`
	SendAt    time.Time `json:"sendAt,omitempty"`
	ReceiveAt time.Time `json:"receiveAt,omitempty"`
}

func (u *File) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
