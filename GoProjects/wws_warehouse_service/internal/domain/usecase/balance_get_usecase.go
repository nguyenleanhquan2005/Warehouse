package usecase

import (
	"app/internal/domain/entity"
	"app/internal/domain/repository"
	appErrors "app/internal/error"
	"app/internal/infrastructure/db/query_builder"
	"context"
)

type balanceGetUsecase struct {
	balanceRepo repository.IBalanceRepo
}

func NewBalanceGetUsecase(balanceRepo repository.IBalanceRepo) *balanceGetUsecase {
	return &balanceGetUsecase{
		balanceRepo: balanceRepo,
	}
}

func (uc *balanceGetUsecase) Do(ctx context.Context, userId, currency string) (*entity.Balance, error) {
	balance, err := uc.balanceRepo.Get(ctx, query_builder.NewBalanceRequestQuery(userId, currency))
	if err != nil {
		return nil, err
	}

	if balance == nil {
		return nil, appErrors.NotFoundError{
			Msg:    "query balance",
			Entity: "balance",
			ID:     nil,
		}
	}

	return balance, nil
}
