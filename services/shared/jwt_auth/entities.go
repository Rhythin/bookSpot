package jwt_auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

// string key to set in claims in context
type TokenKey string

const (
	UserClaimsKey TokenKey = "user_claims"
)

// string key to set in token string in context
type RawTokenKey string

const (
	RawTokenStrKey RawTokenKey = "raw_token"
)

type JWTToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserClaims struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type Tokenizer interface {
	GenerateTokens(claims UserClaims) (JWTToken, error)
	ValidateToken(tokenStr string) (*UserClaims, error)
	ValidateRefreshToken(tokenStr string) (*RefreshClaims, error)
}

// GetClaims retrieves the user claims from a context.
func GetClaims(ctx context.Context) (*UserClaims, bool) {
	claims, ok := ctx.Value(UserClaimsKey).(*UserClaims)
	return claims, ok
}

// WithClaims adds user claims to a context.
func WithClaims(ctx context.Context, claims *UserClaims) context.Context {
	return context.WithValue(ctx, UserClaimsKey, claims)
}

// GetRawToken retrieves the raw token string from a context.
func GetRawToken(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(RawTokenStrKey).(string)
	return token, ok
}

// WithRawToken adds the raw token string to a context.
func WithRawToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, RawTokenStrKey, token)
}
