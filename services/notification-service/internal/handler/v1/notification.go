package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func (h *handlerV1) GetNotifications(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "GetNotifications")
	defer span.End()

	// TODO: get userID from context claims
	userID := ctx.Value("userID").(string)

	// get notifications from service
	notifications, err := h.Service.GetNotifications(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	// write response
	return sendResponse(w, notifications, http.StatusOK)
}

func (h *handlerV1) GetUnreadCount(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "GetUnreadCount")
	defer span.End()

	// TODO: get userID from context claims
	userID := ctx.Value("userID").(string)

	// get unread count from service
	unreadCount, err := h.Service.GetUnreadCount(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	// write response
	return sendResponse(w, unreadCount, http.StatusOK)
}

func (h *handlerV1) MarkAsRead(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "MarkAsRead")
	defer span.End()

	// TODO: get notificationID from url params
	id := chi.URLParam(r, "notificationID")

	if id == "" {
		err := errhandler.NewCustomError(errors.New("notificationID is empty"), http.StatusBadRequest, "notificationID is empty", false)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		customlogger.S().Error("notificationID is empty", "notificationID", id)
		return err
	}

	// mark as read in service
	err = h.Service.MarkAsRead(ctx, id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	// write response
	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) MarkAllAsRead(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("notification-handler")
	ctx, span := tr.Start(r.Context(), "MarkAllAsRead")
	defer span.End()

	// TODO: get userID from context claims
	userID := ctx.Value("userID").(string)

	// mark all as read in service
	err = h.Service.MarkAllAsRead(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	// write response
	return sendResponse(w, nil, http.StatusOK)
}
