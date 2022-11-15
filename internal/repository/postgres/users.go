package postgres

import (
	"context"

	"github.com/achelabov/marketplace-backend/internal/domain"
	"github.com/achelabov/marketplace-backend/pkg/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	_, err := r.pool.Exec(ctx, `insert into user (
		id, first_name, last_name, middle_name, date_of_birth, 
		phone, email, password, created_at, modefied_at)
		values($1, $2, $3, $4, $5, $6, $7, %8, %9, $10)
		)`, user.ID, user.FirstName, user.LastName, user.MiddleName, user.DateOfBirth,
		user.Phone, user.Email, user.Password, user.CreatedAt, user.ModefiedAt)

	return err
}

func (r *UserRepository) GetByCredentials(ctx context.Context, email, password string) (domain.User, error) {
	var user domain.User

	err := r.pool.QueryRow(ctx, `select * form user
	where email=$1, password=$2`, email, password).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.MiddleName, &user.DateOfBirth,
		&user.Phone, &user.Email, &user.Password, &user.CreatedAt, &user.ModefiedAt)

	return user, err
}

func (r *UserRepository) SetSession(ctx context.Context, userID string, session domain.Session) error {
	_, err := r.pool.Exec(ctx, `insert into session (
		id, user_id, refresh_token, expires_at)
		values($1, $2, $3, $4)`, uuid.Generate(), userID, session.RefreshToken, session.ExpiresAt)

	return err
}
