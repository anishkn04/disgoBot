package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	methods "disgoBot/methods"

	"github.com/joho/godotenv"
)

type appsToUse int
const (
	FACEBOOK appsToUse = 1
	DISCORD  appsToUse = 2
	BOTH     appsToUse = -1
)

func main() {
	methods.CheckEnv()
	err := godotenv.Load(".env")
	methods.Check(&err)
	// Create a new ticker that triggers every "WAITTIME" seconds
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"))
	methods.Check(&err)

	apps, err := strconv.Atoi(os.Getenv("APPSTOUSE"));
	methods.Check(&err)

	ticker := time.NewTicker(time.Duration(waittime) * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			senderFunc(appsToUse(apps))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func senderFunc(atu appsToUse) {
	if atu != FACEBOOK && atu != DISCORD && atu != BOTH {
		panic("Wrong congifs, contact developer or remove .env and start again!")
	}
	fetchedData := methods.Fetch()
	publishedTitlesDiscord, publishedTitlesFacebook := methods.ReadJson()
	if atu == FACEBOOK || atu == BOTH {
		fmt.Println("FB: ", publishedTitlesFacebook)
		methods.SendFacebook(*fetchedData, &publishedTitlesFacebook)
	}
	if atu == DISCORD || atu == BOTH {
		fmt.Println("Discord:", publishedTitlesDiscord)
		methods.SendEmbeds(*fetchedData, &publishedTitlesDiscord)
	}
}
