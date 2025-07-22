package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *handler) GetToken(w http.ResponseWriter, r *http.Request) (err error) {
	tempToken := r.URL.Query().Get("tempToken")
	if tempToken == "" {
		customlogger.S().Warnw("tempToken is required", "tempToken", tempToken)
		return errhandler.NewCustomError(errors.New("tempToken is required"), http.StatusBadRequest, "TempToken is required", false)
	}

	resp, err := h.service.GetToken(r.Context(), tempToken)
	if err != nil {
		return err
	}
	return sendResponse(w, resp, http.StatusOK)
}

func (h *handler) RevokeToken(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.RevokeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(&req); err != nil {
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.RevokeToken(r.Context(), &req); err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message": "Token revoked successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.RevokeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	resp, err := h.service.RefreshToken(r.Context(), req.UserID)
	if err != nil {
		return err
	}

	return sendResponse(w, resp, http.StatusOK)
}
