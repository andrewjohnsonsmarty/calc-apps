package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestHTTPHandler_Add(t *testing.T) {
	assertRoute(t, "POST", "/add?a=1&b=3", http.StatusMethodNotAllowed, "Method Not Allowed\n")
	assertRoute(t, "GET", "/add?a=1&b=", http.StatusUnprocessableEntity, "The b parameter must be an integer\n")
	assertRoute(t, "GET", "/add?a=1&b=3", http.StatusOK, "4")
	assertRoute(t, "GET", "/add?a=1&b=3", http.StatusOK, "4")
	assertRoute(t, "GET", "/sub?a=1&b=3", http.StatusOK, "-2")
	assertRoute(t, "GET", "/mul?a=1&b=3", http.StatusOK, "3")
	assertRoute(t, "GET", "/div?a=4&b=2", http.StatusOK, "2")

}

func assertRoute(t *testing.T, method, target string, statusCode int, responseBody string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		request := httptest.NewRequest(method, target, nil)

		recorder := httptest.NewRecorder()
		var logBuffer bytes.Buffer
		logger := log.New(&logBuffer, "test>", 0)
		router := NewHTTPRouter(logger)

		requestDump, err := httputil.DumpRequest(request, true)
		assertErr(t, err, nil)
		t.Log("request dump: \n", string(requestDump))

		router.ServeHTTP(recorder, request)

		responseDump, err := httputil.DumpResponse(recorder.Result(), true)
		assertErr(t, err, nil)
		t.Logf("response dump:\n%s", string(responseDump))

		assertEqual(t, recorder.Code, statusCode)
		assertEqual(t, recorder.Body.String(), responseBody)
	})
}
