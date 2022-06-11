package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type File struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Size      float32   `json:"size"`
	CreateAt  time.Time `json:"create_at"`
	SendAt    time.Time `json:"send_at,omitempty"`
	ReceiveAt time.Time `json:"receive_at,omitempty"`
}

func (u *File) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
