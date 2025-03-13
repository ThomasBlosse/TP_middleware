package main

import (
	"Consumer/internal/helpers"
	"Consumer/internal/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func main() {
	var collections []models.Collection
	//TODO consume collections from mq

	notifications := helpers.GeneratingNotification(collections)
	//TODO send notifiation to alerter
}
