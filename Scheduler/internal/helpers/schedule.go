package helpers

import (
	"context"
	"fmt"
	"github.com/zhashkevych/scheduler"
	"os"
	"os/signal"
	"time"
)

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("waiting...\n")
}

func Schedule() {
	ctx := context.Background()
	sc := scheduler.NewScheduler()
	sc.Add(ctx, parseSubscriptionData, time.Second*2)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}
