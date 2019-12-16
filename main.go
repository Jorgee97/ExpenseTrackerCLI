package main

import "expenses_tracker/src"

import "os"

import "fmt"

func createJSONFile() {
	// create enviromental variable for the location of the file
	filePath := ""
	if os.Getenv("EXPENSES_TRACKER_FILE") == "" {
		os.Setenv("EXPENSES_TRACKER_FILE", "/home/coreman/Documents/")
	}

	filePath = os.Getenv("EXPENSES_TRACKER_FILE")
	fmt.Print(filePath)
	file, err := os.OpenFile(filePath+"expenses.json", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic("The application could create the file at the given location.")
	}
	defer file.Close()

	file.WriteString(`
		{
			payments:[],
			bills: [],
			expenses: [],
			savings: []
		}
	`)
}

func main() {
	createJSONFile()

	src.InitFlags()
}
