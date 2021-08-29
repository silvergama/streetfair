package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/silvergama/streetfair/api/v1"
	"github.com/silvergama/streetfair/app"
	"github.com/silvergama/streetfair/repository"
	"github.com/silvergama/streetfair/service/fair"
	"github.com/sirupsen/logrus"
)

var (
	RootHandlerPath        = "/"
	HealthCheckHandlerPath = "/healthcheck"

	FairHandlerPath = "/v1/fair"

	methodNotAllowedErrMessage = "Invalid request method"
)

func Setup() error {
	port := app.Config.Get("apiPort")
	r := mux.NewRouter()

	r.HandleFunc(RootHandlerPath, RootHandler)
	r.HandleFunc(HealthCheckHandlerPath, HealthCheckHandler)

	repo := repository.NewFairPostgreSQL(repository.GetInstance())
	v1.MakeFairHandler(r, fair.NewService(repo))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	logrus.Infof("listening on %s", port)
	logrus.Fatal(srv.ListenAndServe())
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
