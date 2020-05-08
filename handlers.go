package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dan9186/superman/logins"
)

func handleSomeEndpoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("failed to read message body: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to read message body"))
		return
	}

	defer r.Body.Close()

	var e logins.Event
	err = json.Unmarshal(body, &e)
	if err != nil {
		log.Errorf("failed to unmarshal body as a login event: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cannot translate request body to a login event"))
		return
	}

	err = e.Store(db)
	if err != nil {
		log.Errorf("failed to store login event: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to store login event"))
		return
	}

	a, err := e.Analyze()
	if err != nil {
		log.Errorf("failed to analyze event: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to analyze event"))
		return
	}

	resp, err := json.Marshal(a)
	if err != nil {
		log.Errorf("failed to marshal analysis: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create response"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
