package methods

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

func generatePageToken() (string, error) {
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
		data := token.Data[0].AccessToken
		fmt.Println(data)
		return data, nil
	}

	var erro Error
	json.Unmarshal(respBody, &erro)
	fmt.Println("ERROR")
	return erro.Data.Message, errors.New("cant generate token")

}
