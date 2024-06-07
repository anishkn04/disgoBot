package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	methods "disgoBot/methods"

	"github.com/joho/godotenv"
)

func main() {
	var check string
	fmt.Println("Do you want to generate page token or already have one?")
	fmt.Println("Press 'Y' to generate page token with user_access_token and userId")
	fmt.Println("Press 'N' if you already own page_access_token")
	fmt.Scan(&check)
	checkSwitch := strings.ToLower(check)
	switch checkSwitch {
	case "y":
		methods.GeneratePageToken()

	case "n":
		methods.CheckEnv("")
		err := godotenv.Load(".env")
		methods.Check(&err)
	default:
		fmt.Println("Incorrect Input")
		os.Exit(0)
	}
	// Create a new ticker that triggers every "WAITTIME" seconds
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"))
	methods.Check(&err)

	ticker := time.NewTicker(time.Duration(waittime) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		publishedTitlesDiscord, publishedTitlesFacebook := methods.ReadJson()
		fetchedData := methods.Fetch()
		fmt.Println("Discord:", publishedTitlesDiscord)
		fmt.Println("FB: ", publishedTitlesFacebook)
		methods.SendFacebook(*fetchedData, &publishedTitlesFacebook)
		methods.SendEmbeds(*fetchedData, &publishedTitlesDiscord)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
