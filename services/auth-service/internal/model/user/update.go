package user

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (u *user) UpdateUser(ctx context.Context, user *entities.User) (err error) {

	err = u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", user.ID).
		Updates(user).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to update user", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update user", false)
	}

	return nil
}
