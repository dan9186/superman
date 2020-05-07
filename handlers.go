package main

import (
	"io/ioutil"
	"net/http"
)

func handleSomeEndpoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	}

	defer r.Body.Close()

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
