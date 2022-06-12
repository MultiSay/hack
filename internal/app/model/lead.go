package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Lead struct {
	ID                  int       `json:"id"`
	ClientID            string    `json:"client_id"`
	ProductCategoryName string    `json:"product_category_name"`
	UTMSource           string    `json:"utm_source,omitempty"`
	UTMContent          string    `json:"utm_content,omitempty"`
	UTMCampaing         string    `json:"utm_campaign,omitempty"`
	Date                time.Time `json:"date"`
	CPC                 int       `json:"cpc"`
}

func (u *Lead) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
