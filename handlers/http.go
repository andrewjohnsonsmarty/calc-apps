package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/andrewjohnsonsmarty/calc-lib"
)

func NewHTTPRouter(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /add", NewHTTPHandler(logger, &calc.Addition{}))
	mux.Handle("GET /sub", NewHTTPHandler(logger, &calc.Subtraction{}))
	mux.Handle("GET /mul", NewHTTPHandler(logger, &calc.Multiplication{}))
	mux.Handle("GET /div", NewHTTPHandler(logger, &calc.Division{}))

	return mux
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator calc.Calculator
}

func NewHTTPHandler(logger *log.Logger, calculator calc.Calculator) *HTTPHandler {
	return &HTTPHandler{logger, calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	rawA := query.Get("a")
	rawB := query.Get("b")
	a, err := strconv.Atoi(rawA)
	if err != nil {
		http.Error(response, "The a parameter must be an integer", http.StatusUnprocessableEntity)
		return
	}
	b, err := strconv.Atoi(rawB)
	if err != nil {
		http.Error(response, "The b parameter must be an integer", http.StatusUnprocessableEntity)
		return
	}
	c := this.calculator.Calculate(a, b)
	_, _ = fmt.Fprintf(response, "%d", c)
}
