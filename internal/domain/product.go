package domain

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	Category    string
	Price       uint
	CreatedAt   time.Time
	ModefiedAt  time.Time
}
