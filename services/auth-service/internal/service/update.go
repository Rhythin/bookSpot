package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/jwt_auth"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) UpdateUser(ctx context.Context, request *packets.UpdateUserRequest) error {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "UpdateUser")
	defer span.End()

	// auth check: only the user themselves or an admin can update
	claims, ok := jwt_auth.GetClaims(ctx)
	if !ok {
		err := errhandler.NewCustomError(errors.New("unauthorized"), http.StatusUnauthorized, "Login required", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	if claims.ID != request.UserID && !claims.IsAdmin {
		err := errhandler.NewCustomError(errors.New("forbidden"), http.StatusForbidden, "Insufficient permissions", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	user, err := s.Model.User.GetUserByID(ctx, request.UserID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	if user == nil {
		err := errhandler.NewCustomError(errors.New("user not found"), http.StatusNotFound, "User not found", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName

	err = s.Model.User.UpdateUser(ctx, user)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) DeleteUser(ctx context.Context, userID string) error {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "DeleteUser")
	defer span.End()

	// auth check: only the user themselves or an admin can delete
	claims, ok := jwt_auth.GetClaims(ctx)
	if !ok {
		err := errhandler.NewCustomError(errors.New("unauthorized"), http.StatusUnauthorized, "Login required", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	if claims.ID != userID && !claims.IsAdmin {
		err := errhandler.NewCustomError(errors.New("forbidden"), http.StatusForbidden, "Insufficient permissions", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	err := s.Model.User.DeleteUser(ctx, userID)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) ForgotPassword(ctx context.Context, request *packets.ForgotPasswordRequest) error {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "ForgotPassword")
	defer span.End()

	user, err := s.Model.User.CheckUserExists(ctx, request.Email, request.Username)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	if user == nil {
		// return nil to avoid account enumeration
		return nil
	}

	user.TempToken = generateTempToken()
	err = s.Model.User.UpdateUser(ctx, user)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) ResetPassword(ctx context.Context, request *packets.ResetPasswordRequest) error {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "ResetPassword")
	defer span.End()

	user, err := s.Model.User.GetByTempToken(ctx, request.TempToken)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	if user == nil {
		err := errhandler.NewCustomError(errors.New("invalid token"), http.StatusBadRequest, "Invalid or expired reset token", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	// generate salt
	salt := generateSalt()

	// hash password with salt
	password := request.Password + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		customlogger.S().Errorw("failed to hash password", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to hash password", false)
	}

	user.Password = string(hashedPassword)
	user.Salt = salt
	user.TempToken = "" // clear token after use

	err = s.Model.User.UpdateUser(ctx, user)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) Logout(ctx context.Context) error {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "Logout")
	defer span.End()

	// For JWT, logout often involves blacklisting on the client side or a server-side bloom filter.
	// If we have a session in DB, we'd clear it.
	
	// Since we are using TempToken as a sort of "pre-auth" session, we can clear it if it exists.
	claims, ok := jwt_auth.GetClaims(ctx)
	if !ok {
		return nil
	}

	user, err := s.Model.User.GetUserByID(ctx, claims.ID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	if user == nil {
		return nil
	}

	user.TempToken = ""
	err = s.Model.User.UpdateUser(ctx, user)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}
