package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8000", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
