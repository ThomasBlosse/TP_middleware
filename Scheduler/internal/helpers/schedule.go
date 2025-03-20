package helpers

import (
	"Scheduler/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/scheduler"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func action(ctx context.Context) {
	formattedResources := GetRessource()

	url := fmt.Sprintf("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=2&calType=ical&nbWeeks=8&displayConfigId=128", formattedResources)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Fatalf("Error while fetching calendar data: %s", err.Error())
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {

		logrus.Fatalf("Error while reading calendar data: %s", err.Error())
	}

	eventArray, err := ParseICalEvents(rawData)
	if err != nil {
		logrus.Fatalf("Error parsing calendar: %s", err)
	}

	collections, err := ConvertEventsToCollections(eventArray)
	if err != nil {
		logrus.Fatalf("Error converting events: %s", err)
	}

	jsonData, err := json.Marshal(collections)
	if err != nil {
		logrus.Fatalf("Error while marshalling collections: %s", err.Error())
	}
	var testData []models.Collection
	err = json.Unmarshal(jsonData, &testData)

	err = SendCollection(jsonData)
	if err != nil {
		logrus.Fatalf("Error while sending collections: %s", err.Error())
	}

	fmt.Println("message sent")
}

func Schedule() {
	ctx := context.Background()
	sc := scheduler.NewScheduler()
	sc.Add(ctx, action, time.Second*5)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}
