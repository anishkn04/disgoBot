package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	methods "disgoBot/methods"
)

func main() {
	godotenv.Load()
	TOKEN := os.Getenv("BOTTOKEN")

	discord, err := discordgo.New("Bot " + TOKEN)
	methods.Check(err);

	discord.Open()
	fmt.Println("Bot running....")

	// Create a new ticker that triggers every "WAITTIME" minutes
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"));
	methods.Check(err);

	ticker := time.NewTicker(time.Duration(waittime) * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				// Fetch and send embeds every "WAITTIME" minutes
				sendEmbeds(methods.Fetch(), discord, methods.ReadStoredData())
			}
		}
	}()

	defer discord.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func sendEmbeds(fetchedData methods.ResponseBody, discord *discordgo.Session, sentMessages *[]string) {
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
		fmt.Println(err);
		if(err==nil){
			existingMessages = append(existingMessages, embedMessage.Title)
		}
	}
	*sentMessages = existingMessages
	methods.WriteStoredData(*sentMessages)
}
