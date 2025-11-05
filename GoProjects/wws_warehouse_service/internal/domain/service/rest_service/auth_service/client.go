package auth_service

import (
	"app/pkg/api"
	"app/pkg/api/rest"
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type AuthService interface {
	VerifyToken(ctx context.Context, token string) (*VerifyTokenResponse, error)
}

const (
	verifyTokenPath = "/auth/verify"
)

type authService struct {
	endpoint   string
	apiTimeOut int
	api        rest.API
}

var (
	instance AuthService
	once     sync.Once
)

func NewAuthService(
	authServiceURL string,
	apiTimeOut int,
) AuthService {
	once.Do(func() {
		instance = &authService{
			endpoint:   authServiceURL,
			apiTimeOut: apiTimeOut,
			api:        rest.NewAPI(apiTimeOut),
		}
	})

	return instance
}

func (s *authService) VerifyToken(ctx context.Context, token string) (*VerifyTokenResponse, error) {
	endpoint := s.endpoint + verifyTokenPath

	body, _, err := s.api.Call(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
		api.WithAuthorization(token),
	)
	if err != nil {
		return nil, api.HandleHTTPError(err, "auth")
	}

	var response VerifyTokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
