package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":3000", nil)
}
