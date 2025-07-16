package listners

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/notification-service/internal/events/handler"
	"github.com/rhythin/bookspot/notification-service/internal/events/topics"
	"github.com/rhythin/bookspot/notification-service/internal/service"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/kafkaclient"
)

type Listner interface {
	Listen()
}

func StartListners(s service.Service, v *validator.Validate) {

	eventHandler := handler.NewEventHandler(s, v)
	svcCode := os.Getenv("SVC_CODE")

	var listners []*kafkaclient.Listner

	listnerConfigs := []kafkaclient.ListnerConfig{
		{
			Topic:   topics.NotificationTopic,
			Group:   svcCode,
			Handler: eventHandler.SendNotification,
		},
	}

	for _, config := range listnerConfigs {
		listner, err := kafkaclient.NewListener(&config)
		if err != nil {
			customlogger.S().Errorw("failed to create listner", "error", err)
			continue
		}
		listners = append(listners, listner)
	}

	// register the listners
	kafkaclient.RegisterListeners(listners...)
	customlogger.S().Infow("Listners started", "Listners", listners)
}
