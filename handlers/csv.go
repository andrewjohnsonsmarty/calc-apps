package handlers

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/andrewjohnsonsmarty/calc-lib"
)

type CSVHandler struct {
	logger      *log.Logger
	input       *csv.Reader
	output      *csv.Writer
	calculators map[string]calc.Calculator
}

func NewCSVHandler(logger *log.Logger, input io.Reader, output io.Writer, calculators map[string]calc.Calculator) *CSVHandler {
	return &CSVHandler{
		logger:      logger,
		input:       csv.NewReader(input),
		output:      csv.NewWriter(output),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() error {
	this.input.FieldsPerRecord = 3
	defer this.output.Flush()
	for {
		record, err := this.input.Read()
		if err == io.EOF {
			break
		}
		if errors.Is(err, csv.ErrFieldCount) {
			this.logger.Println(err)
			continue
		}
		// TODO: if err != nil deal with this
		a, err := strconv.Atoi(record[0])
		if err != nil {
			this.logger.Println("invalid arg:", record[0])
			continue
		}
		b, err := strconv.Atoi(record[2])
		if err != nil {
			this.logger.Println("invalid arg:", record[2])
			continue
		}
		op := record[1]
		calculator, ok := this.calculators[op]
		if !ok {
			this.logger.Println("unsupported operator:", op)
			continue
		}
		c := calculator.Calculate(a, b)
		this.output.Write(append(record, strconv.Itoa(c)))
		// TODO: if err != nil...
	}
	return this.output.Error()
}
