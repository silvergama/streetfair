package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/silvergama/streetfair/entity"
	"github.com/silvergama/streetfair/pkg/response"
	"github.com/silvergama/streetfair/service/fair"
)

const v1Fair = "v1/fair"

func MakeFairHandler(r *mux.Router, service fair.UseCase) {

	// swagger:route GET /v1/fair v1 fairsGetV1Req
	// Get street fairs by neighborhood.
	// Responses:
	//   200: successGet
	//   404: notFound
	//   500: internalServerError
	r.Handle(v1Fair, getFair(service)).Methods(http.MethodGet).Name("getFairByNeighborhood")

	// swagger:route POST /v1/fair v1 fairPostV1Req
	// Add a new street fair.
	// Responses:
	//   200: success
	//   500: internalServerError
	r.Handle(v1Fair, addFair(service)).Methods(http.MethodPost).Name("addFair")

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
	r.Handle("/v1/fair/{id:[0-9]+}", deleteFair(service)).Methods(http.MethodDelete).Name("deleteFairByID")
}

func getFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		neighborhood := r.URL.Query().Get("neighborhood")
		fairs, err := service.GetFair(neighborhood)
		if err != nil || len(fairs) == 0 {
			response.WriteNotFound(w, "error finding street fair by neighborhood")
			return
		}

		response.Write(w, response.Fair{
			Total: len(fairs),
			Fairs: fairs,
		}, http.StatusOK)
	})
}

func addFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body := entity.Fair{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.WriteUnprocessableEntity(w, err.Error())
			return
		}

		ID, err := service.CreateFair(&body)
		if err != nil {
			response.WriteServerError(w, "error inserting street fair")
			return
		}
		response.Write(w, response.Success{
			ID: ID,
		}, http.StatusCreated)
	})
}

func updateFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body := entity.Fair{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.WriteServerError(w, err.Error())
			return
		}

		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.WriteUnprocessableEntity(w, err.Error())
			return
		}
		body.ID = ID

		if _, err := service.UpdateFair(&body); err != nil {
			response.WriteServerError(w, "error updating street fair")
			return
		}

		response.Write(w, response.Success{
			ID: ID,
		}, http.StatusOK)
	})
}

func deleteFair(service fair.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.WriteUnprocessableEntity(w, err.Error())
			return
		}

		if err = service.DeleteFair(ID); err != nil {
			response.WriteServerError(w, "error deleting street fair")
			return
		}

		response.Write(w, nil, http.StatusNoContent)
	})
}
