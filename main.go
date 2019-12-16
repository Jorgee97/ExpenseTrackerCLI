package main

import (
	"expenses_tracker/src"
	"os"
	"path"
)

func createJSONFile() string {
	// create enviromental variable for the location of the file
	filePath := ""
	fileName := "expenses.json"
	if os.Getenv("EXPENSES_TRACKER_FILE") == "" {
		os.Setenv("EXPENSES_TRACKER_FILE", "/home/coreman/Documents/")
	}

	filePath = os.Getenv("EXPENSES_TRACKER_FILE")
	if _, err := os.Stat(filePath + "expenses.json"); os.IsNotExist(err) {
		file, err := os.OpenFile(filePath+"expenses.json", os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			panic("The application could create the file at the given location.")
		}
		defer file.Close()

		file.WriteString(`{"payments":[], "bills": [], "expenses": [], "savings": []}`)
	}
	return path.Join(filePath, fileName)
}

func main() {
	fileLocation := createJSONFile()

	src.InitFlags(fileLocation)
}
