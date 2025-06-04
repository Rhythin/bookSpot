package user

import (
	"context"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user *User) (err error)
}

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return &user{
		db: db,
	}
}
