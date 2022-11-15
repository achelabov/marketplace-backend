package domain

import "time"

type PaymentDetails struct {
	ID         string
	OrderID    string
	Provider   string
	Status     string
	Currency   string
	Amount     int
	CreatedAt  time.Time
	ModefiedAt time.Time
}
