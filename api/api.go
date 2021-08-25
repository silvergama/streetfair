package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/silvergama/unico/api/v1"
	"github.com/silvergama/unico/fair"
	"github.com/silvergama/unico/repository"
)

var (
	RootHandlerPath        = "/"
	HealthCheckHandlerPath = "/healthcheck"

	FairHandlerPath = "/v1/fair"

	methodNotAllowedErrMessage = "Invalid request method"
)

func Setup() error {
	r := mux.NewRouter()

	r.HandleFunc(RootHandlerPath, RootHandler)
	r.HandleFunc(HealthCheckHandlerPath, HealthCheckHandler)

	fairRepo := fair.NewService(repository.GetInstance())
	v1.MakeFairHandler(r, fairRepo)

	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
	return nil
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
