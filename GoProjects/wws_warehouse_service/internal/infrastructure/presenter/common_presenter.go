package presenter

import (
	"errors"
	"fmt"
	"net/http"

	appErrors "app/internal/error"
	appLog "app/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Paging  interface{} `json:"paging,omitempty"`
}

// ResponsePaging holds paging response information
type ResponsePaging struct {
	Total uint32 `json:"total"`
}

//nolint:unused
type responseError struct {
	Type    string      `json:"type"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Param   interface{} `json:"param,omitempty"`
}

func RenderErrors(ctx *gin.Context, err error) {
	// handler persistent error
	isPersistentErr := handlerPersistentErrors(ctx, err)
	if isPersistentErr {
		return
	}
	// handle domain error
	var domainErr appErrors.DomainError
	if errors.As(err, &domainErr) {
		log5xxError(ctx, domainErr.Code().Status(), err)

		ctx.JSON(domainErr.Code().Status(), response{
			Code:    domainErr.Code().String(),
			Message: domainErr.Error(),
		})

		return
	}

	// handle unknown error - log all
	log5xxError(ctx, 500, err)

	ctx.JSON(appErrors.CodeInternalServer.Status(), response{
		Code:    appErrors.CodeInternalServer.String(),
		Message: fmt.Sprintf("Uknown Error, Detail: %s", err.Error()),
	})
}

// RenderData returns data response
func RenderData(ctx *gin.Context, data, paging interface{}) {
	ctx.JSON(http.StatusOK, response{
		Code:   "SUCCESS",
		Data:   data,
		Paging: paging,
	})
}

func log5xxError(ctx *gin.Context, status int, err error) {
	if status >= 500 {
		appLog.
			WithField("url", ctx.Request.URL.String()).
			WithField("method", ctx.Request.Method).
			WithField("user_agent", ctx.Request.UserAgent()).
			WithField("ip", ctx.ClientIP()).
			WithError(err).Errorln("internal server error")
	}
}

func handlerPersistentErrors(ctx *gin.Context, err error) bool {
	var status int
	var code string
	var msg string
	var mysqlErr *mysql.MySQLError
	if !errors.As(err, &mysqlErr) {
		return false
	}

	switch mysqlErr.Number {
	case 2006:
		status = appErrors.CodeServiceUnavailable.Status()
		code = appErrors.CodeServiceUnavailable.String()
		msg = "mysql server is down"
	case 1045:
		status = appErrors.CodeInternalServer.Status()
		code = appErrors.CodeInternalServer.String()
		msg = "db access denied"
	case 1064:
		status = appErrors.CodeInternalServer.Status()
		code = appErrors.CodeInternalServer.String()
		msg = "sql syntax error"
	case 1146:
		status = appErrors.CodeInternalServer.Status()
		code = appErrors.CodeInternalServer.String()
		msg = "table does not exist"
	case 1062:
		status = CodeConflict.Status()
		code = CodeConflict.String()
		msg = "Duplicate key"
	case 1451:
		status = CodeConflict.Status()
		code = CodeConflict.String()
		msg = "Cannot delete or update a parent row: a foreign key constraint fails"
	case 1452:
		status = CodeConflict.Status()
		code = CodeConflict.String()
		msg = "Cannot add or update a child row: a foreign key constraint fails"
	case 1048:
		status = CodeConflict.Status()
		code = CodeConflict.String()
		msg = "Column not found"
	default:
		status = appErrors.CodeInternalServer.Status()
		code = appErrors.CodeInternalServer.String()
		msg = "unknown db error"
	}

	log5xxError(ctx, status, err)

	ctx.JSON(status, response{
		Code:    code,
		Message: fmt.Sprintf("Msg: %s Detail: %s", msg, err.Error()),
	})

	return true
}
