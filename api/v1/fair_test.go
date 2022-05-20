package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/silvergama/streetfair/entity"
	"github.com/silvergama/streetfair/pkg/response"
	"github.com/stretchr/testify/assert"
)

var (
	saveFairService   func(f *entity.Fair) (int, error)
	updateFairService func(f *entity.Fair) (int, error)
	removeFairService func(id int) error
	getFairService    func(neighborhood string) ([]*entity.Fair, error)
)

type serviceMock struct{}

func (fs serviceMock) CreateFair(f *entity.Fair) (int, error) {
	return saveFairService(f)
}

func (fs serviceMock) UpdateFair(f *entity.Fair) (int, error) {
	return updateFairService(f)
}

func (fs serviceMock) GetFair(neighborhood string) ([]*entity.Fair, error) {
	return getFairService(neighborhood)
}

func (fs serviceMock) DeleteFair(id int) error {
	return removeFairService(id)
}

/**
* Test coverage AddFair
**/
func TestAddFair(t *testing.T) {
	saveFairService = func(f *entity.Fair) (int, error) {
		return f.ID, nil
	}
	fakeHandler := addFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair", fakeHandler)
	payload := `{
		"id": 10
	}`

	req, err := http.NewRequest("POST", "/v1/fair", bytes.NewBufferString(payload))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var success response.Success
	err = json.Unmarshal(rr.Body.Bytes(), &success)
	assert.Nil(t, err)
	assert.NotNil(t, success)
	assert.EqualValues(t, http.StatusCreated, rr.Code)
	assert.EqualValues(t, 10, success.ID)
}

func TestAddFairUnprocessableEntity(t *testing.T) {
	errMessage := errors.New("json: cannot unmarshal string into Go struct field Fair.id of type int")
	saveFairService = func(f *entity.Fair) (int, error) {
		return f.ID, errMessage
	}
	fakeHandler := addFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair", fakeHandler)

	payload := `{
		"id": "10"
	}`

	req, err := http.NewRequest("PUT", "/v1/fair", bytes.NewBufferString(payload))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var errResponse response.Error
	err = json.Unmarshal(rr.Body.Bytes(), &errResponse)
	assert.Nil(t, err)

	expected := response.Error{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errMessage.Error(),
	}
	assert.NotNil(t, errResponse)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expected, errResponse)
}

/**
* Test coverage UpdateFair
**/
func TestUpdateFairSuccess(t *testing.T) {
	updateFairService = func(f *entity.Fair) (int, error) {
		return int(f.ID), nil
	}
	fakeHandler := updateFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair/{id}", fakeHandler)

	payload := `{
		"id": 10
	}`

	req, err := http.NewRequest("PUT", "/v1/fair/10", bytes.NewBufferString(payload))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var success response.Success
	err = json.Unmarshal(rr.Body.Bytes(), &success)

	assert.Nil(t, err)
	assert.NotNil(t, success)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, response.Success{ID: 10}, success)
}

func TestUpdateFairServerError(t *testing.T) {
	errMessage := errors.New("json: cannot unmarshal string into Go struct field Fair.id of type int")
	updateFairService = func(f *entity.Fair) (int, error) {
		return 0, errMessage
	}
	fakeHandler := updateFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair/{id}", fakeHandler)

	payload := `{
		"id": "10",
		"Lat": "798798798"
	}`

	req, err := http.NewRequest("PUT", "/v1/fair/10", bytes.NewBufferString(payload))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var errResponse response.Error
	err = json.Unmarshal(rr.Body.Bytes(), &errResponse)
	assert.Nil(t, err)

	expected := response.Error{
		StatusCode: http.StatusInternalServerError,
		Message:    errMessage.Error(),
	}
	assert.NotNil(t, errResponse)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, expected, errResponse)
}

func TestUpdateFairUnprocessableEntity(t *testing.T) {
	errMessage := errors.New("strconv.Atoi: parsing \"penha\": invalid syntax")
	updateFairService = func(f *entity.Fair) (int, error) {
		return 0, errMessage
	}
	fakeHandler := updateFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair/{id}", fakeHandler)

	req, err := http.NewRequest("PUT", "/v1/fair/penha", bytes.NewBufferString(`{"ID": 10}`))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var errResponse response.Error
	err = json.Unmarshal(rr.Body.Bytes(), &errResponse)
	assert.Nil(t, err)

	expected := response.Error{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errMessage.Error(),
	}
	assert.NotNil(t, errResponse)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expected, errResponse)
}

/**
* Test coverage GetFair
**/
func TestGetFair(t *testing.T) {
	fairsResponse := []*entity.Fair{{ID: 1}}
	getFairService = func(neighborhood string) ([]*entity.Fair, error) {
		return fairsResponse, nil
	}
	fakeHandler := getFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair", fakeHandler).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", "/v1/fair?neighborhood=Penha", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var fairs response.Fair
	err = json.NewDecoder(rr.Body).Decode(&fairs)
	expected := response.Fair{Total: 1, Fairs: fairsResponse}
	assert.Nil(t, err)
	assert.NotNil(t, fairs)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, fairs)
}

func TestGetFairNotFound(t *testing.T) {
	errMessage := errors.New("error finding street fair by neighborhood")
	getFairService = func(neighborhood string) ([]*entity.Fair, error) {
		return nil, errMessage
	}
	fakeHandler := getFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair", fakeHandler).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", "/v1/fair?neighborhood=Penha", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var errResponse response.Error
	err = json.NewDecoder(rr.Body).Decode(&errResponse)

	expected := response.Error{
		StatusCode: http.StatusNotFound,
		Message:    errMessage.Error(),
	}
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expected, errResponse)
}

/**
* Test coverage DeleteFair
**/
func TestDeleteFair(t *testing.T) {
	removeFairService = func(id int) error {
		return nil
	}
	fakeHandler := deleteFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair/{id}", fakeHandler)

	req, err := http.NewRequest("DELETE", "/v1/fair/20", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteFairUnprocessableEntity(t *testing.T) {
	errMessage := errors.New("strconv.Atoi: parsing \"deleteID\": invalid syntax")
	removeFairService = func(id int) error {
		return errMessage
	}
	fakeHandler := deleteFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair/{id}", fakeHandler)

	req, err := http.NewRequest("DELETE", "/v1/fair/deleteID", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var errResponse response.Error
	err = json.NewDecoder(rr.Body).Decode(&errResponse)

	expected := response.Error{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errMessage.Error(),
	}
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expected, errResponse)
}
