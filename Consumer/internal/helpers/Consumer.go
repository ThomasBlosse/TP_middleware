package helpers

import (
	"Consumer/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	_, err = js.Stream(ctx, "USERS")
	if errors.Is(err, nats.ErrStreamNotFound) {
		_, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:     "USERS",
			Subjects: []string{"USERS.>"},
		})
		if err != nil {
			logrus.Fatalf("Failed to create stream: %v", err)
		}
		logrus.Infof("Created stream: USERS")
	} else if err != nil {
		logrus.Fatalf("Error retrieving stream: %v", err)
	}

	consumer, err := eventConsumer(nc)
	if err != nil {
		logrus.Warnf("Error during NATS consumer creation: %v", err)
		return nil
	}

	return consume(*consumer)
}

func eventConsumer(nc *nats.Conn) (*jetstream.Consumer, error) {

	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := js.Stream(ctx, "USERS")
	if err != nil {
		return nil, err
	}

	consumer, err := stream.Consumer(ctx, "collection_consumer")
	if err != nil {
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Durable:     "collection_consumer",
			Name:        "collection_consumer",
			Description: "Consumer that receive from scheduler",
		})
		if err != nil {
			return nil, err
		}
		logrus.Infof("Created new consumer: collection_consumer")
	} else {
		logrus.Infof("Using existing consumer: collection_consumer")
	}

	return &consumer, nil
}

func consume(consumer jetstream.Consumer) error {
	cc, err := consumer.Consume(func(msg jetstream.Msg) {
		var receivedCollections []models.Collection

		err := json.Unmarshal(msg.Data(), &receivedCollections)
		if err != nil {
			logrus.Fatalf("Error unmarshalling collection data: %v", err)
		}

		fmt.Println("Collection received")

		notifications := GeneratingNotification(receivedCollections)
		if len(notifications) > 0 {
			jsonData, err := json.Marshal(notifications)
			if err != nil {
				logrus.Fatalf("Error marshalling notifications: %v", err)
			}

			err = SendCollection(jsonData)
			if err != nil {
				logrus.Fatalf("Error while sending collections: %s", err.Error())
			}
			fmt.Println("notifications sent")
		} else {
			fmt.Println("No differences were find")
		}

		logrus.Debug(string(msg.Data()))
		_ = msg.Ack()
	})

	<-cc.Closed()
	cc.Stop()

	return err
}
