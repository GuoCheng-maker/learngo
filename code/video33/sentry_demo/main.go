package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "http://51516ca5fc72402bb7eba4ccae6f0150@10.194.18.30:9000/10",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("test error")
}
