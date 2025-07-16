package publishers

import (
	"os"

	"github.com/rhythin/bookspot/notification-service/internal/events/topics"
	"github.com/rhythin/bookspot/services/shared/kafkaclient"
)

type Publisher struct {
	// we define the induvidal publishers here with
	SamplePublisher kafkaclient.Publisher
}

func LoadPublishers() *Publisher {

	svcCode := os.Getenv("SVC_CODE")
	publishers := &Publisher{
		SamplePublisher: kafkaclient.NewPublisher(topics.NotificationTopic, svcCode),
	}

	kafkaclient.RegisterPublisherTopics(
		publishers.SamplePublisher,
	)

	return publishers
}
