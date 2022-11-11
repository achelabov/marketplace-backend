package domain

import "time"

type User struct {
	ID          int
	FirstName   string
	LastName    string
	MiddleName  string
	DateOfBirth time.Time
	Phone       string
	Email       string
	Password    string
	CreatedAt   time.Time
	ModefiedAt  time.Time
}

type UserAddress struct {
	ID           int
	UserID       int
	AddressLine1 string
	AddressLine2 string
	City         string
	PostalCode   string
	Country      string
}

type UserPayment struct {
	ID          int
	UserID      int
	PaymentType string
	Provider    string
}
