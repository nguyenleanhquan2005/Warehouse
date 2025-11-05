package usecase

import (
	"app/internal/domain/entity"
	"app/internal/domain/repository"
	"app/internal/infrastructure/db/query_builder"
	"context"
	"errors"
)

type balanceGetByUserIDUsecase struct {
	balanceRepo repository.IBalanceRepo
}

func NewBalanceListByUserIdUsecase(balanceRepo repository.IBalanceRepo) *balanceGetByUserIDUsecase {
	return &balanceGetByUserIDUsecase{
		balanceRepo: balanceRepo,
	}
}

func (uc *balanceGetByUserIDUsecase) Do(ctx context.Context, userId string) ([]entity.Balance, error) {
	balances, err := uc.balanceRepo.List(ctx, query_builder.NewBalanceAllQuery(userId))
	if err != nil {
		return nil, err
	}
	if balances == nil {
		return nil, errors.New("List balance not found")
	}

	return balances, nil
}
