package postgres

import (
	"context"

	"github.com/achelabov/marketplace-backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepository struct {
	pool *pgxpool.Pool
}

func NewSessionRepository(pool *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{
		pool: pool,
	}
}

func (r *SessionRepository) Create(ctx context.Context, userID string, session domain.Session) error {
	_, err := r.pool.Exec(ctx, `insert into session (
		user_id, refresh_token, expires_at)
		values($1, $2, $3)`, userID, session.RefreshToken, session.ExpiresAt)

	return err
}

func (r *SessionRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Session, error) {
	var session domain.Session

	err := r.pool.QueryRow(ctx, `select * form session
	where refresh_token=$1`, refreshToken).Scan(
		session.ID, session.UserID, session.RefreshToken, session.ExpiresAt)

	return session, err
}

func (r *SessionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `delete from session where id=$1`, id)

	return err
}
