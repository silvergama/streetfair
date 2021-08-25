package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/silvergama/unico/logger"
)

func errorBody(message string, statusCode int) interface{} {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func notFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	responseJSON, _ := responseBodyToJSON(&Error{
		StatusCode: http.StatusNotFound,
		Message:    "not found",
	})
	w.Write(responseJSON)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseJSON, err := responseBodyToJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
}

func responseBodyToJSON(body interface{}) ([]byte, error) {
	responseJSON, err := json.Marshal(body)
	if err != nil {
		logger.WarningLogger.Println(fmt.Sprintf("error marshaling response json, err: %v", err))
		return nil, err
	}

	return responseJSON, nil
}
