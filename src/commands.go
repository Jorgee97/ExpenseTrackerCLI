package src

import (
	"flag"
	"os"
	"strings"
	"time"
)

type Payment struct {
	amount   float64
	note     string
	datetime time.Time
}

var (
	payment = flag.Float64("payment", 0.0, "Add payment")
	expend  = flag.Float64("expend", 0.0, "Add a expended value")
	bill    = flag.Float64("bill", 0.0, "Add a bill payment")
	saved   = flag.Float64("saved", 0.0, "Add an amount of saved money")
)

func addPayment(pay float64, note string) {
	//newPayment := Payment{pay, note, time.Now()}

}

func addExpend() {

}

func addBill() {

}

func addSaving() {

}

// InitFlags Initializes de flags
func InitFlags() {
	flag.Parse()
	if len(os.Args) < 2 {
		panic("You need to pass at least one flag, see package -h for more information.")
	}

	notes := flag.Arg(0)

	switch strings.TrimPrefix(os.Args[1], "-") {
	case "payment":
		addPayment(*payment, notes)
	}
}
