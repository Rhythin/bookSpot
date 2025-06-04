package service

import "github.com/rhythin/bookspot/auth-service/internal/model"

type Service interface {
}

type service struct {
	model model.Model
}

func New(model *model.Model) Service {
	return &service{
		model: *model,
	}
}
