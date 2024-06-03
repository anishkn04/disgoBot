package methods

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Titles struct {
	DiscordTitles  []string `json:"discord"`
	FacebookTitles []string `json:"facebook"`
}

func ReadJson() ([]string, []string) {
	// CreateJsonIfNotExist()
	var title Titles
	file, err := os.Open("events.json")
	if err != nil {
		if err.Error() == "open events.json: The system cannot find the file specified." {
			CreateJsonIfNotExist()
			file, err = os.Open("events.json")
			Check(&err)
		} else {
			Check(&err)
		}
	}
	result, err := io.ReadAll(file)
	Check(&err)
	json.Unmarshal(result, &title)
	tDiscord := title.DiscordTitles[:]
	tFacebook := title.FacebookTitles[:]
	return tDiscord, tFacebook
}

func CreateJsonIfNotExist() {
	if fileExists("events.json") {
		fmt.Println("Exists you idiot!")
		return
	}

	data := map[string][]string{
		"discord":  {},
		"facebook": {},
	}

	file, err := os.Create("events.json")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
