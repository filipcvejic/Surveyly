package survey

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, survey *Survey) error
	FindByID(ctx context.Context, id uuid.UUID) (*Survey, error)
	//FindAll(ctx context.Context, limit, offset int) ([]*Survey, int64, error)
	//Update(ctx context.Context, survey *Survey) error
	//Delete(ctx context.Context, id uuid.UUID) error
}
