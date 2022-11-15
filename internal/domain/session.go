package domain

import "time"

type Session struct {
	ID           string
	UserID       string
	RefreshToken string
	ExpiresAt    time.Time
}
