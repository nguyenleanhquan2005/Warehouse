package repository

import (
	"app/internal/infrastructure/db/query_builder"
	"context"
)

type IBaseRepo[E any] interface {
	Count(ctx context.Context, queryBuilder query_builder.I) (int64, error)
	Get(ctx context.Context, queryBuilder query_builder.I) (*E, error)
	Create(ctx context.Context, entity *E) error
	List(ctx context.Context, queryBuilder query_builder.I) ([]E, error)
	Update(ctx context.Context, queryBuilder query_builder.I, entity *E) error
}
