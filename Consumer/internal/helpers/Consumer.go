package helpers

import (
	"Consumer/internal/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func StartConsumer() []models.Collection {
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

func consume(consumer jetstream.Consumer) []models.Collection {
	var receivedCollections []models.Collection
	collectionsChan := make(chan models.Collection, 100)
	done := make(chan struct{})
	var wg sync.WaitGroup

	cc, err := consumer.Consume(func(msg jetstream.Msg) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var collection models.Collection

			err := json.Unmarshal(msg.Data(), &collection)
			if err != nil {
				logrus.Errorf("Error unmarshalling collection data: %v", err)
				return
			}

			receivedCollections = append(receivedCollections, collection)

			_ = msg.Ack()
		}()
	})

	if err != nil {
		logrus.Errorf("Error consuming messages: %v", err)
		return nil
	}

	go func() {
		wg.Wait()
		close(collectionsChan)
		close(done)
	}()
	for collection := range collectionsChan {
		receivedCollections = append(receivedCollections, collection)
	}

	<-done
	cc.Stop()

	return receivedCollections
}
