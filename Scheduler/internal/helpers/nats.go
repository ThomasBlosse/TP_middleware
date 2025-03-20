package helpers

import (
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func SendCollection(jsonData []byte) error {
	var err error

	// Connect to a server
	// create a nats connection
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatalf("Error connecting to NATS: %v", err)
	}

	// getting Jetstream context
	jsc, err := nc.JetStream()
	if err != nil {
		logrus.Fatalf("Error while getting context JetStream: %v", err)
	}

	//Init stream
	_, err = jsc.AddStream(&nats.StreamConfig{
		Name:     "USERS",             // nom du stream
		Subjects: []string{"USERS.>"}, // tous les sujets sont sous le format "USERS.*"
	})
	if err != nil {
		logrus.Fatalf("Error while initiating Stream: %v", err)
	}

	pubAckFuture, err := jsc.PublishAsync("USERS.create", jsonData)
	if err != nil {
		logrus.Fatalf("Error while publishing data: %s", err.Error())
	}

	select {
	case <-pubAckFuture.Ok():
		return nil
	case <-pubAckFuture.Err():
		return errors.New(string(pubAckFuture.Msg().Data))
	}
}
