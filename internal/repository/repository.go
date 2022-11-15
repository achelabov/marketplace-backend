package repository

import (
	"context"

	"github.com/achelabov/marketplace-backend/internal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.User, error)
	SetSession(ctx context.Context, userID string, session domain.Session)
}

type Products interface {
}

type ShoppingCarts interface {
}

type OrderItems interface {
}

type OrderDetails interface {
}

type PaymantDetails interface {
}

type Sessions interface {
}
