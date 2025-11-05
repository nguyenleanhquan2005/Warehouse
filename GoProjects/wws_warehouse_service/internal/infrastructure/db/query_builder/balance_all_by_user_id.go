package query_builder

import (
	"app/internal/domain/entity"

	"gorm.io/gorm"
)

type balanceByID struct {
	userID string
}

func NewBalanceAllQuery(userID string) I {
	return balanceByID{userID}
}

func (builder balanceByID) Query(db *gorm.DB) *gorm.DB {
	return db.Model(&entity.Balance{}).
		Where("user_id = ?", builder.userID).
		Order("updated_at DESC")
}
