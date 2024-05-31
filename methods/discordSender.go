package methods

import (
	"fmt"
	"os"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func SendEmbeds(fetchedData ResponseBody, sentMessages *[]string) {
	godotenv.Load()
	TOKEN := os.Getenv("BOTTOKEN")
	discord, err := discordgo.New("Bot " + TOKEN)
	Check(err)

	discord.Open()
	fmt.Println("Bot running....")
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
			fmt.Println("Announcement with title:", embedMessage.Title, "already exists. If this was a mistake, edit and try again!")
			continue
		}
		_, err := discord.ChannelMessageSendEmbed(EVENTSCHANNELID, &embedMessage)
		Check(err)
		if err == nil {
			existingMessages = append(existingMessages, embedMessage.Title)
		}
	}
	*sentMessages = existingMessages
	WriteStoredData(*sentMessages)
	defer discord.Close()
}
