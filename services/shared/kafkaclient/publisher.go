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

	// convert the message to json
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		customlogger.S().Errorw("failed to marshal message", "Error", err)
		return
	}

	// prepare the message
	msg := &sarama.ProducerMessage{
		Topic: string(p.Topic),
		Value: sarama.ByteEncoder(jsonMessage),
		Key:   sarama.StringEncoder(p.Group),
	}

	// send the message to the topic and wait for response
	select {
	case publisher.Input() <- msg:
		customlogger.S().Infof("Message request sent to topic cluster '%s'\n", p.Topic)
	case err := <-publisher.Errors():
		customlogger.S().Errorw("failed to produce message", "Error", err)
	}

	// 	wait for success or error acknowledgment from the cluster
	select {
	case <-publisher.Successes():
		customlogger.S().Infof("Message was successfully produced to topic '%s'", p.Topic)
	case err := <-publisher.Errors():
		customlogger.S().Errorw("Failed to produce message", "Error", err)
	}

}
