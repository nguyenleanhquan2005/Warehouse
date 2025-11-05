package mysql

import (
	"app/internal/infrastructure/db/query_builder"
	"context"
	"errors"

	"gorm.io/gorm"
)

type BaseRepo[E any] struct {
	db *gorm.DB
}

func NewRepo[E any](db *gorm.DB) *BaseRepo[E] {
	return &BaseRepo[E]{db: db}
}

func (r *BaseRepo[E]) Count(ctx context.Context, queryBuilder query_builder.I) (int64, error) {
	var countNum int64
	if err := queryBuilder.Query(r.db.WithContext(ctx)).Count(&countNum).Error; err != nil {
		return countNum, err
	}

	return countNum, nil
}

func (r *BaseRepo[E]) Get(ctx context.Context, queryBuilder query_builder.I) (*E, error) {
	var entity E
	if err := queryBuilder.Query(r.db.WithContext(ctx)).Take(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return &entity, err
	}

	return &entity, nil
}

func (r *BaseRepo[E]) Create(ctx context.Context, entity *E) error {
	if err := r.db.WithContext(ctx).Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (r *BaseRepo[E]) List(ctx context.Context, queryBuilder query_builder.I) ([]E, error) {
	var entities []E
	if err := queryBuilder.Query(r.db.WithContext(ctx)).Find(&entities).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return entities, nil
}

func (r *BaseRepo[E]) Update(ctx context.Context, queryBuilder query_builder.I, entity *E) error {
	if err := r.db.WithContext(ctx).Save(entity).Error; err != nil {
		return err
	}

	return nil
}
