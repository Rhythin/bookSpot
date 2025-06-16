package kafkaclient

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/rhythin/bookspot/services/shared/customlogger"
)

type Publisher struct {
	Topic Topic
	Group string
}

func RegisterPublishers(publishers ...Publisher) {

	for _, publisher := range publishers {
		// check if topic exists in the if not create it
		exists, err := topicExists(publisher.Topic)
		if err != nil {
			return
		}

		if !exists {
			if err := createTopic(publisher.Topic); err != nil {
				return
			}
		}
	}

}

func NewPublisher(topic Topic, group string) Publisher {
	return Publisher{
		Topic: topic,
		Group: group,
	}
}

func (p *Publisher) Publish(ctx context.Context, message any) {

	// create a producer from the client
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		customlogger.S().Errorw("failed to create producer", "Error", err)
		return
	}
	defer producer.Close()

	// convert the message to json
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		customlogger.S().Errorw("failed to marshal message", "Error", err)
		return
	}

	// send the message to the topic
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: string(p.Topic),
		Value: sarama.ByteEncoder(jsonMessage),
	})
	if err != nil {
		customlogger.S().Errorw("failed to send message", "Error", err)
		return
	}

}
