package helpers

import (
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"log"
)

func SendCollection(jsonData []byte) error {
	var err error

	// Connect to a server
	// create a nats connection
	nc, _ := nats.Connect(nats.DefaultURL)
	// getting Jetstream context
	jsc, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	//Init stream
	_, err = jsc.AddStream(&nats.StreamConfig{
		Name:     "USERS",             // nom du stream
		Subjects: []string{"USERS.>"}, // tous les sujets sont sous le format "USERS.*"
	})
	if err != nil {
		log.Fatal(err)
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
