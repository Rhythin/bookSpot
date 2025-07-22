package rest

import (
	"encoding/json"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(req); err != nil {
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.UpdateUser(r.Context(), &req); err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message": "User updated successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

func (h *handler) ForgotPassword(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(req); err != nil {
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.ForgotPassword(r.Context(), &req); err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message": "reset password email sent successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

func (h *handler) ResetPassword(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(req); err != nil {
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.ResetPassword(r.Context(), &req); err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message": "Password reset successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}
