package service

import (
	"context"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
)

func (s *service) UpdateUser(ctx context.Context, request *packets.UpdateUserRequest) error {

	return nil
}

func (s *service) DeleteUser(ctx context.Context, userID string) error {
	return nil
}

func (s *service) ForgotPassword(ctx context.Context, request *packets.ForgotPasswordRequest) error {

	return nil
}

func (s *service) ResetPassword(ctx context.Context, request *packets.ResetPasswordRequest) error {
	return nil
}

func (s *service) Logout(ctx context.Context) error {
	return nil
}
