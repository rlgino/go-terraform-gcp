package main

import (
	"net/http"

	"github.com/rlgino/go-gcp-cloud-function/handlers"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.HelloHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error listening " + err.Error())
	}
}
