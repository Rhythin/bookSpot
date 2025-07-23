package service

import (
	"context"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
)

/*
GetToken function is used to generate jwt tokens with userclaims
with a random key  we return both the access and refresh tokens
based on the users which has the tempToken
*/
func (s *service) GetToken(ctx context.Context, tempToken string) (*packets.TokenResponse, error) {
	return nil, nil
}

/*
RefreshToken function is used to refresh the jwt tokens
with a random key  we return both the access and refresh tokens
based on the users which has the refreshToken
*/
func (s *service) RefreshToken(ctx context.Context, refreshToken string) (*packets.TokenResponse, error) {
	return nil, nil
}

/*
RevokeToken function is used to revoke the jwt tokens
this function is used to revoke tokens we will use a bloomfilter to revoke tokens
*/
func (s *service) RevokeToken(ctx context.Context, request *packets.RevokeTokenRequest) error {
	return nil
}
