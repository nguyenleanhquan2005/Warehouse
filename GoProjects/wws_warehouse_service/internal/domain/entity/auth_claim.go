package entity

import "github.com/golang-jwt/jwt"

type AuthClaim struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
