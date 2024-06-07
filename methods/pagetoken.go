package methods

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type Token struct {
	Data []struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

type Error struct {
	Data struct {
		Message string `json:"message"`
	} `json:"error"`
}

func GeneratePageToken() {
	var userId, userAccessToken string
	fmt.Println("Enter UserId")
	fmt.Scan(&userId)
	fmt.Println("Enter User Access Token")
	fmt.Scan(&userAccessToken)

	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/accounts?access_token=%s", userId, userAccessToken)

	recivedData, err := http.Get(url)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	defer recivedData.Body.Close()
	respBody, err := io.ReadAll(recivedData.Body)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	if recivedData.StatusCode == http.StatusOK {
		var token Token
		json.Unmarshal(respBody, &token)
		fmt.Println("SUCCESS")
		fmt.Println(token)
		data := token.Data[0].AccessToken
		fmt.Println(data)
		CheckEnv(data)
		err := godotenv.Load(".env")
		Check(&err)
	}
	if recivedData.StatusCode == 400 {
		var erro Error
		json.Unmarshal(respBody, &erro)
		fmt.Println("ERROR")
		fmt.Println(erro)
	}

}
