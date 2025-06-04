package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/books-service/internal/model"
)

// Service handles business logic for the books service
type service struct {
	Model     model.Model
	Validator *validator.Validate
}

// NewService creates a new Service instance
func NewService(model model.Model, validator *validator.Validate) *service {
	return &service{
		Model:     model,
		Validator: validator,
	}
}

type Service interface {
}
