package methods

import (
	"encoding/json"
	"os"
)

func writeIntoJson(event string, social string) {
	jsonData, err := os.ReadFile("events.json")
	Check(err)
	var title Titles
	err = json.Unmarshal(jsonData, &title)
	Check(err)
	if social == "facebook" {
		title.FacebookTitles = append(title.FacebookTitles, event)
	}
	if social == "discord" {
		title.DiscordTitles = append(title.DiscordTitles, event)
	}
	updatedData, err := json.MarshalIndent(title, "", "    ")
	Check(err)
	err = os.WriteFile("events.json", updatedData, 0644)
	Check(err)

}
