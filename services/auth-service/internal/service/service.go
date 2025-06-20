package service

import (
	"github.com/go-playground/validator/v10"
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
	// CreateUser(ctx context.Context, user *entities.User) error
}
