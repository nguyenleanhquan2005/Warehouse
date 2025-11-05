package presenter

import "app/internal/domain/entity"

type Balance struct {
	UserID   string `json:"user_id"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Version  int    `json:"version"`
}

func PresentBalance(balance *entity.Balance) *Balance {
	return &Balance{
		UserID:   balance.UserID,
		Currency: balance.Currency,
		Amount:   balance.Amount,
		Version:  balance.Version,
	}
}

func PresentBalances(balances []entity.Balance) *[]Balance {
	presents := make([]Balance, len(balances))
	for i, balance := range balances {
		presents[i] = *PresentBalance(&balance)
	}
	return &presents
}
