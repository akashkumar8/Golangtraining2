package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	http.ListenAndServe(":5000", router)

}

func test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"555"})
}
