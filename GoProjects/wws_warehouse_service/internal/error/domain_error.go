package error

import "net/http"

type Code string

const (
	CodeServiceUnavailable Code = "CODE_SERVICE_UNAVAILABLE"
	CodeInternalServer     Code = "CODE_INTERNAL_SERVER"
	CodeTooManyRequest     Code = "CODE_TOO_MANY_REQUEST"
	CodeContextCancelled   Code = "CODE_CONTEXT_CANCELLED"
	CodeUnauthorized       Code = "CODE_UNAUTHORIZED"
	CodeForbidden          Code = "CODE_FORBIDDEN"
	CodeNotFound           Code = "CODE_NOT_FOUND"
	CodeBadRequest         Code = "CODE_BAD_REQUEST"
)

func (c Code) String() string {
	return string(c)
}

func (c Code) Status() int {
	switch c {
	case CodeServiceUnavailable:
		return http.StatusServiceUnavailable
	case CodeInternalServer:
		return http.StatusInternalServerError
	case CodeTooManyRequest:
		return http.StatusTooManyRequests
	case CodeContextCancelled:
		return http.StatusRequestTimeout
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// SystemError define system error
type DomainError interface {
	Code() Code
	Error() string
}
