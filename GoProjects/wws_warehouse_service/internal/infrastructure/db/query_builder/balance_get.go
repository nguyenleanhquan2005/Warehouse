package query_builder

import (
	"app/internal/domain/entity"

	"gorm.io/gorm"
)

type balanceRequest struct {
	userID   string
	currency string
}

func NewBalanceRequestQuery(userID, currency string) balanceRequest {
	return balanceRequest{
		userID:   userID,
		currency: currency}
}

func (builder balanceRequest) Query(db *gorm.DB) *gorm.DB {
	return db.Model(&entity.Balance{}).
		Where("user_id = ? AND currency = ?", builder.userID, builder.currency)
}
