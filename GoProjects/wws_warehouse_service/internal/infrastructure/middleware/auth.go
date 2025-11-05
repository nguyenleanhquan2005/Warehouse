package middleware

import (
	"app/internal/domain/entity"
	"app/internal/domain/service/rest_service/auth_service"
	"app/internal/domain/service/rest_service/user_service"
	"app/internal/infrastructure/presenter"
	"app/pkg/middleware"

	appErrors "app/internal/error"
	apiContext "app/internal/infrastructure/context"

	"github.com/gin-gonic/gin"
)

const (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

type Auth struct {
	authService auth_service.AuthService
	userService user_service.UserService
	jwt         middleware.JWT
	validators  []ClaimsValidator
}

func NewAuth(
	authService auth_service.AuthService,
	userService user_service.UserService,
	jwt middleware.JWT,
	validators ...ClaimsValidator,
) *Auth {
	return &Auth{
		authService: authService,
		userService: userService,
		jwt:         jwt,
		validators:  validators,
	}
}

func (m *Auth) Authenticate(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		presenter.RenderErrors(ctx, appErrors.UnauthorizedError{
			Msg:    "token is empty",
			UserID: nil,
		})
		ctx.Abort()
		return
	}

	var token string
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		token = tokenString[7:]
	}

	claims := entity.AuthClaim{}
	if err := m.jwt.Decrypt(token, &claims, false); err != nil {
		presenter.RenderErrors(ctx, appErrors.UnauthorizedError{
			Msg:    "token is invalid or expired",
			UserID: nil,
		})
		ctx.Abort()

		return
	}

	// execute all validators
	for _, v := range m.validators {
		err := v.validate(ctx, claims)
		if err != nil {
			presenter.RenderErrors(ctx, err)
			ctx.Abort()

			return
		}
	}

	// verify token
	verifyTokenResponse, err := m.authService.VerifyToken(ctx.Request.Context(), token)
	if err != nil {
		return
	}

	if !verifyTokenResponse.IsActive {
		presenter.RenderErrors(ctx, appErrors.UnauthorizedError{
			Msg:    "token is invalid",
			UserID: nil,
		})
		ctx.Abort()

		return
	}

	// get me information from user service
	meResponse, err := m.userService.GetMe(ctx.Request.Context(), token)
	if err != nil {
		return
	}

	// handle vendor id
	meResponse.VendorID = 1

	apiContext.SetVendorID(ctx, meResponse.VendorID)
}
