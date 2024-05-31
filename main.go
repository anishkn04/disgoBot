package main

import (
	"os"
	"os/signal"
	"strconv"
	"time"

	methods "disgoBot/methods"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	// Create a new ticker that triggers every "WAITTIME" seconds
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"))
	methods.Check(err)

	ticker := time.NewTicker(time.Duration(waittime) * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				//Fetch and send embeds every "WAITTIME" minutes
				methods.SendFacebook(methods.Fetch(), methods.ReadStoredData())
				methods.SendEmbeds(methods.Fetch(), methods.ReadStoredData())
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
