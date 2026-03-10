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

func (u *user) DeleteUser(ctx context.Context, userID string) error {
	tr := otel.Tracer("auth-model")
	ctx, span := tr.Start(ctx, "DeleteUser")
	defer span.End()

	err := u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", userID).
		Delete(&entities.User{}).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to delete user", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to delete user", false)
	}

	return nil
}
