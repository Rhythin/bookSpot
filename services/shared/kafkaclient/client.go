package kafkaclient

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/rhythin/bookspot/services/shared/customlogger"
)

var (
	client      sarama.Client
	adminClient sarama.ClusterAdmin
	publisher   sarama.AsyncProducer
)

// kafkaclientConfig is the configuration for the kafka client
type kafkaClientConfig struct {
	Brokers string
}

// initilizes the kafka client using sarama library
func Init(config *kafkaClientConfig) (client sarama.Client, err error) {

	// split the brokers string into a slice of strings
	brokers := strings.Split(config.Brokers, ",")

	saramaConfig := sarama.NewConfig()

	// create a client from the brokers
	client, err = sarama.NewClient(brokers, saramaConfig)
	if err != nil {
		customlogger.S().Errorw("failed to create client", "Error", err)
		return nil, err
	}

	// create an admin client from the client
	adminClient, err = sarama.NewClusterAdminFromClient(client)
	if err != nil {
		customlogger.S().Errorw("failed to create admin client", "Error", err)
		return nil, err
	}

	publisher, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		customlogger.S().Errorw("failed to create publisher", "Error", err)
		return nil, err
	}

	return client, err
}

// createTopic creates a new topic in the kafka cluster
func createTopic(topic Topic) (err error) {

	topicString := string(topic)

	topicDetail := sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	err = adminClient.CreateTopic(topicString, &topicDetail, false)
	if err != nil {
		customlogger.S().Errorw("failed to create topic", "Error", err)
		return err
	}

	return nil
}

// topicExists checks if a topic exists in the kafka cluster
func topicExists(topic Topic) (exists bool, err error) {

	topics, err := client.Topics()
	if err != nil {
		customlogger.S().Errorw("failed to get topics", "Error", err)
		return false, err
	}

	for _, t := range topics {
		if t == string(topic) {
			return true, nil
		}
	}

	return false, nil
}
