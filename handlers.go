package main

import (
	"net/http"
)

func handleSomeEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here be dragons"))
}
