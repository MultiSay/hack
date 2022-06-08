package model

import "time"

type Offer struct {
	ID                     int       `json:"offerId"`
	Date                   time.Time `json:"date"`
	WaveID                 int       `json:"waveId"`
	WaveStart              string    `json:"waveStart"`
	Limit                  int       `json:"limit"`
	LimitFact              int       `json:"limitFact"`
	IsAvailableForTrainees bool      `json:"isAvailableForTrainees"`
	IsAvailableByLimit     bool      `json:"isAvailableByLimit"`
	IsMine                 bool      `json:"isMine"`
	IsEditable             bool      `json:"isEditable"`
}

type ShiftsResponse struct {
	Offers []Offer `json:"offers"` //Массив предложний
}

type ShiftApply struct {
	OfferID int `json:"offerId" validate:"required"`
}
