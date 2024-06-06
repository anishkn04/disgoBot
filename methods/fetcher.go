package methods

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

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
	resp, err := http.Get(os.Getenv("SITE"));
	Check(&err)

	respBody, err := io.ReadAll(resp.Body)
	Check(&err)

	var jsonData ResponseBody
	json.Unmarshal(respBody, &jsonData)

	defer resp.Body.Close()
	return &jsonData
}
