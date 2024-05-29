package methods

import (
	"encoding/csv"
	"os"
)

func WriteStoredData(recievedData []string) {
	file, err := os.OpenFile("storedArray.csv", os.O_WRONLY, 0222)
	Check(err);
	reader := csv.NewWriter(file);
	
	var writableData [][]string

	for _, data := range recievedData{
		writableData = append(writableData, []string{data})
	}

	writeError := reader.WriteAll(writableData);
	Check(writeError)
	file.Close();
}