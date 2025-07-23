package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/auth-service/internal/model"
)

type service struct {
	Model     model.Model
	Validator *validator.Validate
}

func New(model model.Model, validator *validator.Validate) Service {
	return &service{
		Model:     model,
		Validator: validator,
	}
}

type Service interface {
	CreateUser(ctx context.Context, request *packets.RegisterRequest) (string, error)
	Login(ctx context.Context, request *packets.LoginRequest) (string, error)
	GetUsers(ctx context.Context, request *packets.ListUsersRequest) (*packets.ListUsersResponse, error)
	GetUser(ctx context.Context, userID string) (user *packets.UserDetails, err error)
	UpdateUser(ctx context.Context, request *packets.UpdateUserRequest) error
	DeleteUser(ctx context.Context, userID string) error
	ForgotPassword(ctx context.Context, request *packets.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, request *packets.ResetPasswordRequest) error
	Logout(ctx context.Context) error
	GetToken(ctx context.Context, tempToken string) (*packets.TokenResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*packets.TokenResponse, error)
	RevokeToken(ctx context.Context, request *packets.RevokeTokenRequest) error
}
