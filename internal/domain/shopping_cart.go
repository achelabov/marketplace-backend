package domain

import "time"

type ShoppingCart struct {
	ID         string
	SessionID  string
	ProductID  string
	CreatedAt  time.Time
	ModefiedAt time.Time
}
