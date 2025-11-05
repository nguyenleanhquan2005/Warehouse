package mysql

import (
	"app/internal/domain/entity"

	"gorm.io/gorm"
)

type BalanceRepo struct {
    *BaseRepo[entity.Balance]
}

func NewBalanceRepository(db *gorm.DB) BalanceRepo {
    return BalanceRepo{
        BaseRepo: &BaseRepo[entity.Balance]{db: db},
    }
}
