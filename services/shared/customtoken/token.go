package customtoken

type customToken string

type Tokenizer interface {
	GenerateToken(userID string) (customToken, error)
	ValidateToken(token customToken) (string, error)
}
