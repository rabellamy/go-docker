package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Status is the status response struct.
type Status struct {
	Status string `json:"status"`
}

// App loads the API set.
func App() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthCheckHandler).Methods("GET")
	return r
}

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	status := Status{"OK"}

	output, err := json.Marshal(status)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(res, string(output))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", App()))
}
