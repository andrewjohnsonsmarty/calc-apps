package handlers

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/andrewjohnsonsmarty/calc-lib"
)

func TestCSVHandler(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "[TEST] ", 0)
	rawInput := `1,+,2
3,-,2
1
1,2,3,4
NaN,+,3
3,+,NaN
4,+,5
`
	input := strings.NewReader(rawInput)
	var output bytes.Buffer
	handler := NewCSVHandler(logger, input, &output, map[string]calc.Calculator{"+": &calc.Addition{}})
	err := handler.Handle()
	assertErr(t, err, nil)
	expected := `1,+,2,3
4,+,5,9
`
	if output.String() != expected {
		t.Errorf("expected: [%s], got: [%s]", expected, output.String())
	}
	t.Log(logBuffer.String())
}
