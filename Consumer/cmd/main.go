package main

import (
	"Consumer/internal/helpers"
	"fmt"
)

func main() {

	collections := helpers.StartConsumer()
	fmt.Println(collections)
}
