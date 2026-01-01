package repositories

import (
	"context"
	sqlc2 "github.com/filipcvejic/surveyly/db/sqlc"
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRepository struct {
	queries *sqlc2.Queries
}

func NewRefreshTokenRepository(queries *sqlc2.Queries) *RefreshTokenRepository {
	return &RefreshTokenRepository{queries: queries}
}

func (r *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, userID uuid.UUID, ttl time.Duration) (sqlc2.RefreshToken, error) {
	tokenID := uuid.New()
	expiresAt := time.Now().Add(ttl)

	return r.queries.CreateRefreshToken(ctx, sqlc2.CreateRefreshTokenParams{
		ID:        tokenID,
		UserID:    userID,
		Token:     tokenID.String(),
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		Revoked:   false,
	})
}

func (r *RefreshTokenRepository) GetRefreshToken(ctx context.Context, tokenString string) (sqlc2.RefreshToken, error) {
	return r.queries.GetRefreshToken(ctx, tokenString)
}

func (r *RefreshTokenRepository) RevokeRefreshToken(ctx context.Context, tokenString string) error {
	return r.queries.RevokeRefreshToken(ctx, tokenString)
}
