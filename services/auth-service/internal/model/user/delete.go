package user

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (u *user) DeleteUser(ctx context.Context, userID string) error {

	err := u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", userID).
		Delete(&entities.User{}).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to delete user", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to delete user", false)
	}

	return nil
}
