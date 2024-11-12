package main

import (
	"log"
	"os"

	"github.com/andrewjohnsonsmarty/calc-apps/handlers"
	"github.com/andrewjohnsonsmarty/calc-lib"
)

func main() {
	logger := log.New(os.Stderr, ">>> ", 0)
	handler := handlers.NewCSVHandler(logger, os.Stdin, os.Stdout, calculators)
	err := handler.Handle()
	if err != nil {
		logger.Fatal(err)
	}
}

var calculators = map[string]calc.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
