package domain

import "time"

type PaymentDetails struct {
	ID         int
	OrderID    int
	Provider   string
	Status     string
	Currency   string
	Amount     int
	CreatedAt  time.Time
	ModefiedAt time.Time
}
