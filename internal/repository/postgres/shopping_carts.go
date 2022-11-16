package postgres

import (
	"context"

	"github.com/achelabov/marketplace-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShoppingCartRepository struct {
	pool *pgxpool.Pool
}

func NewShoppingCartRepository(pool *pgxpool.Pool) *ShoppingCartRepository {
	return &ShoppingCartRepository{
		pool: pool,
	}
}

func (r *ShoppingCartRepository) Create(ctx context.Context, shoppingCart domain.ShoppingCart) error {
	_, err := r.pool.Exec(ctx, `insert into shopping_cart (
		session_id, product_id, created_at, modefied_at)
		values($1, $2, $3)`,
		shoppingCart.SessionID, shoppingCart.ProductID, shoppingCart.CreatedAt, shoppingCart.ModefiedAt)

	return err
}
