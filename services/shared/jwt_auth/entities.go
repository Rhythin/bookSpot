package jwt_auth

import "github.com/golang-jwt/jwt/v5"

// string key to set in claims in context
type TokenKey string

// string key to set in token string in context
type RawTokenKey string

type JWTToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserClaims struct {
	ID        string
	UserName  string
	Email     string
	IsAdmin   bool
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	ID string
	jwt.RegisteredClaims
}

type Tokenizer interface {
	GenerateToken(userID string) (JWTToken, error)
	ValidateToken(token JWTToken) (string, error)
}
