package response

import (
	"encoding/json"
	"net/http"

	"github.com/silvergama/streetfair/fair"
	"github.com/sirupsen/logrus"
)

type Error struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

type Success struct {
	ID int `json:"id,omitempty"`
}

type Fair struct {
	Total int          `json:"total,omitempty"`
	Fairs []*fair.Fair `json:"fairs,omitempty"`
}

func Write(w http.ResponseWriter, body interface{}, status int) {
	if body == nil {
		w.WriteHeader(status)
		return
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

func WriteServerError(w http.ResponseWriter, message string) {
	logrus.Errorln(message)
	Write(w, Error{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}, http.StatusInternalServerError)
}

func WriteNotFound(w http.ResponseWriter, message string) {
	logrus.Errorln(message)
	Write(w, Error{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}, http.StatusNotFound)
}

func WriteUnprocessableEntity(w http.ResponseWriter, message string) {
	logrus.Errorln(message)
	Write(w, Error{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}, http.StatusUnprocessableEntity)
}
