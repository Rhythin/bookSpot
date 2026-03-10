package jwt_auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/rhythin/bookspot/services/shared/errhandler"
)

type Middleware struct {
	tokenizer Tokenizer
}

func NewMiddleware(tokenizer Tokenizer) *Middleware {
	return &Middleware{
		tokenizer: tokenizer,
	}
}

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errhandler.HttpErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
				return errhandler.NewCustomError(nil, http.StatusUnauthorized, "Missing authorization header", false)
			})(w, r)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			errhandler.HttpErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
				return errhandler.NewCustomError(nil, http.StatusUnauthorized, "Invalid authorization header format", false)
			})(w, r)
			return
		}

		tokenStr := parts[1]
		claims, err := m.tokenizer.ValidateToken(tokenStr)
		if err != nil {
			errhandler.HttpErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
				return errhandler.NewCustomError(err, http.StatusUnauthorized, "Invalid token", false)
			})(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		ctx = context.WithValue(ctx, RawTokenStrKey, tokenStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
