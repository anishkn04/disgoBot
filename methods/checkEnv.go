package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckEnv() {
	_, err := os.Stat(".env")
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		var finalString string
		fmt.Println("PLEASE REFER TO THE GUIDE IF YOU DO NOT UNDERSTAND SOMETHING: https://github.com/anishkn04/goAppCLI")

		fmt.Println("Enter 1 if you want to automate for Facebook, 2 for Discord, and -1 for both: ")
		appsToUse := getIntInput()

		if appsToUse == 1 || appsToUse == -1 {
			fmt.Println("Enter Facebook Page ID: ")
			pageID := getInput()

			fmt.Println("Enter Facebook Page Access Token: ")
			pageAccessToken := getInput()

			finalString += fmt.Sprintf("PAGE_ID='%s'\nPAGE_ACCESS_TOKEN='%s'\n", pageID, pageAccessToken)
		}
		if appsToUse == 2 || appsToUse == -1 {

			fmt.Println("Enter Discord Bot Token: ")
			botToken := getInput()
			fmt.Println("Enter the Channel Id where you want to post changes: ")
			channelId := getInput()

			finalString += fmt.Sprintf("BOTTOKEN='%s'\nCHANNELID='%s'\n", botToken, channelId)
		}

		fmt.Println("Enter the interval (in seconds) in which you want to check for updates: ")
		waitTime := getIntInput()

		fmt.Println("Enter the site or file name (relative) with the events data: ")
		eventSite := getInput()

		finalString += fmt.Sprintf("WAITTIME='%d'\nSITE='%s'\nAPPSTOUSE='%d'", waitTime, eventSite, appsToUse)

		err = os.WriteFile(".env", []byte(finalString), 0777)
		HardCheck(&err)
	}
	HardCheck(&err)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	HardCheck(&err)
	return strings.TrimSpace(input)
}

func getIntInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	HardCheck(&err)
	input = strings.TrimSpace(input)
	intInput, err := strconv.Atoi(input)
	HardCheck(&err)
	return intInput
}
