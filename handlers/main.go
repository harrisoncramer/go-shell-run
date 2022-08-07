package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true }`)
}

type Jobs struct {
	Jobs []string
}

func RestartHandler(w http.ResponseWriter, req *http.Request) {

	var jobs Jobs

	err := json.NewDecoder(req.Body).Decode(&jobs)
	if err != nil {
		http.Error(w, "Must provide jobs.", http.StatusUnauthorized)
		return
	}

	go RunJobs(jobs.Jobs)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"running": true }`)

}
