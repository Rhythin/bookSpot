package jwt_auth

type Token interface {
	GenerateToken(userID string) (JWTToken, error)
	ValidateToken(token JWTToken) (string, error)
}
