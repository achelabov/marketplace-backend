package domain

import "time"

type ShoppingCart struct {
	ID         int
	SessionID  int
	ProductID  int
	CreatedAt  time.Time
	ModefiedAt time.Time
}
