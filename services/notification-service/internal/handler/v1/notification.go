package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/notification-service/internal/entities"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

var _ = entities.Notification{}

// GetNotifications godoc
// @Summary      Get user notifications
// @Description  Get a list of notifications for the current user
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.Notification
// @Failure      500  {object}  map[string]interface{}
// @Router       /notifications [get]
func (h *handlerV1) GetNotifications(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "GetNotifications")
	defer span.End()

	// TODO: get userID from context claims
	userID := ctx.Value("userID").(string)

	// get notifications from service
	notifications, err := h.Service.GetNotifications(ctx, userID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	// write response
	return sendResponse(w, notifications, http.StatusOK)
}

// GetUnreadCount godoc
// @Summary      Get unread notifications count
// @Description  Get the number of unread notifications for the current user
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]int
// @Router       /notifications/unread-count [get]
func (h *handlerV1) GetUnreadCount(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "GetUnreadCount")
	defer span.End()

	// TODO: get userID from context claims
	userID := ctx.Value("userID").(string)

	// get unread count from service
	unreadCount, err := h.Service.GetUnreadCount(ctx, userID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	// write response
	return sendResponse(w, unreadCount, http.StatusOK)
}

// MarkAsRead godoc
// @Summary      Mark notification as read
// @Description  Mark a specific notification as read by its ID
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Param        notificationID  path      string  true  "Notification ID"
// @Success      200             {object}  map[string]interface{}
// @Failure      400             {object}  map[string]interface{}
// @Router       /notifications/{notificationID}/read [put]
func (h *handlerV1) MarkAsRead(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "MarkAsRead")
	defer span.End()

	// TODO: get notificationID from url params
	id := chi.URLParam(r, "notificationID")

	if id == "" {
		err := errhandler.NewCustomError(errors.New("notificationID is empty"), http.StatusBadRequest, "notificationID is empty", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Error("notificationID is empty", "notificationID", id)
		return err
	}

	// mark as read in service
	err = h.Service.MarkAsRead(ctx, id)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	// write response
	return sendResponse(w, nil, http.StatusOK)
}

// MarkAllAsRead godoc
// @Summary      Mark all notifications as read
// @Description  Mark all notifications for the current user as read
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /notifications/read-all [put]
func (h *handlerV1) MarkAllAsRead(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "MarkAllAsRead")
	defer span.End()

	// TODO: get userID from context claims
	userID := ctx.Value("userID").(string)

	// mark all as read in service
	err = h.Service.MarkAllAsRead(ctx, userID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	// write response
	return sendResponse(w, nil, http.StatusOK)
}
