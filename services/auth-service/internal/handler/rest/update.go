package rest

import (
	"encoding/json"
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

// UpdateUser godoc
// @Summary      Update user
// @Description  Update details for a specific user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userID   path      string                    true  "User ID"
// @Param        request  body      packets.UpdateUserRequest  true  "Update User Request"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /user/{userID} [put]
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("auth-handler")
	ctx, span := tr.Start(r.Context(), "UpdateUser")
	defer span.End()

	var req packets.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.UpdateUser(ctx, &req); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	messages := map[string]interface{}{
		"message": "User updated successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

// ForgotPassword godoc
// @Summary      Forgot password
// @Description  Request a password reset email
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request  body      packets.ForgotPasswordRequest  true  "Forgot Password Request"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /user/forgot-password [post]
func (h *handler) ForgotPassword(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("auth-handler")
	ctx, span := tr.Start(r.Context(), "ForgotPassword")
	defer span.End()

	var req packets.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.ForgotPassword(ctx, &req); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	messages := map[string]interface{}{
		"message": "reset password email sent successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Reset user password using a temporary token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request  body      packets.ResetPasswordRequest  true  "Reset Password Request"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /user/reset-password [post]
func (h *handler) ResetPassword(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("auth-handler")
	ctx, span := tr.Start(r.Context(), "ResetPassword")
	defer span.End()

	var req packets.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.validator.Struct(req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	if err := h.service.ResetPassword(ctx, &req); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	messages := map[string]interface{}{
		"message": "Password reset successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}
