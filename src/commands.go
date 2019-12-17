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
	Amount float64 `json:"amount"`
	Note   string  `json:"note"`
	Date   string  `json:"date"`
}

type ExpensesFile struct {
	Payments []Transaction `json:"payments"`
	Bills    []Transaction `json:"bills"`
	Expenses []Transaction `json:"expenses"`
	Savings  []Transaction `json:"savings"`
}

const (
	PAYMENT = "payment"
	EXPEND  = "expend"
	BILL    = "bill"
	SAVED   = "saved"
	TOTAL   = "total"
)

var (
	payment    = flag.Float64("payment", 0.0, "Add payment")
	expend     = flag.Float64("expend", 0.0, "Add a expended value")
	bill       = flag.Float64("bill", 0.0, "Add a bill payment")
	saved      = flag.Float64("saved", 0.0, "Add an amount of saved money")
	totalMoney = flag.String("total", "", "Get the total amount of money, minus expenses, bills and savings")
)

func totalBills(bills []Transaction) float64 {
	total := 0.0
	for _, b := range bills {
		total += b.Amount
	}
	return total
}

func totalExpenses(expenses []Transaction) float64 {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	return total
}

func totalSavings(savings []Transaction) float64 {
	total := 0.0
	for _, s := range savings {
		total += s.Amount
	}
	return total
}

func totalPayments(payments []Transaction) float64 {
	total := 0.0
	for _, p := range payments {
		total += p.Amount
	}
	return total
}

func getTotalMoney(fileLocation string) (float64, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return -1, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	var expensesFile ExpensesFile
	json.Unmarshal(byteValue, &expensesFile)

	sum := totalBills(expensesFile.Bills) + totalExpenses(expensesFile.Expenses) + totalSavings(expensesFile.Savings)

	return totalPayments(expensesFile.Payments) - sum, nil
}

func addTransaction(money float64, note string, fileLocation string, transaction string) {
	t := time.Now()
	newTransaction := Transaction{money, note, t.Format("2006/01/02")}

	file, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)
	if err != nil {
		panic("There was an error trying to read the file.")
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	file.Seek(0, 0) // returns to the beginning o f the file

	var expensesFile ExpensesFile
	json.Unmarshal(byteValue, &expensesFile)

	switch transaction {
	case PAYMENT:
		expensesFile.Payments = append(expensesFile.Payments, newTransaction)
	case EXPEND:
		expensesFile.Expenses = append(expensesFile.Expenses, newTransaction)
	case BILL:
		expensesFile.Bills = append(expensesFile.Bills, newTransaction)
	case SAVED:
		expensesFile.Savings = append(expensesFile.Savings, newTransaction)
	}

	bytes, err := json.Marshal(&expensesFile)
	if err != nil {
		fmt.Printf("error ocurred: %v \n", err)
	}

	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

// InitFlags Initializes de flags
func InitFlags(fileLocation string) {
	flag.Parse()
	if len(os.Args) < 2 {
		panic("You need to pass at least one flag, see package -h for more information.")
	}

	if *totalMoney == "" {
		value, err := getTotalMoney(fileLocation)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Your total amount of money after discounting the saving, expenses and bills is: %.2f", value)

		os.Exit(1)
	}

	notes := flag.Arg(0)
	prefix := strings.TrimPrefix(os.Args[1], "-")

	switch prefix {
	case PAYMENT:
		addTransaction(*payment, notes, fileLocation, PAYMENT)
	case EXPEND:
		addTransaction(*expend, notes, fileLocation, EXPEND)
	case BILL:
		addTransaction(*bill, notes, fileLocation, BILL)
	case SAVED:
		addTransaction(*saved, notes, fileLocation, SAVED)
	default:
		fmt.Printf("There is not a command with the name %s", prefix)
	}

}
