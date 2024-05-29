package methods

import (
	"encoding/csv"
	"fmt"
	"os"
	// "path/filepath"
)

func ReadStoredData() *[]string {
	file, err := os.OpenFile("storedArray.csv", os.O_RDONLY, 0444)
	if err != nil {
		fileCreatedStat := createIfNotExists(err);
		if !fileCreatedStat{
			panic("Couldn't read or create file!")
		}else{
			file, err = os.OpenFile("storedArray.csv", os.O_RDONLY, 0444);
			Check(err);
		}
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	Check(err)

	var dataFromFile []string
	for datumIndex, datumText := range data {
		fmt.Println(datumIndex, datumText)
		dataFromFile = append(dataFromFile, datumText...)
	}
	file.Close()
	return &dataFromFile
}

func createIfNotExists(err error) bool {
	if err.Error() != "open storedArray.csv: The system cannot find the file specified." {
		Check(err)
		return false
	} else {
		fileAgain, err := os.Create("storedArray.csv")
		fileAgain.Close()
		if err == nil{
			return true
		}else{
			return false
		}
	}
}