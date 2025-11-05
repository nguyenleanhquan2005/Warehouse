package handler

import (
	appError "app/internal/error"
	"app/internal/infrastructure/context"

	"github.com/gin-gonic/gin"
)

func ValidateVendorMatch(ctx *gin.Context, vendorID int64) (int64, error) {
	vendorIDFromContext := context.GetVendorID(ctx)
	if vendorIDFromContext != vendorID {
		return 0, appError.ForbiddenError{
			Msg: "vendor ID mismatch",
		}
	}

	return vendorID, nil
}
