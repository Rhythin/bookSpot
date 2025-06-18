package kafkaclient

import "github.com/IBM/sarama"

type Listner struct {
	topic   Topic
	group   string
	client  sarama.Client
	handler func(message *sarama.ConsumerMessage)
}

type ListnerConfig struct {
	Topic   Topic
	Group   string
	Handler func(message *sarama.ConsumerMessage)
}

func NewListener(config *ListnerConfig) Listner {
	return Listner{
		topic: config.Topic,
		group: config.Group,
	}
}

func (l *Listner) Listner() {

}
