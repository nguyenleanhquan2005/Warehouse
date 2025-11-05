package auth_service

type VerifyTokenResponse struct {
	IsActive bool `json:"active"`
}
