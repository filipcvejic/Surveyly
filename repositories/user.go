package repositories

import (
	"context"
	sqlc2 "github.com/filipcvejic/surveyly/db/sqlc"
	"github.com/google/uuid"
	"time"
)

type UserRepository struct {
	queries *sqlc2.Queries
}

func NewUserRepository(queries *sqlc2.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) CreateUser(ctx context.Context, email, username, passwordHash string) (sqlc2.User, error) {
	return r.queries.CreateUser(ctx, sqlc2.CreateUserParams{
		ID:           uuid.New(),
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	})
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (sqlc2.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc2.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}
