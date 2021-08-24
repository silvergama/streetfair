package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type handlerTest struct {
	about              string
	handler            http.HandlerFunc
	handlerPath        string
	method             string
	expectedMsg        string
	expectedStatusCode int
}

var defaultHandlerTests = []handlerTest{
	{
		about:              "Should return welcome message - GET (RootHandler)",
		handler:            RootHandler,
		handlerPath:        RootHandlerPath,
		method:             http.MethodGet,
		expectedMsg:        "street fair",
		expectedStatusCode: http.StatusOK,
	},
	{
		about:              "Should return invalid request - PUT (RootHandler)",
		handler:            RootHandler,
		handlerPath:        RootHandlerPath,
		method:             http.MethodPut,
		expectedMsg:        "Invalid request method",
		expectedStatusCode: http.StatusMethodNotAllowed,
	},
	{
		about:              "Should return OK message - GET (HealthCheckHandler)",
		handler:            HealthCheckHandler,
		handlerPath:        HealthCheckHandlerPath,
		method:             http.MethodGet,
		expectedMsg:        "OK",
		expectedStatusCode: http.StatusOK,
	},
	{
		about:              "Should return invalid request - PUT (HealthCheckHandler)",
		handler:            HealthCheckHandler,
		handlerPath:        HealthCheckHandlerPath,
		method:             http.MethodPut,
		expectedMsg:        "Invalid request method",
		expectedStatusCode: http.StatusMethodNotAllowed,
	},
}

func TestDefaultHandlers(t *testing.T) {
	testHandlers(t, defaultHandlerTests)
}

func testHandlers(t *testing.T, tests []handlerTest) {
	for _, test := range tests {
		t.Log(test.about)
		req, err := http.NewRequest(test.method, test.handlerPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(test.handlerPath, test.handler)
		router.ServeHTTP(recorder, req)
		if status := recorder.Code; status != test.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatusCode)
		}

		responseBody := strings.TrimSuffix(recorder.Body.String(), "\n")
		if test.expectedMsg != "" && responseBody != test.expectedMsg {
			t.Errorf("handler returned unexpected body: got '%v' want '%v'", recorder.Body.String(), test.expectedMsg)
		}
	}
}
