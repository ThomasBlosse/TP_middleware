package main

import (
	"Consumer/internal/helpers"
	"Consumer/internal/models"
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"time"
)

var collections []models.Collection

func main() {
	StartConsumer()

	notifications := helpers.GeneratingNotification(collections)
	//TODO send notifiation to alerter
}

func StartConsumer() {

	consumer, err := EventConsumer()
	if err != nil {
		logrus.Warnf("error during nats consumer creation : %v", err)
	} else {
		err = Consume(*consumer)
		if err != nil {
			logrus.Warnf("error during nats consume : %v", err)
		}
	}
}

func EventConsumer() (*jetstream.Consumer, error) {
	// TODO might have to adapt here
	js, _ := jetstream.New(helpers.NatsConn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// get existing stream handle
	stream, err := js.Stream(ctx, "USERS")
	if err != nil {
		return nil, err
	}

	// retrieve consumer handle from a stream

	// getting durable consumer
	consumer, err := stream.Consumer(ctx, "consumer_name_here")
	if err != nil {
		// if doesn't exist, create durable consumer
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Durable:     "consumer_name_here",
			Name:        "consumer_name_here",
			Description: "consumer_name_here consumer",
		})
		if err != nil {
			return nil, err
		}
		logrus.Infof("Created consumer")
	} else {
		logrus.Infof("Got existing consumer")
	}

	return &consumer, nil
}

func Consume(consumer jetstream.Consumer) (err error) {
	// consume messages from the consumer in callback
	cc, err := consumer.Consume(func(msg jetstream.Msg) {
		// TODO your code here
		logrus.Debug(string(msg.Data()))
		_ = msg.Ack()
	})

	// Important so program lasts until consumer connection is closed
	<-cc.Closed()
	cc.Stop()

	return err
}
