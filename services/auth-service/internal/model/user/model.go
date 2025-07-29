package user

import (
	"context"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user *entities.User) (err error)

	GetUserByID(ctx context.Context, userID string) (user *entities.User, err error)
	GetByUserName(ctx context.Context, username string) (user *entities.User, err error)
	GetUsers(ctx context.Context, request *packets.ListUsersRequest) (*packets.ListUsersResponse, error)
	CheckUserExists(ctx context.Context, email string, username string) (user *entities.User, err error)
	GetByTempToken(ctx context.Context, tempToken string) (user *entities.User, err error)

	UpdateUser(ctx context.Context, user *entities.User) (err error)

	DeleteUser(ctx context.Context, userID string) error
}

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return &user{
		db: db,
	}
}
