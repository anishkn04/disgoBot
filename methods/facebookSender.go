package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"

	"github.com/joho/godotenv"
)

func SendFacebook(fetchedData ResponseBody, sentMessages *[]string) {
	err := godotenv.Load()
	Check(err)
	pageId := os.Getenv("page_id")
	pageAccessToken := os.Getenv("page_access_token")
	fmt.Println("ID: ", pageId)
	fmt.Println("Token: ", pageAccessToken)
	existingMessages := []string{}
	if sentMessages != nil {
		existingMessages = *sentMessages
	}
	fmt.Println("Existing Messages: ", existingMessages)
	for _, event := range fetchedData.Events {
		message := fmt.Sprintf("Event Title: %s\nStart Date: %s\nEnd Date: %s\nLocation: %s\nDescription: %s\n", event.Title, event.Start_date, event.End_date, event.Location, event.Description)
		postData := map[string]interface{}{
			"url":          "https://raw.githubusercontent.com/NepalTekComm/nepal-tek-commuity-website/main/" + event.Banner,
			"message":      message,
			"access_token": pageAccessToken,
		}
		if slices.Contains(existingMessages, event.Title) {
			fmt.Println("Already Announced")
			continue
		}

		jsonData, err := json.Marshal(postData)
		Check(err)

		url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/photos", pageId)

		fmt.Print(url)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		Check(err)

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		Check(err)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Error: received non-200 status code %d\nResponse: %s\n", resp.StatusCode, body)

		}
		// if err == nil {
		// 	existingMessages = append(existingMessages, event.Title)
		// }
		// *sentMessages = existingMessages
		// WriteStoredData(*sentMessages)

	}
}