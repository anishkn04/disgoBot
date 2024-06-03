package methods

import (
	"encoding/json"
	"os"
)

func writeIntoJson(event *string, social string) {
	jsonData, err := os.ReadFile("events.json")
	Check(&err)
	var title Titles
	err = json.Unmarshal(jsonData, &title)
	Check(&err)
	switch social {
	case "facebook":
		title.FacebookTitles = append(title.FacebookTitles, *event)
	case "discord":
		title.DiscordTitles = append(title.DiscordTitles, *event)
	default:
	}
	updatedData, err := json.MarshalIndent(title, "", "    ")
	Check(&err)
	err = os.WriteFile("events.json", updatedData, 0644)
	Check(&err)

}
