package methods

import (
	"fmt"
	"os"
)

func CheckEnv() {
	_, err := os.Stat(".env")
	if err == nil {
		return
	}
	if err.Error() == "CreateFile .env: The system cannot find the file specified." {
		var finalString string
		var token string
		var checker error
		fmt.Println("PLEASE REFER TO THE GUIDE IF YOU DO NOT UNDERSTAND SOMETHING: https://github.com/anishkn04/goAppCLI")
		var appsToUse int = 0
		var generator string
		for {
			if appsToUse != -1 && appsToUse != 1 && appsToUse != 2 {
				fmt.Println("Enter 1 if you want to automate for facebook, 2 for discord and -1 for both: ")
				fmt.Scanln(&appsToUse)
			} else {
				break
			}
		}
		if appsToUse == 1 || appsToUse == -1 {
			fmt.Println("Do you want to generate Facebook Page Token")
			fmt.Println("Press 1 if you want to generate page access token")
			fmt.Println("Press any key if already have page access token")
			fmt.Scanln(&generator)
			if generator == "1" {
				token, checker = generatePageToken()
				fmt.Println(token)
			}

			var pageID string
			fmt.Println("Enter Facebook Page ID: ")
			fmt.Scanln(&pageID)
			var pageAccessToken string
			if checker == nil {
				pageAccessToken = token
			} else {
				fmt.Println("Enter Facebook Page Access Token: ")
				fmt.Scanln(&pageAccessToken)
			}
			finalString += fmt.Sprintf("PAGE_ID='%s'\nPAGE_ACCESS_TOKEN='%s'\n", pageID, pageAccessToken)
		}
		if appsToUse == 2 || appsToUse == -1 {
			var botToken string
			fmt.Println("Enter Discord Bot Token: ")
			fmt.Scanln(&botToken)
			var channelId string
			fmt.Println("Enter the Channel Id where you want to post changes: ")
			fmt.Scanln(&channelId)
			finalString += fmt.Sprintf("BOTTOKEN='%s'\nCHANNELID='%s'\n", botToken, channelId)
		}

		var waitTime string
		fmt.Println("Enter the interval (in seconds) in which you want to check for updates: ")
		fmt.Scanln(&waitTime)

		var eventSite string
		fmt.Println("Enter the site or file name (relative) with the events data: ")
		fmt.Scanln(&eventSite)

		finalString += fmt.Sprintf("WAITTIME='%s'\nSITE='%s'\nAPPSTOUSE='%d'", waitTime, eventSite, appsToUse)

		err = os.WriteFile(".env", []byte(finalString), 0777)
		Check(&err)
	}
	Check(&err)
}
