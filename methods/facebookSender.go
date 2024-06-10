package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func SendFacebook(fetchedData ResponseBody, titles *[]string) {

	err := godotenv.Load()
	Check(&err)
	pageId := os.Getenv("PAGE_ID")
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")

	for _, event := range fetchedData.Events {
		if checkIfExists(*titles, event.Title) {
			continue
		}

		message := fmt.Sprintf("%s\n\nDescription: %s\n\nDate: %s to %s\nLocation: %s", event.Title, event.Description, event.Start_date, event.End_date, event.Location)
		eventBanner := "https://raw.githubusercontent.com/NepalTekComm/nepal-tek-commuity-website/main/" + event.Banner;
		evURL, err := url.Parse(eventBanner);
		Check(&err);
		postData := map[string]interface{}{
			"url":          evURL.String(),
			"message":      message,
			"access_token": pageAccessToken,
		}

		jsonData, err := json.Marshal(postData)
		Check(&err)

		url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/photos", pageId)

		fmt.Print(url)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		Check(&err)

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		Check(&err)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Error: received non-200 status code %d\nResponse: %s\n", resp.StatusCode, body)
			continue
		}

		writeIntoJson(&event.Title, "facebook")
	}

}