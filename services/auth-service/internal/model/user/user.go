package user

import (
	"context"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"go.uber.org/zap"
)

func (u *user) CreateUser(ctx context.Context, user *entities.User) error {

	err := u.db.WithContext(ctx).
		Create(user).
		Error

	if err != nil {
		zap.S().Errorw("failed to create user", "Error", err)
		return err
	}

	return nil
}
