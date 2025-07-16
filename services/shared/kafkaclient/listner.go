package kafkaclient

import (
	"context"
	"sync"

	"github.com/IBM/sarama"
	"github.com/rhythin/bookspot/services/shared/customlogger"
)

type Listner struct {
	topic         Topic
	group         string
	client        sarama.Client
	handler       func(ctx context.Context, headers map[string]string, message *sarama.ConsumerMessage) error
	consumerGroup sarama.ConsumerGroup
	wg            sync.WaitGroup // Used for graceful shutdown - waits for listen() goroutine to finish
	stopChan      chan struct{}
}

type ListnerConfig struct {
	Topic   Topic
	Group   string
	Handler func(ctx context.Context, headers map[string]string, message *sarama.ConsumerMessage) error
}

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct {
	listner *Listner
}

// RegisterListeners registers and starts multiple listeners in goroutines
func RegisterListeners(listeners ...*Listner) error {
	for _, listener := range listeners {
		// Ensure topic exists
		exists, err := topicExists(listener.topic)
		if err != nil {
			customlogger.S().Errorw("Failed to check topic existence", "Error", err, "Topic", listener.topic)
			return err
		}

		if !exists {
			if err := createTopic(listener.topic); err != nil {
				customlogger.S().Errorw("Failed to create topic", "Error", err, "Topic", listener.topic)
				return err
			}
			customlogger.S().Infow("Topic created", "Topic", listener.topic)
		}

		// Start listener in goroutine
		listener.wg.Add(1)
		go func(l *Listner) {
			defer l.wg.Done()
			l.listen()
		}(listener)

		customlogger.S().Infow("Listener started", "Topic", listener.topic, "Group", listener.group)
	}

	customlogger.S().Infow("All listeners registered and started", "ListenerCount", len(listeners))
	return nil
}

// RegisterListenersSlice registers and starts listeners from a slice
func RegisterListenersSlice(listeners []*Listner) error {
	return RegisterListeners(listeners...)
}

// NewListener creates a new listener instance
func NewListener(config *ListnerConfig) (*Listner, error) {
	// Create consumer group
	consumerGroup, err := sarama.NewConsumerGroupFromClient(config.Group, client)
	if err != nil {
		customlogger.S().Errorw("failed to create consumer group", "Error", err, "Group", config.Group)
		return nil, err
	}

	return &Listner{
		topic:         config.Topic,
		group:         config.Group,
		client:        client,
		handler:       config.Handler,
		consumerGroup: consumerGroup,
		stopChan:      make(chan struct{}),
	}, nil
}

// Start starts the listener in a goroutine (deprecated - use RegisterListeners instead)
func (l *Listner) Start() error {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		l.listen()
	}()

	customlogger.S().Infow("Listener started", "Topic", l.topic, "Group", l.group)
	return nil
}

// Stop stops the listener gracefully
// The wait group ensures the listen() goroutine finishes before closing the consumer group
func (l *Listner) Stop() {
	customlogger.S().Infow("Stopping listener", "Topic", l.topic, "Group", l.group)
	close(l.stopChan)
	l.wg.Wait() // Wait for listen() goroutine to finish

	if err := l.consumerGroup.Close(); err != nil {
		customlogger.S().Errorw("failed to close consumer group", "Error", err)
	}
}

// listen is the main listening loop
func (l *Listner) listen() {
	handler := &ConsumerGroupHandler{listner: l}

	for {
		select {
		case <-l.stopChan:
			return
		default:
			// Generate a new context for each consume cycle
			ctx := context.Background()
			topics := []string{string(l.topic)}

			if err := l.consumerGroup.Consume(ctx, topics, handler); err != nil {
				customlogger.S().Errorw("Consumer group consume error", "Error", err, "Topic", l.topic)
				return
			}
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages()
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}

			// Log message reception
			customlogger.S().Infow("Message received",
				"Topic", message.Topic,
				"Partition", message.Partition,
				"Offset", message.Offset,
				"Group", h.listner.group)

			// Extract headers
			headers := make(map[string]string)
			for _, header := range message.Headers {
				headers[string(header.Key)] = string(header.Value)
			}

			// Generate new context for handler
			ctx := context.Background()

			// Call the handler and handle errors
			if err := h.listner.handler(ctx, headers, message); err != nil {
				customlogger.S().Errorw("Handler error",
					"Error", err,
					"Topic", message.Topic,
					"Partition", message.Partition,
					"Offset", message.Offset,
					"Group", h.listner.group)
				// Continue processing other messages even if one fails
			}

			// Mark message as processed
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
