package main

import (
	"flag"
	"log"
	"os"

	"github.com/andrewjohnsonsmarty/calc-apps/handlers"
	"github.com/andrewjohnsonsmarty/calc-lib"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "Operation to calculate")
	flag.Parse()

	handler := handlers.NewHandler(os.Stdout, calculators[operation])
	err := handler.Handle(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]calc.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
