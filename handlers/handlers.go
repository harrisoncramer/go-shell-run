package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/harrisoncramer/go-shell-run/state"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

type Jobs struct {
	Jobs []string
}

func JobHandler(w http.ResponseWriter, req *http.Request) {
	if state.Processing {
		http.Error(w, "Jobs already processing", 503)
		return
	}

	var jobs Jobs

	err := json.NewDecoder(req.Body).Decode(&jobs)
	if err != nil {
		http.Error(w, "Could not read body", http.StatusUnauthorized)
		return
	}

	jobCount := len(jobs.Jobs)
	if jobCount == 0 {
		http.Error(w, "Must provide jobs.", http.StatusUnauthorized)
		return
	}

	state.Processing = true

	go RunJobs(jobs.Jobs)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"running": true}`)

}
