package domain

import "time"

type OrderItems struct {
	ID         string
	OrderID    string
	ProductID  string
	CreatedAt  time.Time
	ModefiedAt time.Time
}
