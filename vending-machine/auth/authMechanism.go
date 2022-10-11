package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"vending-machine/domain"
)

var secret = []byte("BnFdv[DF9>2c`Oq!!(%^")

type claimskey int

var claimsKey claimskey

type AuthMechanism struct {
}

func NewAuthMechanism() *AuthMechanism {
	return &AuthMechanism{}
}
func (a *AuthMechanism) CreateToken(sub string, user domain.User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	expiration := time.Now().Add(1 * time.Hour)
	token.Claims = &domain.JwtClaims{
		user,
		jwt.RegisteredClaims{
			// Set the exp and sub claims.
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   sub,
		},
	}
	val, err := token.SignedString(secret)
	if err != nil {

		return "", err
	}
	return val, nil
}
func (a *AuthMechanism) GetClaimsFromToken(tokenString string) (*domain.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*domain.JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
func (a *AuthMechanism) SetJWTClaimsContext(ctx context.Context, claims domain.JwtClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}
