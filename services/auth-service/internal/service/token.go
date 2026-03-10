package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/jwt_auth"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

/*
GetToken function is used to generate jwt tokens with userclaims
with a random key  we return both the access and refresh tokens
based on the users which has the tempToken
*/
func (s *service) GetToken(ctx context.Context, tempToken string) (*packets.TokenResponse, error) {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "GetToken")
	defer span.End()

	user, err := s.Model.User.GetByTempToken(ctx, tempToken)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	if user == nil {
		err := errhandler.NewCustomError(errors.New("invalid temp token"), http.StatusUnauthorized, "Invalid or expired session", false)
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	claims := jwt_auth.UserClaims{
		ID:        user.ID,
		UserName:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := s.Tokenizer.GenerateTokens(claims)
	if err != nil {
		err := errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to generate tokens", false)
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	// update the user to remove the temp token
	user.TempToken = ""
	err = s.Model.User.UpdateUser(ctx, user)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	return &packets.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

/*
RefreshToken function is used to refresh the jwt tokens
with a random key  we return both the access and refresh tokens
based on the users which has the refreshToken
*/
func (s *service) RefreshToken(ctx context.Context, refreshToken string) (*packets.TokenResponse, error) {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "RefreshToken")
	defer span.End()

	claims, err := s.Tokenizer.ValidateRefreshToken(refreshToken)
	if err != nil {
		err := errhandler.NewCustomError(err, http.StatusUnauthorized, "Invalid refresh token", false)
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	user, err := s.Model.User.GetUserByID(ctx, claims.ID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	if user == nil || user.IsLocked {
		err := errhandler.NewCustomError(errors.New("user not found or locked"), http.StatusUnauthorized, "Account unavailable", false)
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	userClaims := jwt_auth.UserClaims{
		ID:        user.ID,
		UserName:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := s.Tokenizer.GenerateTokens(userClaims)
	if err != nil {
		err := errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to generate tokens", false)
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	return &packets.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

/*
RevokeToken function is used to revoke the jwt tokens
this function is used to revoke tokens we will use a bloomfilter to revoke tokens
*/
func (s *service) RevokeToken(ctx context.Context, request *packets.RevokeTokenRequest) error {
	tr := otel.Tracer("auth-service")
	ctx, span := tr.Start(ctx, "RevokeToken")
	defer span.End()
	return nil
}
