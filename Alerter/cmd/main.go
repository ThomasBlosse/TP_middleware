package main

import "Alerter/internal/helpers"

func main() {

	err := helpers.StartConsumer()
	if err != nil {
		return
	}

}
