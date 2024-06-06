package methods

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Event struct {
	Banner      string `json:"banner"`
	Title       string `json:"title"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type ResponseBody struct {
	Events []Event `json:"events"`
}

func Fetch() *ResponseBody {
	err := godotenv.Load(".env")
	siteLink := os.Getenv("SITE")
	var jsonData ResponseBody
	if !strings.Contains(siteLink, "http") {
		file, err := os.Open(siteLink)
		result, err := io.ReadAll(file)
		Check(&err)
		json.Unmarshal(result, &jsonData)
		return &jsonData
	}
	resp, err := http.Get(siteLink)
	Check(&err)

	respBody, err := io.ReadAll(resp.Body)
	Check(&err)

	json.Unmarshal(respBody, &jsonData)

	defer resp.Body.Close()
	return &jsonData
}
