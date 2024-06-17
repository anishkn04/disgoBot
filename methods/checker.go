package methods

import (
	"fmt"
	"log"
	"os"
	"time"
)

func checkIfExists(arr []string, search string) bool {
	for _, item := range arr {
		if item == search {
			return true
		}
	}
	return false
}

func Check(e *error) {
	if *e == nil {
		return
	}

	fmt.Println(*e)
	if !fileExists("logs.txt") {
		os.Create("logs.txt")
	}
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
	print(err)

	_, err = fmt.Fprintln(file, time.Now().Format(time.DateTime)+" : "+(*e).Error()+"\n")
	print(err)

	defer file.Close()
}

func HardCheck(e *error) {
	if *e == nil {
		return
	}

	if !fileExists("logs.txt") {
		os.Create("logs.txt")
	}
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
	print(err)

	_, err = fmt.Fprintln(file, time.Now().Format(time.DateTime)+" : "+(*e).Error()+"\n")
	print(err)

	defer file.Close()
	log.Fatal(*e)
}

func print(e error) {
	if e != nil {
		fmt.Println("Logger error: ", e)
	}
}
