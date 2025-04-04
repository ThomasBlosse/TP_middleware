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
