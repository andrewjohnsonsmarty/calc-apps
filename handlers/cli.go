package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/andrewjohnsonsmarty/calc-lib"
)

type Handler struct {
	stdout     io.Writer
	calculator *calc.Addition
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{stdout, calculator}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, err)
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", errOutputFailure, err)
	}
	return nil
}

var (
	errWrongArgCount   = errors.New("usage: calculator <a> <b>")
	errInvalidArgument = errors.New("invalid argument")
	errOutputFailure   = errors.New("output failure")
)
