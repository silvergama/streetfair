package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/silvergama/unico/fair"
	"github.com/sirupsen/logrus"
)

func MakeFairHandler(r *mux.Router, service fair.UseCase) {

	// swagger:route GET /v1/fair v1 fairsGetV1Req
	// Get street fairs by neighborhood.
	// Responses:
	//   200: success
	//   500: internalServerError
	r.Handle("/v1/fair", getFair(service)).Methods(http.MethodGet).Name("getFairByNeighborhood")

	// swagger:route POST /v1/fair v1 fairPostV1Req
	// Add a new street fair.
	// Responses:
	//   200: success
	//   500: internalServerError
	r.Handle("/v1/fair", addFair(service)).Methods(http.MethodPost).Name("addFair")

	// swagger:route PUT /v1/fair/{id} v1 fairPutV1Req
	// Update street fair.
	// Responses:
	//   200: success
	//   500: internalServerError
	r.Handle("/v1/fair/{id:[0-9]+}", updateFair(service)).Methods(http.MethodPut).Name("updateFair")

	// swagger:route Delete /v1/fair/{id} v1 fairDeleteV1Req
	// Delete street fair by ID.
	// Responses:
	//   204: noContent
	//   500: internalServerError
	r.Handle("/v1/fair/{id:[0-9]+}", deleteFair(service)).Methods(http.MethodDelete).Name("deleteFair")
}

func getFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		neighborhood := r.URL.Query().Get("neighborhood")
		fairs, err := service.Get(neighborhood)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if fairs == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}

		if err := json.NewEncoder(w).Encode(fairs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
	})
}

func addFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body := fair.Fair{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logrus.Warn(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		if _, err := service.Save(&body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func updateFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			logrus.Warn(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		body := fair.Fair{ID: ID}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logrus.Warn(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		if _, err := service.Update(&body); err != nil {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("internal server error"))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func deleteFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")

		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			logrus.Warn(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err = service.Remove(ID); err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error deliting fair"))
			return
		}
	})
}
