package model

import "time"

type Transaction struct {
	ID          int       `json:"id_transaction"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"` // income | expense
	Category    string    `json:"category"` // personal, work, home, dreams, bullshit, other
	Date        time.Time `json:"date"`
}
