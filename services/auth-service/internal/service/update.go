package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/jwt_auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) UpdateUser(ctx context.Context, request *packets.UpdateUserRequest) error {
	// auth check: only the user themselves or an admin can update
	claims, ok := jwt_auth.GetClaims(ctx)
	if !ok {
		return errhandler.NewCustomError(errors.New("unauthorized"), http.StatusUnauthorized, "Login required", false)
	}

	if claims.ID != request.UserID && !claims.IsAdmin {
		return errhandler.NewCustomError(errors.New("forbidden"), http.StatusForbidden, "Insufficient permissions", false)
	}

	user, err := s.Model.User.GetUserByID(ctx, request.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errhandler.NewCustomError(errors.New("user not found"), http.StatusNotFound, "User not found", false)
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName

	return s.Model.User.UpdateUser(ctx, user)
}

func (s *service) DeleteUser(ctx context.Context, userID string) error {
	// auth check: only the user themselves or an admin can delete
	claims, ok := jwt_auth.GetClaims(ctx)
	if !ok {
		return errhandler.NewCustomError(errors.New("unauthorized"), http.StatusUnauthorized, "Login required", false)
	}

	if claims.ID != userID && !claims.IsAdmin {
		return errhandler.NewCustomError(errors.New("forbidden"), http.StatusForbidden, "Insufficient permissions", false)
	}

	return s.Model.User.DeleteUser(ctx, userID)
}

func (s *service) ForgotPassword(ctx context.Context, request *packets.ForgotPasswordRequest) error {
	user, err := s.Model.User.CheckUserExists(ctx, request.Email, request.Username)
	if err != nil {
		return err
	}
	if user == nil {
		// return nil to avoid account enumeration
		return nil
	}

	user.TempToken = generateTempToken()
	return s.Model.User.UpdateUser(ctx, user)
}

func (s *service) ResetPassword(ctx context.Context, request *packets.ResetPasswordRequest) error {
	user, err := s.Model.User.GetByTempToken(ctx, request.TempToken)
	if err != nil {
		return err
	}
	if user == nil {
		return errhandler.NewCustomError(errors.New("invalid token"), http.StatusBadRequest, "Invalid or expired reset token", false)
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

	return s.Model.User.UpdateUser(ctx, user)
}

func (s *service) Logout(ctx context.Context) error {
	// For JWT, logout often involves blacklisting on the client side or a server-side bloom filter.
	// If we have a session in DB, we'd clear it.
	
	// Since we are using TempToken as a sort of "pre-auth" session, we can clear it if it exists.
	claims, ok := jwt_auth.GetClaims(ctx)
	if !ok {
		return nil
	}

	user, err := s.Model.User.GetUserByID(ctx, claims.ID)
	if err != nil {
		return err
	}
	if user == nil {
		return nil
	}

	user.TempToken = ""
	return s.Model.User.UpdateUser(ctx, user)
}
