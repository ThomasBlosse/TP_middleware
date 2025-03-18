package main

import (
	"Consumer/internal/helpers"
	"fmt"
)

func main() {

	collections := helpers.StartConsumer()

	notifications := helpers.GeneratingNotification(collections)

	fmt.Println(notifications)
	//TODO send notifiation to alerter
}
