package handler

import (
	"app/internal/domain/usecase"
	appError "app/internal/error"
	"app/internal/infrastructure/presenter"

	"github.com/gin-gonic/gin"
)

type BalanceHandler struct {
	listUsecase       usecase.BalanceListUsecase
	balanceGetUsecase usecase.BalanceGetUsecase
}

func NewBalanceHandler(
	listUsecase usecase.BalanceListUsecase,
	balanceGetUsecase usecase.BalanceGetUsecase) *BalanceHandler {
	return &BalanceHandler{
		listUsecase:       listUsecase,
		balanceGetUsecase: balanceGetUsecase}
}

// ListBalanceByUserID godoc
// @Summary List Balances by User ID
// @Description Get a list of balances for a specific user
// @Tags balances
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} []presenter.Balance
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/balances/list [get]
func (h BalanceHandler) ListBalanceByUserID(ctx *gin.Context) {
	userID := ctx.Query("user_id")

	balances, err := h.listUsecase.Do(ctx.Request.Context(), userID)
	if err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	presenter.RenderData(ctx, presenter.PresentBalances(balances), nil)
}

// GetBalance godoc
// @Summary Get Balance by User ID and Currency
// @Description Get a specific balance for a user and currency
// @Tags balances
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param currency path string true "Currency code (e.g., gold, silver, diamond)"
// @Success 200 {object} presenter.Balance
// @Failure 400 {object} map[string]interface{} "Invalid path parameters"
// @Failure 404 {object} map[string]interface{} "Balance not found"
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/balances/{user_id}/{currency} [get]
func (h BalanceHandler) GetBalance(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	currency := ctx.Param("currency")
	if userID == "" || currency == "" {
		presenter.RenderErrors(ctx, appError.InvalidArgumentError{
			Msg:   "missing required path parameter(s)",
			Field: "path parameter",
		})
		return
	}
	balance, err := h.balanceGetUsecase.Do(ctx.Request.Context(), userID, currency)
	if err != nil {
		presenter.RenderErrors(ctx, err)
		return
	}

	presenter.RenderData(ctx, presenter.PresentBalance(balance), nil)
}

// ListBalances godoc
// @Summary List Balances
// @Description Get a list of Balances by user Id
// @Tags balances
// @Accept json
// @Produce json
// @Success 200 {array} []presenter.Balance
// @Router /api/v1/balances [get]
func (h BalanceHandler) UpdateBalanceByUserID(ctx *gin.Context) {

}
