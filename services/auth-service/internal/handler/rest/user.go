package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/auth-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *handler) Register(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	tempToken, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message":   "User registered successfully",
		"tempToken": tempToken,
	}

	return sendResponse(w, messages, http.StatusCreated)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) (err error) {
	var req packets.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		customlogger.S().Warnw("failed to decode request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	tempToken, err := h.service.Login(r.Context(), &req)
	if err != nil {
		return err
	}

	return sendResponse(w, tempToken, http.StatusOK)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) (err error) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	search := r.URL.Query().Get("search")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		customlogger.S().Warnw("failed to convert limit to int", "Error", err)
		customlogger.S().Info("using default limit", "Limit", 10)
		limitInt = 10
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		customlogger.S().Warnw("failed to convert offset to int", "Error", err)
		customlogger.S().Info("using default offset", "Offset", 0)
	}

	req := packets.ListUsersRequest{
		Limit:  limitInt,
		Offset: offsetInt,
		Search: search,
	}

	users, err := h.service.GetUsers(r.Context(), &req)
	if err != nil {
		return err
	}

	return sendResponse(w, users, http.StatusOK)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) (err error) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		return errhandler.NewCustomError(errors.New("userID is required"), http.StatusBadRequest, "UserID is required", false)
	}

	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		return err
	}
	return sendResponse(w, user, http.StatusOK)
}

func (h *handler) Logout(w http.ResponseWriter, r *http.Request) (err error) {

	if err := h.service.Logout(r.Context()); err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message": "User logged out successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) (err error) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		return errhandler.NewCustomError(errors.New("userID is required"), http.StatusBadRequest, "UserID is required", false)
	}

	if err := h.service.DeleteUser(r.Context(), userID); err != nil {
		return err
	}

	messages := map[string]interface{}{
		"message": "User deleted successfully",
	}

	return sendResponse(w, messages, http.StatusOK)
}
