package auth

import (
	"context"
	"net/http"
	"strings"
	"vending-machine/domain"
)

type AuthorizationMw struct {
	claimsDomain domain.JwtClaimsInterface
}
type AuthorizationMechanism interface {
	BuyerAuthorization(next http.Handler) http.Handler
	SellerAuthorization(next http.Handler) http.Handler
}

func NewAuthorizationMiddleware(c domain.JwtClaimsInterface) *AuthorizationMw {
	return &AuthorizationMw{
		claimsDomain: c,
	}
}

func (a *AuthorizationMw) BuyerAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.claimsDomain.GetClaimsFromToken(tokenString)
		if err != nil {
			http.Error(w, "Error getting claims from token", http.StatusUnauthorized)
			return
		}

		if claims.User.Role != "buyer" {
			http.Error(w, "User role must be buyer", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(a.claimsDomain.SetJWTClaimsContext(r.Context(), *claims))
		next.ServeHTTP(w, r)

	})
}

func (a *AuthorizationMw) SellerAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := a.claimsDomain.GetClaimsFromToken(tokenString)
		if err != nil {
			http.Error(w, "Error getting claims from token", http.StatusUnauthorized)
			return
		}
		claimsUser := claims.User
		if claimsUser.Role != "seller" {
			http.Error(w, "User role must be seller", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "requestUser", claimsUser)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
