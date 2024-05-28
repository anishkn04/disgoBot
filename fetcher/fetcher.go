package fetcher

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Event struct{
	Banner string `json:"banner"`
	Title string `json:"title"`
	Start_date string `json:"start_date"`
	End_date string `json:"end_date"`
	Location string `json:"location"`
	Description string `json:"description"`
	Link string `json:"link"`
}

type ResponseBody struct{
	Events []Event`json:"events"`
}

func Fetch() ResponseBody{
	resp, err := http.Get("https://raw.githubusercontent.com/NepalTekComm/nepal-tek-commuity-website/main/resources/json/events.json");
	if err != nil {
		log.Fatal(err)
	}
	respBody, err := io.ReadAll(resp.Body);
	if err != nil{
		log.Fatal("Couldn't Read Data")
	}
	var jsonData ResponseBody;
	json.Unmarshal(respBody, &jsonData)

	defer resp.Body.Close();
	return jsonData;
}