package helpers

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"time"
)

func StartConsumer() error {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		logrus.Fatalf("Failed to initialize JetStream: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = js.Stream(ctx, "NOTIFICATIONS")
	if errors.Is(err, nats.ErrStreamNotFound) {
		_, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:     "NOTIFICATIONS",
			Subjects: []string{"NOTIFICATIONS.>"},
		})
		if err != nil {
			logrus.Fatalf("Failed to create stream: %v", err)
		}
		logrus.Infof("Created stream: NOTIFICATIONS")
	} else if err != nil {
		logrus.Fatalf("Error retrieving stream: %v", err)
	}
	return nil

}

func eventConsumer(nc *nats.Conn) (*jetstream.Consumer, error) {

	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := js.Stream(ctx, "NOTIFICATIONS")
	if err != nil {
		return nil, err
	}

	consumer, err := stream.Consumer(ctx, "collection_consumer")
	if err != nil {
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Durable:     "notification_consumer",
			Name:        "notification_consumer",
			Description: "Alerter that receive from consumer",
		})
		if err != nil {
			return nil, err
		}
		logrus.Infof("Created new consumer: notification_consumer")
	} else {
		logrus.Infof("Using existing consumer: notification_consumer")
	}

	return &consumer, nil
}
