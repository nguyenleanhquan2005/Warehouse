package user_service

import (
	"app/pkg/api"
	"app/pkg/api/rest"
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

const (
	getMePath = "/me"
)

type UserService interface {
	GetMe(ctx context.Context, token string) (*UserInformationResponse, error)
}

type userService struct {
	endpoint   string
	apiTimeOut int
	api        rest.API
}

var (
	instance UserService
	once     sync.Once
)

func NewUserService(
	userServiceURL string,
	apiTimeOut int,
) UserService {
	once.Do(func() {
		instance = &userService{
			endpoint:   userServiceURL,
			apiTimeOut: apiTimeOut,
			api:        rest.NewAPI(apiTimeOut),
		}
	})

	return instance
}

func (s *userService) GetMe(ctx context.Context, token string) (*UserInformationResponse, error) {
	endpoint := s.endpoint + getMePath

	body, _, err := s.api.Call(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
		api.WithAuthorization(token),
	)
	if err != nil {
		return nil, api.HandleHTTPError(err, "user")
	}

	var response UserInformationResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
