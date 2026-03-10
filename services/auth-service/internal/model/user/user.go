package user

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (u *user) CreateUser(ctx context.Context, user *entities.User) error {
	tr := otel.Tracer("auth-model")
	ctx, span := tr.Start(ctx, "CreateUser")
	defer span.End()

	err := u.db.WithContext(ctx).
		Create(user).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		zap.S().Errorw("failed to create user", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to create user", false)
	}

	return nil
}

func (u *user) CheckUserExists(ctx context.Context, email string, username string) (user *entities.User, err error) {
	tr := otel.Tracer("auth-model")
	ctx, span := tr.Start(ctx, "CheckUserExists")
	defer span.End()

	user = &entities.User{}

	err = u.db.WithContext(ctx).
		First(user, "email = ? OR username = ?", email, username).
		Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to check user exists", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to check user exists", false)
	}

	return user, nil
}
