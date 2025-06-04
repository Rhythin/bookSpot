package model

import (
	"github.com/rhythin/bookspot/auth-service/internal/model/user"
	"gorm.io/gorm"
)

type Model struct {
	User user.User
}

func New(db *gorm.DB) *Model {
	return &Model{
		User: user.New(db),
	}
}
