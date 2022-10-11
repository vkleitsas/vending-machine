package domain

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}
type JwtClaimsInterface interface {
	CreateToken(sub string, user User) (string, error)
	GetClaimsFromToken(tokenString string) (*JwtClaims, error)
	SetJWTClaimsContext(ctx context.Context, claims JwtClaims) context.Context
}
