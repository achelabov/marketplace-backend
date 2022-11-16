package repository

import (
	"context"

	"github.com/achelabov/marketplace-backend/internal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
}

type Sessions interface {
	Create(ctx context.Context, userID string, session domain.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Session, error)
	Delete(ctx context.Context, id string) error
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
