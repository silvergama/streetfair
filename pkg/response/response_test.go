package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrite(t *testing.T) {
	body := struct {
		X string `json:"x"`
		Y int    `json:"y"`
	}{
		X: "1",
		Y: 2,
	}
	rw := httptest.NewRecorder()
	Write(rw, body, http.StatusOK)
	rw.Flush()
	if rw.Code != http.StatusOK {
		t.Errorf("unexpected status code %d", rw.Code)
	}
	if rw.Header().Get("content-type") != "application/json" {
		t.Errorf("unexpected content type %s", rw.Header().Get("content-type"))
	}
	if rw.Body.String() != `{"x":"1","y":2}` {
		t.Errorf("unexpected body %s", rw.Body.String())
	}
}

func TestWriteServerError(t *testing.T) {
	body := "internal server error"
	rw := httptest.NewRecorder()
	WriteServerError(rw, body)
	rw.Flush()
	if rw.Code != http.StatusInternalServerError {
		t.Errorf("unexpected status code %d", rw.Code)
	}
	if rw.Header().Get("content-type") != "application/json" {
		t.Errorf("unexpected content type %s", rw.Header().Get("content-type"))
	}
	if rw.Body.String() != `{"status_code":500,"message":"internal server error"}` {
		t.Errorf("unexpected body %s", rw.Body.String())
	}
}

func TestWriteNotFound(t *testing.T) {
	body := "not found"
	rw := httptest.NewRecorder()
	WriteNotFound(rw, body)
	rw.Flush()
	if rw.Code != http.StatusNotFound {
		t.Errorf("unexpected status code %d", rw.Code)
	}
	if rw.Header().Get("content-type") != "application/json" {
		t.Errorf("unexpected content type %s", rw.Header().Get("content-type"))
	}
	if rw.Body.String() != `{"status_code":404,"message":"not found"}` {
		t.Errorf("unexpected body %s", rw.Body.String())
	}
}
func TestWriteUnprocessableEntity(t *testing.T) {
	body := "unprocessable entity"
	rw := httptest.NewRecorder()
	WriteUnprocessableEntity(rw, body)
	rw.Flush()
	if rw.Code != http.StatusUnprocessableEntity {
		t.Errorf("unexpected status code %d", rw.Code)
	}
	if rw.Header().Get("content-type") != "application/json" {
		t.Errorf("unexpected content type %s", rw.Header().Get("content-type"))
	}
	if rw.Body.String() != `{"status_code":422,"message":"unprocessable entity"}` {
		t.Errorf("unexpected body %s", rw.Body.String())
	}
}
