package user

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"gorm.io/gorm"
)

func (u *user) GetUserByID(ctx context.Context, userID string) (user *entities.User, err error) {

	err = u.db.WithContext(ctx).
		First(&user, userID).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		customlogger.S().Errorw("failed to get user by id", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get user by id", false)
	}

	return user, nil
}

func (u *user) GetByUserName(ctx context.Context, username string) (user *entities.User, err error) {

	err = u.db.WithContext(ctx).
		First(&user, "username = ?", username).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		customlogger.S().Errorw("failed to get user by username", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get user by username", false)
	}

	return user, nil
}

func (u *user) GetUsers(ctx context.Context, request *packets.ListUsersRequest) (*packets.ListUsersResponse, error) {

	var users []*packets.UserDetails

	resp := &packets.ListUsersResponse{}

	tx := u.db.WithContext(ctx).Model(&entities.User{})

	if request.Search != "" {
		tx = tx.Where("username LIKE ?", "%"+request.Search+"%")
	}

	tx.Count(&resp.TotalCount)

	if request.Limit != 0 {
		tx = tx.Limit(request.Limit)
	}

	if request.Offset != 0 {
		tx = tx.Offset(request.Offset)
	}

	err := tx.Count(&resp.SearchCount).
		Find(&users).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to get users", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get users", false)
	}

	resp.Users = users
	return resp, nil
}

func (u *user) GetByTempToken(ctx context.Context, tempToken string) (user *entities.User, err error) {

	err = u.db.WithContext(ctx).
		First(&user, "temp_token = ?", tempToken).
		Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		customlogger.S().Errorw("failed to get user by temp token", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get user by temp token", false)
	}

	return user, nil
}
