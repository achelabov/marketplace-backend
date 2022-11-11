package domain

import "time"

type OrderItems struct {
	ID         int
	OrderID    int
	ProductID  int
	CreatedAt  time.Time
	ModefiedAt time.Time
}
