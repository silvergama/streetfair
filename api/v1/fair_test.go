package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/silvergama/streetfair/common/response"
	"github.com/silvergama/streetfair/fair"
	"github.com/stretchr/testify/assert"
)

var (
	saveFairService   func(f *fair.Fair) (int, error)
	updateFairService func(f *fair.Fair) (int64, error)
	removeFairService func(id int) error
	getFairService    func(neighborhood string) ([]*fair.Fair, error)
)

type serviceMock struct{}

func (fs serviceMock) Save(f *fair.Fair) (int, error) {
	return saveFairService(f)
}

func (fs serviceMock) Update(f *fair.Fair) (int64, error) {
	return updateFairService(f)
}

func (fs serviceMock) Get(neighborhood string) ([]*fair.Fair, error) {
	return getFairService(neighborhood)
}

func (fs serviceMock) Remove(id int) error {
	return removeFairService(id)
}

func TestAddFair(t *testing.T) {
	saveFairService = func(f *fair.Fair) (int, error) {
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

func TestUpdateFairSuccess(t *testing.T) {
	updateFairService = func(f *fair.Fair) (int64, error) {
		return int64(f.ID), nil
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
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, 10, success.ID)
}

func TestUpdateFairServerError(t *testing.T) {
	errMessage := errors.New("error updating street fair")
	updateFairService = func(f *fair.Fair) (int64, error) {
		return 0, errMessage
	}
	fakeHandler := updateFair(serviceMock{})

	r := mux.NewRouter()
	r.Handle("/v1/fair/{id}", fakeHandler)

	req, err := http.NewRequest("PUT", "/v1/fair/10", bytes.NewBufferString(`{"ID": 10, "Error": "ServerError"}`))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var errResponse response.Error
	err = json.Unmarshal(rr.Body.Bytes(), &errResponse)
	assert.Nil(t, err)
	assert.NotNil(t, errResponse)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, http.StatusInternalServerError, errResponse.StatusCode)
	assert.Equal(t, errMessage.Error(), errResponse.Message)
}

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

	assert.EqualValues(t, http.StatusNoContent, rr.Code)
}

func TestGetFair(t *testing.T) {
	getFairService = func(neighborhood string) ([]*fair.Fair, error) {
		return []*fair.Fair{{}}, nil
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
	assert.Nil(t, err)
	assert.NotNil(t, fairs)
	assert.Equal(t, 1, fairs.Total)
	assert.EqualValues(t, http.StatusOK, rr.Code)
}
