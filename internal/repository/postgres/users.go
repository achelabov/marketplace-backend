package postgres

import (
	"context"

	"github.com/achelabov/marketplace-backend/internal/domain"
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
		first_name, last_name, middle_name, date_of_birth, 
		phone, email, password, created_at, modefied_at)
		values($1, $2, $3, $4, $5, $6, $7, %8, %9)`,
		user.FirstName, user.LastName, user.MiddleName, user.DateOfBirth,
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

func (r *UserRepository) SetAddress(ctx context.Context, address domain.UserAddress) error {
	_, err := r.pool.Exec(ctx, `insert into user_address (
		user_id, address_line1, address_line2, city, postal_code, country)
		values($1, $2, $3, $4, $5, $6)`,
		address.UserID, address.AddressLine1, address.AddressLine2,
		address.City, address.PostalCode, address.Country)

	return err
}

func (r *UserRepository) SetPayment(ctx context.Context, payment domain.UserPayment) error {
	_, err := r.pool.Exec(ctx, `insert into user_payment (
		user_id, payment_type, provider) 
		values($1, $2, $3)`,
		payment.UserID, payment.PaymentType, payment.Provider)

	return err
}
