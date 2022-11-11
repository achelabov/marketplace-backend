package domain

import "time"

type OrderDetails struct {
	ID         int
	UserID     int
	PaymentID  int
	Total      uint
	CreatedAt  time.Time
	ModefiedAt time.Time
}
