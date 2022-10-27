package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		serverHost := os.Getenv("SERVER_HOST")
		serverPort := os.Getenv("SERVER_PORT")
		serverURL := fmt.Sprintf("http://%s:%s", serverHost, serverPort)
		resp, err := http.DefaultClient.Get(serverURL)
		if err != nil {
			log.Println("failed when calling server: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		var respBody = struct {
			Message string `json:"message"`
		}{}
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			log.Println("failed when decoding response body: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(respBody)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("starting web server at http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
