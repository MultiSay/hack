package model

type Status struct {
	Status string `json:"status,omitempty"` // Статус
}

type ResponseError struct {
	Error string `json:"error,omitempty"`
}

const (
	StatusOK string = "OK"
)
