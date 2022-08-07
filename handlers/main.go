package handlers

import (
	"encoding/json"
	"github.com/harrisoncramer/golang-webhook/state"
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
	if state.Processing {
		http.Error(w, "Jobs already processing", 503)
		return
	}

	state.Processing = true

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
