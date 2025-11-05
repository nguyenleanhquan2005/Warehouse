package usecase

import (
	"app/internal/domain/entity"
	"context"
)

type BalanceListUsecase interface {
	Do(ctx context.Context, userId string) ([]entity.Balance, error)
}

type BalanceGetUsecase interface {
	Do(ctx context.Context, userId, currency string) (*entity.Balance, error)
}
