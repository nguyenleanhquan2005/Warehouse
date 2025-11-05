package context

import (
	"github.com/gin-gonic/gin"
)

func setContextKey(ctx *gin.Context, k key, value interface{}) {
	ctx.Set(k.String(), value)
}

func SetVendorID(ctx *gin.Context, vendorID int64) {
	setContextKey(ctx, keyVendorID, vendorID)
}
