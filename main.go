package main

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	
	methods "disgoBot/methods"
)

func main() {
	godotenv.Load()

	// Create a new ticker that triggers every "WAITTIME" minutes
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"));
	methods.Check(err);

	ticker := time.NewTicker(time.Duration(waittime) * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				// Fetch and send embeds every "WAITTIME" minutes
				methods.SendEmbeds(methods.Fetch(), methods.ReadStoredData())
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}