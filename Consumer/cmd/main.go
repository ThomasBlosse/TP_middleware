package main

import (
	"Consumer/internal/helpers"
)

func main() {

	err := helpers.StartConsumer()
	if err != nil {
		return
	}
}
