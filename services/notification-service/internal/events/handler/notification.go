package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *eventHandler) SendNotification(ctx context.Context, headers map[string]string, message *sarama.ConsumerMessage) error {

	var req packets.CreateNotificationDetails

	// unmarshal the message
	if err := json.Unmarshal(message.Value, &req); err != nil {

		customlogger.S().Errorw("failed to unmarshal message", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "failed to unmarshal message", true)
	}

	// validate the request
	if err := h.validator.Struct(req); err != nil {
		customlogger.S().Errorw("failed to validate request", "error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "failed to validate request", true)
	}

	// create the notification
	if err := h.service.CreateNotification(ctx, &req); err != nil {
		customlogger.S().Errorw("failed to create notification", "error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "failed to create notification", true)
	}

	return nil
}
