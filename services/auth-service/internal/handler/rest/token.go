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
