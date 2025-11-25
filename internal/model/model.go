package model

import "time"

type Transaction struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Amount      int64     `json:"amount" db:"amount"`
	Type        string    `json:"type" db:"type"`
	Category    string    `json:"category" db:"category"`
	EventDate   time.Time `json:"event_date" db:"event_date"`
}
