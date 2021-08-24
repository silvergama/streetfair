package v1

import "net/http"

func FairHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addFair(w, r)
	case http.MethodPut:
		updateFair(w, r)
	case http.MethodDelete:
		deleteFair(w, r)
	case http.MethodGet:
		getFair(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func addFair(w http.ResponseWriter, r *http.Request) {

}

func updateFair(w http.ResponseWriter, r *http.Request) {

}

func deleteFair(w http.ResponseWriter, r *http.Request) {

}

func getFair(w http.ResponseWriter, r *http.Request) {

}
