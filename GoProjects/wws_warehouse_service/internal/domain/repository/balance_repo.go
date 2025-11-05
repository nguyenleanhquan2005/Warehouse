package repository

import (
	"app/internal/domain/entity"
)

type IBalanceRepo interface {
	IBaseRepo[entity.Balance]
}
