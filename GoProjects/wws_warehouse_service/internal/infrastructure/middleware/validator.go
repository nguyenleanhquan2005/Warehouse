package middleware

import (
	"app/internal/domain/entity"
	appErrors "app/internal/error"

	"github.com/gin-gonic/gin"
)

// ClaimsValidator contains methods for verifying claims
type ClaimsValidator interface {
	validate(ctx *gin.Context, claim entity.AuthClaim) error
}

type rolesValidator struct {
	requiredRole string
}

// NewRolesValidator return a claims validator for checking whether claims contains specific role
func NewRolesValidator(requiredRole string) ClaimsValidator {
	return &rolesValidator{
		requiredRole: requiredRole,
	}
}

func (v *rolesValidator) validate(_ *gin.Context, claims entity.AuthClaim) error {
	if claims.Role == "" {
		return appErrors.UnauthorizedError{
			Msg:    "role is required",
			UserID: nil,
		}
	}

	if claims.Role != v.requiredRole {
		return appErrors.UnauthorizedError{
			Msg:    "insufficient permissions",
			UserID: nil,
		}
	}

	return nil
}
