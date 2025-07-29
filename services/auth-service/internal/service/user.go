package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities"
	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, request *packets.RegisterRequest) (string, error) {

	// checkexisting user by email or username
	user, err := s.Model.User.CheckUserExists(ctx, request.Email, request.Username)
	if err != nil {
		return "", err
	}
	if user != nil {
		customlogger.S().Warnw("user already exists", "email", request.Email)
		return "", errhandler.NewCustomError(errors.New("user already exists"), http.StatusBadRequest, "User already exists", false)
	}

	// create user
	newUser := &entities.User{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		IsAdmin:   request.IsAdmin,
	}

	// TODO: check password strength

	// generate salt
	salt := generateSalt()

	// hash password with salt
	password := request.Password + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		customlogger.S().Errorw("failed to hash password", "Error", err)
		return "", errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to hash password", false)
	}

	newUser.Password = string(hashedPassword)
	newUser.Salt = salt
	// create a temp token
	newUser.TempToken = generateTempToken()

	// create user
	err = s.Model.User.CreateUser(ctx, newUser)
	if err != nil {
		customlogger.S().Errorw("failed to create user", "Error", err)
		return "", errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to create user", false)
	}

	return newUser.TempToken, nil
}

func (s *service) Login(ctx context.Context, request *packets.LoginRequest) (string, error) {

	// check existing user by email or username
	user, err := s.Model.User.GetByUserName(ctx, request.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		customlogger.S().Warnw("user not found", "username", request.Username)
		return "", errhandler.NewCustomError(errors.New("user not found"), http.StatusNotFound, "User not found", false)
	}

	// compare password
	password := request.Password + user.Salt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		customlogger.S().Warnw("invalid password", "Error", err)
		return "", errhandler.NewCustomError(err, http.StatusUnauthorized, "Invalid password", false)
	}

	// generate temp token
	user.TempToken = generateTempToken()

	// update user
	err = s.Model.User.UpdateUser(ctx, user)
	if err != nil {
		customlogger.S().Errorw("failed to update user", "Error", err)
		return "", errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update user", false)
	}

	return user.TempToken, nil
}

func (s *service) GetUsers(ctx context.Context, request *packets.ListUsersRequest) (*packets.ListUsersResponse, error) {

	return s.Model.User.GetUsers(ctx, request)
}

func (s *service) GetUser(ctx context.Context, userID string) (UserDetails *packets.UserDetails, err error) {
	user, err := s.Model.User.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &packets.UserDetails{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		IsAdmin:       user.IsAdmin,
		IsLocked:      user.IsLocked,
		LoginAttempts: user.LoginAttempts,
		CreatedAt:     user.CreatedAt,
	}, nil
}

// helper function to generate a salt string for password
func generateSalt() string {

	return ""
}

// helper function to generate random temp token string for user
func generateTempToken() string {
	return ""
}
