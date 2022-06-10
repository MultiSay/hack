package model

import "time"

type File struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Size      float32   `json:"size"`
	CreateAt  time.Time `json:"create_at"`
	SendAt    time.Time `json:"send_at"`
	ReceiveAt time.Time `json:"receive_at"`
}
