package domain

import "time"

type ShoppingSession struct {
	ID          int
	UserID      int
	Total       uint
	CreatedAt   time.Time
	Modefied_At time.Time
}
