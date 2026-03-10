package user

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

func (u *user) UpdateUser(ctx context.Context, user *entities.User) (err error) {
	tr := otel.Tracer("auth-model")
	ctx, span := tr.Start(ctx, "UpdateUser")
	defer span.End()

	err = u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", user.ID).
		Updates(user).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to update user", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update user", false)
	}

	return nil
}
