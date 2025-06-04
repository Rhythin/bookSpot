package user

import (
	"context"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user *entities.User) (err error)
}

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return &user{
		db: db,
	}
}
