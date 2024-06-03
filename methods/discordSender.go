package methods

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func SendEmbeds(fetchedData ResponseBody, title *[]string) {
	godotenv.Load()
	TOKEN := os.Getenv("BOTTOKEN")
	discord, err := discordgo.New("Bot " + TOKEN)
	Check(&err)

	discord.Open()
	fmt.Println("Bot running....")
	EVENTSCHANNELID := os.Getenv("CHANNELID")

	for _, event := range fetchedData.Events {
		if checkIfExists(*title, event.Title) {
			continue
		}

		eventBannerURL := "https://raw.githubusercontent.com/NepalTekComm/nepal-tek-commuity-website/main/" + event.Banner
		embedMessage := discordgo.MessageEmbed{
			URL:         event.Link,
			Image:       &discordgo.MessageEmbedImage{URL: eventBannerURL},
			Title:       event.Title,
			Description: event.Description,
			Timestamp:   event.Start_date,
		}

		_, err := discord.ChannelMessageSendEmbed(EVENTSCHANNELID, &embedMessage)
		Check(&err)
		if err == nil {
			writeIntoJson(&event.Title, "discord")
		}

	}

	defer discord.Close()
}
