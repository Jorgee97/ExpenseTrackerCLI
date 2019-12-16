package src

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Transaction struct {
	Amount   float64   `json:"amount"`
	Note     string    `json:"note"`
	Datetime time.Time `json:"dateTime"`
}

type ExpensesFile struct {
	Payments []Transaction `json:"payments"`
	Bills    []Transaction `json:"bills"`
	Expenses []Transaction `json:"expenses"`
	Savings  []Transaction `json:"savings"`
}

var (
	payment = flag.Float64("payment", 0.0, "Add payment")
	expend  = flag.Float64("expend", 0.0, "Add a expended value")
	bill    = flag.Float64("bill", 0.0, "Add a bill payment")
	saved   = flag.Float64("saved", 0.0, "Add an amount of saved money")
)

func addPayment(pay float64, note string, fileLocation string) {
	newPayment := Transaction{pay, note, time.Now()}

	file, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)
	if err != nil {
		panic("There was an error trying to read the file.")
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var expensesFile ExpensesFile

	json.Unmarshal(byteValue, &expensesFile)

	expensesFile.Payments = append(expensesFile.Payments, newPayment)

	bytes, err := json.Marshal(&expensesFile)
	if err != nil {
		fmt.Printf("error ocurred: %v \n", err)
	}

	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func addExpend() {

}

func addBill() {

}

func addSaving() {

}

// InitFlags Initializes de flags
func InitFlags(fileLocation string) {
	flag.Parse()
	if len(os.Args) < 2 {
		panic("You need to pass at least one flag, see package -h for more information.")
	}

	notes := flag.Arg(0)

	switch strings.TrimPrefix(os.Args[1], "-") {
	case "payment":
		addPayment(*payment, notes, fileLocation)
	}
}
