package service

import (
	"context"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
)

func (s *service) GetToken(ctx context.Context, tempToken string) (*packets.TokenResponse, error) {
	return nil, nil
}

func (s *service) RefreshToken(ctx context.Context, refreshToken string) (*packets.TokenResponse, error) {
	return nil, nil
}

func (s *service) RevokeToken(ctx context.Context, request *packets.RevokeTokenRequest) error {
	return nil
}
