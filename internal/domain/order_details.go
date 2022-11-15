package domain

import "time"

type OrderDetails struct {
	ID         string
	UserID     string
	PaymentID  string
	Total      uint
	CreatedAt  time.Time
	ModefiedAt time.Time
}
