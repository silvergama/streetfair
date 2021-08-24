package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/silvergama/unico/api/v1"
)

var (
	RootHandlerPath        = "/"
	HealthCheckHandlerPath = "/healthcheck"

	FairHandlerPath = "/fair"

	methodNotAllowedErrMessage = "Invalid request method"
)

func Setup() {
	r := mux.NewRouter()

	r.HandleFunc(RootHandlerPath, RootHandler)
	r.HandleFunc(HealthCheckHandlerPath, HealthCheckHandler)
	r.HandleFunc(FairHandlerPath, v1.FairHandler)

	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedErrMessage, http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "street fair")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedErrMessage, http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
