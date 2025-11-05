package context

import "github.com/gin-gonic/gin"

func GetVendorID(ctx *gin.Context) int64 {
	if vendorID, ok := ctx.Get(keyVendorID.String()); ok {
		if vendorID, ok := vendorID.(int64); ok {
			return vendorID
		}
	}
	return 0
}
