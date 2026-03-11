package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

// GetToken godoc
// @Summary      Get auth tokens
// @Description  Exchanges a temporary token for access and refresh tokens
// @Tags         token
// @Accept       json
// @Produce      json
// @Param        tempToken  query     string  true  "Temporary Token"
// @Success      200        {object}  packets.TokenResponse
// @Failure      400        {object}  map[string]interface{}
// @Router       /token [get]
func (h *handler) GetToken(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("auth-handler")
	ctx, span := tr.Start(r.Context(), "GetToken")
	defer span.End()

	tempToken := r.URL.Query().Get("tempToken")
	if tempToken == "" {
		err := errhandler.NewCustomError(errors.New("tempToken is required"), http.StatusBadRequest, "TempToken is required", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("tempToken is required", "tempToken", tempToken)
		return err
	}

	resp, err := h.service.GetToken(ctx, tempToken)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	return sendResponse(w, resp, http.StatusOK)
}

// RevokeToken godoc
// @Summary      Revoke token
// @Description  Revokes tokens for a specific user
// @Tags         token
// @Accept       json
// @Produce      json
// @Param        request  body      packets.RevokeTokenRequest  true  "Revoke Token Request"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /token/revoke [post]
func (h *handler) RevokeToken(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("auth-handler")
	ctx, span := tr.Start(r.Context(), "RevokeToken")
	defer span.End()

	var req packets.RevokeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(&req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.RevokeToken(ctx, &req); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	messages := map[string]interface{}{
		"message": "Token revoked successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

// RefreshToken godoc
// @Summary      Refresh token
// @Description  Refreshes the access token using the refresh token
// @Tags         token
// @Accept       json
// @Produce      json
// @Param        request  body      packets.RevokeTokenRequest  true  "Refresh Token Request (using UserID)"
// @Success      200      {object}  packets.TokenResponse
// @Failure      400      {object}  map[string]interface{}
// @Router       /token/refresh [post]
func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("auth-handler")
	ctx, span := tr.Start(r.Context(), "RefreshToken")
	defer span.End()

	var req packets.RevokeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tracing.RecordSpanError(span, err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	resp, err := h.service.RefreshToken(ctx, req.UserID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	return sendResponse(w, resp, http.StatusOK)
}
