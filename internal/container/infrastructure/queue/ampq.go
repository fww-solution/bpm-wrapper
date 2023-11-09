package queue

import (
	"bpm-wrapper/internal/config"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewSubscriber(cfg *config.MessageQueueConfig) (message.Subscriber, error) {
	ampqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	ampqConfig := amqp.NewDurableQueueConfig(ampqURI)

	subscriber, err := amqp.NewSubscriber(
		ampqConfig,
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		log.Fatal(err)
	}

	return subscriber, err
}

func NewPublisher(cfg *config.MessageQueueConfig) (message.Publisher, error) {
	ampqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	ampqConfig := amqp.NewDurableQueueConfig(ampqURI)

	publisher, err := amqp.NewPublisher(
		ampqConfig,
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		log.Fatal(err)
	}

	return publisher, err
}

func ProcessMessages(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("Got message: %s", string(msg.Payload))
		msg.Ack()
	}
}
