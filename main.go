package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	fetcher "disgoBot/fetcher"
)

func main() {
	godotenv.Load()
	TOKEN := os.Getenv("BOTTOKEN")

	discord, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatal("Couldn't start!")
	}

	discord.Open()
	fmt.Println("Bot running....")

	var sentMessages []string //Slices of Events-Content

	// Create a new ticker that triggers every hour
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"));
	if err != nil {
		log.Fatal("Wrong Wait Time Value")
	}
	ticker := time.NewTicker(time.Duration(waittime) * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				// Fetch and send embeds every hour
				sendEmbeds(fetcher.Fetch(), discord, &sentMessages)
			}
		}
	}()

	defer discord.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func sendEmbeds(fetchedData fetcher.ResponseBody, discord *discordgo.Session, sentMessages *[]string) {
	EVENTSCHANNELID := os.Getenv("CHANNELID")
	existingMessages := []string{}
	if sentMessages != nil {
		existingMessages = *sentMessages
	}
	fmt.Println("Existing Messages: ", existingMessages)
	for _, event := range fetchedData.Events {
		eventBannerURL := "https://raw.githubusercontent.com/NepalTekComm/nepal-tek-commuity-website/main/" + event.Banner
		embedMessage := discordgo.MessageEmbed{
			URL:         event.Link,
			Image:       &discordgo.MessageEmbedImage{URL: eventBannerURL},
			Title:       event.Title,
			Description: event.Description,
			Timestamp:   event.Start_date,
		}
		if slices.Contains(existingMessages, embedMessage.Title) {
			fmt.Println("Announcement with title:", embedMessage.Title, "already exists. If this was a mistake, edit and try again!");
			continue
		}
		_, err := discord.ChannelMessageSendEmbed(EVENTSCHANNELID, &embedMessage)
		if err != nil {
			fmt.Println(err)
		}
		existingMessages = append(existingMessages, embedMessage.Title)
	}
	*sentMessages = existingMessages
}
